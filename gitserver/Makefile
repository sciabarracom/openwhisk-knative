IMG=registry.k8s:5000/knative/hellogo

build:
	docker build . -t $(IMG)
	docker push $(IMG)
	kubectl apply -f git-server.yaml

