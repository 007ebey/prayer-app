import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


BASE = Path(__file__).resolve().parent
CONFIG_PATH = BASE / "add_prayer_group_save.json"


def info(message):
    print(f"[INFO] {message}")


def ok(message):
    print(f"[ OK ] {message}")


def fail(message):
    print(f"[FAIL] {message}")
    print()
    print("No files were changed.")
    sys.exit(1)


def heading(message):
    print()
    print("=" * 64)
    print(message)
    print("=" * 64)
    print()


heading("ADD PRAYER GROUP REPOSITORY SAVE")


# ------------------------------------------------------------
# Configuration
# ------------------------------------------------------------

info(f"Looking for config: {CONFIG_PATH}")

if not CONFIG_PATH.exists():
    fail("Configuration file does not exist")

ok("Configuration file exists")

info("Loading configuration")

try:
    config = json.loads(
        CONFIG_PATH.read_text(encoding="utf-8")
    )
except Exception as error:
    fail(f"Unable to parse configuration: {error}")

ok("Configuration loaded")


# ------------------------------------------------------------
# Project
# ------------------------------------------------------------

project = BASE / config["project"]

info(f"Project: {project}")

if not project.exists():
    fail("Project directory does not exist")

ok("Project directory exists")

go_mod = project / "go.mod"

info("Checking go.mod")

if not go_mod.exists():
    fail("go.mod does not exist")

ok("Go project confirmed")


# ------------------------------------------------------------
# Repository
# ------------------------------------------------------------

relative = Path(config["repository"])
repository_path = project / relative

heading("REPOSITORY DISCOVERY")

info(f"Repository: {relative}")

if not repository_path.exists():
    fail("PrayerGroup repository does not exist")

ok("Repository exists")

info("Reading repository")

source = repository_path.read_text(
    encoding="utf-8"
)

ok(f"Repository read ({len(source)} characters)")


# ------------------------------------------------------------
# Package
# ------------------------------------------------------------

info("Checking package")

package_match = re.search(
    r"^package\s+([A-Za-z0-9_]+)",
    source,
    re.MULTILINE,
)

if not package_match:
    fail("Could not discover package declaration")

package_name = package_match.group(1)

ok(f"Package: {package_name}")


# ------------------------------------------------------------
# Repository struct
# ------------------------------------------------------------

info("Looking for RoleRepository-like PrayerGroup repository struct")

struct_match = re.search(
    r"type\s+PrayerGroupRepository\s+struct\s*\{"
    r"(?P<body>.*?)"
    r"\}",
    source,
    re.DOTALL,
)

if not struct_match:
    fail("PrayerGroupRepository struct was not found")

ok("PrayerGroupRepository struct found")

struct_body = struct_match.group("body")

print()
print("Current repository fields:")
print("-" * 64)
print(struct_body.strip())
print("-" * 64)
print()


# ------------------------------------------------------------
# Existing methods
# ------------------------------------------------------------

info("Discovering repository methods")

methods = re.findall(
    r"func\s+\([^)]*\)\s+"
    r"([A-Za-z0-9_]+)\s*\(",
    source,
)

for method in methods:
    print(f"  - {method}")

print()

if "Save" in methods:
    ok("Save already exists")
    print()
    print("Nothing needs to be changed.")
    sys.exit(0)


# ------------------------------------------------------------
# Discover FindByID implementation
# ------------------------------------------------------------

heading("GROUP STORAGE DISCOVERY")

info("Locating FindByID")

find_match = re.search(
    r"func\s+\("
    r"(?P<receiver>[A-Za-z0-9_]+)\s+\*PrayerGroupRepository"
    r"\)\s+FindByID\s*\("
    r"(?P<args>.*?)"
    r"\)\s*"
    r"(?P<returns>[^{]+)"
    r"\{",
    source,
    re.DOTALL,
)

if not find_match:
    fail("Could not locate PrayerGroupRepository.FindByID")

receiver = find_match.group("receiver")

ok(f"FindByID receiver: {receiver}")

body_start = find_match.end()

depth = 1
position = body_start

while position < len(source) and depth > 0:
    if source[position] == "{":
        depth += 1
    elif source[position] == "}":
        depth -= 1

    position += 1

if depth != 0:
    fail("Could not determine FindByID function boundary")

find_body = source[
    body_start:position - 1
]

print()
print("FindByID implementation:")
print("-" * 64)
print(find_body.strip())
print("-" * 64)
print()


# ------------------------------------------------------------
# Discover map used by FindByID
# ------------------------------------------------------------

info("Discovering PrayerGroup storage field")

patterns = [
    rf"{re.escape(receiver)}\.([A-Za-z0-9_]+)\s*\[\s*id\s*\]",
    rf"{re.escape(receiver)}\.([A-Za-z0-9_]+)\s*\[",
]

storage_field = None

for pattern in patterns:
    match = re.search(
        pattern,
        find_body,
    )

    if match:
        storage_field = match.group(1)
        break

if storage_field is None:
    fail(
        "Could not determine which repository field "
        "FindByID uses to store PrayerGroups"
    )

ok(f"PrayerGroup storage field: {storage_field}")


# ------------------------------------------------------------
# Verify storage field type
# ------------------------------------------------------------

info(
    f"Checking type of repository field "
    f"'{storage_field}'"
)

field_match = re.search(
    rf"\b{re.escape(storage_field)}\s+"
    r"(?P<type>[^\n]+)",
    struct_body,
)

if not field_match:
    fail(
        f"Could not determine type of "
        f"repository field '{storage_field}'"
    )

storage_type = field_match.group("type").strip()

ok(
    f"{storage_field} type: "
    f"{storage_type}"
)

if "map[" not in storage_type:
    fail(
        f"PrayerGroup storage '{storage_field}' "
        f"is not map based. Refusing to guess Save behavior."
    )

if "*prayergroup.PrayerGroup" not in storage_type:
    fail(
        "PrayerGroup storage does not appear to contain "
        "*prayergroup.PrayerGroup values."
    )

ok("PrayerGroup map storage verified")


# ------------------------------------------------------------
# Verify ID field
# ------------------------------------------------------------

info("Verifying PrayerGroup has ID field")

domain_dir = (
    project
    / "internal"
    / "domain"
    / "prayergroup"
)

if not domain_dir.exists():
    fail("PrayerGroup domain directory missing")

domain_source = ""

for file in sorted(domain_dir.glob("*.go")):
    if file.name.endswith("_test.go"):
        continue

    info(
        f"Reading domain file: "
        f"{file.relative_to(project)}"
    )

    domain_source += (
        "\n"
        + file.read_text(encoding="utf-8")
    )

if not re.search(
    r"type\s+PrayerGroup\s+struct\s*\{"
    r"(?P<body>.*?)"
    r"\}",
    domain_source,
    re.DOTALL,
):
    fail("PrayerGroup struct not found")

if not re.search(
    r"\bID\s+ID\b",
    domain_source,
):
    fail("PrayerGroup.ID field could not be verified")

ok("PrayerGroup.ID verified")


# ------------------------------------------------------------
# Generate Save
# ------------------------------------------------------------

heading("GENERATING SAVE METHOD")

method = f'''

func (r *PrayerGroupRepository) Save(
\tctx context.Context,
\tgroup *prayergroup.PrayerGroup,
) error {{
\tr.{storage_field}[group.ID] = group
\treturn nil
}}
'''

print("Generated method:")
print("-" * 64)
print(method.strip())
print("-" * 64)
print()


# ------------------------------------------------------------
# Validate context import
# ------------------------------------------------------------

info("Checking context import")

if '"context"' not in source:
    fail(
        "Repository does not import context. "
        "Unexpected structure; refusing automatic import patch."
    )

ok("context import exists")


# ------------------------------------------------------------
# Validate prayergroup import
# ------------------------------------------------------------

info("Checking prayergroup import")

if '"prayer-api/internal/domain/prayergroup"' not in source:
    fail(
        "Repository does not import prayergroup domain. "
        "Unexpected repository structure."
    )

ok("prayergroup import exists")


# ------------------------------------------------------------
# Prepare update
# ------------------------------------------------------------

heading("PRE-WRITE VALIDATION")

updated = source.rstrip() + method

checks = {
    "Save method":
        "func (r *PrayerGroupRepository) Save(",
    "context argument":
        "ctx context.Context",
    "PrayerGroup argument":
        "group *prayergroup.PrayerGroup",
    "map assignment":
        f"r.{storage_field}[group.ID] = group",
    "nil return":
        "return nil",
}

for name, value in checks.items():
    info(f"Checking: {name}")

    if value not in updated:
        fail(f"Validation failed: {name}")

    ok(f"Validated: {name}")


# ------------------------------------------------------------
# Backup
# ------------------------------------------------------------

timestamp = datetime.now().strftime(
    "%Y%m%d_%H%M%S"
)

backup = (
    project
    / ".generator-backups"
    / timestamp
    / "add-prayer-group-save"
    / relative
)

info(f"Creating backup: {backup}")

backup.parent.mkdir(
    parents=True,
    exist_ok=True,
)

shutil.copy2(
    repository_path,
    backup,
)

ok("Backup created")


# ------------------------------------------------------------
# Write
# ------------------------------------------------------------

info("Writing updated repository")

repository_path.write_text(
    updated,
    encoding="utf-8",
)

ok("Repository written")


# ------------------------------------------------------------
# Read-back verification
# ------------------------------------------------------------

heading("POST-WRITE VERIFICATION")

info("Reading repository back")

written = repository_path.read_text(
    encoding="utf-8"
)

ok("Repository read")

for name, value in checks.items():
    info(f"Verifying written file: {name}")

    if value not in written:
        fail(
            f"Post-write verification failed: {name}"
        )

    ok(f"Verified: {name}")


heading("SAVE METHOD ADDED")

print(f"Repository : {relative}")
print(f"Storage    : {storage_field}")
print(f"Backup     : {backup}")
print()
print("Next:")
print()
print("    python .\\create_prayer_group_api.py")
print()