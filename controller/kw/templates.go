package kw

import "fmt"

// Dockerfile to build an action
func Dockerfile() []byte {
	return []byte(`FROM actionloop/golang-v1.12:knative-whisk
COPY src /action/src
RUN /bin/proxy -compile main </action/src >/action/exec 2>/action/err
ENV OW_AUTOINIT=/action/exec
ENV OW_AUTOINIT_MAIN=main
`)
}

func build(name, source, target string) []byte {
	return []byte(fmt.Sprintf(`apiVersion: 
tekton.dev/v1alpha1
kind: TaskRun
metadata:
  name: %s
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
          value: %s
  outputs:
    resources:
    - name: target
      resourceSpec:
        type: image
        params:
        - name: url
          value: %s
`, name, source, target))

}
