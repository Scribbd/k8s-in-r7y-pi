---
tags: [k8s, volumes]
---

# PV vs PVC vs SC vs Provisioner
[Source](https://rancher.com/blog/2018/2018-09-20-unexpected-kubernetes-part-1/)

When researching how to create persistent storage, I came across a few terms that were used in Kubernetes:
 - Persistent Volume
 - Persistent Volume Claim
 - Storage Class
 - Provisioner

I thought these were all used to provide storage and had all to be used to get a persistent volume. So with that in mind I started searching for an answer. And oh boy was I wrong. So here is a rundown of how to use them:

## Persistent Volume Claim
This is what users / applications use to get a piece of storage. This is different from a normal docker volume as it is maintained through crashes. It is also owned by the namespace. A Persistent Storage Claim has to have a Storage Class or Persistent Volume attached to it.

## Persistent Volume
This is the fixed way (<v1.6) of doing things. Unlike a PVC this can only be made by an administrator. These PVs are pre-carved pieces of storage that can be claimed by a PVC of a user. These PVs are very fixed, if a PVC requires only 1GiB but there is only a claim for 10GiB, K8s will claim the whole 10GiB for that one claim.

The PV and PVC structure was maybe a solution for Google's workload. But it was probably an administrative nightmare.

## Storage Class and Provisioner
With v1.6 came Storage Classes. This is what a PV is, but without a given size. And together with a Provisioner a user could dynamically claim a Persistent Volume. A Provisioner is a piece of software that is able to dynamically carve out a piece of storage, like [[Longhorn]]. A SC is a template for the Priviosoner to use.