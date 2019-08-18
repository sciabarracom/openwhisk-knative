FROM ubuntu
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl ;\
    chmod +x ./kubectl ; mv /usr/bin/kubectl
RUN apt-get update && apt-get -y install git ;\
    git config --global user.email "openwhisk@knative" ;\
    git config --global user.name "Knative Whisk" ;\
    mkdir -p /git
WORKDIR /git
ENV KW_REPO=/git
COPY controller /bin
