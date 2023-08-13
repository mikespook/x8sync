# x8sync

x8sync is a utility designed to facilitate data synchronization between 
two Kubernetes clusters. The name is derived from "cross-great-sync."

# Install

__TODO__

Usage

__TODO__

# Mecanisim

The x8sync tool operates according to the following steps, enabling 
the creation of necessary resources in both Kubernetes clusters and 
the synchronization of data from the source Persistent Volume Claim (PVC)
to the target PVC.

1. Validate the source PVC.
2. Validate the target PVC.
3. Check for name conflicts in the source Kubernetes cluster.
4. Check for name conflicts in the target Kubernetes cluster.
5. Establish a StatefulSet and Service for the OpenSSH server in the source Kubernetes cluster.
6. Generate a Job in the target Kubernetes cluster to synchronize the data.
7. Monitor the synchronization progress.
8. Perform resource cleanup after synchronization.


Certain prerequisites are necessary at the outset:

* Generate SSH private/public key pairs.
* Ensure access to Kubernetes cluster contexts.
* Maintain synchronization task progress.
* Facilitate the creation of multiple tasks.

# Licence

# Authors

* Xing Xing <mikespook@gmail.com> [Blog](http://mikespook.com) 
[@Twitter](http://twitter.com/mikespook)

# Open Source - MIT Software License

See [LICENSE](LICENSE).
