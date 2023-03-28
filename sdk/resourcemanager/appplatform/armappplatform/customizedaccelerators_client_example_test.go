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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/CustomizedAccelerators_List.json
func ExampleCustomizedAcceleratorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomizedAcceleratorsClient().NewListPager("myResourceGroup", "myservice", "default", nil)
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
		// page.CustomizedAcceleratorResourceCollection = armappplatform.CustomizedAcceleratorResourceCollection{
		// 	Value: []*armappplatform.CustomizedAcceleratorResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationAccelerators/customizedAccelerators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationAccelerators/default/customizedAccelerators"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.CustomizedAcceleratorProperties{
		// 				Description: to.Ptr("acc-desc"),
		// 				AcceleratorTags: []*string{
		// 					to.Ptr("tag-a"),
		// 					to.Ptr("tag-b")},
		// 					DisplayName: to.Ptr("acc-name"),
		// 					GitRepository: &armappplatform.AcceleratorGitRepository{
		// 						AuthSetting: &armappplatform.AcceleratorSSHSetting{
		// 							AuthType: to.Ptr("SSH"),
		// 						},
		// 						Branch: to.Ptr("git-branch"),
		// 						Commit: to.Ptr("12345"),
		// 						GitTag: to.Ptr("git-tag"),
		// 						IntervalInSeconds: to.Ptr[int32](70),
		// 						URL: to.Ptr("git-url"),
		// 					},
		// 					IconURL: to.Ptr("acc-icon"),
		// 					ProvisioningState: to.Ptr(armappplatform.CustomizedAcceleratorProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armappplatform.SKU{
		// 					Name: to.Ptr("E0"),
		// 					Capacity: to.Ptr[int32](2),
		// 					Tier: to.Ptr("Enterprise"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/CustomizedAccelerators_Get.json
func ExampleCustomizedAcceleratorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomizedAcceleratorsClient().Get(ctx, "myResourceGroup", "myservice", "default", "acc-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomizedAcceleratorResource = armappplatform.CustomizedAcceleratorResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationAccelerators/customizedAccelerators"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationAccelerators/default/customizedAccelerators/acc-name"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.CustomizedAcceleratorProperties{
	// 		Description: to.Ptr("acc-desc"),
	// 		AcceleratorTags: []*string{
	// 			to.Ptr("tag-a"),
	// 			to.Ptr("tag-b")},
	// 			DisplayName: to.Ptr("acc-name"),
	// 			GitRepository: &armappplatform.AcceleratorGitRepository{
	// 				AuthSetting: &armappplatform.AcceleratorSSHSetting{
	// 					AuthType: to.Ptr("SSH"),
	// 				},
	// 				Branch: to.Ptr("git-branch"),
	// 				Commit: to.Ptr("12345"),
	// 				GitTag: to.Ptr("git-tag"),
	// 				IntervalInSeconds: to.Ptr[int32](70),
	// 				URL: to.Ptr("git-url"),
	// 			},
	// 			IconURL: to.Ptr("acc-icon"),
	// 			ProvisioningState: to.Ptr(armappplatform.CustomizedAcceleratorProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armappplatform.SKU{
	// 			Name: to.Ptr("E0"),
	// 			Capacity: to.Ptr[int32](2),
	// 			Tier: to.Ptr("Enterprise"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/CustomizedAccelerators_CreateOrUpdate.json
func ExampleCustomizedAcceleratorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomizedAcceleratorsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "default", "acc-name", armappplatform.CustomizedAcceleratorResource{
		Properties: &armappplatform.CustomizedAcceleratorProperties{
			Description: to.Ptr("acc-desc"),
			AcceleratorTags: []*string{
				to.Ptr("tag-a"),
				to.Ptr("tag-b")},
			DisplayName: to.Ptr("acc-name"),
			GitRepository: &armappplatform.AcceleratorGitRepository{
				AuthSetting: &armappplatform.AcceleratorSSHSetting{
					AuthType:         to.Ptr("SSH"),
					HostKey:          to.Ptr("git-auth-hostkey"),
					HostKeyAlgorithm: to.Ptr("git-auth-algorithm"),
					PrivateKey:       to.Ptr("git-auth-privatekey"),
				},
				Branch:            to.Ptr("git-branch"),
				Commit:            to.Ptr("12345"),
				GitTag:            to.Ptr("git-tag"),
				IntervalInSeconds: to.Ptr[int32](70),
				URL:               to.Ptr("git-url"),
			},
			IconURL: to.Ptr("acc-icon"),
		},
		SKU: &armappplatform.SKU{
			Name:     to.Ptr("E0"),
			Capacity: to.Ptr[int32](2),
			Tier:     to.Ptr("Enterprise"),
		},
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
	// res.CustomizedAcceleratorResource = armappplatform.CustomizedAcceleratorResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationAccelerators/customizedAccelerators"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationAccelerators/default/customizedAccelerators/acc-name"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.CustomizedAcceleratorProperties{
	// 		Description: to.Ptr("acc-desc"),
	// 		AcceleratorTags: []*string{
	// 			to.Ptr("tag-a"),
	// 			to.Ptr("tag-b")},
	// 			DisplayName: to.Ptr("acc-name"),
	// 			GitRepository: &armappplatform.AcceleratorGitRepository{
	// 				AuthSetting: &armappplatform.AcceleratorSSHSetting{
	// 					AuthType: to.Ptr("SSH"),
	// 				},
	// 				Branch: to.Ptr("git-branch"),
	// 				Commit: to.Ptr("12345"),
	// 				GitTag: to.Ptr("git-tag"),
	// 				IntervalInSeconds: to.Ptr[int32](70),
	// 				URL: to.Ptr("git-url"),
	// 			},
	// 			IconURL: to.Ptr("acc-icon"),
	// 			ProvisioningState: to.Ptr(armappplatform.CustomizedAcceleratorProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armappplatform.SKU{
	// 			Name: to.Ptr("E0"),
	// 			Capacity: to.Ptr[int32](2),
	// 			Tier: to.Ptr("Enterprise"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/CustomizedAccelerators_Delete.json
func ExampleCustomizedAcceleratorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomizedAcceleratorsClient().BeginDelete(ctx, "myResourceGroup", "myservice", "default", "acc-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/CustomizedAccelerators_Validate.json
func ExampleCustomizedAcceleratorsClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomizedAcceleratorsClient().Validate(ctx, "myResourceGroup", "myservice", "default", "acc-name", armappplatform.CustomizedAcceleratorProperties{
		Description: to.Ptr("acc-desc"),
		AcceleratorTags: []*string{
			to.Ptr("tag-a"),
			to.Ptr("tag-b")},
		DisplayName: to.Ptr("acc-name"),
		GitRepository: &armappplatform.AcceleratorGitRepository{
			AuthSetting: &armappplatform.AcceleratorSSHSetting{
				AuthType:         to.Ptr("SSH"),
				HostKey:          to.Ptr("git-auth-hostkey"),
				HostKeyAlgorithm: to.Ptr("git-auth-algorithm"),
				PrivateKey:       to.Ptr("git-auth-privatekey"),
			},
			Branch:            to.Ptr("git-branch"),
			Commit:            to.Ptr("12345"),
			GitTag:            to.Ptr("git-tag"),
			IntervalInSeconds: to.Ptr[int32](70),
			URL:               to.Ptr("git-url"),
		},
		IconURL: to.Ptr("acc-icon"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomizedAcceleratorValidateResult = armappplatform.CustomizedAcceleratorValidateResult{
	// 	ErrorMessage: to.Ptr(""),
	// 	State: to.Ptr(armappplatform.CustomizedAcceleratorValidateResultStateValid),
	// }
}
