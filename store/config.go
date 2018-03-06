package store

type Config interface {
	BuildStore() (Store, error)
}
