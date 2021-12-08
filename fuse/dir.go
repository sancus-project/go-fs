package fuse

import (
	"runtime"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

type Dir struct {
	Node
}

func (fsys *Filesystem) getDir(name string) (*Dir, error) {

	if f, err := fsys.store.Open(name); err != nil {
		return nil, err
	} else if fi, err := f.Stat(); err != nil {
		f.Close()
		return nil, err
	} else if !fi.IsDir() {
		err = &fs.PathError{"getdir", name, types.ENOTDIR}

		f.Close()
		return nil, err
	} else {
		dir := &Dir{Node{
			name: name,
			fs:   fsys,
			f:    f,
		}}

		runtime.SetFinalizer(dir, func(dir *Dir) {
			dir.f.Close()
		})
		return dir, nil
	}
}

func (fsys *Filesystem) Root() (types.Node, error) {

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if fsys.root != nil {
		// cached
	} else if d, err := fsys.getDir("."); err != nil {
		// failed
		return nil, err
	} else {
		// remember
		fsys.root = d
	}

	return fsys.root, nil
}
