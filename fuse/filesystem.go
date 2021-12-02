package fuse

import (
	"io/fs"
	"log"

	"bazil.org/fuse"
)

type Filesystem struct {
	store fs.FS
	conn  *fuse.Conn
	dir   string
}

func (fsys *Filesystem) Close() error {
	// Attempt to close the fuse connection
	if err := fsys.conn.Close(); err != nil {
		return err
	}

	if err := fsys.Unmount(); err != nil {
		log.Println(err)
	}

	// Close store if a Close() is implemented
	if m, ok := fsys.store.(interface{ Close() error }); ok {
		return m.Close()
	}

	return nil
}

func (fsys *Filesystem) Reload() error {
	var err error

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
