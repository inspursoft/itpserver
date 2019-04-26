package models

type VM struct {
	ID   string `json:"vm_id"`
	Name string `json:"vm_name"`
	OS   string `json:"vm_os"`
	Spec VMSpec `json:"vm_spec"`
}

type VMSpec struct {
	CPUs    int32  `json:"cpus"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
	Extras  string `json:"extras"`
}
