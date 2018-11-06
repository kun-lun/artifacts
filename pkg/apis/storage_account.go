package apis

type StorageAccount struct {
	Name     string `yaml:"name"`
	SKU      string `yaml:"sku"`
	Location string `yaml:"location"`
}
