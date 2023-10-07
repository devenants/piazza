package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type DestinationCluster struct {
	Names []string           `yaml:"names,omitempty"`
	Items []*ClusterTemplate `yaml:"items,omitempty"`
}

type DestinationTemplate struct {
	Mask     string               `yaml:"mask,omitempty"`
	Rewrite  string               `yaml:"rewrite,omitempty"`
	Action   string               `yaml:"action,omitempty"`
	Next     string               `yaml:"next,omitempty"`
	Clusters []DestinationCluster `yaml:"clusters,omitempty"`
}

type DestinationSpec struct {
	ForwardTemplate `json:",inline"`
}

type DestinationStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Destination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   ForwardSpec   `yaml:"spec,omitempty"`
	Status ForwardStatus `yaml:"status,omitempty"`
}

type DestinationList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []Destination `yaml:"items"`
}
