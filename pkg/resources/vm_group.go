package resources

// VMGroup is
type VMGroup struct {
	Name  string `yaml:"name"`
	Count int    `yaml:"count"`
	SKU   string `yaml:"sku"`
	Type  string `yaml:"type,omitempty"`
}
