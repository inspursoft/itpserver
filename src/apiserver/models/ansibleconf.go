package models

type Install struct {
	Items []InstallItem `yaml:"install"`
}

type InstallItem struct {
	InsHosts string `yaml:"hosts"`
	Roles    []Role `yaml:",omitempty"`
}

type Role struct {
	Role string `yaml:"role"`
}

func YieldInstall(host string) *Install {
	return &Install{
		[]InstallItem{
			InstallItem{InsHosts: "all"},
			InstallItem{InsHosts: "host"},
		},
	}
}

func (in *Install) AddRoles(pkgList []PackageVO) {
	for _, p := range pkgList {
		in.Items[1].Roles = append(in.Items[1].Roles, Role{Role: p.Name + ":" + p.Tag})
	}
}

type Hosts struct {
	Target HostItem
	Local  HostItem
}

type HostItem struct {
	Name string
	IPs  []string
}

func YieldHosts(localIP string) *Hosts {
	return &Hosts{
		Local: HostItem{
			Name: "local",
			IPs:  []string{localIP},
		},
	}
}

func (in *Hosts) AddTarget(name string, ipList []string) {
	in.Target = HostItem{Name: name, IPs: ipList}
}
