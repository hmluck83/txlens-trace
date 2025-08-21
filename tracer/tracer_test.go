package tracer

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
)

func TestFundFlow(t *testing.T) {
	txHash := w3.H("0xff8d3d66bd1c24130554a61796acccee4f21422ddafd26999138aa41606dba6f")

	fromAddresses := []common.Address{
		common.HexToAddress("0x6CB83a93C763054242a14ee6D51eFDcA881B23D7"),
		common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
		common.HexToAddress("0xb9bA5fBAC900588255CAFFf16AF54a1B316B676f"),
		common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
	}

	toAddresses := []common.Address{
		common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
		common.HexToAddress("0xb9bA5fBAC900588255CAFFf16AF54a1B316B676f"),
		common.HexToAddress("0x6CB83a93C763054242a14ee6D51eFDcA881B23D7"),
		common.HexToAddress("0xa74FA823bC8617fa320A966b3d11B0f722eF09eE"),
	}

	values := []big.Int{
		*big.NewInt(150000000000000000),
		*big.NewInt(149025000000000000),
		*big.NewInt(149025000000000000),
		*big.NewInt(149025000000000000),
		*big.NewInt(39044948453537507),
		*big.NewInt(975000000000000),
	}

	tokens := []common.Address{
		common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x80000FB1474Ed096b72b2e30ae5bcE0eAF3EFb67"),
		common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
	}

	fundflows, err := FundFlowFromTx(txHash)
	if err != nil {
		t.Fatal(err)
	}

	for idx, val := range fundflows {
		if !((val.From.Cmp(fromAddresses[idx]) == 0) &&
			(val.To.Cmp(toAddresses[idx]) == 0) &&
			(val.Value.Cmp(&values[idx]) == 0) &&
			(val.Token.Cmp(tokens[idx]) == 0)) {
			t.Errorf("Value Error")
		}
	}
}
