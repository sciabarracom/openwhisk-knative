TGT=kwhisk/hello-world
mkdir -p workspace/src workspace/bin
cp main.go workspace/src
docker run -ti --entrypoint=/bin/sh \
  -v $PWD/workspace:/workspace \
  actionloop/golang-v1.12:knative-whisk \
  -c "/bin/proxy -compile main </workspace/src/main.go >/workspace/bin/exec"
docker run -ti -v $PWD/workspace:/workspace jib-builder \
  --registry --insecure \
  --credential-helper=docker-credential-simple \
  -E=OW_AUTOINIT=/action/exec \
  registry.kube-system/kwhisk/golang-v1.12 \
  registry.kube-system/$TGT \
  /workspace/bin/exec:/action/exec
echo $TGT


