#!/bin/sh

. "${0%/*}/../../generate.sh"

cat <<EOT

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)
EOT

generate_types bazil.org/fuse \
	Attr

generate_types bazil.org/fuse/fs \
	Node

generate_done
