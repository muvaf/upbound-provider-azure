/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SQLRoleAssignmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLRoleAssignmentParameters struct {

	// +crossplane:generate:reference:type=Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	RoleDefinitionID *string `json:"roleDefinitionId" tf:"role_definition_id,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`
}

// SQLRoleAssignmentSpec defines the desired state of SQLRoleAssignment
type SQLRoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLRoleAssignmentParameters `json:"forProvider"`
}

// SQLRoleAssignmentStatus defines the observed state of SQLRoleAssignment.
type SQLRoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLRoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleAssignment is the Schema for the SQLRoleAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLRoleAssignmentSpec   `json:"spec"`
	Status            SQLRoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleAssignmentList contains a list of SQLRoleAssignments
type SQLRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLRoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	SQLRoleAssignment_Kind             = "SQLRoleAssignment"
	SQLRoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLRoleAssignment_Kind}.String()
	SQLRoleAssignment_KindAPIVersion   = SQLRoleAssignment_Kind + "." + CRDGroupVersion.String()
	SQLRoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(SQLRoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLRoleAssignment{}, &SQLRoleAssignmentList{})
}
