package models

import (
	"time"
)

type Installation struct {
	VMID         int64     `json:"vm_id"`
	CreationTime time.Time `json:"creation_time"`
	Packages     []Package `json:"packages"`
}
