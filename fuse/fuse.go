package fuse

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func (fsys *Filesystem) Unmount() error {
	return fuse.Unmount(fsys.dir)
}

func (fsys *Filesystem) Serve() error {
	return fs.Serve(fsys.conn, fsys)
}
