package models

type OneStepInstallation struct {
	VMWithSpec *VMWithSpec `json:"vm_with_spec"`
	PackageVO  *PackageVO  `json:"package_vo"`
}
