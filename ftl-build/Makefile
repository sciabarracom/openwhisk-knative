help:
	echo "builder | list"

.PHONY: builder
builder:
	multipass transfer kube-master:/var/lib/tftpboot/ca.crt ftl
	docker build ftl -t ftl-builder
	docker tag ftl-builder registry.kube-system/ftl-builder
	docker push registry.kube-system/ftl-builder

.PHONY: preload
preload:
	docker pull docker.io/actionloop/golang-v1.12:knative-whisk
	docker tag docker.io/actionloop/golang-v1.12:knative-whisk \
	  registry.kube-system/kwhisk/golang-v1.12
	docker push registry.kube-system/kwhisk/golang-v1.12

list:
	curl -k -u registry:password https://registry.kube-system/v2/_catalog

