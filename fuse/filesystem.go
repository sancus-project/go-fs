package fuse

import (
	"io/fs"
	"log"
	"sync"

	"bazil.org/fuse"
)

type Filesystem struct {
	mu    sync.RWMutex
	store fs.FS
	conn  *fuse.Conn
	dir   string
	root  *Dir
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
