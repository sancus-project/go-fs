package fuse

import (
	"context"
	"log"

	"bazil.org/fuse"
	"github.com/kr/pretty"

	"go.sancus.dev/core/errors"
	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.HandleReadDirAller = (*Handle)(nil)
)

func (h *Handle) readDir(ctx context.Context, entries []fs.DirEntry) ([]fuse.Dirent, error) {
	log.Printf("%+n: %# v", errors.Here(), pretty.Formatter(entries))

	dirent := make([]fuse.Dirent, 0, len(entries))
	return dirent, nil
}

func (h *Handle) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	log.Printf("%+n: %s", errors.Here(), h)

	if !h.dir {
		return nil, types.ENOTDIR
	} else if rdh, ok := h.f.(fs.ReadDirFile); !ok {
		return nil, types.ENOSYS
	} else if de, err := rdh.ReadDir(0); err != nil {
		return nil, ToErrno(err)
	} else {
		return h.readDir(ctx, de)
	}
}
