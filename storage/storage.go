package storage

type Fetcher interface {
	Fetch(string) string
}

type KV [2]string

type File interface {
	Key() string

	Exist() (bool, string, error)

	Meta() (Fetcher, error)

	Append([]byte, int64, ...KV) (int64, string, error)

	Delete() (string, error)

	Bytes() ([]byte, string, error)

	SetMeta(...KV) error
}

type Iterator interface {
	Next() (File, error)

	LastKey() string
}

type Storage interface {
	File(string) File

	Iterator(prefix string, lastKey string) Iterator

	StoreFile(string, []byte, ...KV) (string, error)
}
