package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func main() {
	// sdk credential chains can get creds from environment https://learn.microsoft.com/en-us/azure/developer/go/sdk/authentication/credential-chains#defaultazurecredential-overview
	// to authenticate using username/password, we need to set AZURE_CLIENT_ID, AZURE_TENANT_ID, AZURE_USERNAME and AZURE_PASSWORD env vars
	// NOTE: if MFA is enabled, username password authentication will fail
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	p := policy.TokenRequestOptions{
		Scopes: []string{"https://management.azure.com//.default"},
	}
	token, err := credential.GetToken(context.TODO(), p)
	if err != nil {
		panic(err)
	}
	fmt.Println(token.Token)
}
