#!/usr/bin/env bats

setup() {
  load "${BATS_LIBRARY_PATH:-../..}/bats-support/load.bash"
  load "${BATS_LIBRARY_PATH:-../..}/bats-assert/load.bash"
}

@test "version reports a version string" {
  run ${PROGRAM_UNDER_TEST:?missing} version
  assert_output --partial "$(basename "$PROGRAM_UNDER_TEST") v"
}

@test "--version exits successfully" {
  run ${PROGRAM_UNDER_TEST:?missing} version
  assert_success
}
