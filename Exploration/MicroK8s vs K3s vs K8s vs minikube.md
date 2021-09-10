---
tags: [exploration, k8s]
aliases: [versus]
---

# MicroK8s vs K3s vs K8s vs minikube vs etc

There are different ways to start a Kubernetes cluster. There is the vanilla way with `kubadm` but there are many other packages out there that claim to be K8s-certified. Here is a summary of what I found so far and what I am in the end choosing to go with.

## MicroK8s
I can be short about MircoK8s. It requires Ubuntu (or an other snap compatible OS) as it needs snap to be installed. As Ubuntu has been frying my SD-cards on a regular basis with simple `apt-get upgrade` commands, I am not going to install Ubuntu on my cluster as base OS.

As far as I can tell, MicroK8s is a great and easy solution to get a cluster going quickly. It does a lot automagically like High-Available file systems and the like and got easy to use modules for different jobs. However, as mentioned before, Ubuntu will stay of my pi's for a bit.

Link: [MicroK8s home](https://microk8s.io/)

## Minikube
Minikube looked like a very promising starter. Yet, I found nothing about setting up a multi-node cluster on hardware. It seems to be great at provisioning virtual multi-node clusters on a local machine.

Link: [Minikube home](https://minikube.sigs.k8s.io/docs/)

## K3s
Rancher's minified K8s. Boasts to be a <40 Mb install that is optimized for ARM processors. It has some included automagic like Traefic for ingress traffic, and built in loadbalancers. Is it K8s? No, but it is certified to be compatible. So kubectl and all programs that run for/on it should just run without trouble.

It also has easier support for [[Longhorn]], my chosen [[Distributed FS]].

As for memory usage it has a minimum of 512mb and a recommended of 1GB. Which is perfect!

Link: [K3s home](https://k3s.io/)
Docs: [K3s docs](https://rancher.com/docs/k3s/latest/en/)

## Kubeadm
The vanilla way. The 'right' way for the purists probably. I have found plenty heated arguments why this is the best way of setting up a cluster. AND why it isn't on a Pi.

It seems it requires a lot of resources comparatively for a Pi. The master node needs 2GB of memory, which my planned master node has only half off. So, this might not be the way I can run the cluster.

Link: [K8s home](https://kubernetes.io)
Docs: [K8s docs](https://kubernetes.io/docs/home/)

## K0s
Jeff Geerling mentioned this flavor of k8s through one of his tutorials. This flavor boast about itself with the following:
> k0s is an all-inclusive Kubernetes distribution, configured with all of the features needed to build a Kubernetes cluster simply by copying and running an executable file on each target host.
\-[source](https://docs.k0sproject.io/v1.22.1+k0s.0/)

It has a 1GB master node requirement which is the minimum. This means that it should be feasible to run this.

Link: [K0s home](https://k0sproject.io/)
Docs: [K0s docs](https://docs.k0sproject.io/v1.22.1+k0s.0/)

## Choice 

The limitations has chosen. [[K3s]] will be my flavor of k8s.

# Links
MicroK8s project page: https://microk8s.io/
K3s project page: https://rancher.com/products/k3s/
MiniKube: https://minikube.sigs.k8s.io/docs/
K0s: https://k0sproject.io/
K8s: https://kubernetes.io/docs/home/