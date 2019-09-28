package config

import (
	"os"
)

var (
	KUBECONFIG = os.Getenv("KUBECONFIG")
)
