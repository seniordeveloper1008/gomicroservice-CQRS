package structures

// Log doc
type Log struct {
	Entry interface{}
}

// LogCommon doc
type LogCommon struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// LogDebug doc
type LogDebug struct {
	LogCommon
}

// LogInfo doc
type LogInfo struct {
	LogCommon
}

// LogWarning doc
type LogWarning struct {
	LogCommon
	Errors interface{} `json:"errors"`
}

// LogError doc
type LogError struct {
	LogCommon
	File   string      `json:"file"`
	Line   string      `json:"line"`
	Errors interface{} `json:"errors"`
}

// LogFatal doc
type LogFatal struct {
	LogCommon
	File   string      `json:"file"`
	Line   string      `json:"line"`
	Errors interface{} `json:"errors"`
}

// LogPanic doc
type LogPanic struct {
	LogCommon
	File   string      `json:"file"`
	Line   string      `json:"line"`
	Errors interface{} `json:"errors"`
}
