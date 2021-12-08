package fuse

import (
	"io/fs"
	"log"
	"sync"

	"bazil.org/fuse"
	"github.com/ancientlore/go-avltree"

	"go.sancus.dev/core/errors"
	"go.sancus.dev/fs/fuse/types"
)

var (
	_ types.FS = (*Filesystem)(nil)
)

type Filesystem struct {
	mu      sync.RWMutex
	store   fs.FS
	handles *avltree.Tree
	last    fuse.HandleID
	conn    *fuse.Conn
	dir     string
	root    types.Node
}

func (fsys *Filesystem) Close() error {
	log.Printf("%+n", errors.Here())

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if err := fsys.Unmount(); err != nil {
		log.Println(err)
	}

	// Attempt to close the fuse connection
	if err := fsys.conn.Close(); err != nil {
		return err
	}

	// Close store if a Close() is implemented
	if m, ok := fsys.store.(interface{ Close() error }); ok {
		return m.Close()
	}

	return nil
}

func (fsys *Filesystem) Reload() error {
	var err error

	log.Printf("%+n", errors.Here())

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if m, ok := fsys.store.(interface{ Reload() error }); ok {
		err = m.Reload()
	}

	// TODO: invalidate cache
	return err
}

func (fsys *Filesystem) Root() (types.Node, error) {
	log.Printf("%+n", errors.Here())

	fsys.mu.Lock()
	defer fsys.mu.Unlock()

	if fsys.root != nil {
		// cached
	} else if d, err := fsys.opendir("."); err != nil {
		// failed
		return nil, ToErrno(err)
	} else {
		// remember
		fsys.root = d
	}

	return fsys.root, nil
}

func New(store fs.FS, dir string, options ...fuse.MountOption) (*Filesystem, error) {
	log.Printf("%+n: %s:%q", errors.Here(), "dir", dir)

	conn, err := fuse.Mount(dir, options...)
	if err != nil {
		return nil, err
	}

	m := &Filesystem{
		store:   store,
		handles: avltree.New(handleCompare, 0),
		conn:    conn,
		dir:     dir,
	}

	return m, nil
}

func handleCompare(a, b interface{}) int {
	ha := a.(*Handle)
	hb := b.(*Handle)

	if ha.id < hb.id {
		return -1
	} else if ha.id > hb.id {
		return 1
	} else {
		return 0
	}
}
