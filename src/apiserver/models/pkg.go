package models

type Package struct {
	ID           int64  `json:"package_id"`
	Name         string `json:"package_name"`
	Tag          string `json:"package_tag"`
	CreationTime string `json:"creation_time"`
}
