# customer-k8s-controller  
## Step 1  Create your CRD type
1. Create a CustomResourceDefinition Object.   
for example, ```resourcedefinition.yaml```'
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: websites.extensions.example.com
spec:
  group: extensions.example.com
  names:
    kind: Website
    singular: website
    plural: websites
    shortNames:
      - ws
  scope: Namespaced
  version: v1
```

2. Create doc.go  
```pkg/apis/extensions.example.com/v1/doc.go```
```go
// +k8s:deepcopy-gen=package,register

// Package v1 is the v1 version of the API.
// +groupName=extensions.example.com
package v1
```

3. Create register.go for group name 
```pkg/apis/extensions.example.com/register.go```  
```go
package extensions_example_com
const (
	GroupName = "extensions.example.com"
)
```

4. Create types.go for your crd
```pkg/apis/extensions.example.com/v1/types.go```
```go
package v1

import (
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=foo

// Website describes a website.
type Website struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WebsiteSpec `json:"spec"`
}

// DatabaseSpec is the spec for a Foo resource
type WebsiteSpec struct {
	Url         string `json:"url"`
	Replicas    *int32 `json:"replicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebsiteList is a list of Website resources
type WebsiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Website `json:"items"`
}
```

