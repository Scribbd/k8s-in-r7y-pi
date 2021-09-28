---
tags: [exploration, k8s, dfs, HA]
aliases: [DFS, Distributed FS]
---
# Distributed File system 
Part of a [[HA concepts|Highly Available]] cluster is to have a HA file system. So that when a stateful pod has to move to another node for whatever reason, it will still have access to all the files it needs.

A distributed file system isn't NFS. NFS is a centralized server to which clients can connect to. If this NFS server goes down it will bring your cluster down.

You can place these systems in several categories:
 - Object Storage
 - Network File System (not NFS)
 - Block Storage

List of possible Distributed File Systems:
 - [Longhorn](https://longhorn.io/) (My choice)
 - [OpenEBS](https://openebs.io/) (Block storage)
	 - [Mayastor](https://mayastor.gitbook.io/introduction/)
	 - [cStor](https://github.com/openebs/cstor-operators/blob/develop/docs/quick.md)
	 - [Jiva](https://github.com/openebs/jiva-operator)
 - [GlusterFS](https://www.gluster.org/) (Networked File System)
 - [Ceph](https://ceph.io/en/)
 - [Rook](https://rook.io/) (orchestrator)
 - [Minio](https://min.io/) (Object storage)