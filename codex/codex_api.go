package codex

// ================================================================================================================
// Common error code definition.
// ================================================================================================================
var (
	CodeNil = simpleCode{-1, "", nil}  // No error code specified.
	CodeOK  = simpleCode{0, "OK", nil} // It is OK.
)

// New creates and returns an error code.
// Note that it returns an interface object of Code.
func New(code int, message string, detail interface{}) Code {
	return simpleCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCode creates and returns a new error code based on given Code.
// The code and message is from given `code`, but the detail if from given `detail`.
func WithCode(code Code, detail interface{}) Code {
	return simpleCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
