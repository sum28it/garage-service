# Other commands to install
# go install github.com/divam/expvarmon@latest

# RSA Keys
# 	To generate a private/public key PEM file.
# 	$ openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# 	$ openssl rsa -pubout -in private.pem -out public.pem

db: 
	go run app/scratch/db/main.go

jwt:
	go run app/scratch/jwt/main.go
auth-local:
	curl -il -H "Authorization: Bearer ${TOKEN}" localhost:3000/auth
run:
	go run app\services\sales-api\main.go | go run app\tooling\logfmt\main.go

tidy:
	go mod tidy
	go mod vendor

metrics-local:
	expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

metrics-view:
	expvarmon -ports="sales-service.sales-system.svc.cluster.local:4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

test-load-local:
	hey -m GET -c 100 -n 10000 http://localhost:3000/status
	
# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 1.0

all: sales

sales:
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		.

# ==============================================================================
# Running from within k8s/kind

GOLANG       := golang:1.19
ALPINE       := alpine:3.17
KIND         := kindest/node:v1.25.3
POSTGRES     := postgres:15-alpine
VAULT        := hashicorp/vault:1.12
ZIPKIN       := openzipkin/zipkin:2.23
TELEPRESENCE := docker.io/datawire/tel2:2.10.4


dev-tel:
	kind load docker-image $(TELEPRESENCE) --name $(KIND_CLUSTER)
	telepresence --context=kind-$(KIND_CLUSTER) helm install
	telepresence --context=kind-$(KIND_CLUSTER) connect

KIND_CLUSTER:=ardan-web-service-cluster

dev-up:
	kind create cluster \
		--image kindest/node:v1.27.2@sha256:3966ac761ae0136263ffdb6cfd4db23ef8a83cba8a463690e98317add2c9ba72 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner
	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)


dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces


dev-load:
	kind load docker-image sales-api:$(VERSION) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/database | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=sales-system --for=condition=Available deployment/database
	kustomize build zarf/k8s/dev/sales | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=sales-system --for=condition=Available deployment/sales

dev-restart:
	kubectl rollout restart deployment sales --namespace=sales-system

dev-logs:
	kubectl logs --namespace=sales-system -l app=sales --all-containers=true -f --tail=100 --max-log-requests=6 | go run app\tooling\logfmt\main.go --service=SALES-API

dev-logs-db:
	kubectl logs --namespace=sales-system -l app=database --all-containers=true -f --tail=100

dev-describe:
	kubectl describe nodes
	kubectl describe svc

dev-describe-deployment:
	kubectl describe deployment --namespace=sales-system sales

dev-describe-sales:
	kubectl describe pod --namespace=sales-system -l app=sales

dev-update: all dev-load dev-restart

dev-update-apply: all dev-load dev-apply

dev-logs-init:
	kubectl logs --namespace=sales-system -l app=sales -f --tail=100 -c init-db