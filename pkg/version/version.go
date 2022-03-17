package version

import "fmt"

var (
	revision     = "$Format:%h$"
	revisionDate = "$Format:%as$"
)

// Version returns version in format - `REVISION (REVISIONDATE)`
// value is assigned in Makefile
func Version() string {
	return fmt.Sprintf("%v (%v)", revision, revisionDate)
}
