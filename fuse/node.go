package fuse

import (
	"context"
	"fmt"
	"path"

	"bazil.org/fuse"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.Node        = (*Node)(nil)
	_ types.NodeMkdirer = (*Node)(nil)
)

type Node struct {
	name string
	fs   *Filesystem
}

func (node *Node) String() string {
	return fmt.Sprintf("node:%p name:%q", node, node.name)
}

func (node *Node) appendName(name string) string {

	if node.name == "." {
		return name
	} else if name == "." {
		return node.name
	} else {
		return path.Join(node.name, name)
	}
}

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

func (node *Node) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (types.Node, error) {

	name := node.appendName(req.Name)

	if fsys, ok := node.fs.store.(fs.MkdirFS); !ok {
		return nil, types.EPERM
	} else if err := fsys.Mkdir(name, req.Mode); err != nil {
		return nil, ToErrno(err)
	} else if cfs, ok := fsys.(fs.ChmodFS); ok {

		// umask
		if fi, err := fs.Stat(node.fs.store, name); err != nil {
			return nil, ToErrno(err)
		} else {
			mode0 := fi.Mode()
			mode1 := mode0 &^ req.Umask

			if mode0 == mode1 {
				// ready
			} else if err = cfs.Chmod(name, mode1); err != nil {
				return nil, ToErrno(err)
			}
		}
	}

	return node.fs.newNode(name)
}

func (fsys *Filesystem) newNode(name string) (types.Node, error) {
	node := &Node{
		name: name,
		fs:   fsys,
	}
	return node, nil
}

func (fsys *Filesystem) open(name string) (types.Node, error) {

	if name == "." && fsys.root != nil {
		return fsys.root, nil
	} else if _, err := fs.Stat(fsys.store, name); err != nil {
		return nil, ToErrno(err)
	} else {
		return fsys.newNode(name)
	}
}

func (fsys *Filesystem) opendir(name string) (types.Node, error) {

	if name == "." && fsys.root != nil {
		return fsys.root, nil
	} else if fi, err := fs.Stat(fsys.store, name); err != nil {
		return nil, ToErrno(err)
	} else if !fi.IsDir() {
		return nil, types.ENOTDIR
	} else {
		return fsys.newNode(name)
	}
}
