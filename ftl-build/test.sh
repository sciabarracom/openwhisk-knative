TGT=ftl/hello-world-3
mkdir -p workspace/src workspace/bin
cp main.go workspace/src
docker run -ti --entrypoint=/bin/sh \
  -v $PWD/workspace:/workspace \
  actionloop/golang-v1.12:knative-whisk \
  -c "/bin/proxy -compile main </workspace/src/main.go >/workspace/bin/exec"
docker run -ti -v $PWD/workspace:/workspace ftl-builder \
  --base registry.kube-system/kwhisk/golang-v1.12 \
  --name registry.kube-system/$TGT \
  --directory /workspace/bin \
  --destination /action


