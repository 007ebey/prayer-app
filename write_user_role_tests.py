import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "user_role_tests.json"


def load_config() -> dict:
    if not CONFIG_FILE.exists():
        print(f"ERROR: {CONFIG_FILE} not found.")
        sys.exit(1)

    print(f"Loading: {CONFIG_FILE}")

    with CONFIG_FILE.open(
        "r",
        encoding="utf-8",
    ) as file:
        return json.load(file)


def resolve_project_root(config: dict) -> Path:
    project_name = config.get("project")

    if not project_name:
        print("ERROR: project missing from JSON.")
        sys.exit(1)

    root = SCRIPT_DIR / project_name

    if not root.exists():
        print(f"ERROR: project not found: {root}")
        sys.exit(1)

    print(f"Using project: {root}")

    return root


def backup_file(
    destination: Path,
    relative_path: str,
    backup_root: Path,
):
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
    created = 0
    updated = 0
    unchanged = 0

    for relative_path, content in files.items():
        destination = root / relative_path

        destination.parent.mkdir(
            parents=True,
            exist_ok=True,
        )

        if not destination.exists():
            destination.write_text(
                content,
                encoding="utf-8",
            )

            print(f"CREATED    {relative_path}")
            created += 1
            continue

        existing = destination.read_text(
            encoding="utf-8"
        )

        if existing == content:
            print(f"UNCHANGED  {relative_path}")
            unchanged += 1
            continue

        backup_file(
            destination,
            relative_path,
            backup_root,
        )

        destination.write_text(
            content,
            encoding="utf-8",
        )

        print(f"UPDATED    {relative_path}")
        updated += 1

    return created, updated, unchanged


def main():
    config = load_config()
    root = resolve_project_root(config)

    files = config.get("files", {})

    if not files:
        print("ERROR: no test files defined.")
        sys.exit(1)

    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "user-role-tests"
    )

    print()

    created, updated, unchanged = write_files(
        root,
        files,
        backup_root,
    )

    print()
    print("User Role tests generated.")
    print(f"Created:   {created}")
    print(f"Updated:   {updated}")
    print(f"Unchanged: {unchanged}")

    if updated:
        print(f"Backups:   {backup_root}")


if __name__ == "__main__":
    main()