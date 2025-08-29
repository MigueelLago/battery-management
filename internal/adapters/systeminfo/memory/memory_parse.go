package memory

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type vmStat struct {
	PageSize uint64
	Counters map[string]uint64
}

func parseVMStatOutuput(out string) (vmStat, error) {

	s := strings.TrimSpace(out)
	if s == "" {
		return vmStat{}, errors.New("empty vm_stat output")
	}
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return vmStat{}, errors.New("no lines in vm_stat output")
	}

	rePage := regexp.MustCompile(`page size of (\d+) bytes`)
	m := rePage.FindStringSubmatch(lines[0])
	if len(m) < 2 {
		return vmStat{}, errors.New("page size not found")
	}
	ps, _ := strconv.ParseUint(m[1], 10, 64)

	reKV := regexp.MustCompile(`^([A-Za-z\s\-"]+):\s+(\d+)\.`)
	counters := map[string]uint64{}
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		kv := reKV.FindStringSubmatch(line)
		if len(kv) == 3 {
			key := strings.TrimSpace(kv[1])
			val, _ := strconv.ParseUint(kv[2], 10, 64)
			counters[key] = val
		}
	}

	return vmStat{PageSize: ps, Counters: counters}, nil
}
