//go:build darwin

package memory

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"system-management/internal/core/models"
)

// getMemoryInfo retrieves memory information on Darwin (macOS) systems.
func GetMemoryInfo() (models.MemoryInfo, error) {

	total, err := darwinTotalMemory()
	if err != nil {
		return models.MemoryInfo{}, err
	}

	// Execute vm_stat command to get memory statistics
	cmd := exec.Command("/usr/bin/vm_stat")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return models.MemoryInfo{}, err
	}

	stats, err := parseVMStatOutuput(out.String())
	if err != nil {
		return models.MemoryInfo{}, err
	}

	page := stats.PageSize
	free := stats.Counters["Pages free"]
	inactive := stats.Counters["Pages active"]
	speculative := stats.Counters["Pages speculative"]

	available := (free + inactive + speculative) * page

	var used uint64
	if total > available {
		used = total - available
	}

	usedPct := (float64(used) / float64(total)) * 100

	return models.MemoryInfo{
		TotalBytes:  total,
		UsedBytes:   used,
		FreeBytes:   available,
		UsedPercent: usedPct,
	}, nil
}

// darwinTotalMemory retrieves the total physical memory on Darwin (macOS) systems.
func darwinTotalMemory() (uint64, error) {
	cmd := exec.Command("/usr/sbin/sysctl", "-n", "hw.memsize")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return 0, err
	}

	s := strings.TrimSpace(out.String())
	return strconv.ParseUint(s, 10, 64)
}
