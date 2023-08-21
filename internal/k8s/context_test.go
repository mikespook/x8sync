package k8s

import (
	"testing"
)

func TestGetCotexts(t *testing.T) {
	home, kubeconfig := mustInit()

	t.Run("getting list of contexts", func(t *testing.T) {
		contexts, err := GetContexts(kubeconfig)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := contexts["foo"]; !ok {
			t.Fatal("context foo is non-existent")
		}

		if _, ok := contexts["bar"]; !ok {
			t.Fatal("context bar is non-existent")
		}
	})

	t.Cleanup(func() {
		mustCleanup(home)
	})
}
