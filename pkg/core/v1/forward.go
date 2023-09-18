package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type TableDescriptor struct {
	Name    string `yaml:"name,omitempty"`
	Groupe  string `yaml:"groupe,omitempty"`
	Segment string `yaml:"segment,omitempty"`
	Weight  string `yaml:"weight,omitempty"`
}

type ForwardTemplate struct {
	Match     string            `yaml:"match,omitempty"`
	Policy    string            `yaml:"policy,omitempty"`
	Endpoints []TableDescriptor `yaml:"endpoints,omitempty"`
}

type ForwardBundle struct {
	Name string `yaml:"name,omitempty"`
}

type ForwardSpec struct {
	Bundles []*ForwardBundle   `yaml:"bundles,omitempty"`
	Config  []*ForwardTemplate `yaml:"config,omitempty"`
}

type ForwardStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Forward struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   ForwardSpec   `yaml:"spec,omitempty"`
	Status ForwardStatus `yaml:"status,omitempty"`
}

type ForwardList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []Forward `yaml:"items"`
}
