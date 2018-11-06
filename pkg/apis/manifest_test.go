package apis_test

import (
	"encoding/json"
	"reflect"

	"github.com/go-test/deep"
	. "github.com/kun-lun/artifacts/pkg/apis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	yaml "gopkg.in/yaml.v2"
)

var _ = Describe("Manifest", func() {

	var (
		m *Manifest
	)
	BeforeEach(func() {

		platform := Platform{
			Type: "php",
		}

		networks := []Network{
			{
				Name:     "vnet-1",
				VNetName: "vnet-1",
				Subnets: []Subnet{
					{
						Range:   "10.10.0.0/24",
						Gateway: "10.10.0.1",
						Name:    "snet-1",
					},
				},
			}}

		loadBalancers := []LoadBalancer{
			{
				Name: "kunlun-wenserver-lb",
				CloudMeta: []yaml.MapItem{
					{
						Key:   "sku",
						Value: LoadBalancerStandardSKU,
					},
				},
			},
		}

		vmGroups := []VMGroup{
			{
				Name:  "jumpbox",
				SKU:   VMStandardDS1V2,
				Count: 1,
				Type:  "VM",
				Storage: &VMStorage{
					Image: &Image{
						Offer:     "offer1",
						Publisher: "ubuntu",
						SKU:       "sku1",
						Version:   "latest",
					},
					OSDisk: &OSDisk{
						Size: 10240,
					},
					DataDisks: []DataDisk{
						{
							Size: 102400,
						},
					},
					AzureFiles: []AzureFile{
						{
							StorageAccount: "storage_account_1",
							Name:           "azure_file_1",
							MountPoint:     "/mnt/azurefile_1",
						},
					},
				},
				NetworkInfos: []VMNetworkInfo{
					{
						Name:             networks[0].Name,
						SubnetName:       networks[0].Subnets[0].Name,
						LoadBalancerName: loadBalancers[0].Name,
					},
				},
				Roles: []Role{},
			},
			{
				Name:  "d2v3_group",
				SKU:   VMStandardDS1V2,
				Count: 2,
				Type:  "VM",
				Storage: &VMStorage{
					OSDisk: &OSDisk{
						Size: 10240,
					},
					DataDisks: []DataDisk{
						{
							Size: 102400,
						},
					},
					AzureFiles: []AzureFile{},
				},
				NetworkInfos: []VMNetworkInfo{
					{
						Name:             networks[0].Name,
						SubnetName:       networks[0].Subnets[0].Name,
						LoadBalancerName: loadBalancers[0].Name,
					},
				},
				Roles: []Role{},
			},
		}

		storageAccounts := []StorageAccount{
			{
				Name: "storage_account_1",
				CloudMeta: []yaml.MapItem{
					{
						Key:   "sku",
						Value: "standard",
					}, {
						Key:   "location",
						Value: "eastus",
					},
				},
			},
		}

		databases := []Database{
			{
				MigrationInformation: &MigrationInformation{
					OriginHost:     "asd",
					OriginDatabase: "asd",
					OriginUsername: "asd",
					OriginPassword: "asd",
				},
				Engine:              MysqlDB,
				EngineVersion:       "5.7",
				Cores:               2,
				Storage:             5,
				BackupRetentionDays: 35,
				Username:            "dbuser",
				Password:            "abcd1234!",
			},
		}

		// The checker add needed resource to manifest
		m = &Manifest{
			Schema: "v0.1",
			IaaS:   "azure",
			CloudMeta: []yaml.MapItem{
				{
					Key:   "location",
					Value: "eastus",
				},
			},
			Platform:        &platform,
			VMGroups:        vmGroups,
			Networks:        networks,
			LoadBalancers:   loadBalancers,
			StorageAccounts: storageAccounts,
			Databases:       databases,
		}

	})
	Describe("ToYAML", func() {
		Context("Everything OK", func() {
			It("should can be deserialize correctly", func() {
				b, err := m.ToYAML()
				Expect(err).To(BeNil())
				mCopy, err := NewManifestFromYAML(b)
				Expect(err).To(BeNil())
				deep_equal := reflect.DeepEqual(m, mCopy)
				if !deep_equal {
					if diff := deep.Equal(m, mCopy); diff != nil {
						diff_bytes, _ := json.Marshal(diff)
						println(string(diff_bytes))
					}
				}
				Expect(deep_equal).To(BeTrue())
			})
		})
	})
})
