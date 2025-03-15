package middlewares

type (
	LogQueries struct{}
)

func (l LogQueries) NewLogQueries() *LogQueries {
	return &LogQueries{}
}
