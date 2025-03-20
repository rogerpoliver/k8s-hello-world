# k8s-hello-world

## Introduction

### What is a Container?

A container is a lightweight, standalone, and executable software package that
includes everything needed to run a piece of software, including the code,
runtime, system tools, libraries, and settings. Containers are isolated from
each other and the host system, which makes them portable and consistent across
different environments.

### What is Kind?

Kind (Kubernetes IN Docker) is a tool for running local Kubernetes clusters
using Docker containers as nodes. It is primarily designed for testing
Kubernetes itself, but it can also be used for local development or CI.

### What is Kubernetes?

Kubernetes is an open-source platform designed to automate deploying, scaling,
and operating application containers. It groups containers that make up an
application into logical units for easy management and discovery.

### What is kubectl?

kubectl is a command-line tool for interacting with a Kubernetes cluster. It
allows you to run commands against Kubernetes clusters to deploy applications,
inspect and manage cluster resources, and view logs.

### What is a Pod?

A pod is the smallest and simplest Kubernetes object. It represents a single
instance of a running process in your cluster. A pod contains one or more
containers, which are tightly coupled and share the same network namespace and
storage.

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
