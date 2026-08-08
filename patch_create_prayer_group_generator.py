import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


BASE = Path(__file__).resolve().parent

PATCH_CONFIG_PATH = (
    BASE / "patch_create_prayer_group_generator.json"
)


def heading(text):
    print()
    print("=" * 64)
    print(text)
    print("=" * 64)
    print()


def info(text):
    print(f"[INFO] {text}")


def ok(text):
    print(f"[ OK ] {text}")


def fail(text):
    print()
    print(f"[FAIL] {text}")
    print()
    print("No files were changed.")
    sys.exit(1)


# ============================================================
# START
# ============================================================

heading("PATCH CREATE PRAYER GROUP GENERATOR")


# ============================================================
# 1. PATCH CONFIG
# ============================================================

heading("1. PATCH CONFIG")

info(f"Config: {PATCH_CONFIG_PATH}")

if not PATCH_CONFIG_PATH.exists():
    fail("Patch config does not exist")

ok("Patch config exists")


try:
    patch_config = json.loads(
        PATCH_CONFIG_PATH.read_text(
            encoding="utf-8"
        )
    )
except Exception as error:
    fail(
        f"Could not parse patch config: {error}"
    )

ok("Patch config loaded")


required_patch_keys = [
    "generator",
    "generator_config",
    "expected_policy",
]

for key in required_patch_keys:
    info(f"Checking patch config key: {key}")

    if key not in patch_config:
        fail(
            f"Missing patch config key: {key}"
        )

    ok(f"Found: {key}")


expected_policy = patch_config[
    "expected_policy"
]

expected_group_type = expected_policy[
    "group_type"
]

expected_source = expected_policy[
    "group_type_source"
]

expected_allow_client = expected_policy[
    "allow_client_group_type"
]


info(
    f"Expected group type: {expected_group_type}"
)

info(
    f"Expected source: {expected_source}"
)

info(
    "Expected client selectable: "
    f"{expected_allow_client}"
)


# ============================================================
# 2. TARGET FILES
# ============================================================

heading("2. TARGET FILES")

generator_path = (
    BASE / patch_config["generator"]
)

generator_config_path = (
    BASE / patch_config["generator_config"]
)


info(f"Generator: {generator_path}")

if not generator_path.exists():
    fail("Generator does not exist")

ok("Generator exists")


info(
    f"Generator config: {generator_config_path}"
)

if not generator_config_path.exists():
    fail("Generator config does not exist")

ok("Generator config exists")


# ============================================================
# 3. CREATION CONFIG
# ============================================================

heading("3. CREATION POLICY")

try:
    create_config = json.loads(
        generator_config_path.read_text(
            encoding="utf-8"
        )
    )
except Exception as error:
    fail(
        f"Could not parse generator config: {error}"
    )

ok("Generator config loaded")


creation = create_config.get("creation")

if not isinstance(creation, dict):
    fail(
        'Missing "creation" policy object'
    )

ok("creation policy found")


actual_group_type = creation.get(
    "group_type"
)

actual_source = creation.get(
    "group_type_source"
)

actual_allow_client = creation.get(
    "allow_client_group_type"
)


info(
    f"group_type = {actual_group_type}"
)

if actual_group_type != expected_group_type:
    fail(
        "group_type does not match expected policy"
    )

ok(
    f"group_type={expected_group_type}"
)


info(
    f"group_type_source = {actual_source}"
)

if actual_source != expected_source:
    fail(
        "group_type_source does not match "
        "expected policy"
    )

ok(
    f"group_type_source={expected_source}"
)


info(
    "allow_client_group_type = "
    f"{actual_allow_client}"
)

if actual_allow_client is not expected_allow_client:
    fail(
        "allow_client_group_type does not match "
        "expected policy"
    )

ok(
    "allow_client_group_type=false"
)


# ============================================================
# 4. READ GENERATOR
# ============================================================

heading("4. GENERATOR DISCOVERY")

source = generator_path.read_text(
    encoding="utf-8"
)

ok(
    f"Generator read ({len(source)} characters)"
)


# ============================================================
# 5. FIND generate_application
# ============================================================

heading("5. GENERATE_APPLICATION DISCOVERY")

function_match = re.search(
    r"^def\s+generate_application\s*\("
    r"\s*constructor\s*,"
    r"\s*id_method\s*,?\s*"
    r"\)\s*:",
    source,
    re.MULTILINE,
)

if not function_match:
    fail(
        "generate_application("
        "constructor, id_method) not found"
    )

ok("generate_application found")


function_start = function_match.start()


next_function = re.search(
    r"^def\s+[A-Za-z_][A-Za-z0-9_]*\s*\(",
    source[function_match.end():],
    re.MULTILINE,
)

if next_function:
    function_end = (
        function_match.end()
        + next_function.start()
    )
else:
    function_end = len(source)


function_source = source[
    function_start:function_end
]


info(
    f"generate_application size: "
    f"{len(function_source)} characters"
)


# ============================================================
# 6. FIND EXISTING TYPE GUARD
# ============================================================

heading("6. GROUP TYPE GUARD")

guard_text = '''    if constructor["type"]:
        fail(
            "Existing PrayerGroup constructor requires groupType. "
            "The requested create API currently only defines name and "
            "description. Add an explicit API default/type decision "
            "before generating production code."
        )
'''


if guard_text in function_source:
    ok(
        'Exact constructor["type"] guard found'
    )

else:
    # Check whether it has already been patched.

    already_patched_tokens = [
        'creation_policy = config.get("creation", {})',
        'configured_group_type = creation_policy.get("group_type")',
        "prayergroup.TypeRegular",
    ]

    already_patched = all(
        token in function_source
        for token in already_patched_tokens
    )

    if already_patched:
        heading("ALREADY PATCHED")

        ok(
            "Generator already contains creation policy"
        )

        ok(
            "Generator already contains TypeRegular"
        )

        print()
        print("No changes required.")
        print()
        sys.exit(0)

    fail(
        "Exact original constructor type guard was not "
        "found and generator does not appear already "
        "patched. Refusing unsafe modification."
    )


# ============================================================
# 7. CREATE NEW GUARD
# ============================================================

heading("7. GENERATING POLICY GUARD")

new_guard = f'''    if constructor["type"]:
        creation_policy = config.get("creation", {{}})

        configured_group_type = creation_policy.get(
            "group_type"
        )

        configured_group_type_source = creation_policy.get(
            "group_type_source"
        )

        allow_client_group_type = creation_policy.get(
            "allow_client_group_type"
        )

        if configured_group_type != "{expected_group_type}":
            fail(
                "PrayerGroup creation policy must use "
                "{expected_group_type}."
            )

        if configured_group_type_source != "{expected_source}":
            fail(
                "PrayerGroup type must be "
                "server-controlled."
            )

        if allow_client_group_type is not False:
            fail(
                "PrayerGroup type must not be "
                "client-controlled."
            )

        ok(
            "PrayerGroup creation policy verified: "
            "{expected_group_type}"
        )
'''


patched_function = function_source.replace(
    guard_text,
    new_guard,
    1,
)


if patched_function == function_source:
    fail("Type guard replacement failed")

ok("Type guard replaced in memory")


# ============================================================
# 8. LOCATE GENERATED CONSTRUCTOR
# ============================================================

heading("8. GENERATED CONSTRUCTOR")

constructor_start = patched_function.find(
    "prayergroup.New("
)

if constructor_start == -1:
    fail(
        "prayergroup.New(...) not found inside "
        "generate_application"
    )

ok("prayergroup.New(...) found")


open_paren = patched_function.find(
    "(",
    constructor_start,
)

depth = 1
position = open_paren + 1


while (
    position < len(patched_function)
    and depth > 0
):
    char = patched_function[position]

    if char == "(":
        depth += 1

    elif char == ")":
        depth -= 1

    position += 1


if depth != 0:
    fail(
        "Could not determine prayergroup.New "
        "closing parenthesis"
    )


constructor_call = patched_function[
    constructor_start:position
]


print()
print("Current generated constructor:")
print("-" * 64)
print(constructor_call)
print("-" * 64)
print()


# ============================================================
# 9. VERIFY CONSTRUCTOR SHAPE
# ============================================================

heading("9. CONSTRUCTOR CONTRACT")

required_constructor_tokens = [
    "groupID",
    "command.Name",
    "command.Description",
]

for token in required_constructor_tokens:
    info(f"Checking constructor token: {token}")

    if token not in constructor_call:
        fail(
            f"Constructor does not contain {token}"
        )

    ok(f"Found: {token}")


type_expression = (
    f"prayergroup.{expected_group_type}"
)


# ============================================================
# 10. ADD TYPE
# ============================================================

heading("10. ADD GROUP TYPE")

if type_expression in constructor_call:
    ok(
        f"{type_expression} already present"
    )

else:
    info(
        f"Adding {type_expression}"
    )

    closing_index = constructor_call.rfind(
        ")"
    )

    before_close = constructor_call[
        :closing_index
    ].rstrip()

    after_close = constructor_call[
        closing_index:
    ]


    if not before_close.endswith(","):
        before_close += ","


    new_constructor_call = (
        before_close
        + "\n\t\t"
        + type_expression
        + ",\n\t"
        + after_close
    )


    patched_function = (
        patched_function[
            :constructor_start
        ]
        + new_constructor_call
        + patched_function[position:]
    )

    ok(
        f"Added {type_expression}"
    )


# ============================================================
# 11. REBUILD GENERATOR
# ============================================================

heading("11. REBUILD GENERATOR")

updated_source = (
    source[:function_start]
    + patched_function
    + source[function_end:]
)

ok("Generator rebuilt in memory")


# ============================================================
# 12. SAFETY VALIDATION
# ============================================================

heading("12. PRE-WRITE VALIDATION")

required_tokens = [
    'config.get("creation", {})',
    'creation_policy.get(',
    '"group_type"',
    '"group_type_source"',
    '"allow_client_group_type"',
    type_expression,
]


for token in required_tokens:
    info(f"Checking: {token}")

    if token not in updated_source:
        fail(
            f"Missing required token: {token}"
        )

    ok(f"Found: {token}")


obsolete_message = (
    "Existing PrayerGroup constructor "
    "requires groupType."
)

info("Checking obsolete hard stop")

if obsolete_message in updated_source:
    fail(
        "Old groupType hard stop remains"
    )

ok("Old groupType hard stop removed")


# ============================================================
# 13. HTTP CONTRACT SAFETY
# ============================================================

heading("13. HTTP CONTRACT SAFETY")

dangerous_tokens = [
    'json:"type"',
    'json:\\"type\\"',
]


for token in dangerous_tokens:
    info(f"Checking absence: {token}")

    if token in updated_source:
        fail(
            "Client-controlled group type detected"
        )

    ok(f"Not present: {token}")


# ============================================================
# 14. BACKUP
# ============================================================

heading("14. BACKUP")

timestamp = datetime.now().strftime(
    "%Y%m%d_%H%M%S"
)

backup_root = (
    BASE
    / ".generator-backups"
    / timestamp
    / "patch-create-prayer-group-generator"
)

backup_path = (
    backup_root
    / generator_path.name
)


info(f"Backup: {backup_path}")

backup_path.parent.mkdir(
    parents=True,
    exist_ok=True,
)

shutil.copy2(
    generator_path,
    backup_path,
)

ok("Backup created")


# ============================================================
# 15. WRITE
# ============================================================

heading("15. WRITE")

info("Writing patched generator")

generator_path.write_text(
    updated_source,
    encoding="utf-8",
)

ok("Generator written")


# ============================================================
# 16. READ-BACK VERIFICATION
# ============================================================

heading("16. POST-WRITE VERIFICATION")

written = generator_path.read_text(
    encoding="utf-8"
)


post_checks = [
    (
        'config.get("creation", {})',
        "creation policy lookup",
    ),
    (
        type_expression,
        "TypeRegular constructor argument",
    ),
]


for token, description in post_checks:
    info(f"Checking: {description}")

    if token not in written:
        fail(
            f"Post-write verification failed: "
            f"{description}"
        )

    ok(f"Verified: {description}")


if obsolete_message in written:
    fail(
        "Old groupType hard stop still exists "
        "after write"
    )

ok("Old hard stop absent")


# ============================================================
# COMPLETE
# ============================================================

heading("PATCH COMPLETE")

print("Creation policy:")
print()
print(
    f"  Group type        : {expected_group_type}"
)
print(
    f"  Source            : {expected_source}"
)
print(
    f"  Client selectable : {expected_allow_client}"
)
print()

print("Generator will emit:")
print()
print(
    f"  {type_expression}"
)
print()

print(f"Backup: {backup_path}")
print()

print("Next:")
print()
print(
    "  python .\\create_prayer_group_api.py"
)
print()