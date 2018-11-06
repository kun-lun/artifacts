package apis

import yaml "gopkg.in/yaml.v2"

type Subnet struct {
	Range     string        `yaml:"range"`
	Gateway   string        `yaml:"gateway"`
	CloudMeta yaml.MapSlice `yaml:"cloud_meta"`
}
