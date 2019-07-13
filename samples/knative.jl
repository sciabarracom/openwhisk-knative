GIT_SERVER = """
apiVersion: v1
kind: Namespace
metadata:
  name: knative-whisk
---
apiVersion: v1
kind: Pod
metadata:
  name: git-server
spec:
  initContainers:
  - name: grep
    image: busybox:1.28
    command: ['sh', '-c', 'cp -v /config/* /build/']
    volumeMounts:
      - name: build
        mountPath: /build
      - name: config
        mountPath: /config
  containers:
  - name: kaniko-build
    image: gcr.io/kaniko-project/executor:latest
    args: ["--dockerfile=/build/Dockerfile",
            "--context=/build",
            "--destination=register.k8s:5000/demo/julia:latest",
            "--skip-tls-verify"]
    volumeMounts:
      - name: build
        mountPath: /build
  restartPolicy: Never
  hostAliases:
  - ip: 192.168.64.3
    hostnames:
    - "register.k8s"
  volumes:
    - name: build
      emptyDir: {}
    - name: config
      configMap:
        name: kaniko-config

"""
