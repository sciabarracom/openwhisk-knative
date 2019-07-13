REG=registry.k8s
IMG=$(REG)/git-hellogo-sources:latest

.PHONY: all
all: secret build deploy 

.PHONY: deploy
deploy:
	kubectl apply -f git-server.yaml

.PHONY: build
build:
	docker build . -t $(IMG)
	docker login -u registry -p password $(REG)
	docker push $(IMG)

.PHONY: secret
secret:
	-kubectl create ns knative-whisk
	-kubectl -n knative-whisk delete secret docker-registry-secret
	kubectl -n knative-whisk create secret docker-registry docker-registry-secret \
	--docker-server=https://registry.k8s \
	--docker-username=registry \
	--docker-password=password \
	--docker-email=knative@openwhisk

proxy:
	kubectl apply -f registry-proxy.yaml

.PHONY: list
list:
	curl  -u registry:password -k https://registry.k8s/v2/_catalog
