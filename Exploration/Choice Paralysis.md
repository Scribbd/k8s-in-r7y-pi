---
tags: [exploration]
---
# Choice Paralysis 
"Kubernetes is popular." is an understatement... Déjà vu? However, this sentence is also a source of my trouble: choice paralysis.

## Kubernetes
As it is massively popular there are many ways to start a cluster. From the -probably seen as- the proper, vanilla, way of installing it with kubeadm. To single executables like microK8s that do a lot of stuff right out of the box. So I am stuck with a choice: Do I want to do it the hard or the easy way? If there is such a thing as easy with Kubernetes.

"But why go through that trouble when you could just rebuild with an other option when you don't like?" you say. I don't know, okay. I like to make things way too complicated and important than actually needed. Still, here is my exploration into my choice of K8s flavor in this [[MicroK8s vs K3s vs K8s vs minikube|versus]] document.

## Automation
As my cluster is growing in size. I am starting to feel the annoyance of logging into every pi and configuring/monitoring them. So, I should start looking into ways of [[Configuring the nodes automagicaly|automating this]] . And boy, are there many options for that as well.

However, for this, somehow, it didn't paralyze me. I chose [[Ansible]] rather quickly.

## Distributed File System
There are some good file systems out there specifically created to run in a cluster environment. Some are block-based, others are object-based. My exploration in that can be found in [[Distributed FS]].

# Choice Paralysis resolved(?)
For Kubernetes I will use: [[K3s]]

For Automation I will use: [[Ansible]]

For DFS I will use: [[Longhorn]]