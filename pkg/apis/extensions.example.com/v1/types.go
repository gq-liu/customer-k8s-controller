package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Website describes a website.
type Website struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WebsiteSpec `json:"spec"`
	Status WebsiteStatus `json:"status"`
}


// WebsiteSpec is the spec for a Website resource
type WebsiteSpec struct {
	Url         string `json:"url"`
	Replicas    *int32 `json:"replicas"`
}

// WebsiteStatus is the status for a Foo resource
type WebsiteStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebsiteList is a list of Website resources
type WebsiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Website `json:"items"`
}


