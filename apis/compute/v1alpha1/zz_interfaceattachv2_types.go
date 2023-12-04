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

type InterfaceAttachV2InitParameters struct {

	// An IP address to assosciate with the port.
	// NOTE: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The ID of the Instance to attach the Port or Network to.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// NOTE: This option and port_id are mutually exclusive.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The ID of the Port to attach to an Instance.
	// NOTE: This option and network_id are mutually exclusive.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to create the interface attachment.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new attachment.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterfaceAttachV2Observation struct {

	// An IP address to assosciate with the port.
	// NOTE: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Instance to attach the Port or Network to.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// NOTE: This option and port_id are mutually exclusive.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The ID of the Port to attach to an Instance.
	// NOTE: This option and network_id are mutually exclusive.
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to create the interface attachment.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new attachment.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InterfaceAttachV2Parameters struct {

	// An IP address to assosciate with the port.
	// NOTE: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// The ID of the Instance to attach the Port or Network to.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// NOTE: This option and port_id are mutually exclusive.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The ID of the Port to attach to an Instance.
	// NOTE: This option and network_id are mutually exclusive.
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The region in which to create the interface attachment.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new attachment.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// InterfaceAttachV2Spec defines the desired state of InterfaceAttachV2
type InterfaceAttachV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InterfaceAttachV2Parameters `json:"forProvider"`
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
	InitProvider InterfaceAttachV2InitParameters `json:"initProvider,omitempty"`
}

// InterfaceAttachV2Status defines the observed state of InterfaceAttachV2.
type InterfaceAttachV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InterfaceAttachV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceAttachV2 is the Schema for the InterfaceAttachV2s API. Attaches a Network Interface to an Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type InterfaceAttachV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	Spec   InterfaceAttachV2Spec   `json:"spec"`
	Status InterfaceAttachV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceAttachV2List contains a list of InterfaceAttachV2s
type InterfaceAttachV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InterfaceAttachV2 `json:"items"`
}

// Repository type metadata.
var (
	InterfaceAttachV2_Kind             = "InterfaceAttachV2"
	InterfaceAttachV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InterfaceAttachV2_Kind}.String()
	InterfaceAttachV2_KindAPIVersion   = InterfaceAttachV2_Kind + "." + CRDGroupVersion.String()
	InterfaceAttachV2_GroupVersionKind = CRDGroupVersion.WithKind(InterfaceAttachV2_Kind)
)

func init() {
	SchemeBuilder.Register(&InterfaceAttachV2{}, &InterfaceAttachV2List{})
}
