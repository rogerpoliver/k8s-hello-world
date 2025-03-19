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

## How to run the server in a Docker container

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

## Contents of the `k8s` folder

- `kind.yaml`: Configuration file to create a Kubernetes cluster using Kind
  (Kubernetes IN Docker). This file defines a cluster with one control-plane
  node and three worker nodes.
