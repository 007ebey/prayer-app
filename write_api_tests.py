import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "api_tests.json"


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
        print("ERROR: project missing from api_tests.json")
        sys.exit(1)

    # Resolve relative to the Python script,
    # not the PowerShell working directory.
    root = SCRIPT_DIR / project_name

    if root.exists():
        print(f"Using project: {root}")
    else:
        print(f"Creating project: {root}")
        root.mkdir(
            parents=True,
            exist_ok=True,
        )

    return root


def write_files(
    root: Path,
    files: dict,
):
    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "api-tests"
    )

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

            print(
                f"CREATED    {relative_path}"
            )

            created += 1
            continue

        existing = destination.read_text(
            encoding="utf-8"
        )

        if existing == content:
            print(
                f"UNCHANGED  {relative_path}"
            )

            unchanged += 1
            continue

        backup = backup_root / relative_path

        backup.parent.mkdir(
            parents=True,
            exist_ok=True,
        )

        shutil.copy2(
            destination,
            backup,
        )

        destination.write_text(
            content,
            encoding="utf-8",
        )

        print(
            f"UPDATED    {relative_path}"
        )

        updated += 1

    return (
        created,
        updated,
        unchanged,
        backup_root,
    )


def main():
    config = load_config()

    root = resolve_project_root(
        config
    )

    files = config.get(
        "files",
        {},
    )

    if not files:
        print("No files defined.")
        return

    print()

    (
        created,
        updated,
        unchanged,
        backup_root,
    ) = write_files(
        root,
        files,
    )

    print()
    print("API test generation complete.")
    print(f"Created:   {created}")
    print(f"Updated:   {updated}")
    print(f"Unchanged: {unchanged}")

    if updated > 0:
        print(
            f"Backups: {backup_root}"
        )


if __name__ == "__main__":
    main()