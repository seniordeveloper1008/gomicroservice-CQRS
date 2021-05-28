package structures

// Response doc
type Response struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}
