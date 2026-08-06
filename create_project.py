import json
import subprocess
from pathlib import Path


CONFIG_FILE = "project.json"


def load_config():
    with open(CONFIG_FILE, "r", encoding="utf-8") as file:
        return json.load(file)


def create_files(root: Path, files: dict[str, str]):
    for relative_path, content in files.items():
        file_path = root / relative_path

        # Create parent directories automatically
        file_path.parent.mkdir(parents=True, exist_ok=True)

        file_path.write_text(
            content,
            encoding="utf-8"
        )

        print(f"Created: {file_path}")


def run_commands(root: Path, commands: list[str]):
    for command in commands:
        print(f"Running: {command}")

        subprocess.run(
            command,
            cwd=root,
            shell=True,
            check=True
        )


def main():
    config = load_config()

    project_name = config["project"]
    files = config.get("files", {})
    commands = config.get("commands", [])

    root = Path(project_name)

    root.mkdir(
        parents=True,
        exist_ok=True
    )

    print(f"\nCreating project: {root}\n")

    create_files(root, files)

    if commands:
        print("\nRunning setup commands...\n")
        run_commands(root, commands)

    print("\nProject created successfully.")
    print(f"\ncd {project_name}")
    print("go run ./cmd/api")


if __name__ == "__main__":
    main()