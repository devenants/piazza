package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type BundleSelector struct {
	Names         []string `json:"names,omitempty"`
	LabelSelector `json:",inline"`
}

type BundleTemplate struct {
	Selector *BundleSelector        `json:"selector,omitempty"`
	Items    []*metav1.ResourceMeta `yaml:"items,omitempty"`
}

type BundleSpec struct {
	BundleTemplate `json:",inline"`
}

type BundleStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Bundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   BundleSpec   `yaml:"spec,omitempty"`
	Status BundleStatus `yaml:"status,omitempty"`
}

type BundleList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []Bundle `yaml:"items"`
}
