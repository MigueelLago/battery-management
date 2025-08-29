//go:build darwin

package systeminfo

import (
	"bytes"
	"os/exec"
)

func getBatteryInfoDarwin() (pct float64, isCharging bool, present bool, err error) {
	cmd := exec.Command("/usr/bin/pmset", "-g", "batt")
	var out bytes.Buffer
	cmd.Stdout = &out
	if e := cmd.Run(); e != nil {
		return 0, false, false, e
	}
	return parsePMSetBattOutput(out.String())
}
