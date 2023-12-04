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

type ACLInitParameters struct {
	Read []ReadInitParameters `json:"read,omitempty" tf:"read,omitempty"`
}

type ACLObservation struct {
	Read []ReadObservation `json:"read,omitempty" tf:"read,omitempty"`
}

type ACLParameters struct {

	// +kubebuilder:validation:Optional
	Read []ReadParameters `json:"read,omitempty" tf:"read,omitempty"`
}

type ConsumersInitParameters struct {
}

type ConsumersObservation struct {

	// The name of the consumer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The consumer URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ConsumersParameters struct {
}

type ContainerV1InitParameters struct {

	// Allows to control an access to a container. Currently only
	// the read operation is supported. If not specified, the container is
	// accessible project wide. The read structure is described below.
	ACL []ACLInitParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// Human-readable name for the Container. Does not have
	// to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// V1 container.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs []SecretRefsInitParameters `json:"secretRefs,omitempty" tf:"secret_refs,omitempty"`

	// Used to indicate the type of container. Must be one of generic, rsa or certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ContainerV1Observation struct {

	// Allows to control an access to a container. Currently only
	// the read operation is supported. If not specified, the container is
	// accessible project wide. The read structure is described below.
	ACL []ACLObservation `json:"acl,omitempty" tf:"acl,omitempty"`

	// The list of the container consumers. The structure is described below.
	Consumers []ConsumersObservation `json:"consumers,omitempty" tf:"consumers,omitempty"`

	// The container reference / where to find the container.
	ContainerRef *string `json:"containerRef,omitempty" tf:"container_ref,omitempty"`

	// The date the container was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The creator of the container.
	CreatorID *string `json:"creatorId,omitempty" tf:"creator_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Human-readable name for the Container. Does not have
	// to be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// V1 container.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs []SecretRefsObservation `json:"secretRefs,omitempty" tf:"secret_refs,omitempty"`

	// The status of the container.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Used to indicate the type of container. Must be one of generic, rsa or certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The date the container was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ContainerV1Parameters struct {

	// Allows to control an access to a container. Currently only
	// the read operation is supported. If not specified, the container is
	// accessible project wide. The read structure is described below.
	// +kubebuilder:validation:Optional
	ACL []ACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// Human-readable name for the Container. Does not have
	// to be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// V1 container.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	// +kubebuilder:validation:Optional
	SecretRefs []SecretRefsParameters `json:"secretRefs,omitempty" tf:"secret_refs,omitempty"`

	// Used to indicate the type of container. Must be one of generic, rsa or certificate.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ReadInitParameters struct {

	// Whether the container is accessible project wide.
	// Defaults to true.
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// The list of user IDs, which are allowed to access the
	// container, when project_access is set to false.
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type ReadObservation struct {

	// The date the container was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Whether the container is accessible project wide.
	// Defaults to true.
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// The date the container was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// The list of user IDs, which are allowed to access the
	// container, when project_access is set to false.
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type ReadParameters struct {

	// Whether the container is accessible project wide.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	ProjectAccess *bool `json:"projectAccess,omitempty" tf:"project_access,omitempty"`

	// The list of user IDs, which are allowed to access the
	// container, when project_access is set to false.
	// +kubebuilder:validation:Optional
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type SecretRefsInitParameters struct {

	// The name of the secret reference. The reference names must correspond the container type, more details are available here.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The secret reference / where to find the secret, URL.
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`
}

type SecretRefsObservation struct {

	// The name of the secret reference. The reference names must correspond the container type, more details are available here.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The secret reference / where to find the secret, URL.
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`
}

type SecretRefsParameters struct {

	// The name of the secret reference. The reference names must correspond the container type, more details are available here.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The secret reference / where to find the secret, URL.
	// +kubebuilder:validation:Optional
	SecretRef *string `json:"secretRef" tf:"secret_ref,omitempty"`
}

// ContainerV1Spec defines the desired state of ContainerV1
type ContainerV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerV1Parameters `json:"forProvider"`
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
	InitProvider ContainerV1InitParameters `json:"initProvider,omitempty"`
}

// ContainerV1Status defines the observed state of ContainerV1.
type ContainerV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerV1 is the Schema for the ContainerV1s API. Manages a V1 Barbican container resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ContainerV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ContainerV1Spec   `json:"spec"`
	Status ContainerV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerV1List contains a list of ContainerV1s
type ContainerV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerV1 `json:"items"`
}

// Repository type metadata.
var (
	ContainerV1_Kind             = "ContainerV1"
	ContainerV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerV1_Kind}.String()
	ContainerV1_KindAPIVersion   = ContainerV1_Kind + "." + CRDGroupVersion.String()
	ContainerV1_GroupVersionKind = CRDGroupVersion.WithKind(ContainerV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerV1{}, &ContainerV1List{})
}
