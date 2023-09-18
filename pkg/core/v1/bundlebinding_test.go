package v1

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestBundleBinding(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: BundleBinding
metadata:
  name: bundlebinding-test
spec:
  type: match
  match: /
  bundle: default
status:
  hits: 1000
`
	var config BundleBinding
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}

func TestBundleBindingList(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: BundleBindingList
metadata:
  name: bundlebinding-list-test
items:
- spec:
    type: match
    match: /
    bundle: default
  status:
    hits: 1000
`
	var config BundleBindingList
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}
