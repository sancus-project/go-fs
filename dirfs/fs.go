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
	if name[0] != '/' || !fs.ValidPath(name[1:]) {
		return "", syscall.EINVAL
	}

	if name == "/" {
		name = fsys.prefix
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

func (fsys *Filesystem) Close() error {
	return nil
}

func New(prefix string) (*Filesystem, error) {
	fsys := &Filesystem{
		prefix: prefix,
	}

	return fsys, nil
}
