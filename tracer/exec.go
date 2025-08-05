package tracer

import "os/exec"

func ExecTrancation(txid string) ([]byte, error) {

	cmd := exec.Command("foundry/cast", "run", "--json", "-r", rpc, txid)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return stdout, nil
}
