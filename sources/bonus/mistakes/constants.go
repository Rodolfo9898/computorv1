//go:build bonus_mode

package mistakes

// The colours of the messages: one per kind of mistake, so that the kind reads at
// a glance.
const (
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)
