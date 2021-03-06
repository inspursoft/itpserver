package models

type Package struct {
	ID   int64  `json:"package_id" orm:"column(id);auto;pk"`
	Name string `json:"package_name" orm:"column(pkg_name)"`
	Tag  string `json:"package_tag" orm:"column(pkg_tag)"`
	VMs  []*VM  `json:"-" orm:"reverse(many)"`
}

type PackageVO struct {
	Name       string `json:"package_name"`
	Tag        string `json:"package_tag"`
	SourceName string `json:"source_name"`
}
