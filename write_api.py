import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


CONFIG_FILE = "api.json"


def load_config() -> dict:
    config_path = Path(CONFIG_FILE)

    if not config_path.exists():
        print(f"ERROR: {CONFIG_FILE} was not found.")
        sys.exit(1)

    with config_path.open("r", encoding="utf-8") as file:
        return json.load(file)


def resolve_project_root(config: dict) -> Path:
    project_name = config.get("project")

    if not project_name:
        print("ERROR: project is missing from api.json")
        sys.exit(1)

    project_root = Path(project_name)

    if project_root.exists():
        print(f"Using existing project: {project_root}")
    else:
        print(f"Creating project: {project_root}")
        project_root.mkdir(parents=True, exist_ok=True)

    return project_root


def backup_file(
    project_root: Path,
    file_path: Path,
    relative_path: str,
    backup_root: Path,
):
    backup_path = backup_root / relative_path

    backup_path.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        file_path,
        backup_path,
    )


def write_files(
    project_root: Path,
    files: dict,
):
    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        project_root /
        ".generator-backups" /
        timestamp
    )

    created = 0
    updated = 0
    unchanged = 0

    for relative_path, content in files.items():
        file_path = project_root / relative_path

        file_path.parent.mkdir(
            parents=True,
            exist_ok=True,
        )

        if not file_path.exists():
            file_path.write_text(
                content,
                encoding="utf-8",
            )

            print(
                f"CREATED    {file_path}"
            )

            created += 1
            continue

        existing = file_path.read_text(
            encoding="utf-8"
        )

        if existing == content:
            print(
                f"UNCHANGED  {file_path}"
            )

            unchanged += 1
            continue

        backup_file(
            project_root,
            file_path,
            relative_path,
            backup_root,
        )

        file_path.write_text(
            content,
            encoding="utf-8",
        )

        print(
            f"UPDATED    {file_path}"
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

    project_root = resolve_project_root(
        config
    )

    files = config.get("files", {})

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
        project_root,
        files,
    )

    print()
    print("API generation complete.")
    print()
    print(f"Created:   {created}")
    print(f"Updated:   {updated}")
    print(f"Unchanged: {unchanged}")

    if updated > 0:
        print()
        print(
            f"Backups: {backup_root}"
        )


if __name__ == "__main__":
    main()