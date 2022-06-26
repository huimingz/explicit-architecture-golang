package persistence

const DEFAULT_MAX_RESULTS = 30

type QueryBuilderInterface interface {
	Build() QueryInterface
}
