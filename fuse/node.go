package fuse

import (
	"context"
	"fmt"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.Node = (*Node)(nil)
)

type Node struct {
	name string
	fs   *Filesystem
}

func (node *Node) String() string {
	return fmt.Sprintf("node:%q (%p)", node.name, node)
}

func (node *Node) Attr(ctx context.Context, attr *types.Attr) error {

	// fs.FileInfo
	fi, err := fs.Stat(node.fs.store, node.name)
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

	if name == "." && fsys.root != nil {
		return fsys.root, nil
	} else if _, err := fs.Stat(fsys.store, name); err != nil {
		return nil, err
	} else {
		node := &Node{
			name: name,
			fs:   fsys,
		}

		return node, nil
	}
}

func (fsys *Filesystem) opendir(name string) (types.Node, error) {

	if name == "." && fsys.root != nil {
		return fsys.root, nil
	} else if fi, err := fs.Stat(fsys.store, name); err != nil {
		return nil, err
	} else if !fi.IsDir() {
		err = &fs.PathError{"opendir", name, types.ENOTDIR}
		return nil, err
	} else {
		node := &Node{
			name: name,
			fs:   fsys,
		}
		return node, nil
	}
}
