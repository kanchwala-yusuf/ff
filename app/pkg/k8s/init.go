package k8s

import (
	"log"
)

func init() {
	clientSet, err := CreateClientSet()
	if err != nil {
		log.Fatalf("Failed to create k8s client. Error: '%s'\n", err.Error())
	}
	ClientSet = clientSet
}
