package k8s

type Metadata struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
}
