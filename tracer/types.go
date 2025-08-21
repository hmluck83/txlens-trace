package tracer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type fundFlow struct {
	From  common.Address
	To    common.Address
	Value big.Int
	Token common.Address
}
