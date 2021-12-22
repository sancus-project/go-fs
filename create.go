package fs

import (
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
