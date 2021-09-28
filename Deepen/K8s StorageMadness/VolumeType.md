---
tags: [k8s, volumes]
---

# Volume Types
When looking into what [types of storage](https://kubernetes.io/docs/concepts/storage/volumes/#volume-types) you can create for a K8s cluster, I was a bit overwhelmed by the list available in the K8s docs. So I simplified things down for a local cluster like mine.

Here is a shortlist extracted from the docs with links:
 - [`configMap`](https://kubernetes.io/docs/concepts/storage/volumes/#configmap)
 - [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)
 - [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)
 - [`local`](https://kubernetes.io/docs/concepts/storage/volumes/#local)
 - [`nfs`](https://kubernetes.io/docs/concepts/storage/volumes/#nfs)
 - [`persistentVolumeClaim`](https://kubernetes.io/docs/concepts/storage/volumes/#persistentvolumeclaim)
 - [`secret`](https://kubernetes.io/docs/concepts/storage/volumes/#projected)

Important possibility to remember: [SubPath](https://kubernetes.io/docs/concepts/storage/volumes/#using-subpath)

And another list of volume types I like to explore more as they seem interesting but confusing:
 - [`downwardAPI`](https://kubernetes.io/docs/concepts/storage/volumes/#downwardapi) Allows K8s api data to be available as files for containers?
 - [`projected`](https://kubernetes.io/docs/concepts/storage/volumes/#projected) Allows certain volume types to be virtually mapped to the same volume?

## `configMap`
To attach a [`kind: ConfigMap`](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) to a container, you can use this volume type. Seems to mount always on the `/etc/config/` path?

## `emptyDir`
Is the basic volume type. Creates an empty directory.

## `hostPath`
> It is not advised doing this as this is a security risk

This is closest thing you can get to a bind-mount from Docker. You do have to specify what kind of bind it will be (Dir, file, socket etc)
## `local`
This will mount a resource on the Node itself as a volume. A `local` cannot be created dynamically. The difference with `hostPath` is that it can only mount directories(?).

Unlike `hostPath`, `local` can have selectors associated with. 
## `nfs`
Is what it says on the bin. It is a storage type that has many-read-write natively  supported.
## `persistentVolumeClaim`
The other side of the `Persistent Volume` architecture.
## `secret`
A secret volume type can be used to attach sensitive data directly in the filesystem of a container. Useful when this secret can not be given through environment variables.

Fun little fact: These volumes are memory backed by default to keep the secret from being written.

# Longhorned
With [[Longhorn]] I will probably only use `persistentVolumeClaim` in combination with `StorageClass`.
