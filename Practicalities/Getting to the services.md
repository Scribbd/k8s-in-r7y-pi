---
tags: [practicals]
---
# Getting to the services

## Kubernetes UI
To open proxy:
 - Linux: `sudo k3s kubectl proxy`
 - Windows: `kubectl proxy`
URL: [http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/](http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/)

To access the Kubernetes dashboard you need to run the following commands to get the access token:  
 - Linux: `kubectl -n kubernetes-dashboard describe secret admin-user-token | grep '^token'`
 - Windows: `kubectl -n kubernetes-dashboard describe secret admin-user-token | Select-String "^token"`

## Longhorn UI:
Through proxy:  `kubectl port-forward -n longhorn-system svc/longhorn-frontend 8002:80`
URL: [Proxy-URL](http://localhost:8002)

On my cluster: [Cluster-URL](https://longhorn-dash.app.cluster.lan)

## Traefik UI:
On my cluster: [Cluster-URL](https://traefik-dash.app.cluster.lan)
