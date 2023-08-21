package k8s

import (
	"errors"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestHasService(t *testing.T) {
	fakeObjs := []runtime.Object{
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "foo",
			},
		},
	}
	clientset := fake.NewSimpleClientset(fakeObjs...)

	t.Run("having service", func(t *testing.T) {
		has, err := HasService(clientset, "foo", "bar")
		if err != nil {
			t.Fatal(err)
		}
		if !has {
			t.Fatal("service is missing")
		}
	})

	t.Run("not having service", func(t *testing.T) {
		has, err := HasService(clientset, "bar", "non-existent")
		if err != nil {
			t.Fatal(err)
		}
		if has {
			t.Fatal("service is existing")
		}
	})

	clientset.CoreV1().(*fakecorev1.FakeCoreV1).PrependReactor("get", "services", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &corev1.Service{}, errors.New("an expected error")
	})

	t.Run("triggering error", func(t *testing.T) {
		has, err := HasService(clientset, "foo", "bar")
		if err == nil {
			t.Fatal("error is expected")
		}
		if has {
			t.Fatal("false is expected on error")
		}
	})

}
