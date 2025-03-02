# Certificates

> This directory contains the certificates used by the pod-webhook-mutator application. The certificates are used to secure the communication between the Kubernetes API server and the pod-webhook-mutator application.

## CA Certificates

Generate private key and certificate signing request

```bash
openssl genrsa -out certs/ca.key 2048
openssl req -new -x509 -days 365 -key certs/ca.key -out certs/ca.crt -subj "/CN=Kubernetes Admission Controller"
```

## Server Certificates

Server certifcates for the pod-webhook-mutator application service

```bash
export SERVICE_NAME=pod-webhook-mutator
openssl genrsa -out certs/server.key 2048
openssl req -new -key certs/server.key -out certs/server.csr -subj "/CN=$SERVICE_NAME"
openssl x509 -req -in certs/server.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/server.crt -days 365 -extfile certs/openssl.cnf -extensions v3_req
```

Check the certificate

```bash
openssl x509 -in certs/server.crt -text -noout
```
It should contain the following subject alternative names

```yaml
X509v3 Subject Alternative Name:
                DNS:pod-webhook-mutator.default.svc.cluster.local, DNS:pod-webhook-mutator.default.svc.cluster, DNS:pod-webhook-mutator.default.svc, DNS:pod-webhook-mutator.default, DNS:pod-webhook-mutator, DNS:localhost
```

## Client Certificates

Client certificates for the pod-webhook-mutator application service

```bash
openssl genrsa -out certs/client.key 2048
openssl req -new -key certs/client.key -out certs/client.csr -subj "/CN=client"
openssl x509 -req -in certs/client.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/client.crt -days 365
```

## Kubernetes Secrets

Create a Kubernetes secret with the certificates

```bash
kubectl create secret generic pod-webhook-mutator-certs --from-file=ca.crt --from-file=server.crt --from-file=server.key --from-file=client.crt --from-file=client.key
```




