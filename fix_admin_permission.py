import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "fix_admin_permission.json"


def fail(message):
    print(f"ERROR: {message}")
    sys.exit(1)


def backup(path, root, backup_root):
    relative = path.relative_to(root)
    destination = backup_root / relative

    destination.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(path, destination)


def find_permission_file(root):
    role_dir = (
        root
        / "internal"
        / "domain"
        / "role"
    )

    matches = []

    for path in role_dir.glob("*.go"):
        if path.name.endswith("_test.go"):
            continue

        content = path.read_text(
            encoding="utf-8"
        )

        if (
            "type Permission " in content
            or "func IsValidPermission" in content
        ):
            matches.append(path)

    if not matches:
        fail(
            "Could not find the Go file containing "
            "Permission or IsValidPermission."
        )

    # Prefer the file containing both.
    for path in matches:
        content = path.read_text(
            encoding="utf-8"
        )

        if (
            "type Permission " in content
            and "func IsValidPermission" in content
        ):
            return path

    return matches[0]


def patch_permission(
    root,
    backup_root,
    constant,
    value,
):
    path = find_permission_file(root)

    print(
        f"Permission domain: "
        f"{path.relative_to(root)}"
    )

    content = path.read_text(
        encoding="utf-8"
    )

    original = content

    #
    # Add constant if missing.
    #
    if constant not in content:
        type_index = content.find(
            "type Permission "
        )

        if type_index == -1:
            fail(
                "Permission type declaration "
                "was not found."
            )

        const_start = content.find(
            "const (",
            type_index,
        )

        if const_start == -1:
            fail(
                "Could not find Permission "
                "const block."
            )

        const_end = content.find(
            ")",
            const_start,
        )

        if const_end == -1:
            fail(
                "Could not find end of "
                "Permission const block."
            )

        declaration = (
            f'\t{constant} Permission = '
            f'"{value}"\n'
        )

        content = (
            content[:const_end]
            + declaration
            + content[const_end:]
        )

        print(
            f"ADDED      {constant}"
        )
    else:
        print(
            f"FOUND      {constant}"
        )

    #
    # Patch IsValidPermission.
    #
    function_index = content.find(
        "func IsValidPermission"
    )

    if function_index == -1:
        fail(
            "Could not find "
            "IsValidPermission()."
        )

    function_content = content[
        function_index:
    ]

    if constant in function_content:
        print(
            "FOUND      permission validation"
        )
    else:
        #
        # Support the common switch:
        #
        # switch permission {
        # case PermissionA,
        #      PermissionB:
        #     return true
        # }
        #
        case_index = function_content.find(
            "case "
        )

        if case_index == -1:
            fail(
                "IsValidPermission exists, "
                "but no switch case was found. "
                "Refusing unsafe patch."
            )

        absolute_case_index = (
            function_index + case_index
        )

        insert_at = (
            absolute_case_index
            + len("case ")
        )

        content = (
            content[:insert_at]
            + constant
            + ",\n\t\t"
            + content[insert_at:]
        )

        print(
            "PATCHED    IsValidPermission"
        )

    if content == original:
        print(
            f"UNCHANGED  "
            f"{path.relative_to(root)}"
        )
        return

    backup(
        path,
        root,
        backup_root,
    )

    path.write_text(
        content,
        encoding="utf-8",
    )

    print(
        f"UPDATED    "
        f"{path.relative_to(root)}"
    )


def patch_role_repository(
    root,
    backup_root,
    admin,
    permission_constant,
):
    path = (
        root
        / "internal"
        / "repository"
        / "memory"
        / "role_repository.go"
    )

    if not path.exists():
        fail(
            "role_repository.go not found"
        )

    content = path.read_text(
        encoding="utf-8"
    )

    original = content

    role_id = admin["id"]
    name = admin["name"]
    description = admin["description"]

    #
    # Add Administrator construction.
    #
    if f'role.ID("{role_id}")' not in content:
        marker = "\treturn &RoleRepository{"

        if marker not in content:
            fail(
                "Could not locate repository "
                "constructor return."
            )

        admin_code = f'''
\tadminRole, err := role.New(
\t\trole.ID("{role_id}"),
\t\t"{name}",
\t\t"{description}",
\t\t[]role.Permission{{
\t\t\trole.{permission_constant},
\t\t}},
\t\ttrue,
\t)

\tif err != nil {{
\t\tpanic(err)
\t}}

'''

        content = content.replace(
            marker,
            admin_code + marker,
            1,
        )

        print(
            "ADDED      Administrator role"
        )
    else:
        print(
            "FOUND      Administrator role"
        )

    #
    # Add Administrator to map.
    #
    if "adminRole.ID: adminRole" not in content:
        marker = (
            "memberRole.ID: memberRole,"
        )

        if marker not in content:
            fail(
                "Could not locate memberRole "
                "repository map entry."
            )

        content = content.replace(
            marker,
            marker
            + "\n\t\t\tadminRole.ID:  adminRole,",
            1,
        )

        print(
            "ADDED      Administrator to role map"
        )

    if content == original:
        print(
            "UNCHANGED  "
            "internal/repository/memory/"
            "role_repository.go"
        )
        return

    backup(
        path,
        root,
        backup_root,
    )

    path.write_text(
        content,
        encoding="utf-8",
    )

    print(
        "UPDATED    "
        "internal/repository/memory/"
        "role_repository.go"
    )


def main():
    if not CONFIG_FILE.exists():
        fail(
            f"{CONFIG_FILE} not found"
        )

    with CONFIG_FILE.open(
        "r",
        encoding="utf-8",
    ) as file:
        config = json.load(file)

    root = (
        SCRIPT_DIR
        / config["project"]
    )

    if not root.exists():
        fail(
            f"Project not found: {root}"
        )

    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "fix-admin-permission"
    )

    print(f"Project: {root}")
    print()

    permission = config["permission"]

    patch_permission(
        root,
        backup_root,
        permission["constant"],
        permission["value"],
    )

    patch_role_repository(
        root,
        backup_root,
        config["administrator"],
        permission["constant"],
    )

    print()
    print(
        "Administrator permission fix complete."
    )
    print(
        f"Backups: {backup_root}"
    )


if __name__ == "__main__":
    main()