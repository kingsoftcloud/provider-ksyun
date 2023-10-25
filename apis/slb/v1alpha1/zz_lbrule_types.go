/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LbRuleHealthCheckObservation struct {

	// ID of the health check.
	HealthCheckID *string `json:"healthCheckId,omitempty" tf:"health_check_id,omitempty"`

	// Status maintained by health examination.Valid Values:'start', 'stop'.
	HealthCheckState *string `json:"healthCheckState,omitempty" tf:"health_check_state,omitempty"`

	// Health threshold.Valid Values:1-10. Default is 5.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The service host name of the health check, which is available only for the HTTP or HTTPS health check.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Interval of health examination.Valid Values:1-3600. Default is 5.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Whether the host name is default or not.
	IsDefaultHostName *bool `json:"isDefaultHostName,omitempty" tf:"is_default_host_name,omitempty"`

	// protocol of the listener.
	ListenerProtocol *string `json:"listenerProtocol,omitempty" tf:"listener_protocol,omitempty"`

	// Health check timeout.Valid Values:1-3600. Default is 4.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Link to HTTP type listener health check.
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	// Unhealthy threshold.Valid Values:1-10. Default is 4.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LbRuleHealthCheckParameters struct {

	// Status maintained by health examination.Valid Values:'start', 'stop'.
	// +kubebuilder:validation:Optional
	HealthCheckState *string `json:"healthCheckState,omitempty" tf:"health_check_state,omitempty"`

	// Health threshold.Valid Values:1-10. Default is 5.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The service host name of the health check, which is available only for the HTTP or HTTPS health check.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// Interval of health examination.Valid Values:1-3600. Default is 5.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Whether the host name is default or not.
	// +kubebuilder:validation:Optional
	IsDefaultHostName *bool `json:"isDefaultHostName,omitempty" tf:"is_default_host_name,omitempty"`

	// Health check timeout.Valid Values:1-3600. Default is 4.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Link to HTTP type listener health check.
	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	// Unhealthy threshold.Valid Values:1-10. Default is 4.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LbRuleObservation struct {

	// The id of backend server group.
	BackendServerGroupID *string `json:"backendServerGroupId,omitempty" tf:"backend_server_group_id,omitempty"`

	// the creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	HealthCheck []LbRuleHealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The id of host header id.
	HostHeaderID *string `json:"hostHeaderId,omitempty" tf:"host_header_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to synchronizethe the health check, the session hold and the forward algorithms of the listener.Valid Values:'on', 'off'. Default is 'on'.
	ListenerSync *string `json:"listenerSync,omitempty" tf:"listener_sync,omitempty"`

	// Forwarding mode of listener.Valid Values:'RoundRobin', 'LeastConnections'. Default is 'RoundRobin'.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The path of rule.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The ID of the rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	Session []LbRuleSessionObservation `json:"session,omitempty" tf:"session,omitempty"`
}

type LbRuleParameters struct {

	// The id of backend server group.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbBackendServerGroup
	// +kubebuilder:validation:Optional
	BackendServerGroupID *string `json:"backendServerGroupId,omitempty" tf:"backend_server_group_id,omitempty"`

	// Reference to a LbBackendServerGroup in slb to populate backendServerGroupId.
	// +kubebuilder:validation:Optional
	BackendServerGroupIDRef *v1.Reference `json:"backendServerGroupIdRef,omitempty" tf:"-"`

	// Selector for a LbBackendServerGroup in slb to populate backendServerGroupId.
	// +kubebuilder:validation:Optional
	BackendServerGroupIDSelector *v1.Selector `json:"backendServerGroupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HealthCheck []LbRuleHealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The id of host header id.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/slb/v1alpha1.LbHostHeader
	// +kubebuilder:validation:Optional
	HostHeaderID *string `json:"hostHeaderId,omitempty" tf:"host_header_id,omitempty"`

	// Reference to a LbHostHeader in slb to populate hostHeaderId.
	// +kubebuilder:validation:Optional
	HostHeaderIDRef *v1.Reference `json:"hostHeaderIdRef,omitempty" tf:"-"`

	// Selector for a LbHostHeader in slb to populate hostHeaderId.
	// +kubebuilder:validation:Optional
	HostHeaderIDSelector *v1.Selector `json:"hostHeaderIdSelector,omitempty" tf:"-"`

	// Whether to synchronizethe the health check, the session hold and the forward algorithms of the listener.Valid Values:'on', 'off'. Default is 'on'.
	// +kubebuilder:validation:Optional
	ListenerSync *string `json:"listenerSync,omitempty" tf:"listener_sync,omitempty"`

	// Forwarding mode of listener.Valid Values:'RoundRobin', 'LeastConnections'. Default is 'RoundRobin'.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// The path of rule.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Session []LbRuleSessionParameters `json:"session,omitempty" tf:"session,omitempty"`
}

type LbRuleSessionObservation struct {

	// The name of cookie.The CookieType is valid and required when it is 'RewriteCookie'; otherwise, this value is ignored.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The type of the cookie.Valid Values:'ImplantCookie', 'RewriteCookie'. Default is 'ImplantCookie'.
	CookieType *string `json:"cookieType,omitempty" tf:"cookie_type,omitempty"`

	// Session hold timeout.Valid Values:1-86400. Default is '7200'.
	SessionPersistencePeriod *float64 `json:"sessionPersistencePeriod,omitempty" tf:"session_persistence_period,omitempty"`

	// The state of session.Valid Values:'start', 'stop'. Default is 'start'.
	SessionState *string `json:"sessionState,omitempty" tf:"session_state,omitempty"`
}

type LbRuleSessionParameters struct {

	// The name of cookie.The CookieType is valid and required when it is 'RewriteCookie'; otherwise, this value is ignored.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The type of the cookie.Valid Values:'ImplantCookie', 'RewriteCookie'. Default is 'ImplantCookie'.
	// +kubebuilder:validation:Optional
	CookieType *string `json:"cookieType,omitempty" tf:"cookie_type,omitempty"`

	// Session hold timeout.Valid Values:1-86400. Default is '7200'.
	// +kubebuilder:validation:Optional
	SessionPersistencePeriod *float64 `json:"sessionPersistencePeriod,omitempty" tf:"session_persistence_period,omitempty"`

	// The state of session.Valid Values:'start', 'stop'. Default is 'start'.
	// +kubebuilder:validation:Optional
	SessionState *string `json:"sessionState,omitempty" tf:"session_state,omitempty"`
}

// LbRuleSpec defines the desired state of LbRule
type LbRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbRuleParameters `json:"forProvider"`
}

// LbRuleStatus defines the observed state of LbRule.
type LbRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbRule is the Schema for the LbRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type LbRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.path)",message="path is a required parameter"
	Spec   LbRuleSpec   `json:"spec"`
	Status LbRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbRuleList contains a list of LbRules
type LbRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbRule `json:"items"`
}

// Repository type metadata.
var (
	LbRule_Kind             = "LbRule"
	LbRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LbRule_Kind}.String()
	LbRule_KindAPIVersion   = LbRule_Kind + "." + CRDGroupVersion.String()
	LbRule_GroupVersionKind = CRDGroupVersion.WithKind(LbRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LbRule{}, &LbRuleList{})
}