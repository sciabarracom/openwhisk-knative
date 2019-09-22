FROM actionloop/golang-v1.12:knative-whisk
RUN curl https://storage.googleapis.com/gcp-container-tools/ftl/node/latest/ftl.par >/bin/ftl && chmod +x /bin/ftl
COPY config.json /root/.docker/config.json
COPY main.go /usr/src/main.go
ENTRYPOINT ["/bin/bash"]
