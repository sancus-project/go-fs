package dirfs

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"go.sancus.dev/core/errors"
	"go.sancus.dev/fs"
)

type Filesystem struct {
	prefix string
}

func (fsys *Filesystem) String() string {
	var s = "dirfs:"

	if fsys != nil {
		s += fmt.Sprintf("%q", fsys.prefix)
	} else {
		s += "-"
	}
	return s
}

func (fsys *Filesystem) fullname(name string) (string, error) {
	name = path.Clean("/" + name)

	if name == "/" {
		name = fsys.prefix
	} else if name[0] != '/' || !fs.ValidPath(name[1:]) {
		return "", syscall.EINVAL
	} else if filepath.Separator != '/' {
		s := strings.Split(name[1:], "/")
		s = append([]string{fsys.prefix}, s...)
		name = filepath.Join(s...)
	} else if fsys.prefix != "/" {
		name = fsys.prefix + name
	}

	return name, nil
}

func (fsys *Filesystem) Create(name string) (fs.File, error) {
	log.Printf("%+n: %s %s:%q", errors.Here(), fsys, "name", name)

	if s, err := fsys.fullname(name); err != nil {
		return nil, fs.AsPathError("create", name, err)
	} else if f, err := os.Create(s); err != nil {
		return nil, fs.AsPathError("create", name, err)
	} else {
		return f, nil
	}
}

func (fsys *Filesystem) Open(name string) (fs.File, error) {
	log.Printf("%+n: %s %s:%q", errors.Here(), fsys, "name", name)

	if s, err := fsys.fullname(name); err != nil {
		return nil, fs.AsPathError("open", name, err)
	} else if f, err := os.Open(s); err != nil {
		return nil, fs.AsPathError("open", name, err)
	} else {
		return f, nil
	}
}

func (fsys *Filesystem) OpenFile(name string, flag int, perm fs.FileMode) (fs.File, error) {
	log.Printf("%+n: %s %s:%q %s:0x%x %s:%v", errors.Here(), fsys,
		"name", name, "flag", flag, "perm", perm)

	if s, err := fsys.fullname(name); err != nil {
		return nil, fs.AsPathError("openfile", name, err)
	} else if f, err := os.OpenFile(s, flag, perm); err != nil {
		return nil, fs.AsPathError("openfile", name, err)
	} else {
		return f, nil
	}
}

func (fsys *Filesystem) ReadDir(dir string) ([]fs.DirEntry, error) {
	log.Printf("%+n: %s %s:%q", errors.Here(), fsys, "dir", dir)

	if fullname, err := fsys.fullname(dir); err != nil {
		return nil, fs.AsPathError("readdir", dir, err)
	} else if de, err := os.ReadDir(fullname); err != nil {
		return nil, fs.AsPathError("readdir", dir, err)
	} else {
		return de, nil
	}
}

func (fsys *Filesystem) Stat(name string) (fs.FileInfo, error) {
	log.Printf("%+n: %s %s:%q", errors.Here(), fsys, "name", name)

	if fullname, err := fsys.fullname(name); err != nil {
		return nil, fs.AsPathError("stat", name, err)
	} else if fi, err := os.Stat(fullname); err != nil {
		return nil, fs.AsPathError("stat", name, err)
	} else {
		return fi, nil
	}
}

func (fsys *Filesystem) Sub(dir string) (fs.FS, error) {
	log.Printf("%+n: %s %s:%q", errors.Here(), fsys, "dir", dir)

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
	log.Printf("%+n: %s", errors.Here(), fsys)
	return nil
}

func New(dir string) (*Filesystem, error) {
	log.Printf("%+n: %s:%q", errors.Here(), "dir", dir)

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
