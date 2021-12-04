#!/bin/sh

. "${0%/*}/generate.sh"

cat <<EOT

import (
	"io/fs"
)
EOT

generate_types io/fs \
	PathError \
	FS \
	File \
	FileInfo \
	FileMode \
	DirEntry \
	GlobFS \
	StatFS \
	SubFS \
	ReadDirFS \
	ReadDirFile \
	ReadFileFS \
	WalkDirFunc

generate_done
