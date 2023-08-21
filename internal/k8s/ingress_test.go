package k8s

import (
	"errors"
	"testing"

	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	fakenetworkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestHasIngress(t *testing.T) {
	fakeObjs := []runtime.Object{
		&networkv1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "foo",
			},
		},
	}
	clientset := fake.NewSimpleClientset(fakeObjs...)

	t.Run("having ingress", func(t *testing.T) {
		has, err := HasIngress(clientset, "foo", "bar")
		if err != nil {
			t.Fatal(err)
		}
		if !has {
			t.Fatal("ingress is missing")
		}
	})

	t.Run("not having ingress", func(t *testing.T) {
		has, err := HasIngress(clientset, "bar", "non-existent")
		if err != nil {
			t.Fatal(err)
		}
		if has {
			t.Fatal("ingress is existing")
		}
	})

	clientset.NetworkingV1().(*fakenetworkingv1.FakeNetworkingV1).PrependReactor("get", "ingresses", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &networkv1.Ingress{}, errors.New("an expected error")
	})

	t.Run("triggering error", func(t *testing.T) {
		has, err := HasIngress(clientset, "foo", "bar")
		if err == nil {
			t.Fatal("error is expected")
		}
		if has {
			t.Fatal("false is expected on error")
		}
	})

}
