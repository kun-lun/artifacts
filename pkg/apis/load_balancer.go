package apis

import (
	"gopkg.in/yaml.v2"
)

// LoadBalancer contains needed information to create a load balancer on Azure.
type LoadBalancer struct {
	Name      string        `yaml:"name"`
	CloudMeta yaml.MapSlice `yaml:"cloud_meta"`
}

const (
	LoadBalancerStandardSKU = "standard"
	LoadBalancerBasicSKU    = "basic"
)
