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

#### Types of Services in Kubernetes

Kubernetes provides four main types of services, each serving a specific
purpose:

- **ClusterIP**: Exposes the service on an internal IP in the cluster. This type
  makes the service only reachable from within the cluster. It is the default
  service type.
- **NodePort**: Exposes the service on the same port of each selected node in
  the cluster using NAT. This type makes the service accessible from outside the
  cluster using `<NodeIP>:<NodePort>`.
- **LoadBalancer**: Exposes the service externally using a cloud provider's load
  balancer. This type is used to balance traffic across multiple pods and
  provide a single point of access.
- **ExternalName**: Maps the service to the contents of the `externalName` field
  (e.g., `foo.bar.example.com`), by returning a CNAME record with its value.
  This type is used to integrate external services into the cluster.

#### What is `targetPort`?

The `targetPort` is the port on the container that the service forwards traffic
to. It allows you to expose a different port on the service than the one the
container is listening on. For example, you can expose port 80 on the service
and forward it to port 8080 on the container.

### How to use Liveness and Readiness Probes

#### What is a Liveness Probe?

A Liveness Probe is used to determine if a container is still running. If the
Liveness Probe fails, Kubernetes will restart the container to try to recover
it. This is useful for detecting and remedying situations where an application
is stuck or in a deadlock state.

To configure a Liveness Probe, you can add the `livenessProbe` field to your
pod's YAML file. For example:

```yaml
livenessProbe:
   httpGet:
      path: /healthz
      port: 8080
   initialDelaySeconds: 3
   periodSeconds: 5
```

#### What is a Readiness Probe?

A Readiness Probe is used to determine if a container is ready to serve
requests. If the Readiness Probe fails, Kubernetes will remove the pod from the
list of endpoints for the service, ensuring that no traffic is sent to it until
it becomes ready again.

To configure a Readiness Probe, you can add the `readinessProbe` field to your
pod's YAML file. For example:

```yaml
readinessProbe:
   httpGet:
      path: /readyz
      port: 8080
   initialDelaySeconds: 3
   periodSeconds: 5
```

#### What is a Startup Probe?

A Startup Probe is used to determine if an application has successfully started.
It is particularly useful for applications that take a long time to initialize.
Without a Startup Probe, combining `livenessProbe` and `readinessProbe` can lead
to issues where the `livenessProbe` restarts the container before it has fully
started, causing a crash loop.

The `startupProbe` is designed to address this problem by delaying the execution
of `livenessProbe` and `readinessProbe` until the application has successfully
started. Once the `startupProbe` succeeds, Kubernetes will begin checking the
`livenessProbe` and `readinessProbe` as usual.

To configure a Startup Probe, you can add the `startupProbe` field to your pod's
YAML file. For example:

```yaml
startupProbe:
   httpGet:
      path: /startupz
      port: 8080
   initialDelaySeconds: 10
   periodSeconds: 5
   failureThreshold: 30
```

#### Why Use a Startup Probe?

- **Long Initialization Times**: Some applications require significant time to
  initialize, and a `livenessProbe` may incorrectly restart them during this
  period.
- **Avoid Crash Loops**: By ensuring the application has fully started before
  other probes are executed, `startupProbe` prevents unnecessary restarts and
  crash loops.

#### How It Works with Other Probes

- The `startupProbe` takes precedence over `livenessProbe` and `readinessProbe`
  during the startup phase.
- Once the `startupProbe` succeeds, Kubernetes begins executing the
  `livenessProbe` and `readinessProbe` as configured.

This ensures a smoother startup process for applications with complex or lengthy
initialization requirements.

#### Key Differences Between Liveness and Readiness Probes

- **Liveness Probe**: Ensures the container is alive and restarts it if
  necessary.
- **Readiness Probe**: Ensures the container is ready to serve traffic and
  temporarily removes it from the service if it is not.

For more details, refer to the Kubernetes documentation on probes.

### How to use the proxy to access the API

Kubernetes provides a proxy that allows you to access the Kubernetes API server.
This can be useful for debugging and managing your cluster. For more details,
refer to the "Useful Commands" section.

[Back to README](../README.md)
