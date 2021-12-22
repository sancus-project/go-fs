package fs

type ChmodFS interface {
	FS

	Chmod(name string, perm FileMode) error
}

type CreateFS interface {
	FS

	Create(name string) (File, error)
}

type OpenFileFS interface {
	FS

	OpenFile(name string, flag int, perm FileMode) (File, error)
}

type MkdirFS interface {
	FS

	Mkdir(name string, perm FileMode) error
}

type WriteFileFS interface {
	FS

	WriteFile(name string, data []byte, perm FileMode) error
}
