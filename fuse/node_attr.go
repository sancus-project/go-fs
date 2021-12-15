package fuse

import (
	"context"

	"bazil.org/fuse"
	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.Node          = (*Node)(nil)
	_ types.NodeGetattrer = (*Node)(nil)
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

func (node *Node) Getattr(ctx context.Context, req *fuse.GetattrRequest, resp *fuse.GetattrResponse) error {
	return node.Attr(ctx, &resp.Attr)
}
