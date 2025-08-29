package systeminfo

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parsePMSetBattOutput(out string) (float64, bool, bool, error) {
	s := strings.TrimSpace(out)

	if s == "" {
		return 0, false, false, errors.New("Battery information not found")
	}

	rePct := regexp.MustCompile(`(\d{1,3})%`)
	m := rePct.FindStringSubmatch(s)
	if len(m) < 2 {

		if strings.Contains(strings.ToLower(s), "no batteries") {
			return 0, false, false, errors.New("No batteries found")
		}
		return 0, false, false, errors.New("Battery percentage not found")
	}

	pcInt, _ := strconv.Atoi(m[1])
	pct := float64(pcInt)

	lower := strings.ToLower(s)
	isCharging := strings.Contains(lower, "charging") || strings.Contains(lower, "charged")

	return pct, isCharging, true, nil
}
