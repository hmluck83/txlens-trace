package tracer

func FundFromTrace(trace *TxTrace) []FundFlow {

	EthZero := "0x0"
	eth := "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	fundFlows := []FundFlow{}
	for _, arena := range trace.Arena {
		if arena.Trace.Value != EthZero {
			appendFundFlow := FundFlow{
				From:  arena.Trace.Caller,
				To:    arena.Trace.Address,
				Value: arena.Trace.Value,
				Token: eth,
			}
			fundFlows = append(fundFlows, appendFundFlow)
		}

		if len(arena.Logs) != 0 {
			for _, log := range arena.Logs {
				if log.Decoded.Name == "Transfer" {
					appendFundFlow := FundFlow{
						From:  log.Decoded.Params[0][1],
						To:    log.Decoded.Params[1][1],
						Value: log.Decoded.Params[2][1],
						Token: arena.Trace.Address,
					}
					fundFlows = append(fundFlows, appendFundFlow)
				}
			}
		}
	}
	return fundFlows
}
