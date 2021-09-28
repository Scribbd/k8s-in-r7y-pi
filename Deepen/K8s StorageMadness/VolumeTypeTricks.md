---
tags: [k8s, volumes]
---

# Fun little volume stuff

## Memory volume
When setting `emptyDir.medium` field to `"Memory"`, Kubernetes mounts a tmpfs (RAM-backed filesystem).

