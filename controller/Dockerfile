FROM ubuntu
RUN apt-get update && apt-get -y install git ;\
    git config --global user.email "openwhisk@knative" ;\ 
    git config --global user.name "Knative Whisk" ;\
    mkdir -p /git
WORKDIR /git
ENV KW_REPO=/git
COPY controller /bin
