package sqlrepov

type ErrorDialect interface {
	NormalizeSQLError(error) error
}

type QueryDialect interface {
	Placeholder(index int) string
}

type Dialect interface {
	ErrorDialect
	QueryDialect
}
