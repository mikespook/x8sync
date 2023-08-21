package k8s

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetContexts(kubeconfig string) (map[string]*api.Context, error) {
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		return nil, err
	}
	return config.Contexts, nil
}
