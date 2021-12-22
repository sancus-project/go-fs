#!/bin/sh

. "${0%/*}/generate.sh"

cat <<EOT

import (
	"io/fs"
	"os"
)
EOT

generate_const io/fs \
	ModeDir

generate_var io/fs \
	ErrInvalid \
	ErrPermission \
	ErrExist \
	ErrNotExist \
	ErrClosed \
	SkipDir

generate_var os \
	ErrNoDeadline \
	ErrDeadlineExceeded

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

generate_types os \
	SyscallError

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

generate_proxies \
	os \
	IsPermission:err:error:bool \
	IsExist:err:error:bool \
	IsNotExist:err:error:bool \
	IsTimeout:err:error:bool \

generate_done
