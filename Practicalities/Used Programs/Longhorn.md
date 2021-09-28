---
tags: [HA, dfs]
---
# Longhorn
Longhorn is a distributed file system and manages this on a volume level. The volumes can be claimed with a [[PV vs PVC vs SC vs Provisioner#Persistent Volume Claim|PVC]] and Longhorn will then provision the storage and will then replicate the data to your specifications in your [[#Storage Class]].

## How will it fit in my cluster
The master node will not run tasks, so it won't be storing 'live' data. I hope to find a way to use the attached storage (1TB of slow HDD) as a backup/snapshot storage of the 256GB working space on the worker nodes.

The 2 worker nodes have both an SSD attached over UASP and, if possible, enable trim on the drive. This is can be configured system-wide with [[Playbooks#Pre-Horn]].

### Drive types
Longhorn allows the use of drive and node tags to annotate drives. I am using two drive tags in the cluster: HDD and SSD.  
My SSD's has very limited space, about 180GB after reservations. I hope to use them for database applications. The first probably being [[Prometheus]]. HDD I hope to use for snapshots, backups, or even object storage. 

How to get to the UI: [[Getting to the services#Longhorn UI|Longhorn UI]]

## How to setup a persistent volume in Longhorn 
Apparently, K8s has 2 ways of defining a persistent volume. Through the fixed [[PV vs PVC vs SC vs Provisioner#Persistent Volume|Persisten Volumes]] or dynamically with [[PV vs PVC vs SC vs Provisioner#Storage Class and Provisioner|Storage Classes]].
More can be read [[PV vs PVC vs SC vs Provisioner|here]]

Longhorn can dynamically provision storage so we will be using the dynamic `StorageClass`.
### Storage Class
Longhorn uses the API for (persistent) volumes of Kubernetes. Longhorn provides a [[PV vs PVC vs SC vs Provisioner#Storage Class and Provisioner|Provisioner]] through which you can define your own `StorageClass`es:
```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: longhorn  # This is the default name, don't use this name
provisioner: driver.longhorn.io
allowVolumeExpansion: true  # Allows resizing of volumes when needed
parameters:
  numberOfReplicas: "3" # For HA.
  staleReplicaTimeout: "2880" # 48 hours in minutes, when a volume enters the ERROR state how long it will take to be deleted by the system
  fromBackup: ""  # This allows for 'imaging' 
  diskSelector: "ssd,fast"  # Will select disk based on tags
  nodeSelector: "storage,fast"  # Will select nodes based on tags
  # Both 'selectors' limit the scope of placement.
  fsType: "ext4" 
  recurringJobs: '[
   {
     "name":"snap", 
     "task":"snapshot",  # This defines the type of of task
     "cron":"*/1 * * * *", 
     "retain":1
   },
   {
     "name":"backup", 
     "task":"backup",  # Backup's are stored externally from the system. Compatible with S3.
     "cron":"*/2 * * * *", 
     "retain":1,
     "labels": {
       "interval":"2m"
      }
   }
  ]'
```


With this template I can define different types of storage with limitations on where they are placed. This is how I can restrict certain tasks to certain types of storage as mentioned [[Longhorn#Drive types|here]].

### Persistent Volume Claim
With the `StorageClass` set you can create a piece of persistent storage.

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: longhorn-vol-pvc
spec:
  accessModes:
    - ReadWriteOnce  # Only mode supported
  volumeMode: FileSystem
  resources:
    requests:
      storage: 2Gi
  volumeName: longhorn-vol-pv
  storageClassName: longhorn
```
[[VolumeMode|VolumeMode explained]]

This will create a persistent piece of storage that is owned by the namespace. So as long as the namespace is available, this piece of storage also will be.

### Deployment
And at the end of this chain is the Deployment. This will contain the named Persistent Volume Claim. Kubernetes will do it's magic and deploy the pod with that volume attached.



## Links
[K8s documentation on volumes](https://kubernetes.io/docs/concepts/storage/volumes/)