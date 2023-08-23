package k8s

import (
	"errors"
	"testing"

	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	"github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/fake"
	fakesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/typed/volumesnapshot/v1/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8stesting "k8s.io/client-go/testing"
)

func TestVolumeSnapshot(t *testing.T) {
	fakeObjs := []runtime.Object{
		&snapshotv1.VolumeSnapshotClass{
			ObjectMeta: metav1.ObjectMeta{},
		},
	}
	clientset := fake.NewSimpleClientset(fakeObjs...)
	vscName := "vsc"
	t.Run("creating volumesnapshot", func(t *testing.T) {
		_, err := CreateVolumeSnapshot(clientset, "foo", "bar", "pvc", vscName)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("having volumesnapshot", func(t *testing.T) {
		has, err := HasVolumeSnapshot(clientset, "foo", "bar")
		if err != nil {
			t.Fatal(err)
		}
		if !has {
			t.Fatal("volumesnapshot is missing")
		}
	})

	t.Run("not having volumesnapshot", func(t *testing.T) {
		has, err := HasVolumeSnapshot(clientset, "bar", "non-existent")
		if err != nil {
			t.Fatal(err)
		}
		if has {
			t.Fatal("volumesnapshot is existing")
		}
	})

	clientset.SnapshotV1().(*fakesnapshotv1.FakeSnapshotV1).PrependReactor("get", "volumesnapshots", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		return true, &snapshotv1.VolumeSnapshot{}, errors.New("an expected error")
	})

	t.Run("triggering error", func(t *testing.T) {
		has, err := HasVolumeSnapshot(clientset, "foo", "bar")
		if err == nil {
			t.Fatal("error is expected")
		}
		if has {
			t.Fatal("false is expected on error")
		}
	})
}
