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

#### What is a Cluster?

A cluster is a set of nodes (machines) that run containerized applications
managed by Kubernetes. A cluster consists of at least one control plane and one
or more worker nodes. The control plane manages the cluster, while the worker
nodes run the applications.

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

#### What is a Pod?

A pod is the smallest and simplest Kubernetes object. It represents a single
instance of a running process in your cluster. A pod contains one or more
containers, which are tightly coupled and share the same network namespace and
storage.

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

#### What is a ReplicaSet?

A ReplicaSet is a Kubernetes resource that ensures a specified number of pod
replicas are running at any given time. If a pod crashes or is deleted, the
ReplicaSet will automatically create a new pod to replace it, ensuring the
desired state is maintained.

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

#### What is a Deployment?

A Deployment is a higher-level Kubernetes resource that manages ReplicaSets and
provides declarative updates to applications. It ensures that the desired number
of replicas of a pod are running and can update the pods to a new version
without manual intervention.

##### Problem with ReplicaSets

When using a ReplicaSet, if you want to update your application to a new version
(e.g., v2), you would need to manually delete each pod and create new ones with
the updated version. This can be cumbersome and error-prone.

##### How Deployments Solve This Problem

Deployments automate the process of updating pods to a new version. When you
apply a Deployment, it creates and manages ReplicaSets, which in turn create the
pods. If you update the Deployment to a new version, it will automatically
create a new ReplicaSet and gradually replace the old pods with new ones.

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

### How to perform a Rollout and view Revisions

#### What is a Rollout?

A Rollout is the process of updating a Deployment to a new version. Kubernetes
supports rolling updates, which means that it will gradually replace the old
pods with new ones, ensuring that the application remains available during the
update.

#### What are Revisions?

Revisions are snapshots of the state of a Deployment. Each time you update a
Deployment, Kubernetes creates a new revision. You can use these revisions to
roll back to a previous state if something goes wrong with the update.

To view the revision history of a Deployment, use the following command:

```sh
kubectl rollout history deployment/<deployment-name>
```

To roll back to a previous revision, use the following command:

```sh
kubectl rollout undo deployment/<deployment-name> --to-revision=<revision-number>
```

### How to create and view Services

#### What is a Service?

A Service in Kubernetes is an abstraction that defines a logical set of pods and
a policy by which to access them. Services enable communication between
different parts of your application and can expose your application to the
internet. A Service also acts as a load balancer, distributing incoming network
traffic across multiple pods to ensure no single pod becomes a bottleneck. This
is particularly useful when you have multiple replicas of a pod running, as it
ensures that the load is evenly distributed among them.

To create a Service, use the following command:

```sh
kubectl apply -f k8s/service.yaml
```

To view the status of your Services, use the following command:

```sh
kubectl get services
```

[Back to README](../README.md)
