# k8s-hello-world

## Requirements

- Docker
- Kind
- kubectl

### Installing Requirements with Homebrew

1. Install Docker:
   ```sh
   brew install --cask docker
   ```

2. Install Kind:
   ```sh
   brew install kind
   ```

3. Install kubectl:
   ```sh
   brew install kubectl
   ```

## Step-by-Step Guide

### How to run the server in a Docker container

1. Build the Docker image:
   ```sh
   docker build -t <your-user>/hello-go .
   ```

2. Run the container and remove on exit:
   ```sh
   docker run --rm -p 8080:8080 <your-user>/hello-go
   ```

3. Push the image to Docker Hub:
   ```sh
   docker push <your-user>/hello-go
   ```

### How to create and view clusters

Before creating a pod, you need to create a Kubernetes cluster. Use the
following command to create a cluster:

```sh
kind create cluster --config=kind.yaml --name=<your-cluster-name>
```

To view the status of your clusters, use the following command:

```sh
kubectl config get-clusters
```

### How to create and view pods

A pod is the smallest and simplest Kubernetes object. It represents a single
instance of a running process in your cluster. A pod contains one container.

To create a pod, use the following command:

```sh
kubectl apply -f k8s/pod.yaml
```

To view the status of your pods, use the following command:

```sh
kubectl get pods
```

## Useful Commands

### How to do port-forwarding

To forward a port from your local machine to a port on a pod, use the following
command:

```sh
kubectl port-forward <pod-name> <local-port>:<pod-port>
```

### How to stop or delete pods, services, and clusters

To stop or delete a pod, use the following command:

```sh
kubectl delete pod <pod-name>
```

To stop or delete a service, use the following command:

```sh
kubectl delete service <service-name>
```

To stop or delete a cluster, use the following command:

```sh
kind delete cluster --name=<your-cluster-name>
```
