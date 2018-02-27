package storage

type MetaProvider interface {
}

type DataProvider interface {
}

type Provider interface {
	MetaProvider
	Provider
}

// TODO: code for register provider etc.
