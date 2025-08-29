package models

// Structure representing a snapshot of the system's state
type SystemSnapshot struct {
	Memory  MemoryInfo  `json:"memory"`
	Disk    DiskInfo    `json:"disk"`
	Network NetworkInfo `json:"network"`
	Battery BatteryInfo `json:"battery"`
}
