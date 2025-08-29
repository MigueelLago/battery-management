package ports

import "system-management/internal/core/models"

type SystemInfoPort interface {
	GetMemory() (models.MemoryInfo, error)
	GetDisk(rootMount string) (models.DiskInfo, error)
	GetNetwork() ([]models.NetworkInfo, error)
	GetBattery() (models.BatteryInfo, error)
}
