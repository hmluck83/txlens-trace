package tracer

import (
	"encoding/json"
	"os"
	"testing"
)

func TestDump(t *testing.T) {

}

func TestTracer(t *testing.T) {
	txString, err := os.ReadFile("../0xff8d3d66bd1c24130554a61796acccee4f21422ddafd26999138aa41606dba6f.json")

	if err != nil {
		t.Fatal(err)
	}

	var txtrace TxTrace
	err = json.Unmarshal(txString, &txtrace)
	if err != nil {
		t.Fatal(err)
	}

	fundflow := FundFromTrace(&txtrace)
	t.Log(len(fundflow))

}
