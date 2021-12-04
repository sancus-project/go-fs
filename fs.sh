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

generate_proxies \
	io/fs \
	ValidPath:name:string:bool \
	Glob:fsys,pattern:fs.FS,string:[]string,error \
	ReadFile:fsys,name:fs.FS,string:[]byte,error \
	ReadDir:fsys,name:fs.FS,string:[]fs.DirEntry,error \
	Stat:fsys,name:fs.FS,string:fs.FileInfo,error \
	Sub:fsys,dir:fs.FS,string:fs.FS,error \
	FileInfoToDirEntry:info:fs.FileInfo:DirEntry \
	WalkDir:fsys,root,fn:fs.FS,string,WalkDirFunc:error \

generate_done
