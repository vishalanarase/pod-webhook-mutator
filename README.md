# Pod Webhook Mutator

> This project implements a Kubernetes Mutating Admission Webhook for pod resources. The webhook mutates the pod objects by modifying the environment variables of containers in the pod.


## Project Structure

`Dockerfile`, `Makefile`, `README.md`: Project build, configuration, and documentation files.
`certs`: SSL/TLS certificate-related files for secure communication.
`client`: Client-side HTTP requests for testing or interacting with the webhook.
`cmd`: The main application entry point (e.g., starting the server).
`docs`: Documentation for certificates and Docker usage.
`manifests`: Kubernetes resources for deploying and managing the webhook in a Kubernetes cluster.
`pkg`: Core application logic, including configuration, webhook handlers, and mutation logic.

## Getting Started

### Prerequisites

- Kind
- Docker
- kubectl
- kustomize
- Go (for development)

### Kind Cluster Setup

1. Kind cluster configuration:

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.32.2
- role: worker
  image: kindest/node:v1.32.2
- role: worker
  image: kindest/node:v1.32.2
``

2. Apply the cluster configuration:

```bash
kind create cluster --config kind.yaml
 ```

### Certificate Generation

Follow the instructions in `docs/certs.md` to generate the necessary certificates for the webhook.

### Docker Build and Push

Follow the instructions in `docs/docker.md` to build and push the Docker image to a container registry.

### Deploying the Webhook

```bash
make deploy
```

### Testing the Webhook

```bash
kubectl apply -f manifests/examples/pod.yaml
```

### Cleanup

```bash
make undeploy
kind delete cluster
```
