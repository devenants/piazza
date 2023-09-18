package labels

type Labels interface {
	Has(label string) (exists bool)
	Get(label string) (value string)
}

type Set map[string]string
