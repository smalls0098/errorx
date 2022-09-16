package errorx

import (
	"github.com/smalls0098/errorx/codex"
)

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code codex.Code) {
	if err == nil {
		return
	}
	err.code = code
}
