package fs

func AsPathError(op string, path string, err error) error {
	if err == nil {
		return nil
	} else if pe, ok := err.(*PathError); ok {
		// reuse
		return pe
	} else {
		return &PathError{op, path, err}
	}
}
