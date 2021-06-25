package structures

// Request doc
type Request struct {
	Message interface{} `json:"message"`
	Headers interface{} `json:"headers"`
	Payload interface{} `json:"payload"`
	Form    interface{} `json:"form"`
	File    interface{} `json:"file"`
}
