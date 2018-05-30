package storage

import (
	"context"
	"errors"
)

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

var Config map[string]string

var instance Storage

var NewStorage func(map[string]string) Storage

func Init() {
	instance = NewStorage(Config)
}

func Register(ns func(map[string]string) Storage) {
	if Config["type"] == "inline" {
		panic("configured to apply inline storage")
	}
	NewStorage = ns
}

func Download(ctx context.Context, key string) ([]byte, string, error) {
	var (
		blob []byte
		rid  string
		err  error
		done chan int = make(chan int)
	)
	go func() {
		blob, rid, err = instance.File(key).Bytes()
		close(done)
	}()
	select {
	case <-ctx.Done():
		return nil, "", errors.New("context timeout")
	case <-done:
		return blob, rid, err
	}
}

func Upload(ctx context.Context, key string, blob []byte, options ...KV) (string, error) {
	return instance.StoreFile(key, blob, options...)
}
