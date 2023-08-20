package context

import (
	"os"
	"path/filepath"
	"testing"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestBuildRESTConfig(t *testing.T) {
	home := mustInitTestHome()
	kubeconfig := filepath.Join(home, ".kube", "config")
	saveTestConfig(kubeconfig, createValidTestConfig)

	t.Run("creating api.Config from the current context", func(t *testing.T) {
		config, err := BuildRESTConfig("", kubeconfig)
		if err != nil {
			t.Fatal(err)
		}
		if config.Host != FooServer {
			t.Fatalf("%s expected, got %s", FooServer, config.Host)
		}
		t.Log(config.String())
	})

	t.Run("creating api.Config from the specific context", func(t *testing.T) {
		config, err := BuildRESTConfig("bar", kubeconfig)
		if err != nil {
			t.Fatal(err)
		}
		if config.Host != BarServer {
			t.Fatalf("%s expected, got %s", BarServer, config.Host)
		}
		t.Log(config.String())
	})

	t.Run("creating api.Config from a non-existent context", func(t *testing.T) {
		config, err := BuildRESTConfig("non-existent", kubeconfig)
		if err == nil {
			t.Fatalf("an error expected")
		}
		if config != nil {
			t.Fatalf("nil value expected")
		}
	})

	t.Cleanup(func() {
		mustCleanup(home)
	})
}

const (
	FooServer = "https://foo:8080"
	FooToken  = "foo-token"
	BarServer = "https://bar:8080"
	BarToken  = "bar=token"
)

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
