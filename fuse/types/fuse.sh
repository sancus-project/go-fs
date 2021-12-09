#!/bin/sh

. "${0%/*}/../../generate.sh"

cat <<EOT

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)
EOT

generate_types bazil.org/fuse \
	AccessRequest \
	Attr \
	CreateRequest \
	CreateResponse \
	Dirent \
	FsyncRequest \
	GetattrRequest \
	GetattrResponse \
	GetxattrRequest \
	GetxattrResponse \
	LinkRequest \
	ListxattrRequest \
	ListxattrResponse \
	LookupRequest \
	LookupResponse \
	MkdirRequest \
	MknodRequest \
	OpenRequest \
	OpenResponse \
	PollRequest \
	PollResponse \
	ReadRequest \
	ReadResponse \
	ReadlinkRequest \
	ReleaseRequest \
	RemoveRequest \
	RemovexattrRequest \
	RenameRequest \
	SetattrRequest \
	SetattrResponse \
	SetxattrRequest \
	StatfsRequest \
	StatfsResponse \
	SymlinkRequest \
	WriteRequest \
	WriteResponse

generate_types bazil.org/fuse/fs \
	FS \
	FSDestroyer \
	FSStatfser \
	Handle \
	HandlePoller \
	HandleReadAller \
	HandleReadDirAller \
	HandleWriter \
	Node \
	NodeAccesser \
	NodeCreater \
	NodeForgetter \
	NodeFsyncer \
	NodeGetattrer \
	NodeGetxattrer \
	NodeLinker \
	NodeListxattrer \
	NodeMkdirer \
	NodeMknoder \
	NodeOpener \
	NodePoller \
	NodeReadlinker \
	NodeRemover \
	NodeRemovexattrer \
	NodeRenamer \
	NodeRequestLookuper \
	NodeSetattrer \
	NodeSetxattrer \
	NodeStringLookuper \
	NodeSymlinker

generate_done
