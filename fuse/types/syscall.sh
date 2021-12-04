#!/bin/sh

. "${0%/*}/../../generate.sh"

cat <<EOT

import (
	"syscall"

	"bazil.org/fuse"
)
EOT

generate_wrapped_const syscall fuse.Errno \
	ENOENT \
	ENOSYS \
	ENOTDIR \

generate_done
