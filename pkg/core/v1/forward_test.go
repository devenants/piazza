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
  config:
  - match: /default/default/default
    weight: 1
    policy: default
    endpoints: 
    - name: default
      groupe: default
      segment: default
      capacity: default
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
- spec:
    config:
    - match: /default/default/default
      weight: 1
      policy: default
      endpoints: 
      - name: default
        groupe: default
        segment: default
        capacity: default
  status:
    hits: 1000
`

	var config ForwardList
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}
