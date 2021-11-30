package fuse

import (
	"go.sancus.dev/fs/fuse/types"
)

func (fsys *Filesystem) Root() (types.Node, error) {
	return nil, types.ENOSYS
}
