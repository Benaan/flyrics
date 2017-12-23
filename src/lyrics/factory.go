package lyrics

type Factory interface {
	CreateProviders() []*LyricProvider
}
