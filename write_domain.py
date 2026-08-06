import json
import sys
from pathlib import Path


CONFIG_FILE = "domain.json"


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
        print("ERROR: project is missing from domain.json")
        sys.exit(1)

    project_root = Path(project_name)

    if project_root.exists():
        print(f"Using existing project: {project_root}")
    else:
        print(f"Creating new project directory: {project_root}")
        project_root.mkdir(parents=True, exist_ok=True)

    return project_root


def write_file(file_path: Path, content: str) -> str:
    file_path.parent.mkdir(
        parents=True,
        exist_ok=True
    )

    if not file_path.exists():
        file_path.write_text(
            content,
            encoding="utf-8"
        )

        return "CREATED"

    existing_content = file_path.read_text(
        encoding="utf-8"
    )

    if existing_content == content:
        return "UNCHANGED"

    file_path.write_text(
        content,
        encoding="utf-8"
    )

    return "UPDATED"


def write_files(project_root: Path, files: dict):
    created = 0
    updated = 0
    unchanged = 0

    print()

    for relative_path, content in files.items():
        file_path = project_root / relative_path

        status = write_file(
            file_path,
            content
        )

        print(
            f"{status:<10} {file_path}"
        )

        if status == "CREATED":
            created += 1

        elif status == "UPDATED":
            updated += 1

        elif status == "UNCHANGED":
            unchanged += 1

    return created, updated, unchanged


def main():
    config = load_config()

    project_root = resolve_project_root(
        config
    )

    files = config.get("files", {})

    if not files:
        print("No files defined.")
        return

    created, updated, unchanged = write_files(
        project_root,
        files
    )

    print()
    print("Domain generation complete.")
    print()
    print(f"Created:   {created}")
    print(f"Updated:   {updated}")
    print(f"Unchanged: {unchanged}")


if __name__ == "__main__":
    main()