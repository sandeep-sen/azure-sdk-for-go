//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdomainservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListOuContainers.json
func ExampleOuContainerClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOuContainerClient().NewListPager("OuContainerResourceGroup", "OuContainer.com", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OuContainerListResult = armdomainservices.OuContainerListResult{
		// 	Value: []*armdomainservices.OuContainer{
		// 		{
		// 			Name: to.Ptr("OuContainer.com/OuContainer1"),
		// 			Type: to.Ptr("Microsoft.AAD/DomainServices/OuContainer"),
		// 			ID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/ouContainerResourceGroup/providers/Microsoft.AAD/domainServices/ouContainer.com/ouContainer/ouContainer1"),
		// 			Properties: &armdomainservices.OuContainerProperties{
		// 				Accounts: []*armdomainservices.ContainerAccount{
		// 					{
		// 						AccountName: to.Ptr("AccountName1"),
		// 						Spn: to.Ptr("Spn1"),
		// 					},
		// 					{
		// 						AccountName: to.Ptr("AccountName2"),
		// 						Spn: to.Ptr("Spn2"),
		// 				}},
		// 				ContainerID: to.Ptr("OuContainer1"),
		// 				DeploymentID: to.Ptr("0FC50BDB-AC45-48E4-BC92-F0651EA0687B"),
		// 				DomainName: to.Ptr("OuContainer.com"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceStatus: to.Ptr("Running"),
		// 				TenantID: to.Ptr("3f8cd22c-7b32-48aa-a01c-f533133b1def"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("OuContainer.com/OuContainer2"),
		// 			Type: to.Ptr("Microsoft.AAD/DomainServices/OuContainer"),
		// 			ID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/ouContainerResourceGroup/providers/Microsoft.AAD/domainServices/ouContainer.com/ouContainer/ouContainer2"),
		// 			Properties: &armdomainservices.OuContainerProperties{
		// 				Accounts: []*armdomainservices.ContainerAccount{
		// 					{
		// 						AccountName: to.Ptr("AccountName3"),
		// 						Spn: to.Ptr("Spn3"),
		// 					},
		// 					{
		// 						AccountName: to.Ptr("AccountName4"),
		// 						Spn: to.Ptr("Spn4"),
		// 				}},
		// 				ContainerID: to.Ptr("OuContainer2"),
		// 				DeploymentID: to.Ptr("0FC50BDB-AC45-48E4-BC92-F0651EA0687B"),
		// 				DomainName: to.Ptr("OuContainer.com"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceStatus: to.Ptr("Running"),
		// 				TenantID: to.Ptr("3f8cd22c-7b32-48aa-a01c-f533133b1def"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetOuContainer.json
func ExampleOuContainerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOuContainerClient().Get(ctx, "OuContainerResourceGroup", "OuContainer.com", "OuContainer1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OuContainer = armdomainservices.OuContainer{
	// 	Name: to.Ptr("OuContainer.com/OuContainer1"),
	// 	Type: to.Ptr("Microsoft.AAD/DomainServices/OuContainer"),
	// 	ID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/ouContainerResourceGroup/providers/Microsoft.AAD/domainServices/ouContainer.com/ouContainer/ouContainer1"),
	// 	Properties: &armdomainservices.OuContainerProperties{
	// 		Accounts: []*armdomainservices.ContainerAccount{
	// 			{
	// 				AccountName: to.Ptr("AccountName1"),
	// 				Spn: to.Ptr("Spn1"),
	// 			},
	// 			{
	// 				AccountName: to.Ptr("AccountName2"),
	// 				Spn: to.Ptr("Spn2"),
	// 		}},
	// 		ContainerID: to.Ptr("OuContainer1"),
	// 		DeploymentID: to.Ptr("0FC50BDB-AC45-48E4-BC92-F0651EA0687B"),
	// 		DomainName: to.Ptr("OuContainer.com"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceStatus: to.Ptr("Running"),
	// 		TenantID: to.Ptr("3f8cd22c-7b32-48aa-a01c-f533133b1def"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/CreateOuContainer.json
func ExampleOuContainerClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOuContainerClient().BeginCreate(ctx, "OuContainerResourceGroup", "OuContainer.com", "OuContainer1", armdomainservices.ContainerAccount{
		AccountName: to.Ptr("AccountName1"),
		Password:    to.Ptr("<password>"),
		Spn:         to.Ptr("Spn1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OuContainer = armdomainservices.OuContainer{
	// 	Name: to.Ptr("OuContainer.com/OuContainer1"),
	// 	Type: to.Ptr("Microsoft.AAD/DomainServices/OuContainer"),
	// 	ID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/ouContainerResourceGroup/providers/Microsoft.AAD/domainServices/ouContainer.com/ouContainer/ouContainer1"),
	// 	Properties: &armdomainservices.OuContainerProperties{
	// 		Accounts: []*armdomainservices.ContainerAccount{
	// 			{
	// 				AccountName: to.Ptr("AccountName1"),
	// 				Spn: to.Ptr("Spn1"),
	// 			},
	// 			{
	// 				AccountName: to.Ptr("AccountName2"),
	// 				Spn: to.Ptr("Spn2"),
	// 		}},
	// 		ContainerID: to.Ptr("OuContainer1"),
	// 		DeploymentID: to.Ptr("0FC50BDB-AC45-48E4-BC92-F0651EA0687B"),
	// 		DomainName: to.Ptr("OuContainer.com"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceStatus: to.Ptr("Running"),
	// 		TenantID: to.Ptr("3f8cd22c-7b32-48aa-a01c-f533133b1def"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/DeleteOuContainer.json
func ExampleOuContainerClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOuContainerClient().BeginDelete(ctx, "OuContainerResourceGroup", "OuContainer.com", "OuContainer1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateOuContainer.json
func ExampleOuContainerClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOuContainerClient().BeginUpdate(ctx, "OuContainerResourceGroup", "OuContainer.com", "OuContainer1", armdomainservices.ContainerAccount{
		AccountName: to.Ptr("AccountName1"),
		Password:    to.Ptr("<password>"),
		Spn:         to.Ptr("Spn1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OuContainer = armdomainservices.OuContainer{
	// 	Name: to.Ptr("OuContainer.com/OuContainer1"),
	// 	Type: to.Ptr("Microsoft.AAD/DomainServices/OuContainer"),
	// 	ID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/ouContainerResourceGroup/providers/Microsoft.AAD/domainServices/ouContainer.com/ouContainer/ouContainer1"),
	// 	Properties: &armdomainservices.OuContainerProperties{
	// 		Accounts: []*armdomainservices.ContainerAccount{
	// 			{
	// 				AccountName: to.Ptr("AccountName1"),
	// 				Spn: to.Ptr("Spn1"),
	// 			},
	// 			{
	// 				AccountName: to.Ptr("AccountName2"),
	// 				Spn: to.Ptr("Spn2"),
	// 		}},
	// 		ContainerID: to.Ptr("OuContainer1"),
	// 		DeploymentID: to.Ptr("0FC50BDB-AC45-48E4-BC92-F0651EA0687B"),
	// 		DomainName: to.Ptr("OuContainer.com"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceStatus: to.Ptr("Running"),
	// 		TenantID: to.Ptr("3f8cd22c-7b32-48aa-a01c-f533133b1def"),
	// 	},
	// }
}