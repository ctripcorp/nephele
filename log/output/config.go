package output

type Config interface {
    BuildConfig() (Output, error)
}
