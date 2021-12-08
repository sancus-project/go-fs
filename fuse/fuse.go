package fuse

import (
	"context"
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func (fsys *Filesystem) Unmount() error {
	return fuse.Unmount(fsys.dir)
}

func (fsys *Filesystem) Serve() error {
	cfg := &fs.Config{
		Debug:       fsys.debug,
		WithContext: fsys.withContext,
	}
	return fs.New(fsys.conn, cfg).Serve(fsys)
}

func (fsys *Filesystem) debug(msg interface{}) {
	log.Println(msg)
}

func (fsys *Filesystem) withContext(ctx context.Context, req fuse.Request) context.Context {
	//log.Println(req)
	return ctx
}
