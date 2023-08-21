package k8s

import (
	"testing"
)

func TestBuildRESTConfig(t *testing.T) {
	home, kubeconfig := mustInit()

	t.Run("creating api.Config from the current context", func(t *testing.T) {
		config, err := BuildRESTConfig("", kubeconfig)
		if err != nil {
			t.Fatal(err)
		}
		if config.Host != FooServer {
			t.Fatalf("%s is expected, got %s", FooServer, config.Host)
		}
		t.Log(config.String())
	})

	t.Run("creating api.Config from the specific context", func(t *testing.T) {
		config, err := BuildRESTConfig("bar", kubeconfig)
		if err != nil {
			t.Fatal(err)
		}
		if config.Host != BarServer {
			t.Fatalf("%s is expected, got %s", BarServer, config.Host)
		}
		t.Log(config.String())
	})

	t.Run("creating api.Config from a non-existent context", func(t *testing.T) {
		config, err := BuildRESTConfig("non-existent", kubeconfig)
		if err == nil {
			t.Fatalf("an error is expected")
		}
		if config != nil {
			t.Fatalf("nil value is expected")
		}
	})

	t.Cleanup(func() {
		mustCleanup(home)
	})
}
