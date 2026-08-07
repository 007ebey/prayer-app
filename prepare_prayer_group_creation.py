import json
import re
import shutil
import sys

from datetime import datetime
from pathlib import Path


BASE = Path(__file__).resolve().parent
CONFIG_FILE = BASE / "prepare_prayer_group_creation.json"


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


def backup(path, relative, backup_root):
    target = backup_root / relative

    info(f"Creating backup: {target}")

    target.parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    shutil.copy2(path, target)

    ok("Backup created")


def write_verified(
    path,
    relative,
    old,
    new,
    backup_root,
):
    if old == new:
        ok(f"No changes required: {relative}")
        return

    backup(
        path,
        relative,
        backup_root,
    )

    info(f"Writing: {relative}")

    path.write_text(
        new,
        encoding="utf-8",
    )

    ok("File written")

    written = path.read_text(
        encoding="utf-8",
    )

    if written != new:
        fail(
            f"Write verification failed: {relative}"
        )

    ok("Write verified")


# ============================================================
# CONFIG
# ============================================================

heading("PREPARE PRAYER GROUP CREATION")

info(f"Config: {CONFIG_FILE}")

if not CONFIG_FILE.exists():
    fail("Configuration file missing")

ok("Configuration file exists")

try:
    config = json.loads(
        CONFIG_FILE.read_text(
            encoding="utf-8",
        )
    )
except Exception as error:
    fail(f"Could not load config: {error}")

ok("Configuration loaded")


# ============================================================
# PROJECT
# ============================================================

project = BASE / config["project"]

info(f"Project: {project}")

if not project.exists():
    fail("Project directory missing")

ok("Project directory exists")

if not (project / "go.mod").exists():
    fail("go.mod missing")

ok("Go project confirmed")

timestamp = datetime.now().strftime(
    "%Y%m%d_%H%M%S"
)

backup_root = (
    project
    / ".generator-backups"
    / timestamp
    / "prepare-prayer-group-creation"
)

info(f"Backup root: {backup_root}")


# ============================================================
# PRAYERGROUP DOMAIN
# ============================================================

heading("1. PRAYER GROUP ID CONTRACT")

domain_dir = project / config["domain_dir"]

if not domain_dir.exists():
    fail("PrayerGroup domain directory missing")

ok("PrayerGroup domain directory exists")

domain_source = ""

for path in sorted(domain_dir.glob("*.go")):
    if path.name.endswith("_test.go"):
        continue

    relative = path.relative_to(project)

    info(f"Reading: {relative}")

    content = path.read_text(
        encoding="utf-8",
    )

    domain_source += "\n" + content

    ok(f"Read: {relative}")


info("Checking prayergroup.ID")

if not re.search(
    r"type\s+ID\s+string",
    domain_source,
):
    fail(
        "Expected type ID string was not found"
    )

ok("prayergroup.ID is string based")


# ============================================================
# ID GENERATOR DISCOVERY
# ============================================================

heading("2. ID GENERATOR DISCOVERY")

id_relative = Path(
    config["id_generator"]
)

id_path = project / id_relative

info(f"ID generator: {id_relative}")

if not id_path.exists():
    fail("ID generator file missing")

ok("ID generator exists")

id_source = id_path.read_text(
    encoding="utf-8",
)

ok(
    f"ID generator read "
    f"({len(id_source)} characters)"
)

print()
print("Current id_generator.go:")
print("-" * 64)
print(id_source.rstrip())
print("-" * 64)
print()


# ============================================================
# DISCOVER NewUserID
# ============================================================

info("Discovering NewUserID")

method_match = re.search(
    r"func\s+\("
    r"(?P<receiver>[A-Za-z0-9_]+)\s+\*IDGenerator"
    r"\)\s+NewUserID\s*\(\s*\)"
    r"\s+user\.ID\s*\{",
    id_source,
)

if not method_match:
    fail(
        "Could not discover "
        "IDGenerator.NewUserID"
    )

receiver = method_match.group("receiver")

ok(f"NewUserID receiver: {receiver}")

body_start = method_match.end()

depth = 1
position = body_start

while (
    position < len(id_source)
    and depth > 0
):
    character = id_source[position]

    if character == "{":
        depth += 1

    elif character == "}":
        depth -= 1

    position += 1

if depth != 0:
    fail(
        "Could not determine NewUserID boundary"
    )

user_body = id_source[
    body_start:position - 1
]

print()
print("NewUserID implementation:")
print("-" * 64)
print(user_body.strip())
print("-" * 64)
print()


# ============================================================
# DISCOVER ACTUAL STRATEGY
# ============================================================

heading("3. ID STRATEGY VERIFICATION")

info("Checking atomic counter strategy")

atomic_match = re.search(
    r"(?P<variable>[A-Za-z0-9_]+)\s*:=\s*"
    r"atomic\.AddUint64\s*\(\s*"
    rf"&{re.escape(receiver)}\."
    r"(?P<field>[A-Za-z0-9_]+)\s*,\s*1\s*\)",
    user_body,
)

if not atomic_match:
    fail(
        "NewUserID does not use the expected "
        "atomic.AddUint64 counter strategy. "
        "Refusing to invent another strategy."
    )

counter_variable = atomic_match.group(
    "variable"
)

counter_field = atomic_match.group(
    "field"
)

ok(
    f"Atomic counter discovered: "
    f"{receiver}.{counter_field}"
)

ok(
    f"Generated value variable: "
    f"{counter_variable}"
)


info("Checking fmt.Sprintf strategy")

format_match = re.search(
    r'user\.ID\s*\(\s*'
    r'fmt\.Sprintf\s*\(\s*'
    r'"(?P<format>[^"]+)"\s*,\s*'
    rf'{re.escape(counter_variable)}'
    r'\s*\)\s*\)',
    user_body,
)

if not format_match:
    fail(
        "Could not verify fmt.Sprintf ID "
        "construction in NewUserID"
    )

user_format = format_match.group(
    "format"
)

ok(
    f"Existing user ID format: {user_format}"
)


info("Checking required imports")

if '"fmt"' not in id_source:
    fail("fmt import missing")

ok("fmt import exists")

if '"sync/atomic"' not in id_source:
    fail("sync/atomic import missing")

ok("sync/atomic import exists")


# ============================================================
# ADD PRAYERGROUP IMPORT
# ============================================================

heading("4. GENERATE PRAYER GROUP ID METHOD")

method_name = config[
    "required_id_method"
]

prefix = config["id_prefix"]

id_updated = id_source


if re.search(
    rf"func\s+\([^)]*\)\s+"
    rf"{re.escape(method_name)}\s*\(",
    id_source,
):
    ok(f"{method_name} already exists")

else:
    info(
        f"{method_name} does not exist"
    )

    prayergroup_import = (
        '"prayer-api/internal/domain/prayergroup"'
    )

    if prayergroup_import in id_updated:
        ok("prayergroup import already exists")

    else:
        info("Adding prayergroup import")

        import_match = re.search(
            r"import\s*\((?P<body>.*?)\)",
            id_updated,
            re.DOTALL,
        )

        if not import_match:
            fail(
                "Expected parenthesized import "
                "block not found"
            )

        import_body = import_match.group(
            "body"
        )

        new_import_body = (
            import_body.rstrip()
            + '\n\n\t'
            + prayergroup_import
            + "\n"
        )

        id_updated = (
            id_updated[
                :import_match.start("body")
            ]
            + new_import_body
            + id_updated[
                import_match.end("body"):
            ]
        )

        ok("prayergroup import added in memory")


    generated_method = f'''

func (g *IDGenerator) {method_name}() prayergroup.ID {{
\tid := atomic.AddUint64(&g.{counter_field}, 1)
\treturn prayergroup.ID(
\t\tfmt.Sprintf("{prefix}%d", id),
\t)
}}
'''

    print()
    print("Generated method:")
    print("-" * 64)
    print(generated_method.strip())
    print("-" * 64)
    print()

    id_updated = (
        id_updated.rstrip()
        + generated_method
    )

    ok(
        f"{method_name} generated in memory"
    )


# ============================================================
# VALIDATE ID GENERATOR
# ============================================================

heading("5. ID GENERATOR VALIDATION")

checks = {
    "PrayerGroup import":
        '"prayer-api/internal/domain/prayergroup"',

    "PrayerGroup method":
        f"{method_name}() prayergroup.ID",

    "Atomic increment":
        f"atomic.AddUint64(&g.{counter_field}, 1)",

    "PrayerGroup conversion":
        "prayergroup.ID(",

    "PrayerGroup prefix":
        f'fmt.Sprintf("{prefix}%d", id)',
}

for name, token in checks.items():
    info(f"Checking: {name}")

    if token not in id_updated:
        fail(
            f"Validation failed: {name}"
        )

    ok(f"Validated: {name}")


# ============================================================
# REPOSITORY SAVE
# ============================================================

heading("6. REPOSITORY SAVE CONCURRENCY")

repo_relative = Path(
    config["prayer_group_repository"]
)

repo_path = project / repo_relative

info(f"Repository: {repo_relative}")

if not repo_path.exists():
    fail("PrayerGroup repository missing")

ok("PrayerGroup repository exists")

repo_source = repo_path.read_text(
    encoding="utf-8",
)

ok(
    f"Repository read "
    f"({len(repo_source)} characters)"
)


save_match = re.search(
    r"func\s+\("
    r"(?P<receiver>[A-Za-z0-9_]+)\s+"
    r"\*PrayerGroupRepository"
    r"\)\s+Save\s*\("
    r"(?P<args>.*?)"
    r"\)\s+error\s*\{",
    repo_source,
    re.DOTALL,
)

if not save_match:
    fail(
        "PrayerGroupRepository.Save "
        "not found"
    )

repo_receiver = save_match.group(
    "receiver"
)

ok(
    f"Save receiver: {repo_receiver}"
)

save_body_start = save_match.end()

depth = 1
position = save_body_start

while (
    position < len(repo_source)
    and depth > 0
):
    character = repo_source[position]

    if character == "{":
        depth += 1

    elif character == "}":
        depth -= 1

    position += 1

if depth != 0:
    fail(
        "Could not determine Save boundary"
    )

save_body = repo_source[
    save_body_start:position - 1
]

print()
print("Current Save implementation:")
print("-" * 64)
print(save_body.strip())
print("-" * 64)
print()


# ============================================================
# VERIFY STORAGE ASSIGNMENT
# ============================================================

assignment = (
    f"{repo_receiver}.groups[group.ID] = group"
)

info("Checking group map assignment")

if assignment not in save_body:
    fail(
        "Expected group map assignment "
        "was not found"
    )

ok("Group map assignment verified")


# ============================================================
# ADD WRITE LOCK
# ============================================================

lock = (
    f"{repo_receiver}.mu.Lock()"
)

unlock = (
    f"defer {repo_receiver}.mu.Unlock()"
)

repo_updated = repo_source

if (
    lock in save_body
    and unlock in save_body
):
    ok("Save already uses write lock")

else:
    info(
        "Save does not use write lock"
    )

    new_body = (
        "\n\t"
        + lock
        + "\n\t"
        + unlock
        + "\n"
        + save_body
    )

    repo_updated = (
        repo_source[:save_body_start]
        + new_body
        + repo_source[position - 1:]
    )

    ok("Write lock inserted in memory")


# ============================================================
# VALIDATE REPOSITORY
# ============================================================

heading("7. REPOSITORY VALIDATION")

for name, token in {
    "Lock": lock,
    "Unlock": unlock,
    "Assignment": assignment,
}.items():

    info(f"Checking: {name}")

    if token not in repo_updated:
        fail(
            f"Repository validation failed: "
            f"{name}"
        )

    ok(f"Validated: {name}")


# ============================================================
# WRITE
# ============================================================

heading("8. WRITE CHANGES")

write_verified(
    id_path,
    id_relative,
    id_source,
    id_updated,
    backup_root,
)

write_verified(
    repo_path,
    repo_relative,
    repo_source,
    repo_updated,
    backup_root,
)


# ============================================================
# FINAL VERIFICATION
# ============================================================

heading("9. FINAL VERIFICATION")

final_id = id_path.read_text(
    encoding="utf-8",
)

final_repo = repo_path.read_text(
    encoding="utf-8",
)

info(f"Checking {method_name}")

if (
    f"{method_name}() prayergroup.ID"
    not in final_id
):
    fail(
        f"{method_name} missing after write"
    )

ok(f"{method_name} verified")


info("Checking atomic ID generation")

if (
    f"atomic.AddUint64(&g.{counter_field}, 1)"
    not in final_id
):
    fail(
        "Atomic ID generation missing"
    )

ok("Atomic ID generation verified")


info("Checking repository locking")

if (
    lock not in final_repo
    or unlock not in final_repo
):
    fail(
        "Repository write locking missing"
    )

ok("Repository write locking verified")


# ============================================================
# COMPLETE
# ============================================================

heading("PRAYER GROUP CREATION DEPENDENCIES READY")

print("ID strategy:")
print()
print(
    f"    atomic.AddUint64(&g.{counter_field}, 1)"
)
print()

print("PrayerGroup IDs:")
print()
print(
    f'    fmt.Sprintf("{prefix}%d", id)'
)
print()

print("Repository:")
print()
print(
    "    PrayerGroupRepository.Save"
)
print(
    "    protected by sync.RWMutex write lock"
)
print()

print("Next:")
print()
print(
    "    python .\\create_prayer_group_api.py"
)
print()