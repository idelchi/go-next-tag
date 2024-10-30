#!/bin/bash

# BLAZINGLY FAST™️ TEST SUITE FR FR
# By Professor of Bustology, PhD

set -e  # FAIL FAST LIKE MY RELATIONSHIPS

# go install ./cmd/go-next-tag

echo "🔥 INITIATING BLAZINGLY FAST™️ TESTING SEQUENCE 🔥"

# BASIC VERSION BUMPING TESTS (MAJOR-MINOR STYLE) NO CAP
echo "🧪 Testing MAJOR-MINOR format (default behavior) SKRRRAHH!"

[ "$(go-next-tag v1.1)" = "v1.2" ] || (echo "❌ test \`v1.1\` (argument) failed" && exit 1)
[ "$(echo "v2.1" | go-next-tag)" = "v2.2" ] || (echo "❌ test \`v2.1\` (piped) failed" && exit 1)
[ "$(echo "v2.1" | go-next-tag v3.1)" = "v3.2" ] || (echo "❌ test \`v2.1/v3.1\` (piped & argument) failed" && exit 1)
[ "$(go-next-tag --format=majorminor)" = "0.0" ] || (echo "❌ test \`no-input\` failed" && exit 1)

# MAJOR BUMPING TESTS FR FR
echo "🧪 Testing MAJOR bump mode (we going UP UP UP!)"

[ "$(go-next-tag --bump=major v1.1)" = "v2.0" ] || (echo "❌ test (--bump=major) \`v1.1\` (argument) failed" && exit 1)
[ "$(echo "v2.1" | go-next-tag --bump=major)" = "v3.0" ] || (echo "❌ test (--bump=major) \`v2.1\` (piped) failed" && exit 1)
[ "$(echo "v2.1" | go-next-tag --bump=major v3.1)" = "v4.0" ] || (echo "❌ test (--bump=major) \`v2.1/v3.1\` (piped & argument) failed" && exit 1)
[ "$(go-next-tag --format=majorminor --bump=major)" = "0.0" ] || (echo "❌ test (--bump=major) \`no-input\` failed" && exit 1)

# PREFIX TESTING WITH MAJOR-MINOR (GOTTA TEST ALL THEM PREFIXES FR FR)
echo "🧪 Testing them SPICY prefixes with MAJOR-MINOR NO CAP"

prefixes=(
    "v"
    "release-"
    "version-"
    "prefix123with-numbers-"
    "Release-"
    "RELEASE-"
    "rc-"
    "alpha-"
    "beta-"
)

for prefix in "${prefixes[@]}"; do
    # Test minor bump (default)
    [ "$(go-next-tag "${prefix}1.1")" = "${prefix}1.2" ] || (echo "❌ test prefix \`${prefix}\` minor bump failed" && exit 1)
    # Test major bump
    [ "$(go-next-tag --bump=major "${prefix}1.1")" = "${prefix}2.0" ] || (echo "❌ test prefix \`${prefix}\` major bump failed" && exit 1)
done

# BASIC SEMVER TESTING (KEEPING IT SEMANTIC FR FR)
echo "🧪 Testing base SEMVER format (semantic vibes NO CAP)"

# Basic semver without prefix
[ "$(go-next-tag --bump=patch v1.2.3)" = "v1.2.4" ] || (echo "❌ test semver patch bump failed" && exit 1)
[ "$(go-next-tag --bump=minor v1.2.3)" = "v1.3.0" ] || (echo "❌ test semver minor bump failed" && exit 1)
[ "$(go-next-tag --bump=major v1.2.3)" = "v2.0.0" ] || (echo "❌ test semver major bump failed" && exit 1)
[ "$(go-next-tag --bump=major)" = "0.0.0" ] || (echo "❌ test (--bump=major) \`no-input\` failed" && exit 1)

# PREFIX TESTING WITH SEMVER (NOW WE REALLY COOKING!)
echo "🧪 Testing them SPICY prefixes with SEMVER NO CAP"

for prefix in "${prefixes[@]}"; do
    # Test patch bump (default)
    [ "$(go-next-tag --bump=patch "${prefix}1.2.3")" = "${prefix}1.2.4" ] || (echo "❌ test prefix \`${prefix}\` semver patch bump failed" && exit 1)
    # Test minor bump
    [ "$(go-next-tag --bump=minor "${prefix}1.2.3")" = "${prefix}1.3.0" ] || (echo "❌ test prefix \`${prefix}\` semver minor bump failed" && exit 1)
    # Test major bump
    [ "$(go-next-tag --bump=major "${prefix}1.2.3")" = "${prefix}2.0.0" ] || (echo "❌ test prefix \`${prefix}\` semver major bump failed" && exit 1)
done

# COMPLEX SEMVER TESTING (GOTTA TEST THEM PRE-RELEASES AND METADATA FR FR)
echo "🧪 Testing COMPLEX SEMVER cases (we going DEEP NO CAP)"

semver_cases=(
    "v1.2.3-alpha.1"
    "v1.2.3-beta.2+build.123"
    "v1.2.3+build.123"
    "v1.2.3-rc.1+build.123"
    "v1.2.3-alpha.beta.1"
)

# Test each complex case with each prefix (NOW WE REALLY REALLY COOKING!)
for prefix in "${prefixes[@]}"; do
    for case in "${semver_cases[@]}"; do
        [ "$(go-next-tag --bump=patch "${prefix}${base_version}")" = "${prefix}1.2.4" ] || (echo "❌ test prefix \`${prefix}\` with complex semver \`${base_version}\` failed" && exit 1)
    done
done

# EDGE CASES (GOTTA TEST THEM EDGES NO CAP)
echo "🧪 Testing EDGE CASES (living life on the edge FR FR)"

edge_cases=(
    "v0.0.0"
    "v99.99.99"
    "prefix-with-multiple-hyphens-v1.2.3"
    "UPPERCASE-PREFIX-v1.2.3"
    "mixed-CASE-prefix-v1.2.3"
    "prefix-with-numbers123-v1.2.3"
)

for case in "${edge_cases[@]}"; do
    # Test both formats with edge cases
    [ "$(go-next-tag "$case")" != "" ] || (echo "❌ test edge case majorminor \`$case\` failed" && exit 1)
    [ "$(go-next-tag --format=semver "$case")" != "" ] || (echo "❌ test edge case semver \`$case\` failed" && exit 1)
done

echo "✨ ALL TESTS PASSED FR FR! THIS IS BLAZINGLY FAST™️! ✨"
echo "🎓 Certified by the International Department of Bustology"
echo "💯 NO VERSIONS WERE HARMED IN THESE TESTS!"
