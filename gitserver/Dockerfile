FROM ubuntu
RUN apt-get update && apt-get -y install git ;\
    git config --global user.email "openwhisk@knative" ;\ 
    git config --global user.name "Knative Whisk" ;\
    mkdir -p /git/knative/hellogo
ADD hellogo/* /git/knative/hellogo/
RUN git init /git/knative/hellogo ;\
    cd /git/knative/hellogo ;\
    git add * ;\ 
    git commit -m "init" -a
WORKDIR /git
CMD git daemon --verbose --export-all --base-path=/git
