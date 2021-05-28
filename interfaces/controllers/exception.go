package interfaces

// Exception doc
type Exception interface {
	NotFound()
	NotAllowed()
	InternalServerError()
	BadGateway()
}
