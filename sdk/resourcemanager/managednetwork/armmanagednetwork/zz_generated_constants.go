//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

const (
	moduleName    = "armmanagednetwork"
	moduleVersion = "v0.1.0"
)

// Kind - Responsibility role under which this Managed Network Group will be created
type Kind string

const (
	KindConnectivity Kind = "Connectivity"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindConnectivity,
	}
}

// ProvisioningState - Provisioning state of the ManagedNetwork resource.
type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// Type - Gets or sets the connectivity type of a network structure policy
type Type string

const (
	TypeHubAndSpokeTopology Type = "HubAndSpokeTopology"
	TypeMeshTopology        Type = "MeshTopology"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeHubAndSpokeTopology,
		TypeMeshTopology,
	}
}