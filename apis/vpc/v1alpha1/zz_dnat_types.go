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

type DnatObservation struct {

	// the description of this dnat rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the Subnet. Notes: Because of there is one resource in the association, conflict with `network_interface_id`.
	DnatID *string `json:"dnatId,omitempty" tf:"dnat_id,omitempty"`

	// the name of dnat rule.
	DnatName *string `json:"dnatName,omitempty" tf:"dnat_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// the protocol of dnat port, Valid Options: `Any`, `TCP` and `UDP`. <br> Notes: `public_port` and `private_port` must be set as `Any`, when `ip_protocol` is `Any`. Instead, you should set ports.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The id of the Nat.
	NATID *string `json:"natId,omitempty" tf:"nat_id,omitempty"`

	// the nat ip of nat.
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip,omitempty"`

	// the private ip of instance in the identical vpc.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// the private port that be accessed in vpc.
	PrivatePort *string `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// the public port of the internet.
	PublicPort *string `json:"publicPort,omitempty" tf:"public_port,omitempty"`
}

type DnatParameters struct {

	// the description of this dnat rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the name of dnat rule.
	// +kubebuilder:validation:Optional
	DnatName *string `json:"dnatName,omitempty" tf:"dnat_name,omitempty"`

	// the protocol of dnat port, Valid Options: `Any`, `TCP` and `UDP`. <br> Notes: `public_port` and `private_port` must be set as `Any`, when `ip_protocol` is `Any`. Instead, you should set ports.
	// +kubebuilder:validation:Optional
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The id of the Nat.
	// +crossplane:generate:reference:type=github.com/kingsoftcloud/provider-ksyun/apis/vpc/v1alpha1.Nat
	// +kubebuilder:validation:Optional
	NATID *string `json:"natId,omitempty" tf:"nat_id,omitempty"`

	// Reference to a Nat in vpc to populate natId.
	// +kubebuilder:validation:Optional
	NATIDRef *v1.Reference `json:"natIdRef,omitempty" tf:"-"`

	// Selector for a Nat in vpc to populate natId.
	// +kubebuilder:validation:Optional
	NATIDSelector *v1.Selector `json:"natIdSelector,omitempty" tf:"-"`

	// the nat ip of nat.
	// +kubebuilder:validation:Optional
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip,omitempty"`

	// the private ip of instance in the identical vpc.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// the private port that be accessed in vpc.
	// +kubebuilder:validation:Optional
	PrivatePort *string `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// the public port of the internet.
	// +kubebuilder:validation:Optional
	PublicPort *string `json:"publicPort,omitempty" tf:"public_port,omitempty"`
}

// DnatSpec defines the desired state of Dnat
type DnatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DnatParameters `json:"forProvider"`
}

// DnatStatus defines the observed state of Dnat.
type DnatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DnatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dnat is the Schema for the Dnats API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ksyun}
type Dnat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ipProtocol)",message="ipProtocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.natIp)",message="natIp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.privateIpAddress)",message="privateIpAddress is a required parameter"
	Spec   DnatSpec   `json:"spec"`
	Status DnatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnatList contains a list of Dnats
type DnatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dnat `json:"items"`
}

// Repository type metadata.
var (
	Dnat_Kind             = "Dnat"
	Dnat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dnat_Kind}.String()
	Dnat_KindAPIVersion   = Dnat_Kind + "." + CRDGroupVersion.String()
	Dnat_GroupVersionKind = CRDGroupVersion.WithKind(Dnat_Kind)
)

func init() {
	SchemeBuilder.Register(&Dnat{}, &DnatList{})
}
