package fs

import (
	"io"
	"io/fs"
	"os"

	"syscall"
)

// Create creates or truncates the named file from the file system.
// If the file already exists, it is truncated.
// If the file does not exist, it is created with mode 0666.
// If successful, methods on the returned File can be used for I/O; the associated file
// descriptor has mode O_RDWR. If there is an error, it will be of type *PathError.
//
// If fs implements CreateFS, Create calls fsys.Create. Otherwise it tries
// fsys.OpenFile
func Create(fsys fs.FS, name string) (File, error) {
	if fsys, ok := fsys.(CreateFS); ok {
		f, err := fsys.Create(name)
		if !IsNotImplemented(err) {
			return f, err
		}
	}

	if fsys, ok := fsys.(OpenFileFS); ok {
		return fsys.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	}

	err := AsSyscallPathError("create", name, syscall.ENOSYS)
	return nil, err
}

// create2 is a variant of Create trying OpenFile first to pass the permissions.
// when fsys.OpenFile is used instead of fsys.Create the associated file
// descriptor has mode O_WRONLY.
func create2(fsys fs.FS, name string, perm fs.FileMode) (File, error) {
	if fsys, ok := fsys.(OpenFileFS); ok {
		f, err := fsys.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
		if !IsNotImplemented(err) {
			return f, err
		}
	}

	if fsys, ok := fsys.(CreateFS); ok {
		return fsys.Create(name)
	}

	err := AsSyscallPathError("create", name, syscall.ENOSYS)
	return nil, err
}

// WriteFile writes data to the named file from the file system, creating it if necessary.
// If the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile
// truncates it before writing, without changing permissions.
// If there is an error, it will be of type *PathError
//
// If fs implements WriteFileFS, WriteFile calls fsys.WriteFile. Otherwise it uses fsys.OpenFile or fsys.Create,
// Writes the data, and closes the file.
func WriteFile(fsys fs.FS, name string, data []byte, perm fs.FileMode) error {

	if fsys, ok := fsys.(WriteFileFS); ok {
		err := fsys.WriteFile(name, data, perm)
		if !IsNotImplemented(err) {
			return err
		}
	}

	f, err := create2(fsys, name, perm)
	if err == nil {
		if f, ok := f.(io.Writer); ok {
			_, err = f.Write(data)
		} else {
			err = AsSyscallPathError("write", name, syscall.EPERM)
		}

		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
	}
	return err
}
