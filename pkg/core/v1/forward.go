package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type DescriptorSpec struct {
	Config metav1.IdentifierListPair `yaml:"config,omitempty"`
	Weight int32                     `yaml:"weight,omitempty"`
	Action string                    `yaml:"action,omitempty"`
	Policy string                    `yaml:"policy,omitempty"`
}

type DescriptorStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Descriptor struct {
	Spec   DescriptorSpec   `yaml:"spec,omitempty"`
	Status DescriptorStatus `yaml:"status,omitempty"`
}

type ForwardTemplate struct {
	Match   string        `yaml:"match,omitempty"`
	Rewrite string        `yaml:"rewrite,omitempty"`
	Action  string        `yaml:"action,omitempty"`
	Next    string        `yaml:"next,omitempty"`
	Policy  string        `yaml:"policy,omitempty"`
	Items   []*Descriptor `yaml:"items,omitempty"`
}

type ForwardSpec struct {
	ForwardTemplate `json:",inline"`
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
