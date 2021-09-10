---
tags: [k8s]
---
# K3s
K3s is a distribution made by [Rancher](https://rancher.com/). It boasts to only require half the memory footprint. For a time I thought K3s stood for Kubes, just like k8s stands for Kubernetes. However, k3s doesn't have an official long-form to it. ^[[source](https://rancher.com/docs/k3s/latest/en/#what-s-with-the-name)]

## Installation
Installation can easily be done with a single command line on the terminal. For those finding this and just wan't to run K3s somewhere, here is that line. 
```bash
curl -sfL https://get.k3s.io | sh -
``` 
Be careful, you will be loading a random script in your system.

Of course, I wan't to do it a bit differently. With Ansible: [[Playbooks#K8s-install]].

## Configuration
