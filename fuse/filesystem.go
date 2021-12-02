package fuse

import (
	"io/fs"

	"bazil.org/fuse"
)

type Filesystem struct {
	store fs.FS
	conn  *fuse.Conn
	dir   string
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
