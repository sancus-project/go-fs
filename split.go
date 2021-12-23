package fs

import (
	"path"
	"strings"
)

// Split splits a path by the final slash.
// Following io.fs rules the root is denoted by "."
// and the returned dir doesn't have a trailing slash
func Split(name string) (dir, base string) {
	switch name {
	case ".", "":
		// root
		return ".", ""
	default:

		i := strings.LastIndex(name, "/")
		if i < 0 {
			// root '/' name
			return ".", name
		} else {
			// name = dir '/' base
			return name[:i], name[i+1:]
		}
	}
}

// Join concatenates and cleans path elements
func Join(elem ...string) string {
	s := path.Join(elem...)
	if s == "" || s[0] != '/' {
		return s
	} else {
		return s[1:]
	}
}

// Clean replaces multiple slashes, and eliminates . and .. elements
func Clean(name string) string {
	s := path.Clean("/" + name)
	if s == "/" {
		return "."
	} else {
		return s[1:]
	}
}
