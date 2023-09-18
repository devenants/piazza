package v1

type TypeMeta struct {
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

type ObjectMeta struct {
	Cluster   string            `yaml:"cluster,omitempty"`
	Namespace string            `yaml:"namespace,omitempty"`
	Name      string            `yaml:"name,omitempty"`
	Labels    map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`
}

type ResourceMeta struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
}
