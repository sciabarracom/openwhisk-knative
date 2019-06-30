
using Base64

indent(file) = replace(open(x -> read(x, String), file), "\n" => "\n    ")

CSR_REQUEST = """
{
  "hosts": [
    "register.knative-whisk.svc.cluster.local"
  ],
  "CN": "register.knative-whisk.svc.cluster.local",
  "key": {
    "algo": "ecdsa",
    "size": 256
  }
}
"""

run(pipeline(IOBuffer(CSR_REQUEST), `cfssl genkey -`, `cfssljson -bare server`))
server_csr = open(f -> read(f, String), "server.csr") 

SIGN_REQUEST(request) = """
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: register.knative-whisk
  namespace: knative-whisk
spec:
  request: $request
  usages:
  - digital signature
  - key encipherment
  - server auth
"""

run(pipeline(IOBuffer(SIGN_REQUEST(base64encode(server_csr))), `kubectl apply -f -`))

run(`kubectl certificate approve register.knative-whisk`)

server_crt = String(base64decode(read(`kubectl -n knative-whisk get csr register.knative-whisk -o jsonpath='{.status.certificate}'`, String)))
open(f -> write(f, server_crt), "server.crt", "w") 


CERT_MAP(keyfile, certfile) = """
kind: ConfigMap 
apiVersion: v1 
metadata:
  name: certs
  namespace: knative-whisk
data:
  tls.key: | 
    $(indent(keyfile))
  tls.crt: |
    $(indent(certfile))
"""

server_key = String(open(f -> read(f), "server-key.pem"))
run(pipeline(IOBuffer(CERT_MAP("server-key.pem", "server.crt")), `kubectl apply -f -`))

DOCKER_SECRET(keydata, certdata) = """
apiVersion: v1
data:
  tls.crt: $(base64encode(certdata))
  tls.key: $(base64encode(keydata))
kind: Secret
metadata:
  name: registry-tls
  namespace: knative-whisk
type: Opaque
"""

run(pipeline(IOBuffer(DOCKER_SECRET(server_key, server_crt)), `kubectl apply -f -`))

REGISTER_DEPLOY = """
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: registry
  namespace: knative-whisk
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: registry
    spec:
      containers:
      - name: proxy
        image: alpine/socat
        command:
        - socat
        - TCP-LISTEN:80,fork
        - TCP:127.0.0.1:5000
        ports:
        - containerPort: 80
      - name: registry
        image: registry:2
        volumeMounts:
        - name: docker-data
          mountPath: /var/lib/registry          
      volumes:
      - name: docker-data
        emptyDir: {}
---
apiVersion: v1
kind: Service
metadata:
  name: register
  namespace: knative-whisk
spec:
  selector:
    app: registry
  ports:
  - name: register
    protocol: TCP
    port: 80
    targetPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: register-nodeport
  namespace: knative-whisk
spec:
  type: NodePort
  selector:
    app: registry
  ports:
  - name: register
    protocol: TCP
    nodePort: 30080
    port: 80
    targetPort: 80
"""

run(pipeline(IOBuffer(REGISTER_DEPLOY), `kubectl delete -f -`))
run(pipeline(IOBuffer(REGISTER_DEPLOY), `kubectl apply -f -`))

