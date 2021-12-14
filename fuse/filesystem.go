package fuse

import (
	"io/fs"
	"log"
	"sync"

	"bazil.org/fuse"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.FS = (*Filesystem)(nil)
)

type Filesystem struct {
	mu    sync.RWMutex
	store fs.FS
	conn  *fuse.Conn
	dir   string
	root  types.Node
}

func (fsys *Filesystem) Close() error {

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if err := fsys.Unmount(); err != nil {
		log.Println(err)
	}

	// Attempt to close the fuse connection
	if err := fsys.conn.Close(); err != nil {
		return err
	}

	// Close store if a Close() is implemented
	if m, ok := fsys.store.(interface{ Close() error }); ok {
		return m.Close()
	}

	return nil
}

func (fsys *Filesystem) Reload() error {
	var err error

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if m, ok := fsys.store.(interface{ Reload() error }); ok {
		err = m.Reload()
	}

	// TODO: invalidate cache
	return err
}

func (fsys *Filesystem) Root() (types.Node, error) {

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if fsys.root != nil {
		// cached
	} else if d, err := fsys.opendir("."); err != nil {
		// failed
		return nil, ToErrno(err)
	} else {
		// remember
		fsys.root = d
	}

	return fsys.root, nil
}

func New(store fs.FS, dir string, options ...fuse.MountOption) (*Filesystem, error) {

	conn, err := fuse.Mount(dir, options...)
	if err != nil {
		return nil, err
	}

	m := &Filesystem{
		store: store,
		conn:  conn,
		dir:   dir,
	}

	return m, nil
}
