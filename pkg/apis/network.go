package apis

type Network struct {
	Name     string   `yaml:"name"`
	VNetName string   `yaml:"vnet_name"`
	Subnets  []Subnet `yaml:"subnets"`
}
