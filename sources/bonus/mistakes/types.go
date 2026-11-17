//go:build bonus_mode

package mistakes

// Error is a refused equation; Kind ("vocabulary" or "structure") chooses the
// colour it is printed in.
type Error struct {
	Kind    string
	Message string
}
