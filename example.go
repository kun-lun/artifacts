package main

import (
	"fmt"

	manifest "github.com/kun-lun/artifacts/pkg/apis"
)

func main() {
	// Below code will be executed by checker, the checker determines what resource will be created

	platform := manifest.Platform{
		Type: "php",
	}

	vNets := []manifest.VirtualNetwork{
		{
			Name: "vnet-1",
			Subnets: []manifest.Subnet{
				{
					Name:    "snet-1",
					Range:   "10.10.0.0/24",
					Gateway: "10.10.0.1",
				},
			},
		},
	}

	loadBalancers := []manifest.LoadBalancer{
		{
			Name: "kunlun-wenserver-lb",
			SKU:  manifest.LoadBalancerStandardSKU,
		},
	}

	vmGroups := []manifest.VMGroup{
		{
			Name: "jumpbox",
			Meta: &manifest.VMGroupMetaData{
				AutoStop: false,
			},
			Count: 1,
			SKU:   manifest.VMStandardDS1V2,
			Type:  "VM",
			Storage: &manifest.VMStorage{
				Image: &manifest.Image{
					Offer:     "offer1",
					Publisher: "ubuntu",
					SKU:       "sku1",
					Version:   "latest",
				},
				OSDisk: &manifest.OSDisk{
					Size: 10240,
				},
				DataDisks: []manifest.DataDisk{
					{
						Size: 102400,
					},
				},
				AzureFiles: []manifest.AzureFile{
					{
						StorageAccount: "storage_account_1",
						Name:           "azure_file_1",
						MountPoint:     "/mnt/azurefile_1",
					},
				},
			},
			Networks: []manifest.VMNetWork{
				{
					Subnet:       &vNets[0].Subnets[0],
					LoadBalancer: &loadBalancers[0],
				},
			},
		},
		{
			Name: "d2v3_group",
			Meta: &manifest.VMGroupMetaData{
				AutoStop: false,
			},
			Count: 2,
			SKU:   manifest.VMStandardD2V3,
			Type:  "VM",
			Storage: &manifest.VMStorage{
				OSDisk: &manifest.OSDisk{
					Size: 10240,
				},
				DataDisks: []manifest.DataDisk{
					{
						Size: 102400,
					},
				},
			},
			Networks: []manifest.VMNetWork{
				{
					Subnet:       &vNets[0].Subnets[0],
					LoadBalancer: &loadBalancers[0],
				},
			},
		},
	}

	storageAccounts := []manifest.StorageAccount{
		{
			Name:     "storage_account_1",
			Location: "eastus",
			SKU:      "standard",
		},
	}

	databases := []manifest.Database{
		{
			MigrationInformation: &manifest.MigrationInformation{
				OriginHost:     "asd",
				OriginDatabase: "asd",
				OriginUsername: "asd",
				OriginPassword: "asd",
			},
			Engine:              manifest.MysqlDB,
			EngineVersion:       "5.7",
			Cores:               2,
			Storage:             5,
			BackupRetentionDays: 35,
			Username:            "binxi",
			Password:            "abcd1234!",
		},
	}

	// The checker add needed resource to manifest
	m := manifest.Manifest{
		Schema:          "v0.1",
		Region:          "eastus",
		IaaS:            "azure",
		Platform:        &platform,
		VMGroups:        vmGroups,
		VNets:           vNets,
		LoadBalancers:   loadBalancers,
		StorageAccounts: storageAccounts,
		Databases:       databases,
	}
	// The checker convert the object to yaml bytes
	b, _ := m.ToYAML()
	fmt.Println(string(b))
	// ...
	// The checker write yaml bytes to the disk
	// ...

	// Below code will be executed by later on modules, here we take infra module as an example

	// ...
	// The infra module read yaml file from disk and get yaml bytes
	// ...

	// The infra module new a manifest object using yaml bytes
	mCopy, err := manifest.NewManifestFromYAML(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	// The infra module access field in the yaml object as needed
	fmt.Println(mCopy.LoadBalancers)
	fmt.Println(mCopy.VMGroups)
	fmt.Println(mCopy.Region)
}
