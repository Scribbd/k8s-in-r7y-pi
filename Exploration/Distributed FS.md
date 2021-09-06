---
tags: [k8s, dfs]
aliases: [DFS, Distributed FS]
---
# Distributed File system 
Part of a [[HA concepts|Highly Available]] cluster is to have a HA file system. So that when a stateful pod has to move to another node for whatever reason, it will still have access to all the files it needs.

