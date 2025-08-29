package models

type Disk struct {
	MountPoint string `json:"mount_point"`
	TotalSize  uint64 `json:"total_size"`
	UsedSize   uint64 `json:"used_size"`
	FreeSize   uint64 `json:"free_size"`
	UsageRate  string `json:"usage_rate"`
}
