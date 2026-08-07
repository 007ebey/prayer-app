import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


SCRIPT_DIR = Path(__file__).resolve().parent
CONFIG_FILE = SCRIPT_DIR / "create_prayer_group_api.json"


# ============================================================
# Logging
# ============================================================

def info(message):
    print(f"[INFO] {message}")


def ok(message):
    print(f"[ OK ] {message}")


def warn(message):
    print(f"[WARN] {message}")


def fail(message):
    print()
    print(f"[FAIL] {message}")
    print()
    print("No further changes will be made.")
    sys.exit(1)


def heading(message):
    print()
    print("=" * 64)
    print(message)
    print("=" * 64)
    print()


def print_matches(title, matches):
    print()
    print(title)
    print("-" * 64)

    if not matches:
        print("  <none>")
    else:
        for item in matches:
            print(f"  {item}")

    print("-" * 64)
    print()


# ============================================================
# Configuration
# ============================================================

def load_config():
    heading("CREATE PRAYER GROUP API")

    info(f"Config: {CONFIG_FILE}")

    if not CONFIG_FILE.exists():
        fail(f"Config file not found: {CONFIG_FILE}")

    ok("Configuration file exists")

    try:
        config = json.loads(
            CONFIG_FILE.read_text(encoding="utf-8")
        )
    except Exception as error:
        fail(f"Could not read JSON: {error}")

    ok("Configuration loaded")

    return config


def resolve_project(config):
    project_name = config.get("project")

    info(f"Configured project: {project_name}")

    if not project_name:
        fail("Missing JSON property: project")

    root = SCRIPT_DIR / project_name

    info(f"Project path: {root}")

    if not root.exists():
        fail(f"Project directory does not exist: {root}")

    ok("Project directory exists")

    go_mod = root / "go.mod"

    info("Checking go.mod")

    if not go_mod.exists():
        fail("go.mod was not found")

    ok("Go project confirmed")

    return root


# ============================================================
# File helpers
# ============================================================

def read_file(root, relative_path, required=True):
    path = root / relative_path

    info(f"Checking: {relative_path}")

    if not path.exists():
        if required:
            fail(f"Required file missing: {relative_path}")

        warn(f"Optional file missing: {relative_path}")
        return path, None

    ok(f"Found: {relative_path}")

    try:
        content = path.read_text(encoding="utf-8")
    except Exception as error:
        fail(f"Could not read {relative_path}: {error}")

    ok(
        f"Read {relative_path} "
        f"({len(content)} characters)"
    )

    return path, content


def backup_file(root, path, relative_path, backup_root):
    destination = backup_root / relative_path

    info(f"Backup: {relative_path}")

    destination.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(path, destination)

    ok(f"Backup created: {destination}")


def write_file(
    root,
    relative_path,
    content,
    backup_root,
):
    path = root / relative_path

    info(f"Preparing: {relative_path}")

    path.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    if path.exists():
        current = path.read_text(encoding="utf-8")

        if current == content:
            ok(f"Already current: {relative_path}")
            return

        backup_file(
            root,
            path,
            relative_path,
            backup_root,
        )

    info(f"Writing: {relative_path}")

    path.write_text(
        content,
        encoding="utf-8",
    )

    ok(f"Written: {relative_path}")

    info(f"Verifying: {relative_path}")

    written = path.read_text(encoding="utf-8")

    if written != content:
        fail(f"Write verification failed: {relative_path}")

    ok(f"Verified: {relative_path}")


def replace_and_write(
    root,
    relative_path,
    original,
    updated,
    backup_root,
):
    if original == updated:
        ok(f"No modification required: {relative_path}")
        return

    path = root / relative_path

    backup_file(
        root,
        path,
        relative_path,
        backup_root,
    )

    info(f"Writing modified file: {relative_path}")

    path.write_text(
        updated,
        encoding="utf-8",
    )

    ok(f"Updated: {relative_path}")

    verification = path.read_text(encoding="utf-8")

    if verification != updated:
        fail(f"Verification failed: {relative_path}")

    ok(f"Verified: {relative_path}")


# ============================================================
# Domain discovery
# ============================================================

def discover_prayer_group_domain(root, config):
    heading("1. PRAYER GROUP DOMAIN DISCOVERY")

    domain_dir = root / config["expected"]["domain_dir"]

    info(f"Scanning: {domain_dir}")

    if not domain_dir.exists():
        fail("PrayerGroup domain directory does not exist")

    ok("PrayerGroup domain directory exists")

    files = sorted(domain_dir.glob("*.go"))

    production_files = [
        path
        for path in files
        if not path.name.endswith("_test.go")
    ]

    info(
        f"Production Go files found: "
        f"{len(production_files)}"
    )

    if not production_files:
        fail("No PrayerGroup production files found")

    combined = ""

    for path in production_files:
        relative = path.relative_to(root)

        info(f"Reading: {relative}")

        content = path.read_text(encoding="utf-8")

        combined += "\n" + content

        ok(f"Read: {relative}")

    # --------------------------------------------------------
    # Struct discovery
    # --------------------------------------------------------

    info("Looking for PrayerGroup struct")

    struct_match = re.search(
        r"type\s+PrayerGroup\s+struct\s*\{(?P<body>.*?)\}",
        combined,
        re.DOTALL,
    )

    if not struct_match:
        fail("type PrayerGroup struct was not found")

    ok("PrayerGroup struct found")

    struct_body = struct_match.group("body")

    fields = []

    for line in struct_body.splitlines():
        line = line.strip()

        if line:
            fields.append(line)

    print_matches(
        "PrayerGroup fields discovered:",
        fields,
    )

    # --------------------------------------------------------
    # Constructor discovery
    # --------------------------------------------------------

    info("Looking for PrayerGroup constructor")

    constructor_match = re.search(
        r"func\s+New\s*\((?P<args>.*?)\)"
        r"\s*\(\s*\*PrayerGroup\s*,\s*error\s*\)",
        combined,
        re.DOTALL,
    )

    if not constructor_match:
        fail(
            "Could not verify constructor "
            "func New(...) (*PrayerGroup, error)"
        )

    constructor_args = constructor_match.group("args")

    ok("PrayerGroup constructor found")

    print_matches(
        "Constructor arguments:",
        [
            line.strip()
            for line in constructor_args.splitlines()
            if line.strip()
        ],
    )

    # --------------------------------------------------------
    # ID
    # --------------------------------------------------------

    info("Checking PrayerGroup ID type")

    if not re.search(
        r"type\s+ID\s+string",
        combined,
    ):
        fail("Expected prayergroup.ID string type not found")

    ok("prayergroup.ID is string based")

    # --------------------------------------------------------
    # Status discovery
    # --------------------------------------------------------

    statuses = re.findall(
        r'(Status[A-Za-z0-9_]+)\s+Status\s*=\s*"([^"]+)"',
        combined,
    )

    if statuses:
        print_matches(
            "PrayerGroup statuses:",
            [
                f"{name} = {value}"
                for name, value in statuses
            ],
        )
    else:
        warn("No typed PrayerGroup statuses discovered")

    # --------------------------------------------------------
    # Type discovery
    # --------------------------------------------------------

    group_types = re.findall(
        r'(Type[A-Za-z0-9_]+)\s+Type\s*=\s*"([^"]+)"',
        combined,
    )

    if group_types:
        print_matches(
            "PrayerGroup types:",
            [
                f"{name} = {value}"
                for name, value in group_types
            ],
        )
    else:
        warn("No typed PrayerGroup types discovered")

    return {
        "combined": combined,
        "fields": fields,
        "constructor_args": constructor_args,
        "statuses": statuses,
        "types": group_types,
    }


# ============================================================
# Constructor contract
# ============================================================

def determine_constructor(domain):
    heading("2. CONSTRUCTOR CONTRACT")

    args = domain["constructor_args"]

    normalized = re.sub(
        r"\s+",
        " ",
        args,
    ).strip()

    info("Normalized constructor:")

    print()
    print(f"    New({normalized})")
    print()

    required_tokens = [
        "id ID",
        "name string",
    ]

    for token in required_tokens:
        info(f"Checking constructor token: {token}")

        if token not in normalized:
            fail(
                f"Constructor does not contain expected "
                f"argument: {token}"
            )

        ok(f"Found: {token}")

    has_description = "description string" in normalized
    has_type = "groupType Type" in normalized
    has_status = "status Status" in normalized
    has_sessions = (
        "sessionIDs" in normalized
        or "sessions" in normalized
    )

    print()
    print("Constructor capabilities:")
    print(f"  description : {has_description}")
    print(f"  group type  : {has_type}")
    print(f"  status      : {has_status}")
    print(f"  sessions    : {has_sessions}")
    print()

    return {
        "normalized": normalized,
        "description": has_description,
        "type": has_type,
        "status": has_status,
        "sessions": has_sessions,
    }


# ============================================================
# Permission
# ============================================================

def ensure_permission(root, config, backup_root):
    heading("3. PRAYER GROUP PERMISSION")

    relative = config["expected"]["permission_file"]

    path, content = read_file(
        root,
        relative,
    )

    constant = config["permission"]["constant"]
    value = config["permission"]["value"]

    info(f"Permission constant: {constant}")
    info(f"Permission value: {value}")

    if constant in content:
        ok(f"{constant} already exists")

        if value not in content:
            fail(
                f"{constant} exists but does not appear "
                f"to use value {value}"
            )

        ok("Permission value verified")

        return

    info("Permission does not exist yet")

    # --------------------------------------------------------
    # Locate permission const block.
    # --------------------------------------------------------

    marker = (
        'PermissionManageAccessGroups'
    )

    if marker not in content:
        fail(
            "Could not find PermissionManageAccessGroups. "
            "Refusing to guess permission.go structure."
        )

    ok("Existing permission structure identified")

    lines = content.splitlines()

    insert_index = None

    for index, line in enumerate(lines):
        if marker in line and "Permission" in line:
            insert_index = index + 1
            break

    if insert_index is None:
        fail("Could not determine permission insertion point")

    new_line = (
        f'\t{constant} Permission = "{value}"'
    )

    info(f"Adding: {new_line.strip()}")

    lines.insert(insert_index, new_line)

    updated = "\n".join(lines) + "\n"

    # --------------------------------------------------------
    # Add to AllPermissions.
    # --------------------------------------------------------

    info("Checking AllPermissions")

    all_permissions_match = re.search(
        r"var\s+AllPermissions\s*=\s*\[\]Permission\s*\{",
        updated,
    )

    if not all_permissions_match:
        fail("AllPermissions slice not found")

    ok("AllPermissions found")

    if re.search(
        rf"\b{re.escape(constant)}\b",
        updated[
            all_permissions_match.end():
        ],
    ):
        # Constant itself will also be elsewhere, so inspect slice.
        slice_start = all_permissions_match.end()
        slice_end = updated.find("}", slice_start)

        if slice_end == -1:
            fail("Could not determine AllPermissions end")

        slice_body = updated[
            slice_start:slice_end
        ]

        if constant not in slice_body:
            info("Adding permission to AllPermissions")

            updated = (
                updated[:slice_end]
                + f"\t{constant},\n"
                + updated[slice_end:]
            )
        else:
            ok("Permission already in AllPermissions")
    else:
        fail("Unexpected permission file structure")

    # Validate.
    if constant not in updated:
        fail("Permission constant missing after patch")

    if value not in updated:
        fail("Permission value missing after patch")

    ok("Permission patch validated")

    replace_and_write(
        root,
        relative,
        content,
        updated,
        backup_root,
    )


# ============================================================
# Admin permission
# ============================================================

def ensure_admin_permission(root, config, backup_root):
    heading("4. ADMIN ROLE PERMISSION")

    relative = config["expected"]["role_repository"]

    path, content = read_file(
        root,
        relative,
    )

    constant = config["permission"]["constant"]

    info("Looking for administrator role")

    if "role_admin" not in content:
        fail(
            "role_admin was not found in role_repository.go"
        )

    ok("Administrator role found")

    # Find role.New containing role_admin.
    admin_position = content.find('"role_admin"')

    start = content.rfind(
        "role.New(",
        0,
        admin_position,
    )

    if start == -1:
        fail(
            "Could not locate role.New for role_admin"
        )

    # Find matching close approximately using parentheses.
    open_pos = content.find("(", start)

    depth = 0
    end = None

    for index in range(open_pos, len(content)):
        char = content[index]

        if char == "(":
            depth += 1

        elif char == ")":
            depth -= 1

            if depth == 0:
                end = index + 1
                break

    if end is None:
        fail("Could not parse administrator role.New call")

    admin_block = content[start:end]

    print()
    print("Administrator role definition:")
    print("-" * 64)
    print(admin_block)
    print("-" * 64)
    print()

    if f"role.{constant}" in admin_block:
        ok(
            f"Administrator already has {constant}"
        )
        return

    info(
        f"Administrator does not yet have {constant}"
    )

    # Find permissions slice inside admin block.
    slice_marker = "[]role.Permission{"

    slice_start = admin_block.find(slice_marker)

    if slice_start == -1:
        fail(
            "Could not locate administrator permission slice"
        )

    slice_end = admin_block.find(
        "}",
        slice_start,
    )

    if slice_end == -1:
        fail(
            "Could not locate end of administrator "
            "permission slice"
        )

    info("Adding prayer-group management permission")

    patched_admin = (
        admin_block[:slice_end]
        + f"\t\t\trole.{constant},\n"
        + admin_block[slice_end:]
    )

    updated = (
        content[:start]
        + patched_admin
        + content[end:]
    )

    if f"role.{constant}" not in updated:
        fail("Admin permission patch validation failed")

    ok("Administrator permission patch validated")

    replace_and_write(
        root,
        relative,
        content,
        updated,
        backup_root,
    )


# ============================================================
# Repository discovery
# ============================================================

def discover_repository(root, config):
    heading("5. PRAYER GROUP REPOSITORY")

    relative = config["expected"]["prayer_group_repository"]

    path, content = read_file(
        root,
        relative,
    )

    methods = re.findall(
        r"func\s+\([^)]*\)\s+"
        r"([A-Za-z0-9_]+)\s*\(",
        content,
    )

    print_matches(
        "Repository methods discovered:",
        methods,
    )

    has_find = "FindByID" in methods
    has_save = "Save" in methods

    info(f"FindByID available: {has_find}")
    info(f"Save available: {has_save}")

    if not has_find:
        fail(
            "PrayerGroupRepository.FindByID is required "
            "before CreatePrayerGroup can safely verify IDs."
        )

    if not has_save:
        fail(
            "PrayerGroupRepository.Save is required "
            "for CreatePrayerGroup."
        )

    ok("Required repository operations exist")

    return content


# ============================================================
# ID generator discovery
# ============================================================

def discover_id_generator(root, config):
    heading("6. ID GENERATOR")

    relative = config["expected"]["id_generator"]

    path, content = read_file(
        root,
        relative,
    )

    methods = re.findall(
        r"func\s+\([^)]*\)\s+"
        r"([A-Za-z0-9_]+)\s*\([^)]*\)"
        r"\s+([^{\n]+)",
        content,
    )

    print_matches(
        "ID generator methods:",
        [
            f"{name} -> {return_type.strip()}"
            for name, return_type in methods
        ],
    )

    candidates = []

    for name, return_type in methods:
        if (
            "PrayerGroup" in name
            or "prayergroup.ID" in return_type
        ):
            candidates.append(
                (name, return_type.strip())
            )

    if not candidates:
        fail(
            "No PrayerGroup ID generator method discovered. "
            "Refusing to invent one."
        )

    if len(candidates) > 1:
        print_matches(
            "PrayerGroup ID candidates:",
            [
                f"{name} -> {return_type}"
                for name, return_type in candidates
            ],
        )

        fail(
            "Multiple PrayerGroup ID generator methods found."
        )

    method_name, return_type = candidates[0]

    ok(
        f"PrayerGroup ID generator: "
        f"{method_name} -> {return_type}"
    )

    return method_name


# ============================================================
# Existing wiring discovery
# ============================================================

def inspect_wiring(root, config):
    heading("7. HTTP / MAIN WIRING DISCOVERY")

    router_relative = config["expected"]["router"]
    main_relative = config["expected"]["main"]

    _, router = read_file(
        root,
        router_relative,
    )

    _, main = read_file(
        root,
        main_relative,
    )

    info("Inspecting NewRouter signature")

    router_signature = re.search(
        r"func\s+NewRouter\s*\((?P<args>.*?)\)"
        r"\s+http\.Handler",
        router,
        re.DOTALL,
    )

    if not router_signature:
        fail("Could not discover NewRouter signature")

    router_args = router_signature.group("args")

    print_matches(
        "NewRouter arguments:",
        [
            line.strip()
            for line in router_args.splitlines()
            if line.strip()
        ],
    )

    info("Inspecting main service wiring")

    interesting = []

    for line in main.splitlines():
        if (
            "New" in line
            or "NewRouter" in line
            or "Handler" in line
        ):
            interesting.append(line.strip())

    print_matches(
        "Relevant main.go lines:",
        interesting,
    )

    return {
        "router": router,
        "main": main,
    }


# ============================================================
# Generate application service
# ============================================================

def generate_application(
    constructor,
    id_method,
):
    heading("8. GENERATING APPLICATION SERVICE")

    if constructor["type"]:
        fail(
            "Existing PrayerGroup constructor requires groupType. "
            "The requested create API currently only defines name and "
            "description. Add an explicit API default/type decision "
            "before generating production code."
        )

    if constructor["status"]:
        fail(
            "Existing PrayerGroup constructor requires status. "
            "We need to determine the correct domain active constant "
            "before generating production code."
        )

    if constructor["sessions"]:
        fail(
            "Existing PrayerGroup constructor requires sessions. "
            "We need to determine the correct empty-session constructor "
            "contract before generating production code."
        )

    if not constructor["description"]:
        fail(
            "PrayerGroup constructor does not accept description. "
            "Refusing to generate an API contract that disagrees "
            "with the domain."
        )

    content = f'''package prayergroup

import (
\t"context"
\t"errors"
\t"strings"

\tdomainprayergroup "prayer-api/internal/domain/prayergroup"
\t"prayer-api/internal/domain/role"
\t"prayer-api/internal/domain/user"
)

var (
\tErrUnauthorized     = errors.New("unauthorized")
\tErrForbidden        = errors.New("forbidden")
\tErrActorNotFound    = errors.New("actor not found")
\tErrPrayerGroupExists = errors.New("prayer group already exists")
)

type UserRepository interface {{
\tFindByExternalID(
\t\tctx context.Context,
\t\texternalID string,
\t) (*user.User, error)
}}

type RoleRepository interface {{
\tFindByID(
\t\tctx context.Context,
\t\tid role.ID,
\t) (*role.Role, error)
}}

type PrayerGroupRepository interface {{
\tFindByID(
\t\tctx context.Context,
\t\tid domainprayergroup.ID,
\t) (*domainprayergroup.PrayerGroup, error)

\tSave(
\t\tctx context.Context,
\t\tgroup *domainprayergroup.PrayerGroup,
\t) error
}}

type IDGenerator interface {{
\t{id_method}() domainprayergroup.ID
}}

type CreateCommand struct {{
\tActorExternalID string
\tName            string
\tDescription     string
}}

type CreateResult struct {{
\tPrayerGroup *domainprayergroup.PrayerGroup
}}

type CreateService struct {{
\tusers  UserRepository
\troles  RoleRepository
\tgroups PrayerGroupRepository
\tids    IDGenerator
}}

func NewCreateService(
\tusers UserRepository,
\troles RoleRepository,
\tgroups PrayerGroupRepository,
\tids IDGenerator,
) *CreateService {{
\treturn &CreateService{{
\t\tusers:  users,
\t\troles:  roles,
\t\tgroups: groups,
\t\tids:    ids,
\t}}
}}

func (s *CreateService) Create(
\tctx context.Context,
\tcommand CreateCommand,
) (*CreateResult, error) {{
\tif strings.TrimSpace(command.ActorExternalID) == "" {{
\t\treturn nil, ErrUnauthorized
\t}}

\tactor, err := s.users.FindByExternalID(
\t\tctx,
\t\tcommand.ActorExternalID,
\t)

\tif err != nil {{
\t\treturn nil, err
\t}}

\tif actor == nil {{
\t\treturn nil, ErrActorNotFound
\t}}

\tallowed := false

\tfor _, roleID := range actor.RoleIDs {{
\t\tresolvedRole, err := s.roles.FindByID(
\t\t\tctx,
\t\t\troleID,
\t\t)

\t\tif err != nil {{
\t\t\treturn nil, err
\t\t}}

\t\tif resolvedRole == nil {{
\t\t\tcontinue
\t\t}}

\t\tif resolvedRole.HasPermission(
\t\t\trole.PermissionManagePrayerGroups,
\t\t) {{
\t\t\tallowed = true
\t\t\tbreak
\t\t}}
\t}}

\tif !allowed {{
\t\treturn nil, ErrForbidden
\t}}

\tgroupID := s.ids.{id_method}()

\texisting, err := s.groups.FindByID(
\t\tctx,
\t\tgroupID,
\t)

\tif err != nil {{
\t\treturn nil, err
\t}}

\tif existing != nil {{
\t\treturn nil, ErrPrayerGroupExists
\t}}

\tgroup, err := domainprayergroup.New(
\t\tgroupID,
\t\tcommand.Name,
\t\tcommand.Description,
\t)

\tif err != nil {{
\t\treturn nil, err
\t}}

\tif err := s.groups.Save(
\t\tctx,
\t\tgroup,
\t); err != nil {{
\t\treturn nil, err
\t}}

\treturn &CreateResult{{
\t\tPrayerGroup: group,
\t}}, nil
}}
'''

    ok("Application service generated in memory")

    return content


# ============================================================
# Generate HTTP handler
# ============================================================

def generate_handler():
    heading("9. GENERATING HTTP HANDLER")

    content = '''package httpapi

import (
\t"encoding/json"
\t"errors"
\t"net/http"

\t"github.com/clerk/clerk-sdk-go/v2"

\tappprayergroup "prayer-api/internal/application/prayergroup"
)

type PrayerGroupHandler struct {
\tcreate *appprayergroup.CreateService
}

func NewPrayerGroupHandler(
\tcreate *appprayergroup.CreateService,
) *PrayerGroupHandler {
\treturn &PrayerGroupHandler{
\t\tcreate: create,
\t}
}

type createPrayerGroupRequest struct {
\tName        string `json:"name"`
\tDescription string `json:"description"`
}

func (h *PrayerGroupHandler) Create(
\tw http.ResponseWriter,
\tr *http.Request,
) {
\tclaims, ok := clerk.SessionClaimsFromContext(
\t\tr.Context(),
\t)

\tif !ok {
\t\twriteJSON(
\t\t\tw,
\t\t\thttp.StatusUnauthorized,
\t\t\tmap[string]any{
\t\t\t\t"error": "unauthorized",
\t\t\t},
\t\t)

\t\treturn
\t}

\tvar request createPrayerGroupRequest

\tif err := json.NewDecoder(
\t\tr.Body,
\t).Decode(&request); err != nil {
\t\twriteJSON(
\t\t\tw,
\t\t\thttp.StatusBadRequest,
\t\t\tmap[string]any{
\t\t\t\t"error": "invalid request body",
\t\t\t},
\t\t)

\t\treturn
\t}

\tresult, err := h.create.Create(
\t\tr.Context(),
\t\tappprayergroup.CreateCommand{
\t\t\tActorExternalID: claims.Subject,
\t\t\tName:            request.Name,
\t\t\tDescription:     request.Description,
\t\t},
\t)

\tif err != nil {
\t\tswitch {
\t\tcase errors.Is(
\t\t\terr,
\t\t\tappprayergroup.ErrUnauthorized,
\t\t):
\t\t\twriteJSON(
\t\t\t\tw,
\t\t\t\thttp.StatusUnauthorized,
\t\t\t\tmap[string]any{
\t\t\t\t\t"error": "unauthorized",
\t\t\t\t},
\t\t\t)

\t\tcase errors.Is(
\t\t\terr,
\t\t\tappprayergroup.ErrActorNotFound,
\t\t):
\t\t\twriteJSON(
\t\t\t\tw,
\t\t\t\thttp.StatusNotFound,
\t\t\t\tmap[string]any{
\t\t\t\t\t"error": "actor not found",
\t\t\t\t},
\t\t\t)

\t\tcase errors.Is(
\t\t\terr,
\t\t\tappprayergroup.ErrForbidden,
\t\t):
\t\t\twriteJSON(
\t\t\t\tw,
\t\t\t\thttp.StatusForbidden,
\t\t\t\tmap[string]any{
\t\t\t\t\t"error": "forbidden",
\t\t\t\t},
\t\t\t)

\t\tcase errors.Is(
\t\t\terr,
\t\t\tappprayergroup.ErrPrayerGroupExists,
\t\t):
\t\t\twriteJSON(
\t\t\t\tw,
\t\t\t\thttp.StatusConflict,
\t\t\t\tmap[string]any{
\t\t\t\t\t"error": "prayer group already exists",
\t\t\t\t},
\t\t\t)

\t\tdefault:
\t\t\twriteJSON(
\t\t\t\tw,
\t\t\t\thttp.StatusBadRequest,
\t\t\t\tmap[string]any{
\t\t\t\t\t"error": err.Error(),
\t\t\t\t},
\t\t\t)
\t\t}

\t\treturn
\t}

\tgroup := result.PrayerGroup

\twriteJSON(
\t\tw,
\t\thttp.StatusCreated,
\t\tmap[string]any{
\t\t\t"prayerGroup": map[string]any{
\t\t\t\t"id":          group.ID,
\t\t\t\t"name":        group.Name,
\t\t\t\t"description": group.Description,
\t\t\t\t"status":      group.Status,
\t\t\t},
\t\t},
\t)
}
'''

    ok("HTTP handler generated in memory")

    return content


# ============================================================
# Wiring patch
# ============================================================

def patch_router(
    root,
    config,
    backup_root,
):
    heading("10. ROUTER PATCH")

    relative = config["expected"]["router"]

    path, content = read_file(
        root,
        relative,
    )

    route = '/api/prayer-groups'

    if route in content:
        ok("PrayerGroup create route already exists")
        return

    info("PrayerGroup route not found")

    # We only support the currently expected simple NewRouter
    # pattern after discovery. If structure differs, stop.
    signature_match = re.search(
        r"func\s+NewRouter\s*\((?P<args>.*?)\)"
        r"\s+http\.Handler\s*\{",
        content,
        re.DOTALL,
    )

    if not signature_match:
        fail("Could not safely patch NewRouter")

    args = signature_match.group("args")

    info("Existing NewRouter arguments:")
    print(args)

    if "prayerGroupHandler" in args:
        ok("PrayerGroupHandler already wired")
        return

    # Insert parameter before closing parenthesis.
    new_args = args.rstrip()

    if new_args and not new_args.rstrip().endswith(","):
        new_args += ","

    new_args += (
        "\n\tprayerGroupHandler *PrayerGroupHandler,\n"
    )

    updated = (
        content[:signature_match.start("args")]
        + new_args
        + content[signature_match.end("args"):]
    )

    # Find a mux HandleFunc registration area.
    return_position = updated.rfind("return ")

    if return_position == -1:
        fail(
            "Could not locate router return statement "
            "for route insertion"
        )

    route_line = (
        '\tmux.HandleFunc('
        '"POST /api/prayer-groups", '
        'prayerGroupHandler.Create'
        ')\n\n'
    )

    # Ensure mux exists.
    if "mux." not in updated:
        fail(
            "Router does not appear to use mux directly. "
            "Refusing unsafe route patch."
        )

    updated = (
        updated[:return_position]
        + route_line
        + updated[return_position:]
    )

    if "prayerGroupHandler.Create" not in updated:
        fail("Router patch validation failed")

    ok("Router patch validated")

    replace_and_write(
        root,
        relative,
        content,
        updated,
        backup_root,
    )


def patch_main(
    root,
    config,
    backup_root,
):
    heading("11. MAIN WIRING PATCH")

    relative = config["expected"]["main"]

    path, content = read_file(
        root,
        relative,
    )

    if "NewPrayerGroupHandler" in content:
        ok("PrayerGroup handler already wired in main")
        return

    info("PrayerGroup handler not yet wired")

    # Add application import.
    auth_import = (
        'appauth "prayer-api/internal/application/auth"'
    )

    if auth_import not in content:
        fail(
            "Expected appauth import not found. "
            "Refusing to guess import layout."
        )

    updated = content.replace(
        auth_import,
        auth_import
        + '\n\tappprayergroup '
          '"prayer-api/internal/application/prayergroup"',
        1,
    )

    ok("PrayerGroup application import added")

    # Find auth handler creation.
    marker = (
        "authHandler := httpapi.NewAuthHandler(loginService)"
    )

    if marker not in updated:
        fail(
            "Expected authHandler wiring not found. "
            "Refusing unsafe main.go patch."
        )

    wiring = '''authHandler := httpapi.NewAuthHandler(loginService)

\tprayerGroupCreateService := appprayergroup.NewCreateService(
\t\tusers,
\t\troles,
\t\tgroups,
\t\tids,
\t)

\tprayerGroupHandler := httpapi.NewPrayerGroupHandler(
\t\tprayerGroupCreateService,
\t)'''

    updated = updated.replace(
        marker,
        wiring,
        1,
    )

    ok("PrayerGroup service and handler wiring added")

    # Patch NewRouter call.
    router_match = re.search(
        r"router\s*:=\s*httpapi\.NewRouter\s*\((?P<args>.*?)\)",
        updated,
        re.DOTALL,
    )

    if not router_match:
        fail("Could not find NewRouter call in main.go")

    args = router_match.group("args").strip()

    info("Current NewRouter call arguments:")
    print()
    print(args)
    print()

    if "prayerGroupHandler" not in args:
        if args and not args.endswith(","):
            args += ","

        args += "\n\t\tprayerGroupHandler,\n\t"

        replacement = (
            "router := httpapi.NewRouter("
            + args
            + ")"
        )

        updated = (
            updated[:router_match.start()]
            + replacement
            + updated[router_match.end():]
        )

        ok("PrayerGroupHandler added to NewRouter call")
    else:
        ok("PrayerGroupHandler already in NewRouter call")

    replace_and_write(
        root,
        relative,
        content,
        updated,
        backup_root,
    )


# ============================================================
# Main
# ============================================================

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
        / "create-prayer-group-api"
    )

    info(f"Backup root: {backup_root}")

    # --------------------------------------------------------
    # DISCOVERY FIRST
    # --------------------------------------------------------

    domain = discover_prayer_group_domain(
        root,
        config,
    )

    constructor = determine_constructor(
        domain
    )

    discover_repository(
        root,
        config,
    )

    id_method = discover_id_generator(
        root,
        config,
    )

    inspect_wiring(
        root,
        config,
    )

    # --------------------------------------------------------
    # MODIFY PERMISSION MODEL
    # --------------------------------------------------------

    ensure_permission(
        root,
        config,
        backup_root,
    )

    ensure_admin_permission(
        root,
        config,
        backup_root,
    )

    # --------------------------------------------------------
    # GENERATE FEATURE
    # --------------------------------------------------------

    application = generate_application(
        constructor,
        id_method,
    )

    handler = generate_handler()

    write_file(
        root,
        config["generated"]["application"],
        application,
        backup_root,
    )

    write_file(
        root,
        config["generated"]["handler"],
        handler,
        backup_root,
    )

    # --------------------------------------------------------
    # WIRING
    # --------------------------------------------------------

    patch_router(
        root,
        config,
        backup_root,
    )

    patch_main(
        root,
        config,
        backup_root,
    )

    # --------------------------------------------------------
    # COMPLETE
    # --------------------------------------------------------

    heading("CREATE PRAYER GROUP API GENERATED")

    print("Endpoint:")
    print()
    print("    POST /api/prayer-groups")
    print()

    print("Permission:")
    print()
    print("    prayer_groups:manage")
    print()

    print("Request:")
    print()
    print('    {')
    print('      "name": "Young Adults",')
    print(
        '      "description": '
        '"Prayer group for young adults"'
    )
    print('    }')
    print()

    print("Expected success:")
    print()
    print("    HTTP 201 Created")
    print()

    print("Next commands:")
    print()
    print("    cd prayer-api")
    print("    go fmt ./...")
    print("    go test -v ./...")
    print()
    print(
        "Do NOT mark the endpoint Implemented yet."
    )
    print(
        "We still need CreatePrayerGroup "
        "application and HTTP tests."
    )


if __name__ == "__main__":
    main()