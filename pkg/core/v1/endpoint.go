package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type EndpointCapacity struct {
	Memory  int64 `yaml:"memory,omitempty"`
	Volume  int64 `yaml:"volume,omitempty"`
	Traffic int64 `yaml:"traffic,omitempty"`
}

type EndpointTemplate struct {
	Slot     string           `yaml:"slot,omitempty"`
	Segment  string           `yaml:"segment,omitempty"`
	Capacity EndpointCapacity `yaml:"capacity,omitempty"`
	Addr     string           `yaml:"addr,omitempty"`
	Weight   int32            `yaml:"weight,omitempty"`
}

type EndpointSpec struct {
	EndpointTemplate `json:",inline"`
}

type EndpointStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   EndpointSpec   `yaml:"spec,omitempty"`
	Status EndpointStatus `yaml:"status,omitempty"`
}

type EndpointList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []Endpoint `yaml:"items"`
}
