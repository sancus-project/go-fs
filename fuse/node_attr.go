package fuse

import (
	"context"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.Node = (*Node)(nil)
)

func (node *Node) Attr(ctx context.Context, attr *types.Attr) error {

	// fs.FileInfo
	fi, err := fs.Stat(node.fs.store, node.name)
	if err != nil {
		return ToErrno(err)
	}

	size := fi.Size()
	if size < 0 {
		size = 0
	}
	modtime := fi.ModTime()

	*attr = types.Attr{
		Size:  uint64(size),
		Mtime: modtime,
		Ctime: modtime,
		Mode:  fi.Mode(),
	}

	return nil
}
