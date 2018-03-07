package store

// Config represents how to build storage and storage configuration.
type Config interface {
	BuildStore() (Store, error)
}
