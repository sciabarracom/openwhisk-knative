kubectl -n knative-whisk create secret docker-registry docker-registry-secret \
	--docker-server=https://registry.kube-system \
	--docker-username=registry \
	--docker-password=password \
	--docker-email=knative@openwhisk
