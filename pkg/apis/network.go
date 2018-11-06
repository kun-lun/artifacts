package apis

import yaml "gopkg.in/yaml.v2"

const (
	VNETType = "vnet"
)

type Network struct {
	Name      string        `yaml:"name"`
	Type      string        `yaml:"type"`
	Subnets   []Subnet      `yaml:"subnets"`
	CloudMeta yaml.MapSlice `yaml:"cloud_meta"`
}
