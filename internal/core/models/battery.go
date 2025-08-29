package models

type Battery struct {
	Percentage     float64 `json:"percentage"`
	IsCharging     bool    `json:"is_charging"`
	DesignCapacity float64 `json:"design_capacity"` // optional (health monitoring)
	Present        bool    `json:"present"`
}
