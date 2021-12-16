package fuse

import (
	"fmt"

	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

type Handle struct {
	node *Node
	id   types.HandleID
	f    fs.File
}

func (h *Handle) String() string {
	return fmt.Sprintf("Handle=%v %s", h.id, h.node)
}
