FROM julia:1.1.0
ARG REPO=https://github.com/sciabarracom/openwhisk-knative-operator
ARG BRANCH=master
WORKDIR /root
RUN curl -sL \
 https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
 -o /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl
RUN curl -sL \
  https://storage.googleapis.com/kubernetes-helm/helm-v2.14.0-linux-amd64.tar.gz \
  | tar xzvf - linux-amd64/helm --strip-components=1 \
  && mv helm /usr/local/bin \
  && helm init --client-only
ADD install.jl installer.jl util.jl /usr/local/bin/
CMD julia /usr/local/bin/install.jl