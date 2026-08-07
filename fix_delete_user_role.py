import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "fix_delete_user_role.json"


def info(message):
    print(f"[INFO] {message}")


def ok(message):
    print(f"[ OK ] {message}")


def warn(message):
    print(f"[WARN] {message}")


def fail(message, content=None):
    print()
    print(f"[FAIL] {message}")

    if content is not None:
        print()
        print("========== SOURCE ==========")

        for number, line in enumerate(
            content.splitlines(),
            start=1,
        ):
            print(f"{number:4}: {line}")

        print("============================")

    sys.exit(1)


def load_config():
    print()
    print("===================================")
    print("DELETE USER ROLE CORRECTIVE PATCH")
    print("===================================")
    print()

    info(f"Looking for config: {CONFIG_FILE}")

    if not CONFIG_FILE.exists():
        fail(
            f"Configuration file not found: "
            f"{CONFIG_FILE}"
        )

    ok("Configuration file exists")

    info("Reading JSON configuration")

    try:
        with CONFIG_FILE.open(
            "r",
            encoding="utf-8",
        ) as file:
            config = json.load(file)

    except json.JSONDecodeError as error:
        fail(
            f"JSON error line {error.lineno}, "
            f"column {error.colno}: "
            f"{error.msg}"
        )

    except Exception as error:
        fail(
            f"Could not read JSON: {error}"
        )

    ok("JSON configuration loaded")

    return config


def resolve_project(config):
    info("Reading project configuration")

    project = config.get("project")

    if not project:
        fail("Missing JSON field: project")

    ok(f"Project configured: {project}")

    root = SCRIPT_DIR / project

    info(f"Checking project directory: {root}")

    if not root.exists():
        fail(
            f"Project directory not found: "
            f"{root}"
        )

    ok("Project directory exists")

    go_mod = root / "go.mod"

    info(f"Checking go.mod: {go_mod}")

    if not go_mod.exists():
        fail(
            "go.mod not found. "
            "Wrong project directory?"
        )

    ok("go.mod found")

    return root


def read_required_file(root, relative_path):
    info(f"Checking file: {relative_path}")

    path = root / relative_path

    if not path.exists():
        fail(
            f"Required file not found: "
            f"{relative_path}"
        )

    ok(f"File exists: {relative_path}")

    info(f"Reading: {relative_path}")

    try:
        content = path.read_text(
            encoding="utf-8",
        )

    except Exception as error:
        fail(
            f"Could not read {relative_path}: "
            f"{error}"
        )

    ok(
        f"Read {relative_path} "
        f"({len(content)} characters)"
    )

    return path, content


def backup_file(
    root,
    path,
    relative_path,
    backup_root,
):
    destination = backup_root / relative_path

    info(
        f"Creating backup: "
        f"{relative_path}"
    )

    destination.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        path,
        destination,
    )

    ok(
        f"Backup: {destination}"
    )


def write_file(
    root,
    path,
    relative_path,
    original,
    updated,
    backup_root,
):
    info(
        f"Checking whether "
        f"{relative_path} changed"
    )

    if updated == original:
        ok(
            f"No changes required: "
            f"{relative_path}"
        )
        return

    ok(
        f"Changes detected: {relative_path}"
    )

    backup_file(
        root,
        path,
        relative_path,
        backup_root,
    )

    info(
        f"Writing: {relative_path}"
    )

    try:
        path.write_text(
            updated,
            encoding="utf-8",
        )

    except Exception as error:
        fail(
            f"Could not write "
            f"{relative_path}: {error}"
        )

    ok(
        f"Updated: {relative_path}"
    )

    info(
        f"Reading {relative_path} "
        f"back for verification"
    )

    verification = path.read_text(
        encoding="utf-8",
    )

    if verification != updated:
        fail(
            f"Write verification failed: "
            f"{relative_path}"
        )

    ok(
        f"Write verified: {relative_path}"
    )


def find_remove_role_method(root, config):
    print()
    print("=== DOMAIN CONTRACT CHECK ===")
    print()

    configured_path = (
        config["domain"]["path"]
    )

    role_dir = (
        root
        / "internal"
        / "domain"
        / "user"
    )

    info(
        "Searching user domain files "
        "for RemoveRole"
    )

    matches = []

    for path in role_dir.glob("*.go"):
        if path.name.endswith("_test.go"):
            continue

        info(
            f"Scanning: "
            f"{path.relative_to(root)}"
        )

        content = path.read_text(
            encoding="utf-8",
        )

        if "RemoveRole(" in content:
            matches.append(
                (path, content)
            )

            ok(
                f"Found RemoveRole in "
                f"{path.relative_to(root)}"
            )

    if not matches:
        fail(
            "No RemoveRole method found "
            "in user domain"
        )

    if len(matches) > 1:
        warn(
            f"RemoveRole text found in "
            f"{len(matches)} files"
        )

    method_pattern = re.compile(
        r"func\s*"
        r"\(\s*u\s+\*User\s*\)\s*"
        r"RemoveRole\s*"
        r"\(\s*"
        r"roleID\s+role\.ID\s*"
        r"\)\s*"
        r"error\s*"
        r"\{",
        re.MULTILINE,
    )

    for path, content in matches:
        info(
            f"Checking RemoveRole signature "
            f"in {path.relative_to(root)}"
        )

        match = method_pattern.search(
            content
        )

        if match:
            ok(
                "Confirmed domain contract:"
            )

            print()
            print(
                "    func (u *User) "
                "RemoveRole(roleID role.ID) error"
            )
            print()

            return path, content

    print()
    warn(
        "RemoveRole exists but expected "
        "error-returning signature was "
        "not detected."
    )

    for path, content in matches:
        print()
        print(
            f"--- {path.relative_to(root)} ---"
        )

        for number, line in enumerate(
            content.splitlines(),
            start=1,
        ):
            if "RemoveRole" in line:
                start = max(
                    0,
                    number - 3,
                )

                end = min(
                    len(content.splitlines()),
                    number + 12,
                )

                lines = content.splitlines()

                for index in range(
                    start,
                    end,
                ):
                    print(
                        f"{index + 1:4}: "
                        f"{lines[index]}"
                    )

    fail(
        "Expected RemoveRole to return error. "
        "Refusing to patch against an "
        "unknown domain contract."
    )


def patch_application(
    root,
    config,
    backup_root,
):
    print()
    print("=== APPLICATION PATCH ===")
    print()

    relative_path = (
        config["application"]["path"]
    )

    path, content = read_required_file(
        root,
        relative_path,
    )

    info(
        "Checking for Service.Remove"
    )

    if "func (s *Service) Remove(" not in content:
        fail(
            "Service.Remove not found",
            content,
        )

    ok("Service.Remove found")

    info(
        "Checking for incorrect "
        "boolean RemoveRole usage"
    )

    incorrect = '''removed := target.RemoveRole(
		resolvedRole.ID,
	)

	if !removed {
		return &RemoveResult{
			User:    target,
			Role:    resolvedRole,
			Removed: false,
		}, nil
	}

	if err := s.users.Save(
		ctx,
		target,
	); err != nil {
		return nil, err
	}

	return &RemoveResult{
		User:    target,
		Role:    resolvedRole,
		Removed: true,
	}, nil'''

    if incorrect in content:
        ok(
            "Found incorrect boolean "
            "RemoveRole implementation"
        )
    else:
        info(
            "Exact incorrect block was "
            "not found"
        )

        if (
            "removed := target.RemoveRole("
            not in content
        ):
            if (
                "if err := target.RemoveRole("
                in content
            ):
                ok(
                    "Application may already "
                    "use error-returning "
                    "RemoveRole"
                )

                return

            fail(
                "Could not identify generated "
                "RemoveRole call. Refusing "
                "unsafe replacement.",
                content,
            )

        fail(
            "Found boolean RemoveRole usage, "
            "but surrounding generated block "
            "differs from expected source.",
            content,
        )

    replacement = '''wasAssigned := false

	for _, current := range target.RoleIDs {
		if current == resolvedRole.ID {
			wasAssigned = true
			break
		}
	}

	if !wasAssigned {
		return &RemoveResult{
			User:    target,
			Role:    resolvedRole,
			Removed: false,
		}, nil
	}

	if err := target.RemoveRole(
		resolvedRole.ID,
	); err != nil {
		return nil, err
	}

	if err := s.users.Save(
		ctx,
		target,
	); err != nil {
		return nil, err
	}

	return &RemoveResult{
		User:    target,
		Role:    resolvedRole,
		Removed: true,
	}, nil'''

    info(
        "Replacing boolean handling with "
        "domain error handling"
    )

    updated = content.replace(
        incorrect,
        replacement,
        1,
    )

    if updated == content:
        fail(
            "Application replacement made "
            "no changes",
            content,
        )

    ok(
        "Application block replaced"
    )

    info(
        "Validating old boolean usage "
        "was removed"
    )

    if (
        "removed := target.RemoveRole("
        in updated
    ):
        fail(
            "Old boolean RemoveRole usage "
            "still exists",
            updated,
        )

    ok(
        "Old boolean usage removed"
    )

    info(
        "Validating error-returning "
        "RemoveRole call"
    )

    if (
        "if err := target.RemoveRole("
        not in updated
    ):
        fail(
            "New error-returning RemoveRole "
            "call missing",
            updated,
        )

    ok(
        "Error-returning RemoveRole "
        "call exists"
    )

    info(
        "Validating idempotency check"
    )

    if "wasAssigned := false" not in updated:
        fail(
            "wasAssigned idempotency "
            "check missing",
            updated,
        )

    ok(
        "Idempotency check exists"
    )

    write_file(
        root,
        path,
        relative_path,
        content,
        updated,
        backup_root,
    )


def patch_domain_test(
    root,
    config,
    backup_root,
):
    print()
    print("=== DOMAIN TEST PATCH ===")
    print()

    relative_path = (
        config["domain_test"]["path"]
    )

    path, content = read_required_file(
        root,
        relative_path,
    )

    info(
        "Checking generated domain tests"
    )

    required_tests = [
        "TestRemoveRoleRemovesAssignedRole",
        "TestRemoveRoleReturnsFalseWhenRoleIsNotAssigned",
        "TestRemoveRoleDoesNotRemoveOtherRoles",
    ]

    for test in required_tests:
        info(
            f"Looking for: {test}"
        )

        if test not in content:
            fail(
                f"Expected generated test "
                f"not found: {test}",
                content,
            )

        ok(
            f"Found: {test}"
        )

    info(
        "Replacing generated domain tests "
        "with error-contract tests"
    )

    updated = '''package user_test

import (
	"testing"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

func TestRemoveRoleRemovesAssignedRole(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	adminRoleID := role.ID("role_admin")

	u.AssignRole(adminRoleID)

	if !hasRoleForRemoveTest(u, adminRoleID) {
		t.Fatal("expected role_admin to be assigned before removal")
	}

	err := u.RemoveRole(adminRoleID)

	if err != nil {
		t.Fatalf("expected role removal to succeed: %v", err)
	}

	if hasRoleForRemoveTest(u, adminRoleID) {
		t.Fatal("expected role_admin to be removed")
	}
}

func TestRemoveRoleNotAssignedDoesNotMutateRoles(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	before := append(
		[]role.ID(nil),
		u.RoleIDs...,
	)

	err := u.RemoveRole(
		role.ID("role_admin"),
	)

	if err != nil {
		t.Fatalf(
			"expected removing an unassigned role to be harmless: %v",
			err,
		)
	}

	if len(u.RoleIDs) != len(before) {
		t.Fatalf(
			"expected role count to remain %d, got %d",
			len(before),
			len(u.RoleIDs),
		)
	}

	for index := range before {
		if u.RoleIDs[index] != before[index] {
			t.Fatalf(
				"roles mutated at index %d: expected %s, got %s",
				index,
				before[index],
				u.RoleIDs[index],
			)
		}
	}
}

func TestRemoveRoleDoesNotRemoveOtherRoles(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	adminRoleID := role.ID("role_admin")
	memberRoleID := role.ID("role_members")

	u.AssignRole(adminRoleID)

	err := u.RemoveRole(adminRoleID)

	if err != nil {
		t.Fatalf(
			"expected role_admin removal to succeed: %v",
			err,
		)
	}

	if hasRoleForRemoveTest(
		u,
		adminRoleID,
	) {
		t.Fatal(
			"expected role_admin to be removed",
		)
	}

	if !hasRoleForRemoveTest(
		u,
		memberRoleID,
	) {
		t.Fatal(
			"removing role_admin must not remove role_members",
		)
	}
}

func newUserForRemoveRoleTest(
	t *testing.T,
) *user.User {
	t.Helper()

	u, err := user.New(
		user.ID("user_remove_test"),
		"clerk_remove_test",
		"Remove Test",
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf(
			"failed to create user: %v",
			err,
		)
	}

	return u
}

func hasRoleForRemoveTest(
	u *user.User,
	roleID role.ID,
) bool {
	for _, current := range u.RoleIDs {
		if current == roleID {
			return true
		}
	}

	return false
}
'''

    info(
        "Validating replacement test file"
    )

    if (
        "removed := u.RemoveRole"
        in updated
    ):
        fail(
            "Boolean RemoveRole usage "
            "still exists in generated tests"
        )

    ok(
        "No boolean RemoveRole usage"
    )

    if (
        "err := u.RemoveRole"
        not in updated
    ):
        fail(
            "Error-returning RemoveRole "
            "tests missing"
        )

    ok(
        "Tests use error-returning "
        "RemoveRole"
    )

    if (
        "TestRemoveRoleNotAssignedDoesNotMutateRoles"
        not in updated
    ):
        fail(
            "Non-mutation test missing"
        )

    ok(
        "Non-mutation test exists"
    )

    write_file(
        root,
        path,
        relative_path,
        content,
        updated,
        backup_root,
    )


def inspect_application_tests(root):
    print()
    print("=== APPLICATION TEST CHECK ===")
    print()

    path = (
        root
        / "internal"
        / "application"
        / "userrole"
        / "remove_test.go"
    )

    info(
        "Checking application remove tests"
    )

    if not path.exists():
        warn(
            "remove_test.go does not exist"
        )
        return

    ok("remove_test.go exists")

    content = path.read_text(
        encoding="utf-8",
    )

    info(
        "Checking application tests for "
        "direct boolean RemoveRole usage"
    )

    if (
        "removed := "
        in content
        and ".RemoveRole(" in content
    ):
        warn(
            "Possible direct boolean "
            "RemoveRole usage found in "
            "application tests"
        )

        for number, line in enumerate(
            content.splitlines(),
            start=1,
        ):
            if "RemoveRole(" in line:
                print(
                    f"{number:4}: {line}"
                )
    else:
        ok(
            "No direct boolean domain "
            "RemoveRole usage detected"
        )


def main():
    config = load_config()

    root = resolve_project(config)

    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "fix-delete-user-role"
    )

    print()
    info(
        f"Backup root: {backup_root}"
    )

    find_remove_role_method(
        root,
        config,
    )

    patch_application(
        root,
        config,
        backup_root,
    )

    patch_domain_test(
        root,
        config,
        backup_root,
    )

    inspect_application_tests(root)

    print()
    print("===================================")
    print("PATCH COMPLETE")
    print("===================================")
    print()
    print(
        "Domain contract preserved:"
    )
    print()
    print(
        "  User.RemoveRole(role.ID) error"
    )
    print()
    print(
        "Application semantics:"
    )
    print()
    print(
        "  assigned     -> remove -> removed=true"
    )
    print(
        "  not assigned -> no-op  -> removed=false"
    )
    print(
        "  domain error -> propagate error"
    )
    print()
    print("Next:")
    print()
    print("  cd prayer-api")
    print("  go fmt ./...")
    print(
        "  go test -v ./internal/domain/user"
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