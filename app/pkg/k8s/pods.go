package k8s

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods returns all the pods present in a namespace
func GetPods(ns string) error {

	// List pods
	p, err := ClientSet.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("Failed to list pods in namespace '%s'. Error: '%s'", ns, err.Error())
		return err
	}

	for _, pod := range p.Items {
		var (
			podName   = pod.Name
			podIP     = pod.Status.PodIP
			podLabels = pod.ObjectMeta.Labels
		)
		podLabels["Name"] = podName
		IPCache[podIP] = podLabels
	}

	// Successful
	return nil
}
