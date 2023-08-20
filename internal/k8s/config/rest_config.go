package context

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// BuildRESTConfig builds a Kubernetes REST configuration based on the provided context name and kubeconfig path.
// If the contextName is empty, it uses the kubeconfig file specified by kubeconfig parameter.
// If contextName is provided, it creates a client configuration using the specified context.
func BuildRESTConfig(contextName, kubeconfig string) (*rest.Config, error) {
	if contextName == "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: contextName,
		}).ClientConfig()
}

