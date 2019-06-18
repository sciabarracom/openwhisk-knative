FROM julia:1.1.0
WORKDIR /root
RUN curl -sL \
 https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
 -o /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl
RUN curl -sL \
  https://storage.googleapis.com/kubernetes-helm/helm-v2.14.0-linux-amd64.tar.gz \
  | tar xzvf - linux-amd64/helm --strip-components=1 -C /usr/local/bin
RUN curl -Ls \
  https://github.com/solo-io/gloo/releases/download/v0.13.29/glooctl-linux-amd64 \
  -o /usr/local/bin/glooctl && chmod +x /usr/local/bin/glooctl
ADD setup.jl .
RUN env BRANCH=$BRANCH julia setup.jl
CMD julia -e 'using Pkg; Pkg.activate("."); using KnativeWhisk'
