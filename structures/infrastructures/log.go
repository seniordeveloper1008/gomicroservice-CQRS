package structures

// Log doc
type Log struct {
	Message string      `json:"message"`
	File    string      `json:"file"`
	Line    string      `json:"line"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}
