package fuse

import (
	"context"
	"log"

	"bazil.org/fuse"

	"go.sancus.dev/core/errors"
	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.NodeMkdirer = (*Node)(nil)
)

func (node *Node) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (types.Node, error) {
	log.Printf("%+n: %s", errors.Here(), req)

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
	log.Printf("%+n: %s:%q", errors.Here(), "name", name)

	node := &Node{
		name: name,
		fs:   fsys,
	}
	return node, nil
}
