#!/bin/sh

. "${0%/*}/../generate.sh"

cat <<EOT

import (
	"bazil.org/fuse"
)
EOT

generate_fuse_mountoption() {
	local fn="$1" vars="${2:-}" types="${3:-}"

	generate_proxy bazil.org/fuse "$fn" \
		"${2:-}" "${3:-}" \
		fuse.MountOption
}

for n in \
	AllowDev \
	AllowNonEmptyMount \
	AllowOther \
	AllowSUID \
	AsyncRead \
	CongestionThreshold:n:uint16 \
	DaemonTimeout:name:string \
	DefaultPermissions \
	FSName:name:string \
	LockingFlock \
	LockingPOSIX \
	MaxBackground:n:uint16 \
	MaxReadahead:n:uint32 \
	ReadOnly \
	Subtype:fstype:string \
	WritebackCache \
	; do

	generate_fuse_mountoption $(echo "$n" | tr ':' ' ')
done

generate_types fuse \
	MountOption

generate_done
