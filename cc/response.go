package cc

// Response is the generic representation of a response from the copy api
type Responder interface {
	String() string
}
