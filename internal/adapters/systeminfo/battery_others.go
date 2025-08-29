//go:build !darwin

package systeminfo

func getBatteryInfoDarwin() (float64, bool, bool, error) {
	return 0, false, false, nil
}
