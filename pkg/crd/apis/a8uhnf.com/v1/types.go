package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// pkg/crd/jinghzhu/v1/types.go

// Jinghzhu is the CRD. Use this command to generate deepcopy for it:
type Jinghzhu struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata"`
	// Specification of the desired behavior of Jinghzhu.
	Spec JinghzhuSpec `json:"spec"`
	// Observed status of Jinghzhu.
	Status JinghzhuStatus `json:"status"`
}

// JinghzhuSpec is a desired state description of Jinghzhu.
type JinghzhuSpec struct {
	Foo string `json:"foo"`
	Bar bool   `json:"bar"`
}

// JinghzhuStatus describes the lifecycle status of Jinghzhu.
type JinghzhuStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}
