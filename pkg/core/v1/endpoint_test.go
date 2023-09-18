package v1

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestEndpoint(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: Endpoint
metadata:
  name: endpoint-test
spec:
  slot: default
  segment: default
  capacity: default
  addr: 127.0.0.1:9000
status:
  hits: 1000
`

	var config Bundle
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}

func TestEndpointList(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: EndpointList
metadata:
  name: endpoint-test
items:
- spec:
    slot: default
    segment: default
    capacity: default
    addr: 127.0.0.1:9000
  status:
    hits: 1000
`

	var config BundleList
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}
