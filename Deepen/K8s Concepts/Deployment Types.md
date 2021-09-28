---
tags: [k8s, deployment]
---
# Deployment Types
An inventory about Kubernetes deployment types
## Kubectl
You can use `kubectl` to deploy everything on the cluster. This is the most tedious way. This is the most scriptable way.

## Yaml (kubectl apply -f)
Still using the command line, this is the way to use infrastructure as code. A yaml file can contain a complete deployment with all resources needed for an application to run.
To remove a deployment you can run kubectl again but with `delete -f`.

## Ansible (kubectl)
[[Ansible]] is also able to automate k8s deployments with the k8s module. Advantages of using Ansible is that you can deploy and delete a deployment with a simple switch in `present:`.

## Helm
A K8s package manager. Is somewhat like docker-compose and tries to simplify the deployment of applications on Kubernetes clusters. A `Helmchart` contains the configuration of the deployment and updates to this Helmchart is automatically applies(?).

## Operator
A meta-way of deployment on a cluster. An operator is a pod that deploys the application-pod based on a configuration provided. Depending on how advanced the operator is, it can also update the resources on configuration changes and/or repair failed deployments on its own.