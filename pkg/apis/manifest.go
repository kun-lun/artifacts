package apis

import (
	yaml "gopkg.in/yaml.v2"
)

// Manifest contains all needed information, all later on modules will
// use this manifest
type Manifest struct {
	Schema          string           `yaml:"schema,omitempty"`
	IaaS            string           `yaml:"iaas,omitempty"`
	CloudMeta       yaml.MapSlice    `yaml:"cloud_meta"`
	Platform        *Platform        `yaml:"platform,omitempty"`
	VMGroups        []VMGroup        `yaml:"vm_groups,omitempty"`
	Networks        []Network        `yaml:"networks,omitempty"`
	LoadBalancers   []LoadBalancer   `yaml:"load_balancers,omitempty"`
	StorageAccounts []StorageAccount `yaml:"storage_accounts,omitempty"`
	Databases       []Database       `yaml:"databases,omitempty"`
}

func (m *Manifest) validate() error {
	return nil
}

func (m *Manifest) FindVNetByName(network_name string) *Network {
	for _, v := range m.Networks {
		if v.Name == network_name {
			return &v
		}
	}
	return nil
}

// ToYAML converts the object to YAML bytes
func (m *Manifest) ToYAML() (b []byte, err error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	return yaml.Marshal(m)
}

// NewManifestFromYAML convert yaml bytes to Manifest object
func NewManifestFromYAML(b []byte) (m *Manifest, err error) {
	var manifest Manifest
	if err := yaml.Unmarshal(b, &manifest); err != nil {
		return nil, err
	}
	if err := manifest.validate(); err != nil {
		return nil, err
	}

	return &manifest, nil
}
