package netplan

type NetplanConfig struct {
	Network struct {
		Version   int                        `yaml:"version"`
		Ethernets map[string]NetplanEthernet `yaml:"ethernets,omitempty"`
		Vlans     map[string]NetplanVlan     `yaml:"vlans,omitempty"`
	}
}

type NetplanEthernet struct {
	Match          NetplanMatch `yaml:"match,omitempty"`
	SetName        string       `yaml:"set-name,omitempty"`
	Addresses      []string     `yaml:"addresses,omitempty"`
	MTU            string       `yaml:"mtu,omitempty"`
	ActivationMode string       `yaml:"activation-mode,omitempty"`
}

type NetplanVlan struct {
	Id             int      `yaml:"id"`
	Link           string   `yaml:"link"`
	Addresses      []string `yaml:"addresses,omitempty"`
	ActivationMode string   `yaml:"activation-mode,omitempty"`
}

type NetplanMatch struct {
	MacAddress string `yaml:"macaddress,omitempty"`
}
