package k8s

import (
	"k8s.io/client-go/kubernetes"
)

var (
	ClientSet *kubernetes.Clientset
	IPCache   = make(map[string]map[string]string)
)
