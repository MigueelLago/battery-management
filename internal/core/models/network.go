package models

type Network struct {
	InterfaceName string   `json:"interface_name"`
	Addresses     []string `json:"addresses"` // IPV4 and/or IPv6 addresses
	IsUp          bool     `json:"is_up"`
	IsWireless    bool     `json:"is_wireless"`
	SSID          string   `json:"ssid"`         // Only for wireless interfaces
	SignalLevel   int      `json:"signal_level"` // Only for wireless interfaces
	MacAddress    string   `json:"mac_address"`
	IPAddress     string   `json:"ip_address"`
}
