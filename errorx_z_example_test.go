package errorx_test

import (
	"errors"
	"fmt"

	"github.com/smalls0098/errorx"
	"github.com/smalls0098/errorx/codex"
)

func ExampleNewCode() {
	err := errorx.NewCode(codex.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(errorx.Code(err))

	// Output:
	// My Error
	// 10000
}

func ExampleNewCodef() {
	err := errorx.NewCodef(codex.New(10000, "", nil), "It's %s", "My Error")
	fmt.Println(err.Error())
	fmt.Println(errorx.Code(err).Code())

	// Output:
	// It's My Error
	// 10000
}

func ExampleWrapCode() {
	err1 := errors.New("permission denied")
	err2 := errorx.WrapCode(codex.New(10000, "", nil), err1, "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(errorx.Code(err2).Code())

	// Output:
	// Custom Error: permission denied
	// 10000
}

func ExampleWrapCodef() {
	err1 := errors.New("permission denied")
	err2 := errorx.WrapCodef(codex.New(10000, "", nil), err1, "It's %s", "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(errorx.Code(err2).Code())

	// Output:
	// It's Custom Error: permission denied
	// 10000
}

func ExampleEqual() {
	err1 := errors.New("permission denied")
	err2 := errorx.New("permission denied")
	err3 := errorx.NewCode(codex.CodeOK, "permission denied")
	fmt.Println(errorx.Equal(err1, err2))
	fmt.Println(errorx.Equal(err2, err3))

	// Output:
	// true
	// false
}

func ExampleIs() {
	err1 := errors.New("permission denied")
	err2 := errorx.Wrap(err1, "operation failed")
	fmt.Println(errorx.Is(err1, err1))
	fmt.Println(errorx.Is(err2, err2))
	fmt.Println(errorx.Is(err2, err1))
	fmt.Println(errorx.Is(err1, err2))

	// Output:
	// false
	// true
	// true
	// false
}
