---
tags: [HA, dfs]
---
# Longhorn
Longhorn is a distributed file system.

## How will it fit in my cluster
The master node will not run tasks, so it won't be storing 'live' data. I hope to find a way to use the attached storage (1TB of slow HDD) as a backup/snapshot storage of the 256GB working space on the worker nodes.

The 2 worker nodes have both an SSD attached over UASP and, if possible, enable trim on the drive. This is can be configured system-wide with [[Playbooks#Pre-Horn]].

## Updates on things:
The use of the slow TB drive as backup: