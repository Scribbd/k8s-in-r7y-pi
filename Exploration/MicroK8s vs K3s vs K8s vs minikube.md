---
tags: [exploration]
aliases: [versus]
---

# MicroK8s vs K3s vs K8s vs minikube

There are different ways to start a Kubernetes cluster. There is the vanilla way with `kubadm` but there are many other packages out there that claim to be K8s-certified. Here is a summary of what I found so far and what I am in the end choosing to go with.

## MicroK8s
I can be short about MircoK8s. It requires Ubuntu (or other compatible OS) as it needs snap to be installed. As Ubuntu has been frying my SD-cards on a regular basis with simple `apt-get upgrade` commands, I am not going to install Ubuntu on my cluster as base OS.

As far as I can tell, MicroK8s is a great and easy solution to get a cluster going quickly. It does a lot automagically like High-Available file systems and the like and got easy to use modules for different jobs.

## Minikube
Minikube looked like a very promising starter. Yet, I found nothing about setting up a multi-node cluster on hardware. It seems to be great at provisioning virtual multi-node clusters

## K3s
Rancher

## Kubeadm
The vanilla way. The 'right' way. I will probably find many heated arguments why this is the best way of setting up a cluster. AND why it isn't.

## Choice 

# Links
MicroK8s project page: https://microk8s.io/
K3s project page: https://rancher.com/products/k3s/
MiniKube: https://minikube.sigs.k8s.io/docs/
K8s: https://kubernetes.io/docs/home/