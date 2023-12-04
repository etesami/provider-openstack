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

type ServiceV2InitParameters struct {

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the service. Changing this updates the name of
	// the existing service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// service.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the router. Changing this creates a new service.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// SubnetID is the ID of the subnet. Default is null.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type ServiceV2Observation struct {

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The read-only external (public) IPv4 address that is used for the VPN service.
	ExternalV4IP *string `json:"externalV4Ip,omitempty" tf:"external_v4_ip,omitempty"`

	// The read-only external (public) IPv6 address that is used for the VPN service.
	ExternalV6IP *string `json:"externalV6Ip,omitempty" tf:"external_v6_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the service. Changing this updates the name of
	// the existing service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// service.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the router. Changing this creates a new service.
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// SubnetID is the ID of the subnet. Default is null.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type ServiceV2Parameters struct {

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the service. Changing this updates the name of
	// the existing service.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// service.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ID of the router. Changing this creates a new service.
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// SubnetID is the ID of the subnet. Default is null.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ServiceV2Spec defines the desired state of ServiceV2
type ServiceV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceV2Parameters `json:"forProvider"`
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
	InitProvider ServiceV2InitParameters `json:"initProvider,omitempty"`
}

// ServiceV2Status defines the observed state of ServiceV2.
type ServiceV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceV2 is the Schema for the ServiceV2s API. Manages a V2 Neutron VPN service resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ServiceV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routerId) || (has(self.initProvider) && has(self.initProvider.routerId))",message="spec.forProvider.routerId is a required parameter"
	Spec   ServiceV2Spec   `json:"spec"`
	Status ServiceV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceV2List contains a list of ServiceV2s
type ServiceV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceV2 `json:"items"`
}

// Repository type metadata.
var (
	ServiceV2_Kind             = "ServiceV2"
	ServiceV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceV2_Kind}.String()
	ServiceV2_KindAPIVersion   = ServiceV2_Kind + "." + CRDGroupVersion.String()
	ServiceV2_GroupVersionKind = CRDGroupVersion.WithKind(ServiceV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceV2{}, &ServiceV2List{})
}
