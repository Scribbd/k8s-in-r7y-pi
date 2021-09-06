---
---
# Docker in Cluster
"Docker is popular." is an understatement. Everybody up and down are using it. I am a regular user of the Docker Engine on my own desktop. Yet, I have never used containers in a cluster. Which is also a "popular thing to do".

This vault is for documenting my journey with docker/containers, and using them in a cluster. Lets get going.

## Why now?
Ubuntu decided to fry my SD-cards in my current raspberry pi's during an update. So, I had to rebuild my cluster anyway. And this time I want it done K8s style.

## Why K8s?
I was advised to just go with Docker Swarm as it is much easier to setup and maintain for home use. And yes, I do see its benefits. It *is* easier to setup than k8s. It *does* a lot for you with networking and interconnecting containers. But it *isn't* the open-source solution that is popular right now. Will I try it out one day, probably. But for now it is all Kubernetes.

## Goals
- [[Setup a Kubernetes cluster]]
- Get familiarity with the HA concepts on a bare metal level
	- Get a [[resilient data]] store
	- Todo when I get more RPI's for the cluster:
		- HA Control Plane
- Get a grip on [[security within the cluster]]
- Get a grip on [[Ansible]]

## Exploration
Like always, I started with [[Choice Paralysis]].

## Available hardware
- 2 x Raspberry Pi 4 8GB
- 1 x Raspberry Pi 3B+ 1GB
- 2 x 256GB SATA SSD with UASP adapter
- 1 x 1TB HDD external w/o UASP for backup / snapshots