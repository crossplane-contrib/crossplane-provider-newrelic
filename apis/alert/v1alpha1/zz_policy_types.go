// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2024 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyInitParameters struct {

	// The New Relic account ID to operate on.  This allows the user to override the account_id attribute set on the provider. Defaults to the environment variable NEW_RELIC_ACCOUNT_ID.
	// The New Relic account ID to operate on.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// DEPRECATED The channel_ids argument is deprecated and will be removed in the next major release of the provider. An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed.
	// An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed.
	ChannelIds []*float64 `json:"channelIds,omitempty" tf:"channel_ids,omitempty"`

	// The rollup strategy for the policy.  Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET.  The default is PER_POLICY.
	// The rollup strategy for the policy. Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET. The default is PER_POLICY.
	IncidentPreference *string `json:"incidentPreference,omitempty" tf:"incident_preference,omitempty"`

	// The name of the policy.
	// The name of the policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PolicyObservation struct {

	// The New Relic account ID to operate on.  This allows the user to override the account_id attribute set on the provider. Defaults to the environment variable NEW_RELIC_ACCOUNT_ID.
	// The New Relic account ID to operate on.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// DEPRECATED The channel_ids argument is deprecated and will be removed in the next major release of the provider. An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed.
	// An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed.
	ChannelIds []*float64 `json:"channelIds,omitempty" tf:"channel_ids,omitempty"`

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The rollup strategy for the policy.  Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET.  The default is PER_POLICY.
	// The rollup strategy for the policy. Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET. The default is PER_POLICY.
	IncidentPreference *string `json:"incidentPreference,omitempty" tf:"incident_preference,omitempty"`

	// The name of the policy.
	// The name of the policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PolicyParameters struct {

	// The New Relic account ID to operate on.  This allows the user to override the account_id attribute set on the provider. Defaults to the environment variable NEW_RELIC_ACCOUNT_ID.
	// The New Relic account ID to operate on.
	// +kubebuilder:validation:Optional
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// DEPRECATED The channel_ids argument is deprecated and will be removed in the next major release of the provider. An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed.
	// An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed.
	// +kubebuilder:validation:Optional
	ChannelIds []*float64 `json:"channelIds,omitempty" tf:"channel_ids,omitempty"`

	// The rollup strategy for the policy.  Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET.  The default is PER_POLICY.
	// The rollup strategy for the policy. Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET. The default is PER_POLICY.
	// +kubebuilder:validation:Optional
	IncidentPreference *string `json:"incidentPreference,omitempty" tf:"incident_preference,omitempty"`

	// The name of the policy.
	// The name of the policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
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
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Create and manage alert policies in New Relic.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,newrelic}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
