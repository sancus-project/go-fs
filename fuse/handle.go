package fuse

import (
	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

type Handle struct {
	node *Node
	id   types.HandleID
	f    fs.File
}
