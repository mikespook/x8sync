package k8s

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

const (
	FooServer = "https://foo:8080"
	FooToken  = "foo-token"
	BarServer = "https://bar:8080"
	BarToken  = "bar=token"
)

func mustInit() (home, kubeconfig string) {
	home = mustInitTestHome()
	kubeconfig = filepath.Join(home, ".kube", "config")
	saveTestConfig(kubeconfig, createValidTestConfig)
	return
}

func createValidTestConfig() *api.Config {
	config := api.NewConfig()
	config.Clusters["foo"] = &api.Cluster{
		Server: FooServer,
	}
	config.AuthInfos["foo"] = &api.AuthInfo{
		Token: FooToken,
	}
	config.Contexts["foo"] = &api.Context{
		Cluster:  "foo",
		AuthInfo: "foo",
	}
	config.Clusters["bar"] = &api.Cluster{
		Server: BarServer,
	}
	config.AuthInfos["bar"] = &api.AuthInfo{
		Token: BarToken,
	}

	config.Contexts["bar"] = &api.Context{
		Cluster:  "bar",
		AuthInfo: "bar",
	}
	config.CurrentContext = "foo"

	return config
}

func saveTestConfig(kubeconfig string, f func() *api.Config) {
	if err := clientcmd.WriteToFile(*f(), kubeconfig); err != nil {
		panic(err)
	}
}

func mustInitTestHome() string {
	home, err := os.MkdirTemp("", "x8sync-*")
	if err != nil {
		panic(err)
	}
	return home
}

func mustCleanup(home string) {
	if err := os.RemoveAll(home); err != nil {
		panic(err)
	}
}
