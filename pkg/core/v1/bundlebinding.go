package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type BundleBindingSpec struct {
	Type   string `yaml:"type,omitempty"`
	Match  string `yaml:"match,omitempty"`
	Bundle string `yaml:"bundle,omitempty"`
}

type BundleBindingStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type BundleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   BundleBindingSpec   `yaml:"spec,omitempty"`
	Status BundleBindingStatus `yaml:"status,omitempty"`
}

type BundleBindingList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []BundleBinding `yaml:"items"`
}
