package fuse

import (
	"context"
	"fmt"
	"log"

	"bazil.org/fuse"
	"go.sancus.dev/core/errors"
	"go.sancus.dev/fs"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.NodeOpener     = (*Node)(nil)
	_ types.HandleReleaser = (*Handle)(nil)
)

type Handle struct {
	node *Node
	id   types.HandleID
	dir  bool
	f    fs.File
}

func (h *Handle) String() string {
	var dir string
	if h.dir {
		dir = " (D)"
	}
	return fmt.Sprintf("Handle=%v%s %s", h.id, dir, h.node)
}

func (node *Node) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (types.Handle, error) {
	log.Printf("%+n: %s", errors.Here(), node)

	fsys := node.fs

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	f, err := fsys.store.Open(node.name)
	if err != nil {
		return nil, ToErrno(err)
	}

	// Handle
	h := &Handle{
		node: node,
		dir:  req.Dir,
		f:    f,
	}

	return fsys.registerHandle(h)
}

func (h *Handle) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
	log.Printf("%+n: %s", errors.Here(), h)

	fsys := h.node.fs

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	return fsys.releaseHandle(h)
}

func (fsys *Filesystem) registerHandle(h *Handle) (types.Handle, error) {
	h.id = fsys.last + 1

	for {
		if _, dupe := fsys.handles.Add(h); !dupe {
			fsys.last = h.id
			return h, nil
		}
		h.id++
	}
}

func (fsys *Filesystem) releaseHandle(h *Handle) error {
	fsys.handles.Remove(h)
	return nil
}
