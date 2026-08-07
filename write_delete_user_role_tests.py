import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "delete_user_role_tests.json"


def info(message):
    print(f"[INFO] {message}")


def ok(message):
    print(f"[ OK ] {message}")


def warn(message):
    print(f"[WARN] {message}")


def fail(message):
    print()
    print(f"[FAIL] {message}")
    sys.exit(1)


def load_config():
    print()
    print("=== DELETE USER ROLE TEST GENERATOR ===")
    print()

    info(f"Looking for config: {CONFIG_FILE}")

    if not CONFIG_FILE.exists():
        fail(
            f"Configuration file not found: "
            f"{CONFIG_FILE}"
        )

    ok("Configuration file exists")

    info("Opening configuration file")

    try:
        with CONFIG_FILE.open(
            "r",
            encoding="utf-8",
        ) as file:
            config = json.load(file)

    except json.JSONDecodeError as error:
        fail(
            f"Invalid JSON at line "
            f"{error.lineno}, "
            f"column {error.colno}: "
            f"{error.msg}"
        )

    except Exception as error:
        fail(
            f"Could not read configuration: "
            f"{error}"
        )

    ok("JSON parsed successfully")

    return config


def validate_config(config):
    info("Validating 'project' field")

    project = config.get("project")

    if not project:
        fail("Missing JSON field: project")

    ok(f"Project: {project}")

    info("Validating 'files' field")

    files = config.get("files")

    if not isinstance(files, list):
        fail(
            "'files' must be a JSON array"
        )

    if len(files) == 0:
        fail(
            "'files' cannot be empty"
        )

    ok(
        f"Found {len(files)} test files "
        f"in configuration"
    )

    for index, item in enumerate(
        files,
        start=1,
    ):
        info(
            f"Validating file entry #{index}"
        )

        if not isinstance(item, dict):
            fail(
                f"File entry #{index} "
                f"is not an object"
            )

        if not item.get("path"):
            fail(
                f"File entry #{index} "
                f"is missing path"
            )

        if "content" not in item:
            fail(
                f"File entry #{index} "
                f"is missing content"
            )

        ok(
            f"File entry #{index}: "
            f"{item['path']}"
        )

    return project, files


def resolve_project(project):
    info(
        "Resolving project directory"
    )

    root = SCRIPT_DIR / project

    info(f"Resolved path: {root}")

    if not root.exists():
        fail(
            f"Project directory does not exist: "
            f"{root}"
        )

    ok("Project directory exists")

    info(
        "Checking for go.mod"
    )

    go_mod = root / "go.mod"

    if not go_mod.exists():
        fail(
            f"go.mod not found in {root}. "
            f"This does not appear to be "
            f"the Go project root."
        )

    ok("go.mod found")

    return root


def validate_source_dependencies(root):
    print()
    print("=== DEPENDENCY CHECKS ===")
    print()

    checks = [
        (
            "User.RemoveRole",
            root / "internal/domain/user/user.go",
            "func (u *User) RemoveRole(",
        ),
        (
            "userrole.Service.Remove",
            root
            / "internal/application/userrole/remove.go",
            "func (s *Service) Remove(",
        ),
        (
            "ErrProtectedRole",
            root
            / "internal/application/userrole/remove.go",
            "ErrProtectedRole",
        ),
        (
            "Administrator repository role",
            root
            / "internal/repository/memory/role_repository.go",
            'role.ID("role_admin")',
        ),
        (
            "PermissionManageUsers",
            root
            / "internal/domain/role/permission.go",
            "PermissionManageUsers",
        ),
    ]

    for name, path, marker in checks:
        info(
            f"Checking dependency: {name}"
        )

        info(
            f"Reading: "
            f"{path.relative_to(root)}"
        )

        if not path.exists():
            fail(
                f"Dependency file missing: "
                f"{path.relative_to(root)}"
            )

        content = path.read_text(
            encoding="utf-8",
        )

        ok(
            f"Read "
            f"{path.relative_to(root)}"
        )

        info(
            f"Searching for: {marker}"
        )

        if marker not in content:
            print()
            print(
                f"[FAIL] Dependency missing: "
                f"{name}"
            )
            print(
                f"[FAIL] Expected marker: "
                f"{marker}"
            )
            print(
                f"[FAIL] File: "
                f"{path.relative_to(root)}"
            )
            print()
            print(
                "Run the DELETE user-role "
                "API generator first."
            )
            sys.exit(1)

        ok(
            f"Dependency satisfied: {name}"
        )


def create_backup(
    root,
    destination,
    backup_root,
):
    relative = destination.relative_to(
        root
    )

    backup = backup_root / relative

    info(
        f"Creating backup for: {relative}"
    )

    backup.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        destination,
        backup,
    )

    ok(
        f"Backup created: {backup}"
    )


def write_test_file(
    root,
    item,
    backup_root,
):
    relative_path = item["path"]
    expected_content = item["content"]

    print()
    print(
        f"--- {relative_path} ---"
    )

    destination = root / relative_path

    info(
        f"Destination: {destination}"
    )

    info(
        "Checking destination parent"
    )

    destination.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    ok(
        "Destination parent exists"
    )

    if destination.exists():
        warn(
            "Test file already exists"
        )

        info(
            "Reading existing file"
        )

        existing = destination.read_text(
            encoding="utf-8",
        )

        ok(
            f"Existing file read "
            f"({len(existing)} chars)"
        )

        info(
            "Comparing existing content "
            "with generated content"
        )

        if existing == expected_content:
            ok(
                "Existing file is identical"
            )
            ok(
                "No write required"
            )
            return

        print()
        print(
            "[WARN] Existing test file differs "
            "from generated version."
        )

        create_backup(
            root,
            destination,
            backup_root,
        )

        info(
            "Existing file will be replaced "
            "after backup"
        )

    else:
        ok(
            "Test file does not exist; "
            "safe to create"
        )

    info(
        "Validating generated test content"
    )

    if not expected_content.strip():
        fail(
            f"Generated content for "
            f"{relative_path} is empty"
        )

    if "package " not in expected_content:
        fail(
            f"Generated Go test "
            f"{relative_path} has no "
            f"package declaration"
        )

    if "func Test" not in expected_content:
        fail(
            f"Generated Go test "
            f"{relative_path} contains "
            f"no Test functions"
        )

    ok(
        "Basic Go test validation passed"
    )

    info(
        f"Writing {relative_path}"
    )

    try:
        destination.write_text(
            expected_content,
            encoding="utf-8",
        )

    except Exception as error:
        fail(
            f"Could not write "
            f"{relative_path}: {error}"
        )

    ok(
        f"Written: {relative_path}"
    )

    info(
        "Reading file back for verification"
    )

    written = destination.read_text(
        encoding="utf-8",
    )

    if written != expected_content:
        fail(
            f"Verification failed for "
            f"{relative_path}"
        )

    ok(
        "Written content verified"
    )


def main():
    config = load_config()

    project, files = validate_config(
        config
    )

    root = resolve_project(
        project
    )

    validate_source_dependencies(
        root
    )

    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "delete-user-role-tests"
    )

    print()
    info(
        f"Backup root: {backup_root}"
    )

    print()
    print("=== WRITING TEST FILES ===")

    for index, item in enumerate(
        files,
        start=1,
    ):
        info(
            f"Processing file "
            f"{index}/{len(files)}"
        )

        write_test_file(
            root,
            item,
            backup_root,
        )

    print()
    print("==============================")
    print("TEST GENERATION COMPLETE")
    print("==============================")
    print()

    for item in files:
        print(
            f"[ OK ] {item['path']}"
        )

    print()
    print("Run:")
    print()
    print("  cd prayer-api")
    print("  go fmt ./...")
    print(
        "  go test -v "
        "./internal/domain/user"
    )
    print(
        "  go test -v "
        "./internal/application/userrole"
    )


if __name__ == "__main__":
    main()