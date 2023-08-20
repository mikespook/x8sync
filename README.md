# x8sync

x8sync is a utility designed to facilitate data synchronization between 
two Kubernetes clusters. The name is derived from "cross-great-sync."

# Install

__TODO__

# Usage

```
x8sync [flags] <source-pvc-path> <target-pvc-path>
```

## PVC Path

The PVC path contains three crucial elements that denote the PVC resources in the cluster:

1. Cluster context name
2. Namespace
3. PVC name

For example: `$(cluster-context-name)/$(namespace)/$(pvc-name)`

## Flags

### `--name`


This flag assigns a name to a newly created syncing task.
The default value is a randomly generated UUID.

If the provided name already exists within the clusters, 
x8sync will identify the ongoing progress and continue accordingly. 
In case any errors are detected, it will provide a status update for 
the syncing process.

### `--force`

Using this flag will forcibly remove all data from the target PVC.
Please note that this action could be potentially risky and result in data loss.

### `--dry-run`

Enabling this flag will initiate a simulated syncing process that doesn't 
modify any data. However, all the necessary resources for the syncing 
operation will be allocated. This feature is useful for identifying potential
issues without affecting actual data.

### `--verbose`

Enabling this flag will enhance the program's verbosity, providing additional
details throughout the process. This increased level of detail proves invaluable
for identifying potential issues within the Kubernetes environment.

# Examples

* Create syncing task
```
x8sync <source-uri> <target-uri>
```

* Create syncing task with specific name
```
x8sync --name="foobar" <source-uri> <target-uri>
```

# Mecanisim

The x8sync tool operates according to the following steps, enabling 
the creation of necessary resources in both Kubernetes clusters and 
the synchronization of data from the source Persistent Volume Claim (PVC)
to the target PVC.

1. Validate the source PVC.
2. Validate the target PVC.
3. Check for name conflicts in the source Kubernetes cluster.
4. Check for name conflicts in the target Kubernetes cluster.
5. Establish a StatefulSet and Service for the OpenSSH server in the source 
Kubernetes cluster.
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
