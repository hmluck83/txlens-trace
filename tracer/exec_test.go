package tracer

import (
	"os"
	"testing"
)

func TestExec(t *testing.T) {
	t.Log(os.Getwd())
	out, err := ExecTrancation("0x8efefd6d1b09fca29c9cae52223e63e1d21f3a15a6f9285b857ebb8ddc13c8d7")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(out))
}
