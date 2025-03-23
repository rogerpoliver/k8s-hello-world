## Useful Commands

### How to do port-forwarding

Once you have the cluster and pods running, you can use port-forwarding to
access your application on the desired port. For example, to forward port 8080
from your local machine to port 8080 on the pod, use the following command:

```sh
kubectl port-forward <pod-name> 8080:8080
```

### How to use the proxy to access the API

To start the proxy and specify the port, use the following command:

```sh
kubectl proxy --port=8001
```

Once the proxy is running, you can access the API server at
`http://localhost:8001`.

### How to stop or delete pods, services, and clusters

To stop or delete a pod, use the following command:

```sh
kubectl delete pod <pod-name>
```

To stop or delete a service, use the following command:

```sh
kubectl delete service <service-name>
```

To stop or delete a ReplicaSet, use the following command:

```sh
kubectl delete replicaset <replicaset-name>
```

To stop or delete a cluster, use the following command:

```sh
kind delete cluster --name=<your-cluster-name>
```

[Back to README](../README.md)
