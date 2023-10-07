package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type ClusterTemplate struct {
	Slot              string     `yaml:"slot,omitempty"`
	Segment           string     `yaml:"segment,omitempty"`
	ConenctTimeout    int32      `yaml:"connecttimeout,omitempty"`
	LoadBalancePolicy int32      `yaml:"loadbalancepolicy,omitempty"`
	Hosts             []HostAddr `yaml:"hosts,omitempty"`
}

type ClusterSpec struct {
	ClusterTemplate `json:",inline"`
}

type ClusterStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   ClusterSpec   `yaml:"spec,omitempty"`
	Status ClusterStatus `yaml:"status,omitempty"`
}

type ClusterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []Cluster `yaml:"items"`
}
