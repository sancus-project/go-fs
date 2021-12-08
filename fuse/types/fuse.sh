#!/bin/sh

. "${0%/*}/../../generate.sh"

cat <<EOT

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)
EOT

generate_types bazil.org/fuse \
	Attr \
	AccessRequest \
	CreateRequest \
	CreateResponse \
	FsyncRequest \
	GetattrRequest \
	GetattrResponse \
	GetxattrRequest \
	GetxattrResponse \
	LinkRequest \
	ListxattrRequest \
	ListxattrResponse \
	MkdirRequest \
	MknodRequest \
	OpenRequest \
	OpenResponse \
	PollRequest \
	PollResponse \
	ReadlinkRequest \
	RemoveRequest \
	RemovexattrRequest \
	RenameRequest \
	LookupRequest \
	LookupResponse \
	SetattrRequest \
	SetattrResponse \
	SetxattrRequest \
	SymlinkRequest


generate_types bazil.org/fuse/fs \
	Node \
	NodeForgetter \
	NodeAccesser \
	NodeCreater \
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
