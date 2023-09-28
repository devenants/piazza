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
  match: /
  items:
  - spec:
      weight: 100
      config:
        head:
        - metadata:
            name: default
            cluster: default
            namespace: default
        tail:
        - metadata:
            name: default
            cluster: default
            namespace: default
    status:
      hits: 1000
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
    match: /
    items:
    - spec:
        weight: 100
        config:
          head:
          - metadata:
              name: default
              cluster: default
              namespace: default
          tail:
          - metadata:
              name: default
              cluster: default
              namespace: default
      status:
        hits: 1000
  status:
    hits: 1000
`

	var config ForwardList
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		t.Fatalf("Failed to Unmarshal: %v", err)
	}
}
