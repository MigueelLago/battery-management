// go:build linux
package memory

func getMemoryInfo() (models.MemoryInfo, error) {
	m, err := readMeminfo()
	if err != nil {
		return models.MemoryInfo{}, err
	}

	total := m["MemTotal"] * 1024
	var available uint64

	if v, ok := m["MemAvailable"]; ok && v > 0 {
		available = v * 1024
	} else {

		free := m["MemFree"]
		buffers := m["Buffers"]
		cached := m["Cached"]
		sreclaim := m["SReclaimable"]
		shmem := m["Shmem"]
		if cached > shmem {
			cached -= shmem
		}
		available = (free + buffers + cached + sreclaim) * 1024
	}

	var used uint64
	if total > available {
		used = total - available
	}

	usedPct := (float64(used) / float64(total)) * 100.0

	return models.MemoryInfo{
		TotalBytes: total,
		UsedBytes:  used,
		FreeBytes:  available,
		UsedPct:    usedPct,
	}, nil
}

func readMeminfo() (map[string]uint64, error) {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	want := map[string]struct{}{
		"MemTotal":     {},
		"MemAvailable": {},
		"MemFree":      {},
		"Buffers":      {},
		"Cached":       {},
		"SReclaimable": {},
		"Shmem":        {},
	}

	out := make(map[string]uint64, len(want))
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()

		colon := strings.IndexByte(line, ':')
		if colon <= 0 {
			continue
		}
		key := line[:colon]
		if _, keep := want[key]; !keep {
			continue
		}
		rest := strings.TrimSpace(line[colon+1:]) // "123456 kB"
		fields := strings.Fields(rest)
		if len(fields) == 0 {
			continue
		}
		val, err := strconv.ParseUint(fields[0], 10, 64) // em kB
		if err == nil {
			out[key] = val
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
