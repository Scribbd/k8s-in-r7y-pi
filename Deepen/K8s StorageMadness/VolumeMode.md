---
tags: [k8s, volumes]
---

# Volume Mode
When looking through [[Longhorn]] examples, [here](https://longhorn.io/docs/1.2.0/references/examples/). I saw `VolumeMode` as a parameter, which were either `Block` or `Filesystem`. Trying to understand the difference I took a dive into the [K8s documentation](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#volume-mode) on `VolumeMode`s

There are two modes a Volume can be in:
 - `Filesystem`
 - `Block`

## `Filesystem`
`Filesystem` is the simplest: It is basically a folder mounted in the container's filesystem. This is the closest to what a Docker Volume is.

## `Block`
`Block` is a bit more advanced. This uses a piece of block storage as-is. Where `Filesystem` will create an empty filesystem when none is present. `Block` will opt out on this. This requires the application that will use this volume to be able to handle raw block storage. 

A block volume can also be provisioned through either a [[PV vs PVC vs SC vs Provisioner#Persistent Volume|Persistent Volume]] or through a [[PV vs PVC vs SC vs Provisioner#Storage Class and Provisioner|Storage Class]]. How Kubernetes does this is unkown to me
