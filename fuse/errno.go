package fuse

import (
	"errors"
	"io/fs"
	"syscall"

	"bazil.org/fuse"
)

func ToErrno(err error) fuse.Errno {
	switch v := err.(type) {
	case syscall.Errno:
		// direct wrapping
		return fuse.Errno(v)
	case *fs.PathError:
		// fs.FS requires *fs.PathError, so we will get them often
		return ToErrno(v.Err)
	default:
		var errnum fuse.ErrorNumber
		if errors.As(err, &errnum) {
			return errnum.Errno()
		} else {
			return fuse.DefaultErrno
		}
	}
}
