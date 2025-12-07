package colors

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Grey   = "\033[90m"
	Reset  = "\033[0m"
)

func Wrap(color, msg string) string {
	return color + msg + Reset
}
