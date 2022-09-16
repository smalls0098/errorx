package codex_test

import (
	"testing"

	"github.com/smalls0098/errorx/codex"
)

func Test_Nil(t *testing.T) {
	c := codex.New(1, "custom error", "detailed description")
	if c.Code() != 1 {
		t.Fatal()
	}
	if c.Message() != "custom error" {
		t.Fatal()
	}
	if c.Detail() != "detailed description" {
		t.Failed()
	}
	t.Log("ok")
}
