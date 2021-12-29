package fs

import (
	"io/fs"
	"sort"
	"syscall"
)

// ReadDir reads the named directory and returns a list of directory entries sorted by filename.
//
// If fs implements ReadDirFS, ReadDir calls fs.ReadDir. Otherwise ReadDir calls fs.Open and uses
// ReadDir and Close on the returned file.
func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error) {

	if fsys, ok := fsys.(ReadDirFS); ok {
		list, err := fsys.ReadDir(name)
		if !IsNotImplemented(err) {
			return list, err
		}
	}

	file, err := fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dir, ok := file.(ReadDirFile)
	if !ok {
		err := AsSyscallPathError("readdir", name, syscall.ENOSYS)
		return nil, err
	}

	list, err := dir.ReadDir(-1)
	if err != nil {
		return list, err
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
