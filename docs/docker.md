# Dockerfile

> This Dockerfile is used to build a Docker image for the pod-webhook-mutator application. The image is based on the official Golang image and uses a multi-stage build to reduce the final image size.

## Build

To build the Docker image, run the following command:

```bash
docker build -t pod-webhook-mutator .
```

## Run

To run the Docker container, run the following command:

```bash
docker run -d -p 443:443 --name pod-webhook-mutator pod-webhook-mutator
```

## Logs

To view the logs of the Docker container, run the following command:

```bash
docker logs pod-webhook-mutator
```

## Shell

To open a shell in the Docker container, run the following command:

```bash
docker exec -it pod-webhook-mutator sh
```

## Stop

To stop the Docker container, run the following command:

```bash
docker stop pod-webhook-mutator
```

## Remove

To remove the Docker container, run the following command:

```bash
docker rm pod-webhook-mutator
```
