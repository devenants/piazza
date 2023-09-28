package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type BundleBindingSelector struct {
	Names         []string `json:"names,omitempty"`
	LabelSelector `json:",inline"`
	Mask          string `yaml:"mask,omitempty"`
}

type BundleBindingSpec struct {
	Selector *BundleBindingSelector `json:"selector,omitempty"`
	Bundles  []*metav1.ResourceMeta `yaml:"bundles,omitempty"`
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
