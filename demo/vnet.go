package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

func main() {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	networkClient, err := armnetwork.NewClientFactory(os.Getenv("SUBSCRIPTION_ID"), credential, nil)
	if err != nil {
		panic(err)
	}

	vnetClient := networkClient.NewVirtualNetworksClient()
	pager := vnetClient.NewListAllPager(nil)
	ctx := context.Background()
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		for _, v := range page.Value {
			fmt.Println(*v.ID)
			fmt.Println(*v.Name)
		}
	}
}
