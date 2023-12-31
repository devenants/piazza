package v1

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBundle(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: Bundle
metadata:
  name: bundle-test
status:
  hits: 1000
`

	var config Bundle
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}

func TestBundleList(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: BundleList
metadata:
  name: bundle-list-test
items:
- status:
    hits: 1000
`

	var config BundleList
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
		return
	}
}
