package fuse

import (
	"context"
	"runtime"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

type Node struct {
	name string
	fs   *Filesystem
	f    fs.File
}

func (node *Node) Attr(ctx context.Context, attr *types.Attr) error {

	// fs.FileInfo
	fi, err := node.f.Stat()
	if err != nil {
		return nil
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

func (fsys *Filesystem) open(name string) (types.Node, error) {

	if name == "." {
		return fsys.Root()
	}

	if f, err := fsys.store.Open(name); err != nil {
		return nil, err
	} else {
		node := &Node{
			name: name,
			fs:   fsys,
			f:    f,
		}

		runtime.SetFinalizer(node, func(node *Node) {
			node.f.Close()
		})
		return node, nil
	}
}
