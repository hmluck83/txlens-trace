package main

import (
	"fmt"
	"os"

	"github.com/hmluck83/txlens-trace/tracer"
	"github.com/lmittmann/w3"
)

func main() {
	if len(os.Args) != 2 {
		panic("TXID is required")
	}
	txArgs := os.Args[1]

	txHash := w3.H(txArgs)

	fundFlows, err := tracer.FundFlowFromTx(txHash)
	if err != nil {
		panic(err)
	}

	for _, ff := range fundFlows {
		fmt.Printf("From: %s To: %s | Value %s | Token %s\n", ff.From.Hex(), ff.To.Hex(), ff.Value.String(), ff.Token.Hex())
	}
}
