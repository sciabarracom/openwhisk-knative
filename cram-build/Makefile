help:
	echo "build | list"

.PHONY: builder
builder:
	multipass transfer kube-master:/var/lib/tftpboot/ca.crt cram
	docker build cram -t jib-builder
	docker tag jib-builder registry.kube-system/jib-builder
	docker push registry.kube-system/jib-builder

.PHONY: preload
preload:
	docker pull docker.io/actionloop/golang-v1.12:knative-whisk
	docker tag docker.io/actionloop/golang-v1.12:knative-whisk \
	  registry.kube-system/kwhisk/golang-v1.12
	docker push registry.kube-system/kwhisk/golang-v1.12

list:
	curl -k -u registry:password https://registry.kube-system/v2/_catalog

