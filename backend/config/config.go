package config

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Registry
const RegistryAddress = "localhost:5000"
const RegistryUsername = "projecthub"
const RegistryPassword = "password"

// Database
const DatabaseUri = "mongodb://projecthub:password@localhost:27017/"

// Kubernetes
const K8sIngressHost = "kubernetes.docker.internal"
const K8sProjectObjPath = "/var/lib/projecthub/k8sObj"

var k8sRestConfig *rest.Config

// returns k8s config for creating client
func GetK8sRestConfig() (*rest.Config, error) {
	if k8sRestConfig == nil {
		var k8sConfigPath *string
		if home := homedir.HomeDir(); home != "" { // check if machine has home directory.
			// read kubeconfig flag. if not provided use config file $HOME/.kube/config
			k8sConfigPath = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			k8sConfigPath = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// build configuration from the config file
		k8sRestConfig, _ = clientcmd.BuildConfigFromFlags("", *k8sConfigPath)
		/* 		if err != nil { //TODO catch err
			panic(err)
		} */
	}

	return k8sRestConfig, nil
}
