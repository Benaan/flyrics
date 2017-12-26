package lyrics

type Factory interface {
	CreateProviders() []*BestMatchProvider
}
