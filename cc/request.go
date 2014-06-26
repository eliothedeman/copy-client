package cc

// Request is the generic representation of a request to the copy api
type Requester interface {
	Do() (*Response, error)
}
