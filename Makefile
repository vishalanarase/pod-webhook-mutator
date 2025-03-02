
.PHONY: run build docker-build docker-push kustomize

run:
	go run cmd/main.go

build:
	go build -o pod-webhook-mutator cmd/main.go

docker-build:
	docker build -t vishalanarase/pod-webhook-mutator:latest .

docker-push:
	docker push vishalanarase/pod-webhook-mutator:latest

deploy:
	kustomize build manifests/ | kubectl apply -f -

undeploy:
	kustomize build manifests/ | kubectl delete -f -