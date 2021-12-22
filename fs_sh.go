package fs

//go:generate ./fs.sh

// Code generated by ./fs.sh. DO NOT EDIT.

import (
	"io/fs"
	"os"
)

// ModeDir is an alias of the io/fs.ModeDir constant.
const ModeDir = fs.ModeDir

var (
	// ErrInvalid is an alias of the io/fs.ErrInvalid constant.
	ErrInvalid = fs.ErrInvalid
	// ErrPermission is an alias of the io/fs.ErrPermission constant.
	ErrPermission = fs.ErrPermission
	// ErrExist is an alias of the io/fs.ErrExist constant.
	ErrExist = fs.ErrExist
	// ErrNotExist is an alias of the io/fs.ErrNotExist constant.
	ErrNotExist = fs.ErrNotExist
	// ErrClosed is an alias of the io/fs.ErrClosed constant.
	ErrClosed = fs.ErrClosed
	// SkipDir is an alias of the io/fs.SkipDir constant.
	SkipDir = fs.SkipDir
)

var (
	// ErrNoDeadline is an alias of the os.ErrNoDeadline constant.
	ErrNoDeadline = os.ErrNoDeadline
	// ErrDeadlineExceeded is an alias of the os.ErrDeadlineExceeded constant.
	ErrDeadlineExceeded = os.ErrDeadlineExceeded
)

type (
	// PathError is an alias of the io/fs.PathError type.
	PathError = fs.PathError
	// FS is an alias of the io/fs.FS type.
	FS = fs.FS
	// File is an alias of the io/fs.File type.
	File = fs.File
	// FileInfo is an alias of the io/fs.FileInfo type.
	FileInfo = fs.FileInfo
	// FileMode is an alias of the io/fs.FileMode type.
	FileMode = fs.FileMode
	// DirEntry is an alias of the io/fs.DirEntry type.
	DirEntry = fs.DirEntry
	// GlobFS is an alias of the io/fs.GlobFS type.
	GlobFS = fs.GlobFS
	// StatFS is an alias of the io/fs.StatFS type.
	StatFS = fs.StatFS
	// SubFS is an alias of the io/fs.SubFS type.
	SubFS = fs.SubFS
	// ReadDirFS is an alias of the io/fs.ReadDirFS type.
	ReadDirFS = fs.ReadDirFS
	// ReadDirFile is an alias of the io/fs.ReadDirFile type.
	ReadDirFile = fs.ReadDirFile
	// ReadFileFS is an alias of the io/fs.ReadFileFS type.
	ReadFileFS = fs.ReadFileFS
	// WalkDirFunc is an alias of the io/fs.WalkDirFunc type.
	WalkDirFunc = fs.WalkDirFunc
)

// SyscallError is an alias of the os.SyscallError type.
type SyscallError = os.SyscallError

// ValidPath is a proxy function to io/fs.ValidPath(),
// see https://pkg.go.dev/io/fs#ValidPath for details.
func ValidPath(name string) bool {
	return fs.ValidPath(name)
}

// Glob is a proxy function to io/fs.Glob(),
// see https://pkg.go.dev/io/fs#Glob for details.
func Glob(fsys fs.FS, pattern string) ([]string, error) {
	return fs.Glob(fsys, pattern)
}

// ReadFile is a proxy function to io/fs.ReadFile(),
// see https://pkg.go.dev/io/fs#ReadFile for details.
func ReadFile(fsys fs.FS, name string) ([]byte, error) {
	return fs.ReadFile(fsys, name)
}

// ReadDir is a proxy function to io/fs.ReadDir(),
// see https://pkg.go.dev/io/fs#ReadDir for details.
func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, name)
}

// Stat is a proxy function to io/fs.Stat(),
// see https://pkg.go.dev/io/fs#Stat for details.
func Stat(fsys fs.FS, name string) (fs.FileInfo, error) {
	return fs.Stat(fsys, name)
}

// Sub is a proxy function to io/fs.Sub(),
// see https://pkg.go.dev/io/fs#Sub for details.
func Sub(fsys fs.FS, dir string) (fs.FS, error) {
	return fs.Sub(fsys, dir)
}

// FileInfoToDirEntry is a proxy function to io/fs.FileInfoToDirEntry(),
// see https://pkg.go.dev/io/fs#FileInfoToDirEntry for details.
func FileInfoToDirEntry(info fs.FileInfo) DirEntry {
	return fs.FileInfoToDirEntry(info)
}

// WalkDir is a proxy function to io/fs.WalkDir(),
// see https://pkg.go.dev/io/fs#WalkDir for details.
func WalkDir(fsys fs.FS, root string, fn WalkDirFunc) error {
	return fs.WalkDir(fsys, root, fn)
}

// IsPermission is a proxy function to os.IsPermission(),
// see https://pkg.go.dev/os#IsPermission for details.
func IsPermission(err error) bool {
	return os.IsPermission(err)
}

// IsExist is a proxy function to os.IsExist(),
// see https://pkg.go.dev/os#IsExist for details.
func IsExist(err error) bool {
	return os.IsExist(err)
}

// IsNotExist is a proxy function to os.IsNotExist(),
// see https://pkg.go.dev/os#IsNotExist for details.
func IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// IsTimeout is a proxy function to os.IsTimeout(),
// see https://pkg.go.dev/os#IsTimeout for details.
func IsTimeout(err error) bool {
	return os.IsTimeout(err)
}
