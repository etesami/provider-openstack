// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetRouteV2InitParameters struct {

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetRouteV2Observation struct {

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetRouteV2Parameters struct {

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	// +kubebuilder:validation:Optional
	DestinationCidr *string `json:"destinationCidr,omitempty" tf:"destination_cidr,omitempty"`

	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	// +kubebuilder:validation:Optional
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// routing entry.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

// SubnetRouteV2Spec defines the desired state of SubnetRouteV2
type SubnetRouteV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetRouteV2Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SubnetRouteV2InitParameters `json:"initProvider,omitempty"`
}

// SubnetRouteV2Status defines the observed state of SubnetRouteV2.
type SubnetRouteV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetRouteV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetRouteV2 is the Schema for the SubnetRouteV2s API. Creates a routing entry on a OpenStack V2 subnet.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type SubnetRouteV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationCidr) || (has(self.initProvider) && has(self.initProvider.destinationCidr))",message="spec.forProvider.destinationCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nextHop) || (has(self.initProvider) && has(self.initProvider.nextHop))",message="spec.forProvider.nextHop is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetId) || (has(self.initProvider) && has(self.initProvider.subnetId))",message="spec.forProvider.subnetId is a required parameter"
	Spec   SubnetRouteV2Spec   `json:"spec"`
	Status SubnetRouteV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetRouteV2List contains a list of SubnetRouteV2s
type SubnetRouteV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetRouteV2 `json:"items"`
}

// Repository type metadata.
var (
	SubnetRouteV2_Kind             = "SubnetRouteV2"
	SubnetRouteV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetRouteV2_Kind}.String()
	SubnetRouteV2_KindAPIVersion   = SubnetRouteV2_Kind + "." + CRDGroupVersion.String()
	SubnetRouteV2_GroupVersionKind = CRDGroupVersion.WithKind(SubnetRouteV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetRouteV2{}, &SubnetRouteV2List{})
}
