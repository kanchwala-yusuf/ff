package k8s

import (
	"log"
)

func init() {

	// Create k8s client
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Fatalf("Failed to create k8s client. Error: '%s'", err.Error())
	}
	ClientSet = clientSet

	// Periodically update IPPodMap
	go PeriodicUpdateIPPodMap()
}
