REG=registry.k8s:5000
IMG=$(REG)/knative/hellogo:latest

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
	-kubectl -n knative-whisk delete secret docker-registry-secret
	kubectl -n knative-whisk create secret docker-registry docker-registry-secret \
	--docker-server=registry.k8s:5000 \
	--docker-username=registry \
	--docker-password=password \
	--docker-email=knative@openwhisk

.PHONY: list
list:
	curl  -u registry:password -k https://registry.k8s:5000/v2/_catalog