package storage

type MetaProvider interface {
}

type DataProvider interface {
}

type Provider interface {
	MetaProvider
	DataProvider
}

// TODO: code for register provider etc.
