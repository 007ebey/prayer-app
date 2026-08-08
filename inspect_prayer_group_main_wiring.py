from pathlib import Path
import re
import sys


BASE = Path(__file__).resolve().parent
PROJECT = BASE / "prayer-api"

MAIN = PROJECT / "cmd/api/main.go"
GENERATOR = BASE / "create_prayer_group_api.py"


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
    print(f"[FAIL] {text}")
    sys.exit(1)


heading("INSPECT PRAYER GROUP MAIN WIRING")


# ============================================================
# 1. FILE CHECKS
# ============================================================

info(f"main.go: {MAIN}")

if not MAIN.exists():
    fail("main.go does not exist")

ok("main.go exists")


info(f"Generator: {GENERATOR}")

if not GENERATOR.exists():
    fail("Generator does not exist")

ok("Generator exists")


main = MAIN.read_text(encoding="utf-8")
generator = GENERATOR.read_text(encoding="utf-8")

ok(f"main.go read ({len(main)} chars)")
ok(f"Generator read ({len(generator)} chars)")


# ============================================================
# 2. CURRENT MAIN
# ============================================================

heading("1. CURRENT MAIN.GO")

print(main)


# ============================================================
# 3. HANDLER CONSTRUCTION
# ============================================================

heading("2. HANDLER CONSTRUCTION")

handler_pattern = re.compile(
    r'(?P<variable>[A-Za-z_][A-Za-z0-9_]*)\s*:=\s*'
    r'httpapi\.New[A-Za-z0-9_]*Handler\s*\('
)

handlers = list(handler_pattern.finditer(main))

if not handlers:
    fail("No HTTP handler constructors discovered")


for match in handlers:
    variable = match.group("variable")

    start = match.start()

    # Print useful surrounding context.
    snippet = main[
        max(0, start - 100):
        min(len(main), start + 500)
    ]

    print()
    print(f"Handler variable: {variable}")
    print("-" * 64)
    print(snippet)
    print("-" * 64)


# ============================================================
# 4. AUTH HANDLER EXACT BLOCK
# ============================================================

heading("3. AUTH HANDLER BLOCK")

auth_pos = main.find(
    "authHandler := httpapi.NewAuthHandler("
)

if auth_pos == -1:
    fail(
        "authHandler := httpapi.NewAuthHandler(...) "
        "not found"
    )

ok("authHandler constructor found")


# Find balanced constructor call.
open_pos = main.find("(", auth_pos)

depth = 1
pos = open_pos + 1

while pos < len(main) and depth > 0:
    if main[pos] == "(":
        depth += 1
    elif main[pos] == ")":
        depth -= 1

    pos += 1


if depth != 0:
    fail("Could not determine authHandler block")


# Include rest of assignment line.
line_end = main.find("\n", pos)

if line_end == -1:
    line_end = len(main)


auth_block = main[
    auth_pos:line_end
]


print()
print("Exact authHandler block:")
print("-" * 64)
print(auth_block)
print("-" * 64)


# ============================================================
# 5. SERVICE CONSTRUCTION
# ============================================================

heading("4. APPLICATION SERVICES")

service_pattern = re.compile(
    r'(?P<variable>[A-Za-z_][A-Za-z0-9_]*)\s*:=\s*'
    r'[A-Za-z_][A-Za-z0-9_]*\.NewService\s*\('
)

services = list(service_pattern.finditer(main))

if not services:
    info("No NewService assignments discovered")
else:
    for match in services:
        print(
            f"[ OK ] {match.group('variable')}"
        )


# ============================================================
# 6. ROUTER CALL
# ============================================================

heading("5. ROUTER CALL")

router_pos = main.find(
    "httpapi.NewRouter("
)

if router_pos == -1:
    fail("httpapi.NewRouter(...) call not found")

ok("NewRouter call found")


open_pos = main.find("(", router_pos)

depth = 1
pos = open_pos + 1

while pos < len(main) and depth > 0:
    if main[pos] == "(":
        depth += 1
    elif main[pos] == ")":
        depth -= 1

    pos += 1


if depth != 0:
    fail("Could not determine NewRouter call")


router_call = main[
    router_pos:pos
]


print()
print("Current router call:")
print("-" * 64)
print(router_call)
print("-" * 64)


# ============================================================
# 7. GENERATOR FAILURE LOCATION
# ============================================================

heading("6. GENERATOR MAIN PATCH LOGIC")

needle = (
    "Expected authHandler wiring not found"
)

failure_pos = generator.find(needle)

if failure_pos == -1:
    fail(
        "Could not locate generator failure message"
    )

ok("Generator failure location found")


snippet = generator[
    max(0, failure_pos - 2500):
    min(len(generator), failure_pos + 1500)
]


print()
print("Generator logic around failure:")
print("-" * 64)
print(snippet)
print("-" * 64)


# ============================================================
# 8. LIKELY MISMATCHES
# ============================================================

heading("7. CONTRACT COMPARISON")

checks = [
    (
        "authHandler := httpapi.NewAuthHandler(",
        "auth handler constructor",
    ),
    (
        "userHandler := httpapi.NewUserHandler(",
        "user handler constructor",
    ),
    (
        "userRoleHandler := httpapi.NewUserRoleHandler(",
        "user-role handler constructor",
    ),
    (
        "router := httpapi.NewRouter(",
        "router assignment",
    ),
]


for token, description in checks:
    info(f"Checking {description}")

    if token in main:
        ok(f"Found: {token}")
    else:
        print(f"[WARN] Missing: {token}")


# ============================================================
# COMPLETE
# ============================================================

heading("INSPECTION COMPLETE")

print(
    "No project files were modified."
)
print()