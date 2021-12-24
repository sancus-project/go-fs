package fs

import (
	"errors"
	"io/fs"
	"os"
	"syscall"
)

// AsSyscallPathError() returns a fs.PathError indicating a syscall.Errno.
// Returns nil if err == 0 (aka success)
func AsSyscallPathError(op string, path string, err syscall.Errno) error {
	if err == 0 {
		return nil
	}

	return &fs.PathError{op, path, err}
}

// AsPathError() wraps an error in fs.PathError, unwrapping known error
// types, and returning nil if nil is given.
func AsPathError(op string, path string, err error) error {
	if err == nil {
		return nil
	}

	switch e2 := err.(type) {
	case syscall.Errno:
		return AsSyscallPathError(op, path, e2)
	case *fs.PathError:
		// reuse op and err, not path
		op, err = e2.Op, e2.Err
	case *os.LinkError:
		err = e2.Err
	case *os.SyscallError:
		err = e2.Err
	}

	return &fs.PathError{op, path, err}
}

// IsNotImplemented tests if an error is ENOSYS
func IsNotImplemented(err error) bool {
	return errors.Is(err, syscall.ENOSYS)
}
