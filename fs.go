package fs

type CreateFS interface {
	FS

	Create(name string) (*File, error)
}

type OpenFileFS interface {
	FS

	OpenFile(name string, flag int, perm FileMode) (File, error)
}
