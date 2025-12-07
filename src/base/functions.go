package base

// Error satisfies the error interface for ValidationError.
func (e ValidationError) Error() string {
	return e.Message
}
