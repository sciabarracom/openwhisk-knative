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
      - name: registry
        image: registry:2
        ports:
        - containerPort: 5000
        command:
        - tail 
        - -f
        - /dev/null
        volumeMounts:
        - name: docker-data
          mountPath: /var/lib/registry
        - name: certs
          mountPath: /certs
        env:
        - name: REGISTRY_HTTP_ADDR
          value: 0.0.0.0:5000
        - name: REGISTRY_HTTP_TLS_CERTIFICATE
          value: /certs/tls.crt
        - name: REGISTRY_HTTP_TLS_KEY
          value: /certs/tls.key
      volumes:
      - name: docker-data
        emptyDir: {}
      - name: certs
        configMap:
          name: certs
"""
