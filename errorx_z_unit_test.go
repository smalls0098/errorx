package errorx_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/smalls0098/errorx"
	"github.com/smalls0098/errorx/codex"
)

func Test_Nil(t *testing.T) {
	if errorx.New("") == nil {
		t.Fatal()
	}
	if errorx.Wrap(nil, "test") != nil {
		t.Fatal()
	}
	if errorx.Wrap(errorx.New(""), "test") == nil {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_New(t *testing.T) {
	err := errorx.New("1")
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "1" {
		t.Fatal()
	}
	err = errorx.Newf("this is %s, number %d", "test", 1)
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "this is test, number 1" {
		t.Fatal()
	}
	err = errorx.NewSkip(1, "1")
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "1" {
		t.Fatal()
	}
	if err != nil {
		if err, ok := err.(*errorx.Error); ok {
			t.Log(err.Stack())
		}
	}
	err = errorx.NewSkipf(1, "this is %s, number %d", "test", 1)
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "this is test, number 1" {
		t.Fatal()
	}
	if err != nil {
		if err, ok := err.(*errorx.Error); ok {
			t.Log(err.Stack())
		}
	}
	t.Log("ok")
}

func Test_Wrap(t *testing.T) {
	err := errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}
	err = errorx.New("1")
	err = errorx.Wrap(err, "")
	if err.Error() != "1" {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Wrapf(t *testing.T) {
	err := errorx.New("1")
	err = errorx.Wrapf(err, "%d", 2)
	err = errorx.Wrapf(err, "%d", 3)
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}
	err = errorx.New("1")
	err = errorx.Wrapf(err, "")
	if err.Error() != "1" {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_WrapSkip(t *testing.T) {
	err := errorx.New("1")
	err = errorx.WrapSkip(1, err, "2")
	err = errorx.WrapSkip(1, err, "3")
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}
	err = errorx.New("1")
	err = errorx.WrapSkip(1, err, "")
	if err.Error() != "1" {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_WrapSkipf(t *testing.T) {
	err := errorx.New("1")
	err = errorx.WrapSkipf(1, err, "%d", 2)
	err = errorx.WrapSkipf(1, err, "%d", 3)
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}
	err = errorx.New("1")
	err = errorx.WrapSkipf(1, err, "")
	if err.Error() != "1" {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Cause(t *testing.T) {
	err := errors.New("1")
	if errorx.Cause(err) != err {
		t.Fatal()
	}
	err = errors.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if errorx.Cause(err).Error() != "1" {
		t.Fatal()
	}
	err = errorx.New("1")
	if errorx.Cause(err).Error() != "1" {
		t.Fatal()
	}
	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if errorx.Cause(err).Error() != "1" {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Format(t *testing.T) {
	err := errors.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if fmt.Sprintf("%s", err) != "3: 2: 1" {
		t.Fatal(fmt.Sprintf("%s", err))
	}
	if fmt.Sprintf("%v", err) != "3: 2: 1" {
		t.Fatal(fmt.Sprintf("%v", err))
	}
	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if fmt.Sprintf("%s", err) != "3: 2: 1" {
		t.Fatal(fmt.Sprintf("%s", err))
	}
	if fmt.Sprintf("%v", err) != "3: 2: 1" {
		t.Fatal(fmt.Sprintf("%v", err))
	}
	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if fmt.Sprintf("%-s", err) != "3" {
		t.Fatal(fmt.Sprintf("%s", err))
	}
	if fmt.Sprintf("%-v", err) != "3" {
		t.Fatal(fmt.Sprintf("%v", err))
	}
	t.Log("ok")
}

func Test_Stack(t *testing.T) {
	err := errors.New("1")
	if fmt.Sprintf("%+v", err) != "1" {
		t.Fatal()
	}
	t.Log(err)
	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	t.Log(fmt.Printf("%+v", err))
	t.Log("ok")
}

func Test_Current(t *testing.T) {
	var err error
	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if err.Error() != "3: 2: 1" {
		t.Fatal(err)
	}
	if errorx.Current(err).Error() != "3" {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Unwrap(t *testing.T) {
	var err error
	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.Wrap(err, "3")
	if err.Error() != "3: 2: 1" {
		t.Fatal(err)
	}
	err = errorx.Unwrap(err)
	if err.Error() != "2: 1" {
		t.Fatal(err)
	}
	err = errorx.Unwrap(err)
	if err.Error() != "1" {
		t.Fatal(err)
	}
	err = errorx.Unwrap(err)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func codeError() error {
	return errorx.NewCode(codex.New(1000, "ID不存在", nil), "match id is empty.")
}

func Test_Code2(t *testing.T) {
	err := codeError()
	e, ok := err.(errorx.ICode)
	t.Log(e, ok)
}

func Test_Code(t *testing.T) {
	err := errors.New("123")
	if errorx.Code(err) != codex.CodeNil {
		t.Fatal()
	}

	err = errorx.NewCode(codex.CodeOK, "123")
	if errorx.Code(err) != codex.CodeOK {
		t.Fatal()
	}
	if err.Error() != "123" {
		t.Fatal()
	}

	err = errorx.NewCodef(codex.CodeOK, "%s %d", "123", 1)
	if errorx.Code(err) != codex.CodeOK {
		t.Fatal()
	}
	if err.Error() != "123 1" {
		t.Fatal()
	}

	err = errorx.NewCodeSkip(codex.CodeOK, 0, "123")
	if errorx.Code(err).Code() != codex.CodeOK.Code() {
		t.Fatal()
	}
	if err.Error() != "123" {
		t.Fatal()
	}

	err = errorx.NewCodeSkipf(codex.CodeOK, 0, "%s %d", "123", 1)
	if errorx.Code(err).Code() != codex.CodeOK.Code() {
		t.Fatal()
	}
	if err.Error() != "123 1" {
		t.Fatal()
	}

	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.WrapCode(codex.CodeOK, err, "3")
	if errorx.Code(err).Code() != codex.CodeOK.Code() {
		t.Fatal()
	}
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}

	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.WrapCodef(codex.CodeOK, err, "%d", 3)
	if errorx.Code(err).Code() != codex.CodeOK.Code() {
		t.Fatal()
	}
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}

	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.WrapCodeSkip(codex.CodeOK, 1, err, "3")
	if errorx.Code(err).Code() != codex.CodeOK.Code() {
		t.Fatal()
	}
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}

	err = errorx.New("1")
	err = errorx.Wrap(err, "2")
	err = errorx.WrapCodeSkipf(codex.CodeOK, 1, err, "%d", 3)
	if errorx.Code(err).Code() != codex.CodeOK.Code() {
		t.Fatal()
	}
	if err.Error() != "3: 2: 1" {
		t.Fatal()
	}

	t.Log("ok")
}

func Test_SetCode(t *testing.T) {
	err := errorx.New("123")
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "123" {
		t.Fatal()
	}
	err.(*errorx.Error).SetCode(codex.CodeOK)
	if errorx.Code(err) != codex.CodeOK {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Json(t *testing.T) {
	err := errorx.Wrap(errorx.New("1"), "2")
	b, e := json.Marshal(err)
	if e != nil {
		t.Fatal()
	}
	if string(b) != `"2: 1"` {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Equal(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("1")
	err3 := errorx.New("1")
	err4 := errorx.New("4")

	if errorx.Equal(err1, err2) {
		t.Fatal()
	}
	if errorx.Equal(err3, err4) {
		t.Fatal()
	}
	if !errorx.Equal(err1, err3) {
		t.Fatal()
	}
	if errorx.Equal(err2, err4) {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_Is(t *testing.T) {
	err1 := errors.New("1")
	err2 := errorx.Wrap(err1, "2")
	err2 = errorx.Wrap(err2, "3")

	if errorx.Is(err2, err1) == false {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_HashError(t *testing.T) {
	err1 := errors.New("1")
	err2 := errorx.Wrap(err1, "2")
	err2 = errorx.Wrap(err2, "3")

	if errorx.Haerrorx(err2, err1) == false {
		t.Fatal()
	}
	t.Log("ok")
}

func Test_HashCode(t *testing.T) {
	err1 := errors.New("1")
	err2 := errorx.WrapCode(codex.CodeOK, err1, "2")
	err3 := errorx.Wrap(err2, "3")
	err4 := errorx.Wrap(err3, "4")

	if errorx.Hacodex(err1, codex.CodeOK) == true {
		t.Fatal()
	}
	if errorx.Hacodex(err2, codex.CodeOK) == false {
		t.Fatal()
	}
	if errorx.Hacodex(err3, codex.CodeOK) == false {
		t.Fatal()
	}
	if errorx.Hacodex(err4, codex.CodeOK) == false {
		t.Fatal()
	}
	t.Log("ok")
}
