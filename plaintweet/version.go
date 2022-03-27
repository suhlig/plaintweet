package plaintweet

import "fmt"

// ldflags will be set by goreleaser
var version = "vDEV"
var commit = "NONE"
var date = "UNKNOWN"

func VersionString() string {
	return fmt.Sprintf("plaintweet %s (%s), built on %s", version, commit, date)
}
