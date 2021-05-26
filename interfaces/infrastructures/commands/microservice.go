package interfaces

// Microservice doc
type Microservice interface {
	Ensure()
	Migrate()
	Serve()
	Test()
}
