#!/bin/bash
# shellcheck disable=all

set -e

trap 'echo "❌❌ Tests failed! ❌❌"' ERR # Exit if any command fails

go install ./cmd/go-next-tag

echo "🧪 Testing MAJOR-MINOR format"

[[ "$(go-next-tag v1.1)" == "v1.2" ]] || (echo '❌ test `v1.1` (argument) failed' && exit 1)
[[ "$(echo "v2.1" | go-next-tag)" == "v2.2" ]] || (echo '❌ test `v2.1` (piped) failed' && exit 1)
[[ "$(echo "v2.1" | go-next-tag v3.1)" == "v3.2" ]] || (echo '❌ test `v2.1/v3.1` (piped & argument) failed' && exit 1)
[[ "$(go-next-tag --format=majorminor)" == "0.0" ]] || (echo '❌ test `no-input` failed' && exit 1)

echo "🧪 Testing MAJOR bump mode"

[[ "$(go-next-tag --bump=major v1.1)" == "v2.0" ]] || (echo '❌ test (--bump=major) `v1.1` (argument) failed' && exit 1)
[[ "$(echo "v2.1" | go-next-tag --bump=major)" == "v3.0" ]] || (echo '❌ test (--bump=major) `v2.1` (piped) failed' && exit 1)
[[ "$(echo "v2.1" | go-next-tag --bump=major v3.1)" == "v4.0" ]] || (echo '❌ test (--bump=major) `v2.1/v3.1` (piped & argument) failed' && exit 1)
[[ "$(go-next-tag --format=majorminor --bump=major)" == "0.0" ]] || (echo '❌ test (--bump=major) `no-input` failed' && exit 1)

echo "🧪 Testing prefixes with MAJOR-MINOR"

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
  [[ "$(go-next-tag "${prefix}1.1")" == "${prefix}1.2" ]] || (echo "❌ test prefix \`${prefix}\` minor bump failed" && exit 1)
  # Test major bump
  [[ "$(go-next-tag --bump=major "${prefix}1.1")" == "${prefix}2.0" ]] || (echo "❌ test prefix \`${prefix}\` major bump failed" && exit 1)
done

echo "🧪 Testing base SEMVER format"

# Basic semver without prefix
[[ "$(go-next-tag --bump=patch v1.2.3)" == "v1.2.4" ]] || (echo "❌ test semver patch bump failed" && exit 1)
[[ "$(go-next-tag --bump=minor v1.2.3)" == "v1.3.0" ]] || (echo "❌ test semver minor bump failed" && exit 1)
[[ "$(go-next-tag --bump=major v1.2.3)" == "v2.0.0" ]] || (echo "❌ test semver major bump failed" && exit 1)
[[ "$(go-next-tag --bump=major)" == "0.0.0" ]] || (echo '❌ test (--bump=major) `no-input` failed' && exit 1)

echo "🧪 Testing prefixes with SEMVER"

for prefix in "${prefixes[@]}"; do
  # Test patch bump (default)
  [[ "$(go-next-tag --bump=patch "${prefix}1.2.3")" == "${prefix}1.2.4" ]] || (echo "❌ test prefix \`${prefix}\` semver patch bump failed" && exit 1)
  # Test minor bump
  [[ "$(go-next-tag --bump=minor "${prefix}1.2.3")" == "${prefix}1.3.0" ]] || (echo "❌ test prefix \`${prefix}\` semver minor bump failed" && exit 1)
  # Test major bump
  [[ "$(go-next-tag --bump=major "${prefix}1.2.3")" == "${prefix}2.0.0" ]] || (echo "❌ test prefix \`${prefix}\` semver major bump failed" && exit 1)
done

echo "🧪 Testing complex SEMVER cases"

semver_cases=(
  "1.2.3-alpha.1"
  "1.2.3-beta.2+build.123"
  "1.2.3+build.123"
  "1.2.3-rc.1+build.123"
  "1.2.3-alpha.beta.1"
)

for prefix in "${prefixes[@]}"; do
  for case in "${semver_cases[@]}"; do
    [[ "$(go-next-tag --bump=minor "${prefix}${case}")" == "${prefix}1.3.0" ]] || (echo "❌ test prefix \`${prefix}\` with complex semver \`${case}\` failed" && exit 1)
  done
done

echo "🧪 Testing EDGE CASES"

edge_cases=(
  "v0.0.0"
  "v99.99.99"
  "prefix-with-multiple-hyphens-v1.2.3"
  "prefix.with.multiple.dots.v1.2.3-alpha.1"
  "UPPERCASE-PREFIX-v1.2.3"
  "mixed-CASE-prefix-v1.2.3"
  "prefix-with-numbers123-v1.2.3"
)

for case in "${edge_cases[@]}"; do
  # Test both formats with edge cases
  [[ "$(go-next-tag --format=majorminor "${case}")" != "" ]] || (echo "❌ test edge case majorminor \`${case}\` failed" && exit 1)
  [[ "$(go-next-tag --format=semver "${case}")" != "" ]] || (echo "❌ test edge case semver \`${case}\` failed" && exit 1)
done

echo "✨ ALL TESTS PASSED ! ✨"
