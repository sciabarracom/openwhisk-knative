/bin/proxy -compile main </usr/src/main.go >/action/exec
ftl --name docker.io/actionloop/kwhisk-hello-world --base docker.io/actionloop/golang-v1.12:knative-whisk --directory /action --destination docker.io/actionloop/kwhisk-hello-world
