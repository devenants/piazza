package v1

type IdentifierPair struct {
	Head ResourceMeta `yaml:"head,omitempty"`
	Tail ResourceMeta `yaml:"tail,omitempty"`
}

type IdentifierListPair struct {
	Head []*ResourceMeta `yaml:"head,omitempty"`
	Tail []*ResourceMeta `yaml:"tail,omitempty"`
}
