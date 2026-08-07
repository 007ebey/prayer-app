import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


BASE = Path(__file__).resolve().parent
CONFIG_FILE = BASE / "add_prayer_group_permission.json"


def info(message):
    print(f"[INFO] {message}")


def ok(message):
    print(f"[ OK ] {message}")


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


heading("ADD PRAYER GROUP PERMISSION")


# ============================================================
# 1. CONFIG
# ============================================================

info(f"Config: {CONFIG_FILE}")

if not CONFIG_FILE.exists():
    fail("Configuration file missing")

ok("Configuration file exists")

try:
    config = json.loads(
        CONFIG_FILE.read_text(encoding="utf-8")
    )
except Exception as error:
    fail(f"Could not load configuration: {error}")

ok("Configuration loaded")


# ============================================================
# 2. PROJECT
# ============================================================

project = BASE / config["project"]

info(f"Project: {project}")

if not project.exists():
    fail("Project directory missing")

ok("Project directory exists")

go_mod = project / "go.mod"

info("Checking go.mod")

if not go_mod.exists():
    fail("go.mod missing")

ok("Go project confirmed")


# ============================================================
# 3. PERMISSION FILE
# ============================================================

heading("PERMISSION DOMAIN DISCOVERY")

relative = Path(config["permission_file"])
path = project / relative

info(f"Permission file: {relative}")

if not path.exists():
    fail("Permission file missing")

ok("Permission file exists")

source = path.read_text(encoding="utf-8")

ok(f"Permission file read ({len(source)} characters)")

print()
print("Current permission.go:")
print("-" * 64)
print(source.rstrip())
print("-" * 64)
print()


# ============================================================
# 4. VERIFY TYPE
# ============================================================

info("Checking Permission type")

if not re.search(
    r"type\s+Permission\s+string",
    source,
):
    fail("type Permission string not found")

ok("Permission type verified")


# ============================================================
# 5. DISCOVER CONST BLOCK
# ============================================================

heading("CONST BLOCK DISCOVERY")

info("Looking for permission const block")

const_match = re.search(
    r"const\s*\((?P<body>.*?)\)",
    source,
    re.DOTALL,
)

if not const_match:
    fail("Permission const block not found")

ok("Permission const block found")

const_body = const_match.group("body")

print()
print("Current constants:")
print("-" * 64)
print(const_body.strip())
print("-" * 64)
print()


# ============================================================
# 6. DISCOVER EXISTING PERMISSIONS
# ============================================================

info("Discovering permission constants")

permissions = re.findall(
    r"(Permission[A-Za-z0-9_]+)"
    r"\s+Permission\s*=\s*"
    r'"([^"]+)"',
    const_body,
)

if not permissions:
    fail("No permission constants discovered")

for name, value in permissions:
    print(f"  {name} = {value}")

print()

ok(f"Discovered {len(permissions)} permissions")


# ============================================================
# 7. TARGET PERMISSION
# ============================================================

heading("TARGET PERMISSION")

constant = config["permission_constant"]
value = config["permission_value"]

info(f"Constant: {constant}")
info(f"Value   : {value}")


constant_exists = any(
    name == constant
    for name, _ in permissions
)

value_exists = any(
    current_value == value
    for _, current_value in permissions
)


if constant_exists:
    ok(f"{constant} already exists")

if value_exists:
    ok(f'Permission value "{value}" already exists')


if constant_exists != value_exists:
    fail(
        "Permission domain is inconsistent: "
        "constant/value only partially exists"
    )


# ============================================================
# 8. DISCOVER AllPermissions
# ============================================================

heading("ALL PERMISSIONS DISCOVERY")

info("Looking for AllPermissions")

all_match = re.search(
    r"var\s+AllPermissions\s*=\s*"
    r"\[\]Permission\s*\{"
    r"(?P<body>.*?)"
    r"\}",
    source,
    re.DOTALL,
)

if not all_match:
    fail(
        "Expected structure not found: "
        "var AllPermissions = []Permission{...}"
    )

ok("AllPermissions found")

all_body = all_match.group("body")

all_entries = re.findall(
    r"\b(Permission[A-Za-z0-9_]+)\b",
    all_body,
)

print()
print("Current AllPermissions:")
print("-" * 64)

for entry in all_entries:
    print(f"  {entry}")

print("-" * 64)
print()


if not all_entries:
    fail("AllPermissions contains no permission constants")

ok(
    f"AllPermissions contains "
    f"{len(all_entries)} entries"
)


# ============================================================
# 9. CHECK DOMAIN CONSISTENCY
# ============================================================

heading("DOMAIN CONSISTENCY")

declared_names = {
    name
    for name, _ in permissions
}

registered_names = set(all_entries)


info("Checking declared permissions are registered")

missing = declared_names - registered_names

if missing:
    print()
    print("Missing from AllPermissions:")

    for name in sorted(missing):
        print(f"  - {name}")

    fail(
        "Existing permission domain is inconsistent. "
        "Refusing to hide the problem while adding "
        "another permission."
    )

ok("All declared permissions are registered")


info("Checking AllPermissions entries are declared")

unknown = registered_names - declared_names

if unknown:
    print()
    print("Unknown AllPermissions entries:")

    for name in sorted(unknown):
        print(f"  - {name}")

    fail(
        "AllPermissions references unknown constants"
    )

ok("AllPermissions entries are valid")


# ============================================================
# 10. IDEMPOTENCY
# ============================================================

if (
    constant_exists
    and constant in registered_names
):
    heading("NO CHANGE REQUIRED")

    print(
        f"{constant} already exists and is "
        f"registered in AllPermissions."
    )
    print()
    print("Next:")
    print()
    print(
        "    python .\\create_prayer_group_api.py"
    )
    print()

    sys.exit(0)


# ============================================================
# 11. ADD CONSTANT
# ============================================================

heading("GENERATE PATCH")

info("Adding permission constant")

new_constant_line = (
    f'\n\t{constant} '
    f'Permission = "{value}"'
)

new_const_body = (
    const_body.rstrip()
    + new_constant_line
    + "\n"
)

updated = (
    source[:const_match.start("body")]
    + new_const_body
    + source[const_match.end("body"):]
)

ok("Permission constant added in memory")


# ============================================================
# 12. REDISCOVER AllPermissions AFTER CONST CHANGE
# ============================================================

info("Rediscovering AllPermissions in updated source")

updated_all_match = re.search(
    r"var\s+AllPermissions\s*=\s*"
    r"\[\]Permission\s*\{"
    r"(?P<body>.*?)"
    r"\}",
    updated,
    re.DOTALL,
)

if not updated_all_match:
    fail(
        "Could not rediscover AllPermissions "
        "after const patch"
    )

ok("AllPermissions rediscovered")


# ============================================================
# 13. ADD TO AllPermissions
# ============================================================

info("Adding permission to AllPermissions")

updated_all_body = updated_all_match.group("body")

if re.search(
    rf"\b{re.escape(constant)}\b",
    updated_all_body,
):
    ok("Permission already registered")

else:
    new_all_body = (
        updated_all_body.rstrip()
        + f"\n\t{constant},\n"
    )

    updated = (
        updated[:updated_all_match.start("body")]
        + new_all_body
        + updated[updated_all_match.end("body"):]
    )

    ok("Permission registered in memory")


# ============================================================
# 14. PRE-WRITE VALIDATION
# ============================================================

heading("PRE-WRITE VALIDATION")

checks = {
    "permission constant":
        constant,

    "permission value":
        f'"{value}"',

    "typed declaration":
        f'{constant} Permission = "{value}"',

    "AllPermissions":
        "var AllPermissions = []Permission{",
}

for name, token in checks.items():
    info(f"Checking: {name}")

    if token not in updated:
        fail(f"Validation failed: {name}")

    ok(f"Validated: {name}")


info("Checking target appears exactly twice")

occurrences = len(
    re.findall(
        rf"\b{re.escape(constant)}\b",
        updated,
    )
)

# One declaration + one AllPermissions entry.
if occurrences != 2:
    fail(
        f"Expected {constant} exactly twice, "
        f"found {occurrences}"
    )

ok(
    "Target occurs exactly twice "
    "(declaration + registry)"
)


# ============================================================
# 15. BACKUP
# ============================================================

timestamp = datetime.now().strftime(
    "%Y%m%d_%H%M%S"
)

backup_path = (
    project
    / ".generator-backups"
    / timestamp
    / "add-prayer-group-permission"
    / relative
)

info(f"Creating backup: {backup_path}")

backup_path.parent.mkdir(
    parents=True,
    exist_ok=True,
)

shutil.copy2(
    path,
    backup_path,
)

ok("Backup created")


# ============================================================
# 16. WRITE
# ============================================================

heading("WRITE")

info(f"Writing: {relative}")

path.write_text(
    updated,
    encoding="utf-8",
)

ok("Permission file written")


# ============================================================
# 17. READ-BACK
# ============================================================

heading("POST-WRITE VERIFICATION")

info("Reading permission file back")

written = path.read_text(
    encoding="utf-8",
)

ok("Permission file read")


info("Checking constant")

if (
    f'{constant} Permission = "{value}"'
    not in written
):
    fail(
        "Permission constant missing "
        "after write"
    )

ok("Permission constant verified")


info("Checking AllPermissions registration")

written_all = re.search(
    r"var\s+AllPermissions\s*=\s*"
    r"\[\]Permission\s*\{"
    r"(?P<body>.*?)"
    r"\}",
    written,
    re.DOTALL,
)

if not written_all:
    fail(
        "AllPermissions missing after write"
    )

if not re.search(
    rf"\b{re.escape(constant)}\b",
    written_all.group("body"),
):
    fail(
        "Permission missing from "
        "AllPermissions after write"
    )

ok("AllPermissions registration verified")


# ============================================================
# COMPLETE
# ============================================================

heading("PRAYER GROUP PERMISSION ADDED")

print(f"Constant : {constant}")
print(f"Value    : {value}")
print(f"File     : {relative}")
print(f"Backup   : {backup_path}")
print()

print("Next:")
print()
print(
    "    python .\\create_prayer_group_api.py"
)
print()