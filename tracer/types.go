package tracer

const ()

type TxTrace struct {
	Arena []Arena `json:"arena"`
}

// INFO: null을 모두 any로 지정한 것으로 보임. 다만 null이 아닌 return 값을 아직 확인하지 못하였음
type Arena struct {
	Parent   *int  `json:"parent,omitempty"`
	Children []int `json:"children"`
	Idx      int   `json:"idx"`
	Trace    struct {
		Depth                        int    `json:"depth"`
		Success                      bool   `json:"success"`
		Caller                       string `json:"caller"`
		Address                      string `json:"address"`
		MaybePrecompile              any    `json:"maybe_precompile"`
		SelfdestructAddress          any    `json:"selfdestruct_address"`
		SelfdestructRefundTarget     any    `json:"selfdestruct_refund_target"`
		SelfdestructTransferredValue any    `json:"selfdestruct_transferred_value"`
		Kind                         string `json:"kind"`
		Value                        string `json:"value"`
		Data                         string `json:"data"`
		Output                       string `json:"output"`
		GasUsed                      int    `json:"gas_used"`
		GasLimit                     int    `json:"gas_limit"`
		Status                       string `json:"status"`
		Steps                        []any  `json:"steps"`
		Decoded                      struct {
			Label      any `json:"label"`
			ReturnData any `json:"return_data"`
			CallData   struct {
				Signature string   `json:"signature"`
				Args      []string `json:"args"`
			} `json:"call_data"`
		} `json:"decoded"`
	} `json:"trace"`
	Logs     []Logs `json:"logs"`
	Ordering []struct {
		Call int `json:"Call"`
	} `json:"ordering"`
}

type Logs struct {
	RawLog struct {
		Topics []string `json:"topics"`
		Data   string   `json:"data"`
	} `json:"raw_log"`
	Decoded struct {
		Name   string     `json:"name"`
		Params [][]string `json:"params"`
	} `json:"decoded"`
	Position int `json:"position"`
}

type FundFlow struct {
	From  string
	To    string
	Value string
	Token string
}
