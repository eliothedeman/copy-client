package cc

// Request is the generic representation of a request to the copy api
type Request interface {
	Do() (*Response, error)
}
