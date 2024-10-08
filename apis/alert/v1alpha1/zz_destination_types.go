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

type AuthBasicInitParameters struct {

	// Specifies an authentication password for use with a destination.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username of the basic auth.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type AuthBasicObservation struct {

	// The username of the basic auth.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type AuthBasicParameters struct {

	// Specifies an authentication password for use with a destination.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username of the basic auth.
	// +kubebuilder:validation:Optional
	User *string `json:"user" tf:"user,omitempty"`
}

type AuthCustomHeaderInitParameters struct {

	// The key of the header.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The secret value of the header.
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type AuthCustomHeaderObservation struct {

	// The key of the header.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type AuthCustomHeaderParameters struct {

	// The key of the header.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The secret value of the header.
	// +kubebuilder:validation:Optional
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type AuthTokenInitParameters struct {

	// The prefix of the token auth.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the token for integrating.
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

type AuthTokenObservation struct {

	// The prefix of the token auth.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type AuthTokenParameters struct {

	// The prefix of the token auth.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the token for integrating.
	// +kubebuilder:validation:Optional
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`
}

type DestinationInitParameters struct {

	// Determines the New Relic account where the notification destination will be created. Defaults to the account associated with the API key used.
	// The account ID under which to put the destination.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Indicates whether the destination is active.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// A nested block that describes a basic username and password authentication credentials. Only one auth_basic block is permitted per notification destination definition.  See Nested auth_basic blocks below for details.
	// Basic username and password authentication credentials.
	AuthBasic []AuthBasicInitParameters `json:"authBasic,omitempty" tf:"auth_basic,omitempty"`

	// A nested block that describes a custom header authentication credentials. Multiple blocks are permitted per notification destination definition. Nested auth_custom_header blocks below for details.
	// Custom header based authentication
	AuthCustomHeader []AuthCustomHeaderInitParameters `json:"authCustomHeader,omitempty" tf:"auth_custom_header,omitempty"`

	// A nested block that describes a token authentication credentials. Only one auth_token block is permitted per notification destination definition.  See Nested auth_token blocks below for details.
	// Token authentication credentials.
	AuthToken []AuthTokenInitParameters `json:"authToken,omitempty" tf:"auth_token,omitempty"`

	// The name of the destination.
	// (Required) The name of the destination.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A nested block that describes a notification destination property. See Nested property blocks below for details.
	// Notification destination property type.
	Property []DestinationPropertyInitParameters `json:"property,omitempty" tf:"property,omitempty"`

	// A nested block that describes a URL that contains sensitive data at the path or parameters. Only one secure_url block is permitted per notification destination definition. See Nested secure_url blocks below for details.
	// URL in secure format
	SecureURL []SecureURLInitParameters `json:"secureUrl,omitempty" tf:"secure_url,omitempty"`

	// The type of destination.  One of: EMAIL, SERVICE_NOW, SERVICE_NOW_APP, WEBHOOK, JIRA, MOBILE_PUSH, EVENT_BRIDGE, PAGERDUTY_ACCOUNT_INTEGRATION or PAGERDUTY_SERVICE_INTEGRATION.
	// (Required) The type of the destination. One of: (WEBHOOK, EMAIL, SERVICE_NOW, SERVICE_NOW_APP, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA, SLACK, SLACK_COLLABORATION, SLACK_LEGACY, MOBILE_PUSH, EVENT_BRIDGE).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DestinationObservation struct {

	// Determines the New Relic account where the notification destination will be created. Defaults to the account associated with the API key used.
	// The account ID under which to put the destination.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Indicates whether the destination is active.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// A nested block that describes a basic username and password authentication credentials. Only one auth_basic block is permitted per notification destination definition.  See Nested auth_basic blocks below for details.
	// Basic username and password authentication credentials.
	AuthBasic []AuthBasicObservation `json:"authBasic,omitempty" tf:"auth_basic,omitempty"`

	// A nested block that describes a custom header authentication credentials. Multiple blocks are permitted per notification destination definition. Nested auth_custom_header blocks below for details.
	// Custom header based authentication
	AuthCustomHeader []AuthCustomHeaderObservation `json:"authCustomHeader,omitempty" tf:"auth_custom_header,omitempty"`

	// A nested block that describes a token authentication credentials. Only one auth_token block is permitted per notification destination definition.  See Nested auth_token blocks below for details.
	// Token authentication credentials.
	AuthToken []AuthTokenObservation `json:"authToken,omitempty" tf:"auth_token,omitempty"`

	// The unique entity identifier of the destination in New Relic.
	// Destination entity GUID
	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	// The ID of the destination.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The last time a notification was sent.
	LastSent *string `json:"lastSent,omitempty" tf:"last_sent,omitempty"`

	// The name of the destination.
	// (Required) The name of the destination.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A nested block that describes a notification destination property. See Nested property blocks below for details.
	// Notification destination property type.
	Property []DestinationPropertyObservation `json:"property,omitempty" tf:"property,omitempty"`

	// A nested block that describes a URL that contains sensitive data at the path or parameters. Only one secure_url block is permitted per notification destination definition. See Nested secure_url blocks below for details.
	// URL in secure format
	SecureURL []SecureURLObservation `json:"secureUrl,omitempty" tf:"secure_url,omitempty"`

	// The status of the destination.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The type of destination.  One of: EMAIL, SERVICE_NOW, SERVICE_NOW_APP, WEBHOOK, JIRA, MOBILE_PUSH, EVENT_BRIDGE, PAGERDUTY_ACCOUNT_INTEGRATION or PAGERDUTY_SERVICE_INTEGRATION.
	// (Required) The type of the destination. One of: (WEBHOOK, EMAIL, SERVICE_NOW, SERVICE_NOW_APP, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA, SLACK, SLACK_COLLABORATION, SLACK_LEGACY, MOBILE_PUSH, EVENT_BRIDGE).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DestinationParameters struct {

	// Determines the New Relic account where the notification destination will be created. Defaults to the account associated with the API key used.
	// The account ID under which to put the destination.
	// +kubebuilder:validation:Optional
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Indicates whether the destination is active.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// A nested block that describes a basic username and password authentication credentials. Only one auth_basic block is permitted per notification destination definition.  See Nested auth_basic blocks below for details.
	// Basic username and password authentication credentials.
	// +kubebuilder:validation:Optional
	AuthBasic []AuthBasicParameters `json:"authBasic,omitempty" tf:"auth_basic,omitempty"`

	// A nested block that describes a custom header authentication credentials. Multiple blocks are permitted per notification destination definition. Nested auth_custom_header blocks below for details.
	// Custom header based authentication
	// +kubebuilder:validation:Optional
	AuthCustomHeader []AuthCustomHeaderParameters `json:"authCustomHeader,omitempty" tf:"auth_custom_header,omitempty"`

	// A nested block that describes a token authentication credentials. Only one auth_token block is permitted per notification destination definition.  See Nested auth_token blocks below for details.
	// Token authentication credentials.
	// +kubebuilder:validation:Optional
	AuthToken []AuthTokenParameters `json:"authToken,omitempty" tf:"auth_token,omitempty"`

	// The name of the destination.
	// (Required) The name of the destination.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A nested block that describes a notification destination property. See Nested property blocks below for details.
	// Notification destination property type.
	// +kubebuilder:validation:Optional
	Property []DestinationPropertyParameters `json:"property,omitempty" tf:"property,omitempty"`

	// A nested block that describes a URL that contains sensitive data at the path or parameters. Only one secure_url block is permitted per notification destination definition. See Nested secure_url blocks below for details.
	// URL in secure format
	// +kubebuilder:validation:Optional
	SecureURL []SecureURLParameters `json:"secureUrl,omitempty" tf:"secure_url,omitempty"`

	// The type of destination.  One of: EMAIL, SERVICE_NOW, SERVICE_NOW_APP, WEBHOOK, JIRA, MOBILE_PUSH, EVENT_BRIDGE, PAGERDUTY_ACCOUNT_INTEGRATION or PAGERDUTY_SERVICE_INTEGRATION.
	// (Required) The type of the destination. One of: (WEBHOOK, EMAIL, SERVICE_NOW, SERVICE_NOW_APP, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA, SLACK, SLACK_COLLABORATION, SLACK_LEGACY, MOBILE_PUSH, EVENT_BRIDGE).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DestinationPropertyInitParameters struct {

	// The notification property display value.
	// Notification property display key.
	DisplayValue *string `json:"displayValue,omitempty" tf:"display_value,omitempty"`

	// The notification property key.
	// Notification property key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The notification property label.
	// Notification property label.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The notification property value.
	// Notification property value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DestinationPropertyObservation struct {

	// The notification property display value.
	// Notification property display key.
	DisplayValue *string `json:"displayValue,omitempty" tf:"display_value,omitempty"`

	// The notification property key.
	// Notification property key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The notification property label.
	// Notification property label.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The notification property value.
	// Notification property value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DestinationPropertyParameters struct {

	// The notification property display value.
	// Notification property display key.
	// +kubebuilder:validation:Optional
	DisplayValue *string `json:"displayValue,omitempty" tf:"display_value,omitempty"`

	// The notification property key.
	// Notification property key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The notification property label.
	// Notification property label.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The notification property value.
	// Notification property value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type SecureURLInitParameters struct {

	// The prefix of the URL.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The suffix of the URL, which contains sensitive data.
	SecureSuffixSecretRef v1.SecretKeySelector `json:"secureSuffixSecretRef" tf:"-"`
}

type SecureURLObservation struct {

	// The prefix of the URL.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SecureURLParameters struct {

	// The prefix of the URL.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// The suffix of the URL, which contains sensitive data.
	// +kubebuilder:validation:Optional
	SecureSuffixSecretRef v1.SecretKeySelector `json:"secureSuffixSecretRef" tf:"-"`
}

// DestinationSpec defines the desired state of Destination
type DestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DestinationParameters `json:"forProvider"`
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
	InitProvider DestinationInitParameters `json:"initProvider,omitempty"`
}

// DestinationStatus defines the observed state of Destination.
type DestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Destination is the Schema for the Destinations API. Create and manage a notification destination for notifications in New Relic.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,newrelic}
type Destination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.property) || (has(self.initProvider) && has(self.initProvider.property))",message="spec.forProvider.property is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   DestinationSpec   `json:"spec"`
	Status DestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DestinationList contains a list of Destinations
type DestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Destination `json:"items"`
}

// Repository type metadata.
var (
	Destination_Kind             = "Destination"
	Destination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Destination_Kind}.String()
	Destination_KindAPIVersion   = Destination_Kind + "." + CRDGroupVersion.String()
	Destination_GroupVersionKind = CRDGroupVersion.WithKind(Destination_Kind)
)

func init() {
	SchemeBuilder.Register(&Destination{}, &DestinationList{})
}
