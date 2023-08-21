package k8s

import (
	"errors"
	"testing"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	fakebatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1/fake"
	k8stesting "k8s.io/client-go/testing"
)

func TestHasJob(t *testing.T) {
	fakeObjs := []runtime.Object{
		&batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "foo",
			},
		},
	}
	clientset := fake.NewSimpleClientset(fakeObjs...)

	t.Run("having job", func(t *testing.T) {
		has, err := HasJob(clientset, "foo", "bar")
		if err != nil {
			t.Fatal(err)
		}
		if !has {
			t.Fatal("job is missing")
		}
	})

	t.Run("not having job", func(t *testing.T) {
		has, err := HasJob(clientset, "bar", "non-existent")
		if err != nil {
			t.Fatal(err)
		}
		if has {
			t.Fatal("job is existing")
		}
	})

	clientset.BatchV1().(*fakebatchv1.FakeBatchV1).PrependReactor("get", "jobs", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &batchv1.Job{}, errors.New("an expected error")
	})

	t.Run("triggering error", func(t *testing.T) {
		has, err := HasJob(clientset, "foo", "bar")
		if err == nil {
			t.Fatal("error is expected")
		}
		if has {
			t.Fatal("false is expected on error")
		}
	})

}
