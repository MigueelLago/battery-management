package usecases

import "time"

type SnapshotRequest struct {
	RootMount       string        // pointer to the root mount disk
	OnlyUpInterface bool          // flag to indicate if only active network interfaces should be included
	EnrichWiFiSSID  bool          // flag to indicate if WiFi SSID information should be enriched
	Timeout         time.Duration // timeout for the snapshot operation
}

func (r SnapshotRequest) normalized() SnapshotRequest {

	our := r
	if our.RootMount == "" {
		our.RootMount = "/"
	}
	if our.Timeout <= 0 {
		our.Timeout = 5 * time.Second
	}

	return our
}
