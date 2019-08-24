TGT=kwhisk/hello-world
time docker run -ti -v $PWD/workspace:/workspace jib-builder \
  --registry --insecure \
  --credential-helper=docker-credential-simple \
  -E=OW_AUTOINIT=/action/exec \
  registry.kube-system/kwhisk/golang-v1.12 \
  registry.kube-system/$TGT \
  /workspace/bin/exec:/action/exec



