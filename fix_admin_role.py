import json
import shutil
import sys

from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "fix_admin_role.json"


def log(message: str):
    print(f"[INFO] {message}")


def ok(message: str):
    print(f"[ OK ] {message}")


def warn(message: str):
    print(f"[WARN] {message}")


def print_source(content: str):
    print()
    print("========== CURRENT SOURCE ==========")

    for number, line in enumerate(
        content.splitlines(),
        start=1,
    ):
        print(f"{number:4}: {line}")

    print("====================================")
    print()


def fail(message: str, content: str | None = None):
    print()
    print(f"[FAIL] {message}")

    if content is not None:
        print_source(content)

    sys.exit(1)


def load_config():
    log(f"Looking for config: {CONFIG_FILE}")

    if not CONFIG_FILE.exists():
        fail(
            f"Config file does not exist: "
            f"{CONFIG_FILE}"
        )

    ok("Config file exists")

    log("Reading JSON configuration")

    try:
        with CONFIG_FILE.open(
            "r",
            encoding="utf-8",
        ) as file:
            config = json.load(file)
    except Exception as error:
        fail(
            f"Could not parse JSON: {error}"
        )

    ok("JSON configuration loaded")

    return config


def resolve_project(config):
    log("Reading project from configuration")

    project = config.get("project")

    if not project:
        fail(
            "JSON does not contain 'project'"
        )

    ok(f"Project configured as: {project}")

    root = SCRIPT_DIR / project

    log(f"Checking project path: {root}")

    if not root.exists():
        fail(
            f"Project directory does not exist: "
            f"{root}"
        )

    ok("Project directory exists")

    return root


def load_repository(root, config):
    relative_path = config.get("path")

    log(
        "Reading repository path from "
        "configuration"
    )

    if not relative_path:
        fail(
            "JSON does not contain 'path'"
        )

    ok(
        f"Repository path: {relative_path}"
    )

    path = root / relative_path

    log(f"Checking file: {path}")

    if not path.exists():
        fail(
            f"Repository file does not exist: "
            f"{path}"
        )

    ok("Repository file exists")

    log("Reading repository source")

    try:
        content = path.read_text(
            encoding="utf-8"
        )
    except Exception as error:
        fail(
            f"Could not read repository: "
            f"{error}"
        )

    ok(
        f"Repository loaded "
        f"({len(content)} characters)"
    )

    return path, relative_path, content


def validate_admin_config(config):
    log("Validating Administrator configuration")

    admin = config.get("admin")

    if not admin:
        fail(
            "JSON does not contain 'admin'"
        )

    required = [
        "id",
        "name",
        "description",
        "permissions",
    ]

    for field in required:
        log(
            f"Checking admin field: {field}"
        )

        if field not in admin:
            fail(
                f"Missing admin field: {field}"
            )

        ok(
            f"Admin field exists: {field}"
        )

    if not admin["permissions"]:
        fail(
            "Administrator permissions "
            "cannot be empty"
        )

    ok(
        f"Administrator has "
        f"{len(admin['permissions'])} "
        f"configured permissions"
    )

    return admin


def inspect_source(content, admin):
    log(
        "Checking whether role_admin "
        "already exists"
    )

    admin_id_literal = (
        f'role.ID("{admin["id"]}")'
    )

    admin_exists = (
        admin_id_literal in content
    )

    if admin_exists:
        ok(
            f"Found existing "
            f"{admin_id_literal}"
        )
    else:
        log(
            f"{admin_id_literal} "
            f"does not exist yet"
        )

    log(
        "Checking for NewRoleRepository"
    )

    constructor_marker = (
        "func NewRoleRepository() "
        "*RoleRepository {"
    )

    if constructor_marker not in content:
        fail(
            "Could not find "
            "NewRoleRepository constructor",
            content,
        )

    ok("Found NewRoleRepository")

    log(
        "Checking for memberRole creation"
    )

    if "memberRole, err := role.New(" not in content:
        fail(
            "Could not find memberRole "
            "construction",
            content,
        )

    ok("Found memberRole construction")

    log(
        "Checking for memberRole error guard"
    )

    member_start = content.find(
        "memberRole, err := role.New("
    )

    return_start = content.find(
        "return &RoleRepository{",
        member_start,
    )

    if return_start == -1:
        fail(
            "Could not find "
            "'return &RoleRepository{' "
            "after memberRole construction",
            content,
        )

    ok(
        "Found RoleRepository return block"
    )

    log(
        "Checking repository roles map"
    )

    map_marker = (
        "memberRole.ID: memberRole,"
    )

    if map_marker not in content:
        fail(
            "Could not find memberRole "
            "inside repository roles map",
            content,
        )

    ok(
        "Found memberRole in roles map"
    )

    return {
        "admin_exists": admin_exists,
        "return_start": return_start,
        "map_marker": map_marker,
    }


def build_admin_code(admin):
    log(
        "Generating Administrator Go code"
    )

    permission_lines = []

    for permission in admin["permissions"]:
        log(
            f"Adding permission to generated "
            f"role: {permission}"
        )

        permission_lines.append(
            f"\t\t\trole.{permission},"
        )

    permissions = "\n".join(
        permission_lines
    )

    code = (
        "\n"
        "\tadminRole, err := role.New(\n"
        f'\t\trole.ID("{admin["id"]}"),\n'
        f'\t\t"{admin["name"]}",\n'
        f'\t\t"{admin["description"]}",\n'
        "\t\t[]role.Permission{\n"
        f"{permissions}\n"
        "\t\t},\n"
        "\t\ttrue,\n"
        "\t)\n"
        "\n"
        "\tif err != nil {\n"
        "\t\tpanic(err)\n"
        "\t}\n"
        "\n"
    )

    ok("Administrator Go code generated")

    return code


def patch_source(
    content,
    admin,
    inspection,
):
    updated = content

    if inspection["admin_exists"]:
        log(
            "Administrator construction "
            "already exists"
        )

        if (
            "adminRole.ID:  adminRole"
            in updated
            or
            "adminRole.ID: adminRole"
            in updated
        ):
            ok(
                "Administrator already exists "
                "in roles map"
            )

            return updated

        warn(
            "Administrator is constructed but "
            "not present in roles map"
        )

    else:
        log(
            "Administrator does not exist; "
            "inserting before repository return"
        )

        admin_code = build_admin_code(
            admin
        )

        return_marker = (
            "\treturn &RoleRepository{"
        )

        if return_marker not in updated:
            # gofmt/spacing fallback
            return_marker = (
                "return &RoleRepository{"
            )

        if return_marker not in updated:
            fail(
                "Could not locate repository "
                "return insertion point",
                updated,
            )

        updated = updated.replace(
            return_marker,
            admin_code + return_marker,
            1,
        )

        ok(
            "Administrator construction inserted"
        )

    log(
        "Checking whether Administrator "
        "is in roles map"
    )

    if (
        "adminRole.ID:  adminRole"
        not in updated
        and
        "adminRole.ID: adminRole"
        not in updated
    ):
        map_marker = (
            "memberRole.ID: memberRole,"
        )

        log(
            f"Looking for map marker: "
            f"{map_marker}"
        )

        if map_marker not in updated:
            fail(
                "memberRole map entry disappeared "
                "during patching",
                updated,
            )

        updated = updated.replace(
            map_marker,
            map_marker
            + "\n"
            + "\t\t\tadminRole.ID:  adminRole,",
            1,
        )

        ok(
            "Administrator added to roles map"
        )

    else:
        ok(
            "Administrator already present "
            "in roles map"
        )

    return updated


def validate_result(content, admin):
    print()
    log("Validating generated source")

    checks = {
        "Administrator ID":
            f'role.ID("{admin["id"]}")',

        "Administrator variable":
            "adminRole",

        "Administrator map entry":
            "adminRole.ID",

        "manage users permission":
            "role.PermissionManageUsers",
    }

    for name, marker in checks.items():
        log(
            f"Checking result: {name}"
        )

        if marker not in content:
            fail(
                f"Generated source is missing "
                f"{name}: {marker}",
                content,
            )

        ok(
            f"Validated: {name}"
        )

    for permission in admin["permissions"]:
        marker = f"role.{permission}"

        log(
            f"Checking generated permission: "
            f"{marker}"
        )

        if marker not in content:
            fail(
                f"Generated Administrator is "
                f"missing permission {marker}",
                content,
            )

        ok(
            f"Permission present: {marker}"
        )

    ok("Generated source passed validation")


def create_backup(
    root,
    path,
    relative_path,
):
    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_path = (
        root
        / ".generator-backups"
        / timestamp
        / "fix-admin-role"
        / relative_path
    )

    log(
        f"Creating backup: {backup_path}"
    )

    backup_path.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        path,
        backup_path,
    )

    ok("Backup created")

    return backup_path


def main():
    print()
    print(
        "=== Administrator Role Fix ==="
    )
    print()

    config = load_config()

    root = resolve_project(config)

    path, relative_path, content = (
        load_repository(
            root,
            config,
        )
    )

    admin = validate_admin_config(
        config
    )

    print()
    log("Inspecting current Go source")

    inspection = inspect_source(
        content,
        admin,
    )

    print()
    log("Applying in-memory patch")

    updated = patch_source(
        content,
        admin,
        inspection,
    )

    validate_result(
        updated,
        admin,
    )

    if updated == content:
        print()
        ok(
            "Source already has the required "
            "Administrator configuration"
        )
        print(
            "No file changes were necessary."
        )
        return

    print()
    log(
        "Patch validation succeeded. "
        "Preparing to write."
    )

    backup_path = create_backup(
        root,
        path,
        relative_path,
    )

    log(
        "Writing updated repository source"
    )

    try:
        path.write_text(
            updated,
            encoding="utf-8",
        )
    except Exception as error:
        fail(
            f"Could not write repository: "
            f"{error}"
        )

    ok("Repository source written")

    print()
    print("=== SUCCESS ===")
    print(
        f"Updated: {relative_path}"
    )
    print(
        f"Backup:  {backup_path}"
    )
    print()
    print(
        "Next:"
    )
    print(
        "  cd prayer-api"
    )
    print(
        "  go fmt ./..."
    )
    print(
        "  go test -v "
        "./internal/application/userrole"
    )
    print(
        "  go test -v ./internal/http"
    )
    print(
        "  go test -v ./..."
    )


if __name__ == "__main__":
    main()