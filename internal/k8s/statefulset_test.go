package k8s

import (
	"errors"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	fakeappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestHasStatefulSet(t *testing.T) {
	fakeObjs := []runtime.Object{
		&appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "foo",
			},
		},
	}
	clientset := fake.NewSimpleClientset(fakeObjs...)

	t.Run("having statefulset", func(t *testing.T) {
		has, err := HasStatefulSet(clientset, "foo", "bar")
		if err != nil {
			t.Fatal(err)
		}
		if !has {
			t.Fatal("statefulset is missing")
		}
	})

	t.Run("not having statefulset", func(t *testing.T) {
		has, err := HasStatefulSet(clientset, "bar", "non-existent")
		if err != nil {
			t.Fatal(err)
		}
		if has {
			t.Fatal("statefulset is existing")
		}
	})

	clientset.AppsV1().(*fakeappsv1.FakeAppsV1).PrependReactor("get", "statefulsets", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &appsv1.StatefulSet{}, errors.New("an expected error")
	})

	t.Run("triggering error", func(t *testing.T) {
		has, err := HasStatefulSet(clientset, "foo", "bar")
		if err == nil {
			t.Fatal("error is expected")
		}
		if has {
			t.Fatal("false is expected on error")
		}
	})

}
