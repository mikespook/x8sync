package k8s

import (
	"context"

	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	snapshotclientset "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
)

// HasVolumeSnapshot checks if a VolumeSnapshot with the given name exists in the specified namespace.
func HasVolumeSnapshot(clientset snapshotclientset.Interface, namespace, name string) (bool, error) {
	_, err := clientset.SnapshotV1().VolumeSnapshots(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err == nil {
		return true, nil
	}
	if errors.IsNotFound(err) {
		return false, nil
	}
	return false, err
}

// CreateVolumeSnapshot creates a new VolumeSnapshot resource using the provided parameters.
func CreateVolumeSnapshot(clientset snapshotclientset.Interface, namespace, snapshotName, pvcName, vscName string) (*snapshotv1.VolumeSnapshot, error) {
	snapshot := &snapshotv1.VolumeSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      snapshotName,
			Namespace: namespace,
		},
		Spec: snapshotv1.VolumeSnapshotSpec{
			VolumeSnapshotClassName: &vscName,
			Source: snapshotv1.VolumeSnapshotSource{
				PersistentVolumeClaimName: &pvcName,
			},
		},
	}

	result, err := clientset.SnapshotV1().
		VolumeSnapshots(namespace).
		Create(context.TODO(), snapshot, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// IsSupportingVolumeSnapshot checks if the Kubernetes cluster supports VolumeSnapshot functionality.
func IsSupportingVolumeSnapshot(clientset discovery.DiscoveryInterface) (bool, error) {
	apiResourceList, err := clientset.ServerResourcesForGroupVersion("snapshot.storage.k8s.io/v1")
	if err != nil {
		return false, err
	}

	for _, res := range apiResourceList.APIResources {
		if res.Kind == "VolumeSnapshot" {
			return true, nil
		}
	}
	return false, nil
}
