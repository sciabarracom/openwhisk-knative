ID=$(date +%s)
kubectl -n knative-whisk apply -f - <<EOF
apiVersion: tekton.dev/v1alpha1
kind: TaskRun
metadata:
  name: kwhisk-build-$ID
  namespace: knative-whisk
spec:
  taskRef:
    name: kwhisk-builder
  inputs:
    resources:
    - name: source
      resourceSpec:
        type: git
        params:
        - name: url
          value: git://kwhisk-git-server/kwhisk/hello/world
  outputs:
    resources:
    - name: target
      resourceSpec:
        type: image
        params:
        - name: url
          value: registry.kube-system/kwhisk/hello/world/$ID
EOF
kubectl -n knative-whisk get po -w

