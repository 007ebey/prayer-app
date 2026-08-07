import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


BASE = Path(__file__).resolve().parent
PATCH_CONFIG = BASE / "patch_create_prayer_group_generator.json"


def info(message):
    print(f"[INFO] {message}")


def ok(message):
    print(f"[ OK ] {message}")


def fail(message):
    print()
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


heading("PATCH CREATE PRAYER GROUP GENERATOR")


# ============================================================
# 1. CONFIG
# ============================================================

info(f"Patch config: {PATCH_CONFIG}")

if not PATCH_CONFIG.exists():
    fail("Patch configuration missing")

ok("Patch configuration exists")

try:
    patch_config = json.loads(
        PATCH_CONFIG.read_text(encoding="utf-8")
    )
except Exception as error:
    fail(f"Could not load patch configuration: {error}")

ok("Patch configuration loaded")


generator_path = (
    BASE / patch_config["generator"]
)

generator_config_path = (
    BASE / patch_config["generator_config"]
)

expected_policy = patch_config["expected_policy"]

expected_group_type = expected_policy["group_type"]
expected_source = expected_policy["group_type_source"]
expected_allow_client = expected_policy["allow_client_group_type"]


# ============================================================
# 2. GENERATOR
# ============================================================

heading("1. GENERATOR DISCOVERY")

info(f"Generator: {generator_path}")

if not generator_path.exists():
    fail("create_prayer_group_api.py missing")

ok("Generator exists")

source = generator_path.read_text(
    encoding="utf-8"
)

ok(
    f"Generator read ({len(source)} characters)"
)


# ============================================================
# 3. VERIFY FAILURE GUARD
# ============================================================

heading("2. GROUP TYPE GUARD DISCOVERY")

needle = (
    "Existing PrayerGroup constructor requires groupType."
)

info("Looking for current groupType failure guard")

if needle not in source:
    fail(
        "Expected groupType failure message was not found. "
        "Generator differs from the version this patch targets."
    )

ok("Current groupType failure guard found")


# ============================================================
# 4. SHOW SURROUNDING SOURCE
# ============================================================

index = source.find(needle)

start = max(
    0,
    source.rfind("\n", 0, max(0, index - 1000))
)

end = source.find(
    "\n",
    index + len(needle) + 1000,
)

if end == -1:
    end = len(source)


print()
print("Current generator section:")
print("-" * 64)
print(source[start:end])
print("-" * 64)
print()


# ============================================================
# 5. CONFIG UPDATE
# ============================================================

heading("3. GENERATOR CONFIG")

info(
    f"Config file: {generator_config_path}"
)

if not generator_config_path.exists():
    fail("create_prayer_group_api.json missing")

ok("Generator config exists")

try:
    create_config = json.loads(
        generator_config_path.read_text(
            encoding="utf-8"
        )
    )
except Exception as error:
    fail(
        f"Could not load generator config: {error}"
    )

ok("Generator config loaded")


current_type = create_config.get(
    "create_group_type"
)

if current_type is None:
    info(
        "create_group_type is not configured"
    )

    create_config[
        "create_group_type"
    ] = group_type

    ok(
        f"Will configure create_group_type="
        f"{group_type}"
    )

elif current_type == group_type:
    ok(
        f"create_group_type already equals "
        f"{group_type}"
    )

else:
    fail(
        f"create_group_type already has unexpected "
        f"value: {current_type}"
    )


# ============================================================
# 6. FIND EXACT FAILURE STATEMENT
# ============================================================

heading("4. FAILURE STATEMENT DISCOVERY")

# We deliberately anchor on the exact known failure message
# instead of guessing function names.

failure_pattern = re.compile(
    r'(?P<indent>[ \t]*)fail\(\s*'
    r'(?P<quote>["\'])'
    r'Existing PrayerGroup constructor requires groupType\.'
    r'.*?'
    r'(?P=quote)\s*'
    r'\)',
    re.DOTALL,
)

failure_match = failure_pattern.search(
    source
)

if not failure_match:
    fail(
        "Could not safely identify the fail(...) "
        "statement containing the groupType guard."
    )

ok("groupType fail(...) statement located")

old_failure = failure_match.group(0)

print()
print("Current failure statement:")
print("-" * 64)
print(old_failure)
print("-" * 64)
print()


# ============================================================
# 7. DISCOVER SURROUNDING IF
# ============================================================

info("Checking surrounding constructor capability guard")

before_failure = source[
    max(0, failure_match.start() - 500):
    failure_match.start()
]

if "group_type" not in before_failure.lower():
    info(
        "No lowercase group_type token immediately "
        "before failure; checking broader context"
    )

    broader = source[
        max(0, failure_match.start() - 1500):
        failure_match.start()
    ]

    if (
        "group type" not in broader.lower()
        and "group_type" not in broader.lower()
        and "grouptype" not in broader.lower()
    ):
        fail(
            "Could not verify that this failure belongs "
            "to the constructor group-type guard."
        )

ok("Failure belongs to group-type decision path")


# ============================================================
# 8. PATCH FAILURE INTO POLICY VALIDATION
# ============================================================

heading("5. PATCH GENERATOR POLICY")

indent = failure_match.group("indent")

replacement = (
    f'{indent}configured_group_type = '
    f'config.get("create_group_type")\n'
    f'\n'
    f'{indent}if configured_group_type != "{group_type}":\n'
    f'{indent}    fail(\n'
    f'{indent}        "PrayerGroup constructor requires groupType. "\n'
    f'{indent}        "Set create_group_type to {group_type}."\n'
    f'{indent}    )\n'
    f'\n'
    f'{indent}ok(\n'
    f'{indent}    "Create PrayerGroup policy: "\n'
    f'{indent}    "{group_type}"\n'
    f'{indent})'
)

updated_source = (
    source[:failure_match.start()]
    + replacement
    + source[failure_match.end():]
)

ok("Failure guard replaced with explicit policy validation")


# ============================================================
# 9. FIND PRAYERGROUP.New GENERATION
# ============================================================

heading("6. CONSTRUCTOR GENERATION DISCOVERY")

info("Looking for generated prayergroup.New call")

# The generator may contain the Go source in a Python
# multiline string, so inspect all occurrences.

new_occurrences = [
    match.start()
    for match in re.finditer(
        r'prayergroup\.New\s*\(',
        updated_source,
    )
]

if not new_occurrences:
    fail(
        "No prayergroup.New(...) generation was found. "
        "The generator knows the constructor requires a type "
        "but does not appear to generate the constructor call."
    )

ok(
    f"Found {len(new_occurrences)} "
    f"prayergroup.New occurrence(s)"
)


for number, position in enumerate(
    new_occurrences,
    start=1,
):
    section = updated_source[
        max(0, position - 250):
        min(len(updated_source), position + 700)
    ]

    print()
    print(
        f"prayergroup.New occurrence #{number}:"
    )
    print("-" * 64)
    print(section)
    print("-" * 64)


# ============================================================
# 10. PATCH CONSTRUCTOR ARGUMENT
# ============================================================

heading("7. ADD TypeRegular TO GENERATED CONSTRUCTOR")

# We target calls that currently contain the known creation
# arguments and do not already contain TypeRegular.

constructor_pattern = re.compile(
    r'prayergroup\.New\(\s*'
    r'(?P<id>[^,\n]+),\s*'
    r'(?P<name>[^,\n]+),\s*'
    r'(?P<description>[^,\n]+),?\s*'
    r'\)',
    re.MULTILINE,
)

constructor_matches = list(
    constructor_pattern.finditer(
        updated_source
    )
)

if not constructor_matches:
    # It may already have been updated or use a different
    # alias. Check before refusing.
    if (
        "prayergroup.TypeRegular"
        in updated_source
    ):
        ok(
            "Generator already contains "
            "prayergroup.TypeRegular"
        )

    else:
        fail(
            "Could not safely identify a 3-argument "
            "generated prayergroup.New call. "
            "Refusing to guess its formatting."
        )

else:
    info(
        f"Found {len(constructor_matches)} "
        f"3-argument constructor candidate(s)"
    )

    # Work backwards so offsets remain valid.
    for match in reversed(
        constructor_matches
    ):
        original = match.group(0)

        info("Candidate constructor:")
        print()
        print(original)
        print()

        replacement_constructor = (
            "prayergroup.New(\n"
            f"\t\t{match.group('id').strip()},\n"
            f"\t\t{match.group('name').strip()},\n"
            f"\t\t{match.group('description').strip()},\n"
            "\t\tprayergroup.TypeRegular,\n"
            "\t)"
        )

        updated_source = (
            updated_source[:match.start()]
            + replacement_constructor
            + updated_source[match.end():]
        )

        ok(
            "Added prayergroup.TypeRegular "
            "to constructor generation"
        )


# ============================================================
# 11. VERIFY POLICY DOES NOT EXPOSE HTTP TYPE
# ============================================================

heading("8. HTTP CONTRACT SAFETY")

info(
    "Checking generator does not introduce "
    "client-controlled group type"
)

dangerous_patterns = [
    r'json:"type"',
    r'GroupType\s+string\s+`json:',
    r'Type\s+string\s+`json:"type"',
]

for pattern in dangerous_patterns:
    if re.search(
        pattern,
        updated_source,
    ):
        fail(
            "Generator appears to expose prayer-group "
            "type through the HTTP request. "
            "Creation type must remain server-controlled."
        )

ok(
    "No client-controlled prayer-group type "
    "request field detected"
)


# ============================================================
# 12. PRE-WRITE VALIDATION
# ============================================================

heading("9. PRE-WRITE VALIDATION")

checks = {
    "config lookup":
        'config.get("create_group_type")',

    "TypeRegular policy":
        group_type,

    "constructor type":
        "prayergroup.TypeRegular",
}

for name, token in checks.items():
    info(f"Checking: {name}")

    if token not in updated_source:
        fail(
            f"Generator validation failed: {name}"
        )

    ok(f"Validated: {name}")


if needle in updated_source:
    fail(
        "Original hard-stop message still exists "
        "after patch"
    )

ok("Original hard-stop removed")


# ============================================================
# 13. BACKUPS
# ============================================================

heading("10. BACKUPS")

timestamp = datetime.now().strftime(
    "%Y%m%d_%H%M%S"
)

backup_root = (
    BASE
    / ".generator-backups"
    / timestamp
    / "patch-create-prayer-group-generator"
)

generator_backup = (
    backup_root
    / generator_path.name
)

config_backup = (
    backup_root
    / generator_config_path.name
)


info(f"Backing up generator: {generator_backup}")

generator_backup.parent.mkdir(
    parents=True,
    exist_ok=True,
)

shutil.copy2(
    generator_path,
    generator_backup,
)

ok("Generator backup created")


info(f"Backing up config: {config_backup}")

shutil.copy2(
    generator_config_path,
    config_backup,
)

ok("Config backup created")


# ============================================================
# 14. WRITE GENERATOR
# ============================================================

heading("11. WRITE GENERATOR")

info("Writing patched generator")

generator_path.write_text(
    updated_source,
    encoding="utf-8",
)

ok("Generator written")


info("Reading generator back")

written_generator = generator_path.read_text(
    encoding="utf-8",
)

if written_generator != updated_source:
    fail("Generator write verification failed")

ok("Generator write verified")


# ============================================================
# 15. WRITE CONFIG
# ============================================================

heading("12. WRITE CONFIG")

new_config_text = json.dumps(
    create_config,
    indent=2,
)

new_config_text += "\n"

info("Writing generator config")

generator_config_path.write_text(
    new_config_text,
    encoding="utf-8",
)

ok("Generator config written")


info("Reading config back")

try:
    written_config = json.loads(
        generator_config_path.read_text(
            encoding="utf-8",
        )
    )
except Exception as error:
    fail(
        f"Written config is invalid JSON: {error}"
    )

if (
    written_config.get("create_group_type")
    != group_type
):
    fail(
        "create_group_type missing after write"
    )

ok(
    f"create_group_type={group_type} verified"
)


# ============================================================
# 16. FINAL VERIFICATION
# ============================================================

heading("13. FINAL VERIFICATION")

final_generator = generator_path.read_text(
    encoding="utf-8",
)

info("Checking TypeRegular constructor policy")

if (
    "prayergroup.TypeRegular"
    not in final_generator
):
    fail(
        "TypeRegular constructor policy missing"
    )

ok("TypeRegular constructor policy verified")


info("Checking old hard stop")

if needle in final_generator:
    fail(
        "Old groupType hard stop remains"
    )

ok("Old hard stop removed")


# ============================================================
# COMPLETE
# ============================================================

heading("GENERATOR PATCH COMPLETE")

print("Creation policy:")
print()
print("    POST /api/prayer-groups")
print("            ↓")
print("    name + description")
print("            ↓")
print("    server selects TypeRegular")
print("            ↓")
print("    prayergroup.New(..., TypeRegular)")
print()

print("Protected invariant:")
print()
print(
    "    TypeVisitor is NOT client-creatable"
)
print()

print("Backups:")
print()
print(f"    {backup_root}")
print()

print("Next:")
print()
print(
    "    python .\\create_prayer_group_api.py"
)
print()