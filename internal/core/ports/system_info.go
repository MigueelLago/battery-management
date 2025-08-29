package ports

import "system-management/internal/core/models"

type SystemInfoPort interface {
	GetMemory(info models.MemoryInfo, err error)
	GetDisk(info models.Disk, err error)
	GetNetwork(info models.Network, err error)
	GetBattery(info models.Battery, err error)
}
