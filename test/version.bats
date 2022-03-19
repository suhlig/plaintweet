#!/usr/bin/env bats

# load "${BATS_LIBRARY_PATH:-$(dirname "${BASH_SOURCE[0]}")/../..}/bats-support/load.bash"
# load "${BATS_LIBRARY_PATH:-$(dirname "${BASH_SOURCE[0]}")/../..}/bats-assert/load.bash"

# setup() {
#     load 'bats-support/load'
#     load 'bats-assert/load'
# }

@test "version reports a version string" {
  run ${PROGRAM_UNDER_TEST:?missing} version
  assert_output --partial "$(basename "$PROGRAM_UNDER_TEST") v"
}

@test "--version exits successfully" {
  run ${PROGRAM_UNDER_TEST:?missing} version
  assert_success
}
