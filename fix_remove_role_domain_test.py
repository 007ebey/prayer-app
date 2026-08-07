import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = (
    SCRIPT_DIR
    / "fix_remove_role_domain_test.json"
)


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


def print_source(path, content):
    print()
    print(
        f"========== {path.name} =========="
    )

    for number, line in enumerate(
        content.splitlines(),
        start=1,
    ):
        print(
            f"{number:4}: {line}"
        )

    print("==============================")
    print()


def load_config():
    print()
    print(
        "================================="
    )
    print(
        "REMOVE ROLE DOMAIN TEST FIX"
    )
    print(
        "================================="
    )
    print()

    info(
        f"Looking for config: "
        f"{CONFIG_FILE}"
    )

    if not CONFIG_FILE.exists():
        fail(
            "Configuration file not found"
        )

    ok(
        "Configuration file exists"
    )

    info(
        "Loading JSON configuration"
    )

    try:
        with CONFIG_FILE.open(
            "r",
            encoding="utf-8",
        ) as file:
            config = json.load(file)

    except Exception as error:
        fail(
            f"Could not load JSON: {error}"
        )

    ok(
        "Configuration loaded"
    )

    return config


def resolve_project(config):
    project = config.get("project")

    info(
        f"Configured project: {project}"
    )

    if not project:
        fail(
            "Missing project in configuration"
        )

    root = SCRIPT_DIR / project

    info(
        f"Checking project: {root}"
    )

    if not root.exists():
        fail(
            f"Project not found: {root}"
        )

    ok(
        "Project directory exists"
    )

    go_mod = root / "go.mod"

    info(
        "Checking go.mod"
    )

    if not go_mod.exists():
        fail(
            "go.mod not found"
        )

    ok(
        "Go project confirmed"
    )

    return root


def discover_role_error(
    root,
    config,
):
    print()
    print(
        "=== DOMAIN ERROR DISCOVERY ==="
    )
    print()

    relative_dir = config["domain_dir"]

    domain_dir = root / relative_dir

    info(
        f"Domain directory: "
        f"{relative_dir}"
    )

    if not domain_dir.exists():
        fail(
            f"Domain directory not found: "
            f"{domain_dir}"
        )

    ok(
        "Domain directory exists"
    )

    go_files = sorted(
        path
        for path in domain_dir.glob("*.go")
        if not path.name.endswith(
            "_test.go"
        )
    )

    info(
        f"Found {len(go_files)} "
        f"production Go files"
    )

    if not go_files:
        fail(
            "No production Go files found"
        )

    expected_message = config.get(
        "expected_error_message"
    )

    info(
        "Expected error message:"
    )

    print(
        f'       "{expected_message}"'
    )

    candidates = []

    #
    # Matches things like:
    #
    # ErrRoleNotAssigned =
    #     errors.New(
    #         "role is not assigned to user",
    #     )
    #
    # and:
    #
    # ErrRoleNotAssigned = errors.New(
    #     "role is not assigned to user",
    # )
    #
    pattern = re.compile(
        r'(?P<name>Err[A-Za-z0-9_]+)'
        r'\s*=\s*'
        r'errors\.New\s*\(\s*'
        r'"(?P<message>[^"]+)"'
        r'\s*\)',
        re.MULTILINE,
    )

    for path in go_files:
        relative = path.relative_to(root)

        info(
            f"Scanning: {relative}"
        )

        content = path.read_text(
            encoding="utf-8",
        )

        ok(
            f"Read {relative}"
        )

        matches = list(
            pattern.finditer(content)
        )

        info(
            f"Found {len(matches)} "
            f"errors.New declarations"
        )

        for match in matches:
            name = match.group("name")
            message = match.group(
                "message"
            )

            info(
                f"Discovered: "
                f"{name} = "
                f'"{message}"'
            )

            if (
                message
                == expected_message
            ):
                candidates.append(
                    {
                        "name": name,
                        "message": message,
                        "path": path,
                    }
                )

                ok(
                    "Error message matches "
                    "RemoveRole contract"
                )

    if len(candidates) == 0:
        print()
        warn(
            "Could not find an error "
            "constant matching:"
        )

        print(
            f'    "{expected_message}"'
        )

        print()
        print(
            "Relevant domain lines:"
        )

        for path in go_files:
            content = path.read_text(
                encoding="utf-8",
            )

            lines = (
                content.splitlines()
            )

            for number, line in enumerate(
                lines,
                start=1,
            ):
                lower = line.lower()

                if (
                    "role" in lower
                    and (
                        "error" in lower
                        or
                        "assigned" in lower
                    )
                ):
                    print(
                        f"{path.name}:"
                        f"{number}: "
                        f"{line}"
                    )

        fail(
            "Domain error could not "
            "be verified. No files changed."
        )

    if len(candidates) > 1:
        print()

        warn(
            "Multiple domain errors have "
            "the expected message"
        )

        for candidate in candidates:
            print(
                f"  {candidate['name']} "
                f"in "
                f"{candidate['path'].name}"
            )

        fail(
            "Ambiguous domain contract. "
            "No files changed."
        )

    candidate = candidates[0]

    print()
    ok(
        "DOMAIN ERROR VERIFIED"
    )

    print()
    print(
        f"  Constant : "
        f"{candidate['name']}"
    )

    print(
        f"  Message  : "
        f"{candidate['message']}"
    )

    print(
        f"  File     : "
        f"{candidate['path'].relative_to(root)}"
    )

    return candidate["name"]


def verify_remove_role_contract(
    root,
    config,
    error_name,
):
    print()
    print(
        "=== REMOVE ROLE CONTRACT ==="
    )
    print()

    domain_dir = (
        root / config["domain_dir"]
    )

    method_found = False
    error_used = False

    for path in domain_dir.glob(
        "*.go"
    ):
        if path.name.endswith(
            "_test.go"
        ):
            continue

        content = path.read_text(
            encoding="utf-8",
        )

        if "func (u *User) RemoveRole(" not in content:
            continue

        method_found = True

        relative = path.relative_to(root)

        ok(
            f"Found RemoveRole in "
            f"{relative}"
        )

        info(
            "Checking return type"
        )

        signature_pattern = re.compile(
            r"func\s*"
            r"\(\s*u\s+\*User\s*\)\s*"
            r"RemoveRole\s*"
            r"\([^)]*\)\s*"
            r"error\s*"
            r"\{",
            re.MULTILINE,
        )

        if not signature_pattern.search(
            content
        ):
            print_source(
                path,
                content,
            )

            fail(
                "RemoveRole does not have "
                "the expected error return type"
            )

        ok(
            "RemoveRole returns error"
        )

        info(
            f"Checking whether "
            f"{error_name} is used by "
            f"RemoveRole source"
        )

        if error_name in content:
            error_used = True

            ok(
                f"{error_name} appears in "
                f"RemoveRole source file"
            )

    if not method_found:
        fail(
            "User.RemoveRole method "
            "was not found"
        )

    if not error_used:
        fail(
            f"{error_name} was discovered, "
            f"but it does not appear in the "
            f"RemoveRole source file"
        )

    ok(
        "RemoveRole domain contract "
        "verified"
    )


def patch_test(
    root,
    config,
    error_name,
):
    print()
    print(
        "=== TEST PATCH ==="
    )
    print()

    relative_path = (
        config["test_path"]
    )

    path = root / relative_path

    info(
        f"Checking test file: "
        f"{relative_path}"
    )

    if not path.exists():
        fail(
            f"Test file not found: "
            f"{relative_path}"
        )

    ok(
        "Test file exists"
    )

    content = path.read_text(
        encoding="utf-8",
    )

    ok(
        f"Test file read "
        f"({len(content)} chars)"
    )

    test_name = (
        "TestRemoveRoleNotAssigned"
        "DoesNotMutateRoles"
    )

    info(
        f"Checking for test: "
        f"{test_name}"
    )

    if test_name not in content:
        print_source(
            path,
            content,
        )

        fail(
            f"Expected test not found: "
            f"{test_name}"
        )

    ok(
        "Target test found"
    )

    info(
        "Checking imports"
    )

    updated = content

    #
    # Add errors import.
    #
    if '"errors"' in updated:
        ok(
            "errors package already imported"
        )

    else:
        info(
            "errors package not imported"
        )

        import_marker = (
            'import (\n'
        )

        if import_marker not in updated:
            fail(
                "Could not locate Go "
                "import block"
            )

        updated = updated.replace(
            import_marker,
            'import (\n\t"errors"\n',
            1,
        )

        ok(
            "Added errors import"
        )

    #
    # Find target function using brace
    # matching instead of brittle exact
    # multiline replacement.
    #
    info(
        "Locating target test function"
    )

    marker = (
        f"func {test_name}("
    )

    start = updated.find(marker)

    if start == -1:
        fail(
            "Could not locate target "
            "test function declaration"
        )

    ok(
        "Target function declaration found"
    )

    brace_start = updated.find(
        "{",
        start,
    )

    if brace_start == -1:
        fail(
            "Could not locate target "
            "test opening brace"
        )

    depth = 0
    end = None

    for index in range(
        brace_start,
        len(updated),
    ):
        character = updated[index]

        if character == "{":
            depth += 1

        elif character == "}":
            depth -= 1

            if depth == 0:
                end = index + 1
                break

    if end is None:
        fail(
            "Could not determine end "
            "of target test function"
        )

    ok(
        "Target function boundaries found"
    )

    old_function = updated[
        start:end
    ]

    print()
    print(
        "Current target test:"
    )
    print(
        "------------------------------"
    )
    print(old_function)
    print(
        "------------------------------"
    )
    print()

    new_function = f'''func {test_name}(t *testing.T) {{
\tu := newUserForRemoveRoleTest(t)

\tbefore := append(
\t\t[]role.ID(nil),
\t\tu.RoleIDs...,
\t)

\terr := u.RemoveRole(
\t\trole.ID("role_admin"),
\t)

\tif !errors.Is(
\t\terr,
\t\tuser.{error_name},
\t) {{
\t\tt.Fatalf(
\t\t\t"expected {error_name}, got %v",
\t\t\terr,
\t\t)
\t}}

\tif len(u.RoleIDs) != len(before) {{
\t\tt.Fatalf(
\t\t\t"expected role count to remain %d, got %d",
\t\t\tlen(before),
\t\t\tlen(u.RoleIDs),
\t\t)
\t}}

\tfor index := range before {{
\t\tif u.RoleIDs[index] != before[index] {{
\t\t\tt.Fatalf(
\t\t\t\t"roles mutated at index %d: expected %s, got %s",
\t\t\t\tindex,
\t\t\t\tbefore[index],
\t\t\t\tu.RoleIDs[index],
\t\t\t)
\t\t}}
\t}}
}}'''

    info(
        f"Generating test using "
        f"user.{error_name}"
    )

    updated = (
        updated[:start]
        + new_function
        + updated[end:]
    )

    ok(
        "Target test replaced in memory"
    )

    #
    # Validate before write.
    #
    print()
    print(
        "=== PRE-WRITE VALIDATION ==="
    )
    print()

    checks = [
        (
            "errors import",
            '"errors"',
        ),
        (
            "errors.Is assertion",
            "errors.Is(",
        ),
        (
            "domain error constant",
            f"user.{error_name}",
        ),
        (
            "non-mutation snapshot",
            "before := append(",
        ),
        (
            "role count check",
            "len(u.RoleIDs) != len(before)",
        ),
    ]

    for name, value in checks:
        info(
            f"Checking: {name}"
        )

        if value not in updated:
            fail(
                f"Validation failed: "
                f"{name}"
            )

        ok(
            f"Validated: {name}"
        )

    if updated == content:
        ok(
            "Test already matches "
            "required contract"
        )

        print(
            "No write required."
        )

        return

    #
    # Backup.
    #
    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup = (
        root
        / ".generator-backups"
        / timestamp
        / "fix-remove-role-domain-test"
        / relative_path
    )

    info(
        f"Creating backup: {backup}"
    )

    backup.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        path,
        backup,
    )

    ok(
        "Backup created"
    )

    #
    # Write.
    #
    info(
        "Writing updated test"
    )

    path.write_text(
        updated,
        encoding="utf-8",
    )

    ok(
        "Updated test written"
    )

    #
    # Read-back verification.
    #
    info(
        "Reading updated test back"
    )

    written = path.read_text(
        encoding="utf-8",
    )

    if written != updated:
        fail(
            "Written file does not match "
            "generated content"
        )

    ok(
        "Write verification passed"
    )

    print()
    print(
        f"Backup: {backup}"
    )


def main():
    config = load_config()

    root = resolve_project(
        config
    )

    #
    # Python now performs the verification
    # that we previously did manually with
    # Get-ChildItem / Select-String.
    #
    error_name = discover_role_error(
        root,
        config,
    )

    verify_remove_role_contract(
        root,
        config,
        error_name,
    )

    patch_test(
        root,
        config,
        error_name,
    )

    print()
    print(
        "================================="
    )
    print(
        "FIX COMPLETE"
    )
    print(
        "================================="
    )
    print()

    print(
        "Verified domain behavior:"
    )

    print(
        f"  unassigned role -> "
        f"user.{error_name}"
    )

    print()
    print(
        "Updated test behavior:"
    )

    print(
        "  expects domain error"
    )

    print(
        "  verifies RoleIDs are unchanged"
    )

    print()
    print(
        "Next:"
    )

    print()
    print(
        "  cd prayer-api"
    )

    print(
        "  go fmt ./..."
    )

    print(
        "  go test -v "
        "./internal/domain/user"
    )

    print(
        "  go test -v "
        "./internal/application/userrole"
    )

    print(
        "  go test -v ./..."
    )


if __name__ == "__main__":
    main()