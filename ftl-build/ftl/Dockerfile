FROM ubuntu
RUN apt-get update && apt-get -y install curl python zip unzip \
    && curl https://storage.googleapis.com/gcp-container-tools/ftl/node/latest/ftl.par \
       >/usr/local/bin/ftl.par && \
    chmod +x /usr/local/bin/ftl.par && \
    mkdir -p /root/.credhelper /root/.docker /usr/local/share/ca-certificates/registry.kube-system
#ADD ca.crt /usr/local/share/ca-certificates/registry.kube-system/ca.crt
#RUN update-ca-certificates
ADD docker-credential-simple /usr/bin/docker-credential-simple
ADD creds.json /root/.credhelper/creds.json
ADD config.json /root/.docker/config.json
# patch ftl.par to disable ssl validation
RUN cd /usr/local/bin && \
unzip ftl.par containerregistry/transport/transport_pool_.py && \
sed -i.bak \
  -e's/transport_factory()/transport_factory(disable_ssl_certificate_validation=True)/' \
  containerregistry/transport/transport_pool_.py && \
zip -u ftl.par containerregistry/transport/transport_pool_.py
ENTRYPOINT ["/usr/local/bin/ftl.par"]
