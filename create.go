package fs

import (
	"io"
	"io/fs"
	"os"

	"syscall"
)

func Create(fsys fs.FS, name string) (File, error) {
	if fsys, ok := fsys.(CreateFS); ok {
		return fsys.Create(name)
	}

	if fsys, ok := fsys.(OpenFileFS); ok {
		return fsys.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	}

	err := &PathError{"create", name, syscall.ENOSYS}
	return nil, err
}

func create2(fsys fs.FS, name string, perm fs.FileMode) (File, error) {
	if fsys, ok := fsys.(OpenFileFS); ok {
		return fsys.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	}

	if fsys, ok := fsys.(CreateFS); ok {
		return fsys.Create(name)
	}

	err := &PathError{"create", name, syscall.ENOSYS}
	return nil, err
}

func WriteFile(fsys fs.FS, name string, data []byte, perm fs.FileMode) error {

	if fsys, ok := fsys.(WriteFileFS); ok {
		return fsys.WriteFile(name, data, perm)
	}

	f, err := create2(fsys, name, perm)
	if err == nil {
		if f, ok := f.(io.Writer); ok {
			_, err = f.Write(data)
		} else {
			err = &PathError{"write", name, syscall.EPERM}
		}

		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
	}
	return err
}
