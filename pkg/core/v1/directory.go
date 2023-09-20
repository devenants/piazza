package v1

import (
	metav1 "github.com/devenants/piazza/pkg/meta/v1"
)

type DirectoryDescriptor struct {
	Name    string   `yaml:"name,omitempty"`
	Targets []string `yaml:"targets,omitempty"`
}

type DirectoryTemplate struct {
	Mapping     string                 `yaml:"mapping,omitempty"`
	Policy      string                 `yaml:"policy,omitempty"`
	Descriptors []*DirectoryDescriptor `yaml:"descriptors,omitempty"`
}

type DirectorySpec struct {
	Identifier string                 `yaml:"identifier,omitempty"`
	Bundles    []*metav1.ResourceMeta `yaml:"bundles,omitempty"`
	Config     DirectoryTemplate      `yaml:"config,omitempty"`
}

type DirectoryStatus struct {
	Hits string `yaml:"hits,omitempty"`
}

type Directory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`

	Spec   DirectorySpec   `yaml:"spec,omitempty"`
	Status DirectoryStatus `yaml:"status,omitempty"`
}

type DirectoryList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `yaml:"metadata,omitempty"`
	Items             []Directory `yaml:"items"`
}
