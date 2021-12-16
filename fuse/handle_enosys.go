package fuse

import (
	"context"

	"bazil.org/fuse"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.HandleFlusher      = (*Handle)(nil)
	_ types.HandlePoller       = (*Handle)(nil)
	_ types.HandleReadAller    = (*Handle)(nil)
	_ types.HandleReadDirAller = (*Handle)(nil)
	_ types.HandleReader       = (*Handle)(nil)
	_ types.HandleReleaser     = (*Handle)(nil)
	_ types.HandleWriter       = (*Handle)(nil)
)

func (h *Handle) Flush(ctx context.Context, req *fuse.FlushRequest) error {
	return types.ENOSYS
}

func (h *Handle) Poll(ctx context.Context, req *fuse.PollRequest, resp *fuse.PollResponse) error {
	return types.ENOSYS
}

func (h *Handle) ReadAll(ctx context.Context) ([]byte, error) {
	return nil, types.ENOSYS
}

func (h *Handle) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return nil, types.ENOSYS
}

func (h *Handle) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	return types.ENOSYS
}

func (h *Handle) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
	return types.ENOSYS
}

func (h *Handle) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	return types.ENOSYS
}
