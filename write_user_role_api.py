import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "user_role_api.json"


def load_config() -> dict:
    if not CONFIG_FILE.exists():
        print(f"ERROR: {CONFIG_FILE} not found.")
        sys.exit(1)

    with CONFIG_FILE.open(
        "r",
        encoding="utf-8",
    ) as file:
        return json.load(file)


def backup_file(
    root: Path,
    destination: Path,
    relative_path: str,
    backup_root: Path,
):
    if not destination.exists():
        return

    backup = backup_root / relative_path

    backup.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        destination,
        backup,
    )


def write_files(
    root: Path,
    files: dict,
    backup_root: Path,
):
    for relative_path, content in files.items():
        destination = root / relative_path

        destination.parent.mkdir(
            parents=True,
            exist_ok=True,
        )

        if destination.exists():
            existing = destination.read_text(
                encoding="utf-8"
            )

            if existing == content:
                print(
                    f"UNCHANGED  {relative_path}"
                )
                continue

            backup_file(
                root,
                destination,
                relative_path,
                backup_root,
            )

            destination.write_text(
                content,
                encoding="utf-8",
            )

            print(
                f"UPDATED    {relative_path}"
            )

        else:
            destination.write_text(
                content,
                encoding="utf-8",
            )

            print(
                f"CREATED    {relative_path}"
            )


def patch_role_domain(
    root: Path,
    backup_root: Path,
):
    relative_path = (
        "internal/domain/role/role.go"
    )

    path = root / relative_path

    if not path.exists():
        print(
            f"WARNING: cannot patch {relative_path}"
        )
        return

    content = path.read_text(
        encoding="utf-8"
    )

    if "PermissionManageUsers" in content:
        print(
            f"UNCHANGED  {relative_path} "
            "(manage_users already exists)"
        )
        return

    backup_file(
        root,
        path,
        relative_path,
        backup_root,
    )

    # Locate the Permission const block.
    marker = "type Permission string"

    marker_index = content.find(marker)

    if marker_index == -1:
        print(
            "ERROR: Permission type not found in "
            f"{relative_path}"
        )
        sys.exit(1)

    const_index = content.find(
        "const (",
        marker_index,
    )

    if const_index == -1:
        print(
            "ERROR: Permission const block not found."
        )
        sys.exit(1)

    closing_index = content.find(
        ")",
        const_index,
    )

    if closing_index == -1:
        print(
            "ERROR: Permission const block "
            "closing ')' not found."
        )
        sys.exit(1)

    new_constant = (
        '\tPermissionManageUsers '
        'Permission = "manage_users"\n'
    )

    content = (
        content[:closing_index]
        + new_constant
        + content[closing_index:]
    )

    # Try to locate permission validation.
    #
    # Most versions of our domain contain a
    # switch over permissions. Add the new
    # permission beside the existing ones.
    if (
        "case PermissionViewPrayerSessions,"
        in content
    ):
        content = content.replace(
            "case PermissionViewPrayerSessions,",
            "case PermissionManageUsers, "
            "PermissionViewPrayerSessions,",
            1,
        )

    elif (
        "PermissionViewPrayerSessions:"
        in content
    ):
        # Map-based validation.
        marker = "PermissionViewPrayerSessions:"

        map_index = content.find(marker)

        line_start = content.rfind(
            "\n",
            0,
            map_index,
        ) + 1

        indentation = (
            content[line_start:map_index]
        )

        content = (
            content[:line_start]
            + indentation
            + "PermissionManageUsers: true,\n"
            + content[line_start:]
        )

    else:
        print()
        print(
            "WARNING: PermissionManageUsers "
            "constant was added, but the writer "
            "could not safely detect your "
            "permission validation structure."
        )
        print(
            "Open role.go and add "
            "PermissionManageUsers to the "
            "allowed permission set."
        )
        print()

    path.write_text(
        content,
        encoding="utf-8",
    )

    print(
        f"PATCHED    {relative_path}"
    )


def patch_role_repository(
    root: Path,
    backup_root: Path,
):
    relative_path = (
        "internal/repository/memory/"
        "role_repository.go"
    )

    path = root / relative_path

    if not path.exists():
        print(
            f"WARNING: cannot patch {relative_path}"
        )
        return

    content = path.read_text(
        encoding="utf-8"
    )

    if 'role.ID("role_admin")' in content:
        print(
            f"UNCHANGED  {relative_path} "
            "(Administrator already exists)"
        )
        return

    backup_file(
        root,
        path,
        relative_path,
        backup_root,
    )

    # Insert Administrator construction before
    # the repository return.
    marker = "\treturn &RoleRepository{"

    if marker not in content:
        print(
            "ERROR: could not locate "
            "RoleRepository constructor return."
        )
        sys.exit(1)

    admin_code = '''
\tadminRole, err := role.New(
\t\trole.ID("role_admin"),
\t\t"Administrator",
\t\t"Administrative user management access.",
\t\t[]role.Permission{
\t\t\trole.PermissionManageUsers,
\t\t},
\t\tfalse,
\t)

\tif err != nil {
\t\tpanic(err)
\t}

'''

    content = content.replace(
        marker,
        admin_code + marker,
        1,
    )

    # Add role to the existing map.
    #
    # We know memberRole.ID is currently
    # inserted into the roles map.
    map_marker = (
        "memberRole.ID: memberRole,"
    )

    if map_marker not in content:
        print(
            "ERROR: could not locate roles map."
        )
        sys.exit(1)

    content = content.replace(
        map_marker,
        map_marker
        + "\n\t\t\tadminRole.ID:  adminRole,",
        1,
    )

    path.write_text(
        content,
        encoding="utf-8",
    )

    print(
        f"PATCHED    {relative_path}"
    )


def main():
    config = load_config()

    project_name = config.get("project")

    if not project_name:
        print("ERROR: project missing.")
        sys.exit(1)

    root = SCRIPT_DIR / project_name

    if not root.exists():
        print(
            f"ERROR: project not found: {root}"
        )
        sys.exit(1)

    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "user-role-api"
    )

    print(f"Project: {root}")
    print()

    write_files(
        root,
        config.get("files", {}),
        backup_root,
    )

    patch_role_domain(
        root,
        backup_root,
    )

    patch_role_repository(
        root,
        backup_root,
    )

    print()
    print(
        "User Role API generation complete."
    )
    print(
        f"Backups: {backup_root}"
    )


if __name__ == "__main__":
    main()