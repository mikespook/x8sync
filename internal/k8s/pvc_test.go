package k8s

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestHasPVC(t *testing.T) {
	fakeObjs := []runtime.Object{
		&corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "foo",
			},
		},
	}
	clientset := fake.NewSimpleClientset(fakeObjs...)

	t.Run("having pvc", func(t *testing.T) {
		pvc, err := GetPVC(clientset, "foo", "bar")
		if errors.IsNotFound(err) {
			t.Fatal("pvc is missing")
		}
		if err != nil {
			t.Fatal(err)
		}

		if IsPVCBound(pvc) {
			t.Fatal("pvc is bound")
		}
	})

	t.Run("not having pvc", func(t *testing.T) {
		_, err := GetPVC(clientset, "bar", "non-existent")
		if err != nil {
			if errors.IsNotFound(err) {
				return
			} else {
				t.Fatal(err)
			}
		}
		t.Fatal("pvc is existing")
	})

	clientset.CoreV1().(*fakecorev1.FakeCoreV1).PrependReactor("get", "persistentvolumeclaims", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &corev1.PersistentVolumeClaim{}, errors.NewResourceExpired("an expected error")
	})

	t.Run("triggering error", func(t *testing.T) {
		_, err := GetPVC(clientset, "foo", "bar")
		if err == nil {
			t.Fatal("error is expected")
		}
	})

}
