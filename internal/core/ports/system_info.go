package ports

import "system-management/internal/core/models"

type SystemInfoPort interface {
	GetMemory(info models.MemoryInfo, err error)
	GetDisk(info models.DiskInfo, err error)
	GetNetwork(info models.NetworkInfo, err error)
	GetBattery(info models.BatteryInfo, err error)
}
