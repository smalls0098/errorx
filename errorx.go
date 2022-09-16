package errorx

import (
	"github.com/smalls0098/errorx/codex"
)

// IIs is the interface for Is feature.
type IIs interface {
	Error() string
	Is(target error) bool
}

// IEqual is the interface for Equal feature.
type IEqual interface {
	Error() string
	Equal(target error) bool
}

// ICode is the interface for Code feature.
type ICode interface {
	Error() string
	Code() codex.Code
}

// IStack is the interface for Stack feature.
type IStack interface {
	Error() string
	Stack() string
}

// ICause is the interface for Cause feature.
type ICause interface {
	Error() string
	Cause() error
}

// ICurrent is the interface for Current feature.
type ICurrent interface {
	Error() string
	Current() error
}

// IUnwrap is the interface for Unwrap feature.
type IUnwrap interface {
	Error() string
	Unwrap() error
}
