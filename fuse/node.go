package fuse

import (
	"fmt"
	"path"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
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
