//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/Skus_List.json
func ExampleSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(nil)
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
		// page.ResourceSKUCollection = armappplatform.ResourceSKUCollection{
		// 	Value: []*armappplatform.ResourceSKU{
		// 		{
		// 			Name: to.Ptr("B0"),
		// 			Capacity: &armappplatform.SKUCapacity{
		// 				Default: to.Ptr[int32](1),
		// 				Maximum: to.Ptr[int32](20),
		// 				Minimum: to.Ptr[int32](1),
		// 				ScaleType: to.Ptr(armappplatform.SKUScaleTypeAutomatic),
		// 			},
		// 			LocationInfo: []*armappplatform.ResourceSKULocationInfo{
		// 				{
		// 					Location: to.Ptr("eastus"),
		// 					ZoneDetails: []*armappplatform.ResourceSKUZoneDetails{
		// 					},
		// 					Zones: []*string{
		// 					},
		// 			}},
		// 			Locations: []*string{
		// 				to.Ptr("eastus")},
		// 				ResourceType: to.Ptr("Spring"),
		// 				Restrictions: []*armappplatform.ResourceSKURestrictions{
		// 				},
		// 				Tier: to.Ptr("Basic"),
		// 		}},
		// 	}
	}
}
