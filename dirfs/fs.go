package dirfs

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"go.sancus.dev/fs"
)

type Filesystem struct {
	prefix string
}

func (fsys *Filesystem) fullname(name string) (string, error) {
	name = path.Clean("/" + name)

	if name == "/" {
		name = fsys.prefix
	} else if name[0] != '/' || !fs.ValidPath(name[1:]) {
		return "", syscall.EINVAL
	} else if filepath.Separator == '/' {
		name = fsys.prefix + name
	} else {
		s := strings.Split(name[1:], "/")
		s = append([]string{fsys.prefix}, s...)
		name = filepath.Join(s...)
	}

	return name, nil
}

func (fsys *Filesystem) Open(name string) (fs.File, error) {
	if s, err := fsys.fullname(name); err != nil {
		return nil, fs.AsPathError("open", name, err)
	} else {
		return os.Open(s)
	}
}

func (fsys *Filesystem) ReadDir(dir string) ([]fs.DirEntry, error) {
	if fullname, err := fsys.fullname(dir); err != nil {
		return nil, fs.AsPathError("readdir", dir, err)
	} else {
		return os.ReadDir(fullname)
	}
}

func (fsys *Filesystem) Stat(name string) (fs.FileInfo, error) {
	if fullname, err := fsys.fullname(name); err != nil {
		return nil, fs.AsPathError("stat", name, err)
	} else if fi, err := os.Stat(fullname); err != nil {
		return nil, fs.AsPathError("stat", name, err)
	} else {
		return fi, nil
	}
}

func (fsys *Filesystem) Sub(dir string) (fs.FS, error) {
	if fullname, err := fsys.fullname(dir); err != nil {
		return nil, fs.AsPathError("sub", dir, err)
	} else if fi, err := os.Stat(fullname); err != nil {
		return nil, fs.AsPathError("sub", dir, err)
	} else if !fi.IsDir() {
		return nil, fs.AsPathError("sub", dir, syscall.ENOTDIR)
	} else {
		sub := &Filesystem{
			prefix: fullname,
		}
		return sub, nil
	}
}

func (fsys *Filesystem) Close() error {
	return nil
}

func New(dir string) (*Filesystem, error) {
	if fi, err := os.Stat(dir); err != nil {
		return nil, err
	} else if !fi.IsDir() {
		return nil, syscall.ENOTDIR
	} else {

		fsys := &Filesystem{
			prefix: dir,
		}

		return fsys, nil
	}
}
