package manifests

import yaml "gopkg.in/yaml.v2"
import "resources"

type InfraManifest struct {
	LoadBalancer resources.LoadBalancer `yaml:"load_balancer,omitempty"`
	VMGroups     []resources.VMGroup    `yaml:"vm_groups,omitempty"`
	Region       string                 `yaml:"region,omitempty"`
}

// ToYAML converts the object to YAML bytes
func (im InfraManifest) ToYAML() (b []byte, err error) {
	return yaml.Marshal(im)
}

// NewInfraManifestFromYAML convert yaml bytes to InfraManifest object
func NewInfraManifestFromYAML(b []byte) (m InfraManifest, err error) {
	var im InfraManifest
	err = yaml.Unmarshal(b, &im)
	return im, err
}
