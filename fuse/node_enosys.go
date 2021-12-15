package fuse

import (
	"context"

	"bazil.org/fuse"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.NodeAccesser       = (*Node)(nil)
	_ types.NodeCreater        = (*Node)(nil)
	_ types.NodeForgetter      = (*Node)(nil)
	_ types.NodeGetxattrer     = (*Node)(nil)
	_ types.NodeLinker         = (*Node)(nil)
	_ types.NodeListxattrer    = (*Node)(nil)
	_ types.NodeMknoder        = (*Node)(nil)
	_ types.NodePoller         = (*Node)(nil)
	_ types.NodeOpener         = (*Node)(nil)
	_ types.NodeReadlinker     = (*Node)(nil)
	_ types.NodeRemover        = (*Node)(nil)
	_ types.NodeRemovexattrer  = (*Node)(nil)
	_ types.NodeRenamer        = (*Node)(nil)
	_ types.NodeSetattrer      = (*Node)(nil)
	_ types.NodeSetxattrer     = (*Node)(nil)
	_ types.NodeStringLookuper = (*Node)(nil)
	_ types.NodeSymlinker      = (*Node)(nil)
)

func (node *Node) Access(ctx context.Context, req *fuse.AccessRequest) error {
	return types.ENOSYS
}

func (node *Node) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (types.Node, types.Handle, error) {
	return nil, nil, types.ENOSYS
}

func (node *Node) Forget() {
}

func (node *Node) Fsyncer(ctx context.Context, req *fuse.FsyncRequest) error {
	return types.ENOSYS
}

func (node *Node) Getxattr(ctx context.Context, req *fuse.GetxattrRequest, resp *fuse.GetxattrResponse) error {
	return types.ENOSYS
}

func (node *Node) Link(ctx context.Context, req *fuse.LinkRequest, old types.Node) (types.Node, error) {
	return nil, types.ENOSYS
}

func (node *Node) Listxattr(ctx context.Context, req *fuse.ListxattrRequest, resp *fuse.ListxattrResponse) error {
	return types.ENOSYS
}

func (node *Node) Mknod(ctx context.Context, req *fuse.MknodRequest) (types.Node, error) {
	return nil, types.ENOSYS
}

func (node *Node) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (types.Handle, error) {
	return nil, types.ENOSYS
}

func (node *Node) Poll(ctx context.Context, req *fuse.PollRequest, resp *fuse.PollResponse) error {
	return types.ENOSYS
}

func (node *Node) Readlink(ctx context.Context, req *fuse.ReadlinkRequest) (string, error) {
	return "", types.ENOSYS
}

func (node *Node) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	return types.ENOSYS
}

func (node *Node) Removexattr(ctx context.Context, req *fuse.RemovexattrRequest) error {
	return types.ENOSYS
}

func (node *Node) Rename(ctx context.Context, req *fuse.RenameRequest, newDir types.Node) error {
	return types.ENOSYS
}

func (node *Node) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
	return types.ENOSYS
}

func (node *Node) Setxattr(ctx context.Context, req *fuse.SetxattrRequest) error {
	return types.ENOSYS
}

func (node *Node) Lookup(ctx context.Context, name string) (types.Node, error) {
	return nil, types.ENOSYS
}

func (node *Node) Symlink(ctx context.Context, req *fuse.SymlinkRequest) (types.Node, error) {
	return nil, types.ENOSYS
}
