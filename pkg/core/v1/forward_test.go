package v1

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestForward(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: Forward
metadata:
  name: forward-test
spec:
  bundles:
  - name: default
  config:
  - match: /default/default/default
    policy: default
    endpoints: 
    - name: default
      slot: default
      segment: default
      weight: 10
status:
  hits: 1000
`

	var config Forward
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}

func TestForwardList(t *testing.T) {
	yamlStr := `
apiVersion: v1
kind: ForwardList
metadata:
  name: forward-list-test
items:
- apiVersion: v1
  kind: Forward
  metadata:
    name: forward-test
  spec:
    bundles:
    - name: default
    config:
    - match: /default/default/default
      policy: default
      endpoints: 
      - name: default
        slot: default
        segment: default
        weight: 10
  status:
      hits: 1000
`

	var config ForwardList
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}
