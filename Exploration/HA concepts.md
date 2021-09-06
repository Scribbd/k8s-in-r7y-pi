---
tags: [HA]
aliases: [HA, HA Concepts, Highly Available, Highly Available Concepts]
---
# Highly Available K8s
One of my goals is to mess around with [[HA concepts]]. So whenever I redeploy an app, or whenever a Pi goes down, it won't make the whole cluster go down.

However, my cluster at it's current size is limited. I got 2 8GB pi4's and an older 1GB pi3+. To get to a HA cluster it is advised to have the following according to [Rancher](https://rancher.com/docs/k3s/latest/en/installation/ha/):
1.   Two or more **server nodes** that will serve the Kubernetes API and run other control plane services
2.   Zero or more **agent nodes** that are designated to run your apps and services
3.   An **external datastore** / or embedded etcd (as opposed to the embedded SQLite datastore used in single-server setups)
4.  A **fixed registration address** that is placed in front of the server nodes to allow agent nodes to register with the cluster

I can comply with point 2, 3, and 4. But I can't with 1. It would leave 