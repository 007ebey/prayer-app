import json
import shutil
import sys
from datetime import datetime
from pathlib import Path
import os


CONFIG_FILE = "test.json"


def load_config() -> dict:
    path = Path(CONFIG_FILE)
    print(f"Loading configuration from {path}")
    print(f"Current working directory: {os.getcwd()}")

    if not path.exists():
        print(f"ERROR: {CONFIG_FILE} not found!")
        sys.exit(1)

    with path.open("r", encoding="utf-8") as file:
        return json.load(file)


def resolve_project_root(config: dict) -> Path:
    project_name = config.get("project")

    if not project_name:
        print(f"ERROR: project missing from {CONFIG_FILE}")
        sys.exit(1)

    root = Path(project_name)

    if root.exists():
        print(f"Using existing project: {root}")
    else:
        print(f"Creating project: {root}")
        root.mkdir(parents=True)

    return root


def write_files(root: Path, files: dict):
    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "tests"
    )

    created = 0
    updated = 0
    unchanged = 0

    for relative_path, content in files.items():
        destination = root / relative_path

        destination.parent.mkdir(
            parents=True,
            exist_ok=True
        )

        if not destination.exists():
            destination.write_text(
                content,
                encoding="utf-8"
            )

            print(f"CREATED    {destination}")
            created += 1
            continue

        existing = destination.read_text(
            encoding="utf-8"
        )

        if existing == content:
            print(f"UNCHANGED  {destination}")
            unchanged += 1
            continue

        backup = backup_root / relative_path

        backup.parent.mkdir(
            parents=True,
            exist_ok=True
        )

        shutil.copy2(
            destination,
            backup
        )

        destination.write_text(
            content,
            encoding="utf-8"
        )

        print(f"UPDATED    {destination}")
        updated += 1

    return created, updated, unchanged, backup_root


def main():
    config = load_config()

    root = resolve_project_root(config)

    files = config.get("files", {})

    if not files:
        print("No test files defined.")
        return

    print()

    created, updated, unchanged, backup_root = write_files(
        root,
        files
    )

    print()
    print("Test generation complete.")
    print(f"Created:   {created}")
    print(f"Updated:   {updated}")
    print(f"Unchanged: {unchanged}")

    if updated:
        print(f"Backups:   {backup_root}")


if __name__ == "__main__":
    main()