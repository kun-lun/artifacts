package apis

import yaml "gopkg.in/yaml.v2"

type StorageAccount struct {
	Name      string        `yaml:"name"`
	CloudMeta yaml.MapSlice `yaml:"cloud_meta"`
}
