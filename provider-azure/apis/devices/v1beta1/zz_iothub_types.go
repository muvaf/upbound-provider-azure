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

type CloudToDeviceObservation struct {
}

type CloudToDeviceParameters struct {

	// +kubebuilder:validation:Optional
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Feedback []FeedbackParameters `json:"feedback,omitempty" tf:"feedback,omitempty"`

	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
}

type EndpointObservation struct {
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	BatchFrequencyInSeconds *float64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`

	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	EntityPath *string `json:"entityPath,omitempty" tf:"entity_path,omitempty"`

	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`

	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	MaxChunkSizeInBytes *float64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EndpointParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Container
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name"`

	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type EnrichmentObservation struct {
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnrichmentParameters struct {
}

type FallbackRouteObservation struct {
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type FallbackRouteParameters struct {
}

type FeedbackObservation struct {
}

type FeedbackParameters struct {

	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// +kubebuilder:validation:Optional
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type FileUploadObservation struct {
}

type FileUploadParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// +kubebuilder:validation:Optional
	Notifications *bool `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// +kubebuilder:validation:Optional
	SASTTL *string `json:"sasTtl,omitempty" tf:"sas_ttl,omitempty"`
}

type IOTHubObservation struct {
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	Enrichment []EnrichmentObservation `json:"enrichment,omitempty" tf:"enrichment,omitempty"`

	EventHubEventsEndpoint *string `json:"eventHubEventsEndpoint,omitempty" tf:"event_hub_events_endpoint,omitempty"`

	EventHubEventsNamespace *string `json:"eventHubEventsNamespace,omitempty" tf:"event_hub_events_namespace,omitempty"`

	EventHubEventsPath *string `json:"eventHubEventsPath,omitempty" tf:"event_hub_events_path,omitempty"`

	EventHubOperationsEndpoint *string `json:"eventHubOperationsEndpoint,omitempty" tf:"event_hub_operations_endpoint,omitempty"`

	EventHubOperationsPath *string `json:"eventHubOperationsPath,omitempty" tf:"event_hub_operations_path,omitempty"`

	FallbackRoute []FallbackRouteObservation `json:"fallbackRoute,omitempty" tf:"fallback_route,omitempty"`

	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	Route []RouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	SharedAccessPolicy []SharedAccessPolicyObservation `json:"sharedAccessPolicy,omitempty" tf:"shared_access_policy,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IOTHubParameters struct {

	// +kubebuilder:validation:Optional
	CloudToDevice []CloudToDeviceParameters `json:"cloudToDevice,omitempty" tf:"cloud_to_device,omitempty"`

	// +kubebuilder:validation:Required
	Endpoint []EndpointParameters `json:"endpoint" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	EventHubPartitionCount *float64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count,omitempty"`

	// +kubebuilder:validation:Optional
	EventHubRetentionInDays *float64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	FileUpload []FileUploadParameters `json:"fileUpload,omitempty" tf:"file_upload,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRuleSet []NetworkRuleSetParameters `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IPRuleObservation struct {
}

type IPRuleParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	IPMask *string `json:"ipMask" tf:"ip_mask,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkRuleSetObservation struct {
}

type NetworkRuleSetParameters struct {

	// +kubebuilder:validation:Optional
	ApplyToBuiltinEventHubEndpoint *bool `json:"applyToBuiltinEventhubEndpoint,omitempty" tf:"apply_to_builtin_eventhub_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	IPRule []IPRuleParameters `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
}

type RouteObservation struct {
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type RouteParameters struct {
}

type SharedAccessPolicyObservation struct {
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type SharedAccessPolicyParameters struct {
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// IOTHubSpec defines the desired state of IOTHub
type IOTHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubParameters `json:"forProvider"`
}

// IOTHubStatus defines the observed state of IOTHub.
type IOTHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHub is the Schema for the IOTHubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubSpec   `json:"spec"`
	Status            IOTHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubList contains a list of IOTHubs
type IOTHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHub `json:"items"`
}

// Repository type metadata.
var (
	IOTHub_Kind             = "IOTHub"
	IOTHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHub_Kind}.String()
	IOTHub_KindAPIVersion   = IOTHub_Kind + "." + CRDGroupVersion.String()
	IOTHub_GroupVersionKind = CRDGroupVersion.WithKind(IOTHub_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHub{}, &IOTHubList{})
}
