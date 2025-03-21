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

### What is a ReplicaSet?

A ReplicaSet is a Kubernetes resource that ensures a specified number of pod
replicas are running at any given time. If a pod crashes or is deleted, the
ReplicaSet will automatically create a new pod to replace it, ensuring the
desired state is maintained.

### What is a Deployment?

A Deployment is a higher-level Kubernetes resource that manages ReplicaSets and
provides declarative updates to applications. It ensures that the desired number
of replicas of a pod are running and can update the pods to a new version
without manual intervention.

#### Problem with ReplicaSets

When using a ReplicaSet, if you want to update your application to a new version
(e.g., v2), you would need to manually delete each pod and create new ones with
the updated version. This can be cumbersome and error-prone.

#### How Deployments Solve This Problem

Deployments automate the process of updating pods to a new version. When you
apply a Deployment, it creates and manages ReplicaSets, which in turn create the
pods. If you update the Deployment to a new version, it will automatically
create a new ReplicaSet and gradually replace the old pods with new ones.

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

The name of the pod is defined in the `k8s/pod.yaml` file.

To view the status of your pods, use the following command:

```sh
kubectl get pods
```

Once the pod is running, you can test it by using port-forwarding to access the
application in your browser. The commands for port-forwarding can be found in
the "Useful Commands" section.

### How to create and view ReplicaSets

To create a ReplicaSet, use the following command:

```sh
kubectl apply -f k8s/replicaset.yaml
```

To view the status of your ReplicaSets, use the following command:

```sh
kubectl get replicasets
```

To test the ReplicaSet, you can view the pods, delete one of them, and then view
the pods again to see that the ReplicaSet has automatically created a new pod.
The commands for these actions can be found in the "Useful Commands" section.

### How to create and view Deployments

Before applying a Deployment, you need to delete any existing ReplicaSets and
Pods. The commands for these actions can be found in the "Useful Commands"
section.

To create a Deployment, use the following command:

```sh
kubectl apply -f k8s/deployment.yaml
```

To view the status of your Deployments, use the following command:

```sh
kubectl get deployments
```

To test the Deployment, you can update the number of replicas in the
`k8s/deployment.yaml` file, apply the file, and then use `kubectl get pods` to
see the changes. The commands for these actions can be found in the "Useful
Commands" section.

## Useful Commands

### How to do port-forwarding

Once you have the cluster and pods running, you can use port-forwarding to
access your application on the desired port. For example, to forward port 8080
from your local machine to port 8080 on the pod, use the following command:

```sh
kubectl port-forward <pod-name> 8080:8080
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

To stop or delete a ReplicaSet, use the following command:

```sh
kubectl delete replicaset <replicaset-name>
```

To stop or delete a cluster, use the following command:

```sh
kind delete cluster --name=<your-cluster-name>
```
