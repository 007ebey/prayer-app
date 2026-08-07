from pathlib import Path
import sys


BASE = Path(__file__).resolve().parent

path = BASE / "patch_create_prayer_group_generator.py"

print()
print("=" * 64)
print("VERIFY PRAYER GROUP GENERATOR PATCHER")
print("=" * 64)
print()

print(f"[INFO] File: {path}")

if not path.exists():
    print("[FAIL] File does not exist")
    sys.exit(1)

print("[ OK ] File exists")

source = path.read_text(encoding="utf-8")

print(f"[ OK ] Read {len(source)} characters")
print()


checks_required = [
    (
        'patch_config["expected_policy"]',
        "new expected_policy config contract",
    ),
    (
        'expected_policy["group_type"]',
        "group_type policy",
    ),
    (
        'expected_policy["group_type_source"]',
        "group_type_source policy",
    ),
    (
        'expected_policy["allow_client_group_type"]',
        "allow_client_group_type policy",
    ),
    (
        'if constructor["type"]:',
        "actual generator constructor type guard",
    ),
    (
        'config.get("creation", {})',
        "creation configuration lookup",
    ),
]


checks_forbidden = [
    (
        'patch_config["create_group_type"]',
        "obsolete patch-config create_group_type",
    ),
    (
        'create_config.get("create_group_type")',
        "obsolete generator-config create_group_type lookup",
    ),
    (
        'create_config["create_group_type"]',
        "obsolete generator-config create_group_type assignment",
    ),
    (
        '] = group_type',
        "obsolete group_type assignment",
    ),
    (
        '== group_type',
        "obsolete group_type comparison",
    ),
]


print("REQUIRED CONTRACT")
print("-" * 64)

failed = False

for token, description in checks_required:
    print(f"[INFO] Checking: {description}")

    if token not in source:
        print(f"[FAIL] Missing: {token}")
        failed = True
    else:
        print(f"[ OK ] Found: {token}")


print()
print("OBSOLETE CONTRACT")
print("-" * 64)

for token, description in checks_forbidden:
    print(f"[INFO] Checking absence: {description}")

    if token in source:
        print(f"[FAIL] Obsolete code still exists:")
        print(f"       {token}")
        failed = True
    else:
        print(f"[ OK ] Not present: {token}")


print()
print("CONFIG VERIFICATION")
print("-" * 64)

config_path = BASE / "patch_create_prayer_group_generator.json"

print(f"[INFO] Config: {config_path}")

if not config_path.exists():
    print("[FAIL] Config missing")
    failed = True

else:
    import json

    try:
        config = json.loads(
            config_path.read_text(encoding="utf-8")
        )

        print("[ OK ] Config is valid JSON")

        policy = config.get("expected_policy")

        if not isinstance(policy, dict):
            print("[FAIL] expected_policy object missing")
            failed = True

        else:
            expected = {
                "group_type": "TypeRegular",
                "group_type_source": "server",
                "allow_client_group_type": False,
            }

            for key, expected_value in expected.items():
                actual = policy.get(key)

                print(
                    f"[INFO] {key}: "
                    f"expected={expected_value!r}, "
                    f"actual={actual!r}"
                )

                if actual != expected_value:
                    print(f"[FAIL] {key} mismatch")
                    failed = True
                else:
                    print(f"[ OK ] {key}")

    except Exception as error:
        print(f"[FAIL] Invalid config: {error}")
        failed = True


print()
print("=" * 64)

if failed:
    print("VERIFICATION FAILED")
    print("=" * 64)
    print()
    print(
        "patch_create_prayer_group_generator.py "
        "is still using the old contract."
    )
    print()
    sys.exit(1)

print("VERIFICATION PASSED")
print("=" * 64)
print()
print("The patcher is using the new creation-policy contract.")
print()

print()
print("VARIABLE CONTRACT")
print("-" * 64)

if re.search(r"\bgroup_type\b", source):
    print("[INFO] Bare group_type references discovered")

    lines = source.splitlines()

    bad_lines = []

    for number, line in enumerate(lines, start=1):
        if re.search(r"\bgroup_type\b", line):
            if (
                "expected_group_type" not in line
                and '"group_type"' not in line
                and "'group_type'" not in line
            ):
                bad_lines.append(
                    (number, line.strip())
                )

    if bad_lines:
        print("[FAIL] Suspicious stale group_type references:")

        for number, line in bad_lines:
            print(
                f"       line {number}: {line}"
            )

        failed = True
    else:
        print("[ OK ] No stale bare group_type variables")
else:
    print("[ OK ] No bare group_type references")