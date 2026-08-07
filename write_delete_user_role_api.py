import json
import shutil
import sys
from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "delete_user_role_api.json"


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
    info(f"Looking for config: {CONFIG_FILE}")

    if not CONFIG_FILE.exists():
        fail("Configuration file not found")

    ok("Configuration file exists")

    info("Loading JSON")

    try:
        with CONFIG_FILE.open(
            "r",
            encoding="utf-8",
        ) as file:
            config = json.load(file)
    except Exception as error:
        fail(f"Invalid JSON: {error}")

    ok("JSON loaded")

    return config


def resolve_project(config):
    info("Resolving project directory")

    project = config.get("project")

    if not project:
        fail("Missing 'project' in JSON")

    root = SCRIPT_DIR / project

    info(f"Project path: {root}")

    if not root.exists():
        fail(f"Project does not exist: {root}")

    ok("Project exists")

    return root


def read_file(root, relative_path):
    info(f"Reading: {relative_path}")

    path = root / relative_path

    if not path.exists():
        fail(f"File does not exist: {relative_path}")

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
        f"({len(content)} chars)"
    )

    return path, content


def make_backup(
    root,
    path,
    relative_path,
    backup_root,
):
    info(f"Backing up: {relative_path}")

    destination = (
        backup_root / relative_path
    )

    destination.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(
        path,
        destination,
    )

    ok(f"Backup created: {destination}")


def write_changed_file(
    root,
    path,
    relative_path,
    original,
    updated,
    backup_root,
):
    if updated == original:
        ok(f"No changes: {relative_path}")
        return

    make_backup(
        root,
        path,
        relative_path,
        backup_root,
    )

    info(f"Writing: {relative_path}")

    path.write_text(
        updated,
        encoding="utf-8",
    )

    ok(f"Updated: {relative_path}")


def patch_user_domain(
    root,
    config,
    backup_root,
):
    print()
    print("=== USER DOMAIN ===")

    relative_path = config["domain"]["path"]

    path, content = read_file(
        root,
        relative_path,
    )

    info("Checking for RemoveRole")

    if "func (u *User) RemoveRole(" in content:
        ok("RemoveRole already exists")
        return

    info("Checking for AssignRole")

    marker = "func (u *User) AssignRole("

    start = content.find(marker)

    if start == -1:
        fail(
            "Could not locate User.AssignRole. "
            "Refusing to guess insertion point.",
            content,
        )

    ok("Found User.AssignRole")

    info(
        "Finding end of AssignRole method"
    )

    brace_start = content.find(
        "{",
        start,
    )

    if brace_start == -1:
        fail(
            "AssignRole opening brace not found",
            content,
        )

    depth = 0
    end = None

    for index in range(
        brace_start,
        len(content),
    ):
        character = content[index]

        if character == "{":
            depth += 1

        elif character == "}":
            depth -= 1

            if depth == 0:
                end = index + 1
                break

    if end is None:
        fail(
            "Could not determine end of "
            "AssignRole",
            content,
        )

    ok("Found end of AssignRole")

    method = '''

func (u *User) RemoveRole(roleID role.ID) bool {
	for index, current := range u.RoleIDs {
		if current != roleID {
			continue
		}

		u.RoleIDs = append(
			u.RoleIDs[:index],
			u.RoleIDs[index+1:]...,
		)

		return true
	}

	return false
}
'''

    updated = (
        content[:end]
        + method
        + content[end:]
    )

    info(
        "Validating generated RemoveRole"
    )

    if (
        "func (u *User) RemoveRole("
        not in updated
    ):
        fail(
            "RemoveRole validation failed",
            updated,
        )

    ok("RemoveRole generated")

    write_changed_file(
        root,
        path,
        relative_path,
        content,
        updated,
        backup_root,
    )


def create_application_file(
    root,
    config,
):
    print()
    print("=== APPLICATION SERVICE ===")

    relative_path = (
        config["application"]["path"]
    )

    path = root / relative_path

    info(
        f"Checking application file: "
        f"{relative_path}"
    )

    if path.exists():
        content = path.read_text(
            encoding="utf-8",
        )

        if (
            "func (s *Service) Remove("
            in content
        ):
            ok(
                "Remove application service "
                "already exists"
            )
            return

        fail(
            "remove.go exists but does not "
            "contain the expected Remove service. "
            "Refusing to overwrite it.",
            content,
        )

    ok(
        "remove.go does not exist; "
        "safe to create"
    )

    path.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    content = '''package userrole

import (
	"context"
	"errors"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

var ErrProtectedRole = errors.New("protected role")

type RemoveCommand struct {
	ActorExternalID string
	TargetUserID    user.ID
	RoleID          role.ID
}

type RemoveResult struct {
	User    *user.User
	Role    *role.Role
	Removed bool
}

func (s *Service) Remove(
	ctx context.Context,
	cmd RemoveCommand,
) (*RemoveResult, error) {
	actor, err := s.users.FindByExternalID(
		ctx,
		cmd.ActorExternalID,
	)

	if err != nil {
		return nil, err
	}

	if actor == nil {
		return nil, ErrForbidden
	}

	allowed, err := s.hasPermission(
		ctx,
		actor,
		role.PermissionManageUsers,
	)

	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrForbidden
	}

	target, err := s.users.FindByID(
		ctx,
		cmd.TargetUserID,
	)

	if err != nil {
		return nil, err
	}

	if target == nil {
		return nil, ErrUserNotFound
	}

	resolvedRole, err := s.roles.FindByID(
		ctx,
		cmd.RoleID,
	)

	if err != nil {
		return nil, err
	}

	if resolvedRole == nil {
		return nil, ErrRoleNotFound
	}

	if resolvedRole.ID == role.ID("role_members") {
		return nil, ErrProtectedRole
	}

	removed := target.RemoveRole(
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
	}, nil
}
'''

    info(
        "Writing application remove service"
    )

    path.write_text(
        content,
        encoding="utf-8",
    )

    ok(f"Created: {relative_path}")


def patch_handler(
    root,
    config,
    backup_root,
):
    print()
    print("=== HTTP HANDLER ===")

    relative_path = (
        config["http"]["handler_path"]
    )

    path, content = read_file(
        root,
        relative_path,
    )

    info("Checking for Remove handler")

    if (
        "func (h *UserRoleHandler) Remove("
        in content
    ):
        ok("Remove handler already exists")
        return

    info(
        "Checking handler imports errors"
    )

    if '"errors"' not in content:
        fail(
            "Expected existing user role "
            "handler to import errors. "
            "Current POST handler should "
            "already use errors.Is.",
            content,
        )

    ok("errors import exists")

    handler = '''

func (h *UserRoleHandler) Remove(
	w http.ResponseWriter,
	r *http.Request,
) {
	claims, ok := clerk.SessionClaimsFromContext(
		r.Context(),
	)

	if !ok {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]any{
				"error": "unauthorized",
			},
		)
		return
	}

	targetUserID := user.ID(
		r.PathValue("userID"),
	)

	roleID := role.ID(
		r.PathValue("roleID"),
	)

	result, err := h.service.Remove(
		r.Context(),
		userrole.RemoveCommand{
			ActorExternalID: claims.Subject,
			TargetUserID:    targetUserID,
			RoleID:          roleID,
		},
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			userrole.ErrForbidden,
		):
			writeJSON(
				w,
				http.StatusForbidden,
				map[string]any{
					"error": "forbidden",
				},
			)

		case errors.Is(
			err,
			userrole.ErrUserNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "user not found",
				},
			)

		case errors.Is(
			err,
			userrole.ErrRoleNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "role not found",
				},
			)

		case errors.Is(
			err,
			userrole.ErrProtectedRole,
		):
			writeJSON(
				w,
				http.StatusConflict,
				map[string]any{
					"error": "protected role",
				},
			)

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": "internal server error",
				},
			)
		}

		return
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]any{
			"user": map[string]any{
				"id":     result.User.ID,
				"name":   result.User.Name,
				"status": result.User.Status,
				"roles":  result.User.RoleIDs,
			},
			"role": map[string]any{
				"id":          result.Role.ID,
				"name":        result.Role.Name,
				"permissions": result.Role.Permissions,
			},
			"removed": result.Removed,
		},
	)
}
'''

    updated = content + handler

    info("Validating Remove handler")

    if (
        "userrole.ErrProtectedRole"
        not in updated
    ):
        fail(
            "Generated handler does not "
            "handle protected role",
            updated,
        )

    ok("Remove handler validated")

    write_changed_file(
        root,
        path,
        relative_path,
        content,
        updated,
        backup_root,
    )


def patch_router(
    root,
    config,
    backup_root,
):
    print()
    print("=== ROUTER ===")

    relative_path = (
        config["http"]["router_path"]
    )

    path, content = read_file(
        root,
        relative_path,
    )

    route = (
        'DELETE /api/users/'
        '{userID}/roles/{roleID}'
    )

    info(
        f"Checking route: {route}"
    )

    if route in content:
        ok("DELETE route already exists")
        return

    info(
        "Looking for POST user-role route"
    )

    post_route = (
        '"POST /api/users/'
        '{userID}/roles/{roleID}"'
    )

    start = content.find(post_route)

    if start == -1:
        fail(
            "POST user-role route not found. "
            "Cannot safely determine where "
            "DELETE route belongs.",
            content,
        )

    ok("Found POST user-role route")

    info(
        "Finding end of POST mux.Handle block"
    )

    handle_start = content.rfind(
        "mux.Handle(",
        0,
        start,
    )

    if handle_start == -1:
        fail(
            "Could not locate mux.Handle "
            "for POST user-role route",
            content,
        )

    depth = 0
    paren_start = content.find(
        "(",
        handle_start,
    )

    end = None

    for index in range(
        paren_start,
        len(content),
    ):
        char = content[index]

        if char == "(":
            depth += 1

        elif char == ")":
            depth -= 1

            if depth == 0:
                end = index + 1
                break

    if end is None:
        fail(
            "Could not determine end of "
            "POST route block",
            content,
        )

    ok("Found end of POST route block")

    delete_block = '''

	mux.Handle(
		"DELETE /api/users/{userID}/roles/{roleID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(userRoles.Remove),
		),
	)
'''

    updated = (
        content[:end]
        + delete_block
        + content[end:]
    )

    info("Validating DELETE route")

    if route not in updated:
        fail(
            "DELETE route insertion failed",
            updated,
        )

    ok("DELETE route generated")

    write_changed_file(
        root,
        path,
        relative_path,
        content,
        updated,
        backup_root,
    )


def validate_main(
    root,
    config,
):
    print()
    print("=== MAIN WIRING ===")

    relative_path = config["main"]["path"]

    _, content = read_file(
        root,
        relative_path,
    )

    checks = [
        "userrole.NewService(",
        "httpapi.NewUserRoleHandler(",
        "userRoleHandler",
    ]

    for marker in checks:
        info(
            f"Checking main.go for: {marker}"
        )

        if marker not in content:
            fail(
                f"Missing existing user-role "
                f"wiring: {marker}",
                content,
            )

        ok(f"Found: {marker}")

    ok(
        "No main.go changes required. "
        "DELETE uses existing UserRoleHandler."
    )


def main():
    print()
    print(
        "=== DELETE USER ROLE GENERATOR ==="
    )
    print()

    config = load_config()
    root = resolve_project(config)

    timestamp = datetime.now().strftime(
        "%Y%m%d_%H%M%S"
    )

    backup_root = (
        root
        / ".generator-backups"
        / timestamp
        / "delete-user-role-api"
    )

    info(
        f"Backup root: {backup_root}"
    )

    patch_user_domain(
        root,
        config,
        backup_root,
    )

    create_application_file(
        root,
        config,
    )

    patch_handler(
        root,
        config,
        backup_root,
    )

    patch_router(
        root,
        config,
        backup_root,
    )

    validate_main(
        root,
        config,
    )

    print()
    print("==============================")
    print("GENERATION COMPLETE")
    print("==============================")
    print()
    print(
        "Endpoint:"
    )
    print(
        "DELETE /api/users/"
        "{userID}/roles/{roleID}"
    )
    print()
    print(
        "Next commands:"
    )
    print(
        "  cd prayer-api"
    )
    print(
        "  go fmt ./..."
    )
    print(
        "  go test -v ./..."
    )


if __name__ == "__main__":
    main()