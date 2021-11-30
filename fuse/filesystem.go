package fuse

import (
	"io/fs"

	"bazil.org/fuse"
)

type Filesystem struct {
	fs   fs.FS
	conn *fuse.Conn
	dir  string
}

func New(fsys fs.FS, dir string, options ...fuse.MountOption) (*Filesystem, error) {
	conn, err := fuse.Mount(dir, options...)
	if err != nil {
		return nil, err
	}

	m := &Filesystem{
		fs:   fsys,
		conn: conn,
		dir:  dir,
	}

	return m, nil
}
