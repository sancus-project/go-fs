package fs

//go:generate ./fs.sh

// Code generated by ./fs.sh. DO NOT EDIT.

import (
	"io/fs"
)

type (
	// PathError is an alias of the io/fs.PathError type
	PathError = fs.PathError
	// FS is an alias of the io/fs.FS type
	FS = fs.FS
	// File is an alias of the io/fs.File type
	File = fs.File
	// FileInfo is an alias of the io/fs.FileInfo type
	FileInfo = fs.FileInfo
	// FileMode is an alias of the io/fs.FileMode type
	FileMode = fs.FileMode
	// DirEntry is an alias of the io/fs.DirEntry type
	DirEntry = fs.DirEntry
	// GlobFS is an alias of the io/fs.GlobFS type
	GlobFS = fs.GlobFS
	// StatFS is an alias of the io/fs.StatFS type
	StatFS = fs.StatFS
	// SubFS is an alias of the io/fs.SubFS type
	SubFS = fs.SubFS
	// ReadDirFS is an alias of the io/fs.ReadDirFS type
	ReadDirFS = fs.ReadDirFS
	// ReadDirFile is an alias of the io/fs.ReadDirFile type
	ReadDirFile = fs.ReadDirFile
	// ReadFileFS is an alias of the io/fs.ReadFileFS type
	ReadFileFS = fs.ReadFileFS
	// WalkDirFunc is an alias of the io/fs.WalkDirFunc type
	WalkDirFunc = fs.WalkDirFunc
)

// ValidPath is a proxy function to io/fs.ValidPath()
func ValidPath(name string) bool {
	return fs.ValidPath(name)
}

// Glob is a proxy function to io/fs.Glob()
func Glob(fsys fs.FS, pattern string) ([]string, error) {
	return fs.Glob(fsys, pattern)
}

// ReadFile is a proxy function to io/fs.ReadFile()
func ReadFile(fsys fs.FS, name string) ([]byte, error) {
	return fs.ReadFile(fsys, name)
}

// ReadDir is a proxy function to io/fs.ReadDir()
func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, name)
}

// Stat is a proxy function to io/fs.Stat()
func Stat(fsys fs.FS, name string) (fs.FileInfo, error) {
	return fs.Stat(fsys, name)
}

// Sub is a proxy function to io/fs.Sub()
func Sub(fsys fs.FS, dir string) (fs.FS, error) {
	return fs.Sub(fsys, dir)
}

// FileInfoToDirEntry is a proxy function to io/fs.FileInfoToDirEntry()
func FileInfoToDirEntry(info fs.FileInfo) DirEntry {
	return fs.FileInfoToDirEntry(info)
}

// WalkDir is a proxy function to io/fs.WalkDir()
func WalkDir(fsys fs.FS, root string, fn WalkDirFunc) error {
	return fs.WalkDir(fsys, root, fn)
}
