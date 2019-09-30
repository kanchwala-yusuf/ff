package k8s

import (
	"os"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/ff/app/pkg/config"
)

// CreateClientSet create kubernetes client using the kubeconfig path
func CreateClientSet() (*kubernetes.Clientset, error) {

	// Check for incluster config
	kubeConfig, err := rest.InClusterConfig()
	if err != nil {

		// Check KUBECONFIG variable
		kubeconfigPath := config.KUBECONFIG
		if kubeconfigPath == "" {

			// Default KUBECONFIG path
			kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
		}

		// Build k8s config from kubeconfig file path
		kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, err
		}

		// Create clientset
		clientset, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			return nil, err
		}

		// Successfull
		return clientset, nil

	} else {

		// Create client set
		clientset, err := kubernetes.NewForConfig(kubeConfig)
		if err != nil {
			return nil, err
		}

		// Successfull
		return clientset, nil
	}
}
