# OpenWhisk Knative 

Warning: this is still work in progress.

This is a proof-of-concept implementation of OpenWhisk on top of Knative. It tries to be compatible with the current user experience of OpenWhisk users implementing it on top of Knative.

It actually uses Knative Serving version 0.6 with Istio and Tekton Pipelines. An installer of the prerequisites is included in the kit.

In the current state it has many [limitations](#limitations) See below.

Once you have installed both the cluster and the runtime, you can just do:

```
$ wsk action create hello hello.js
```

It will build and create an action that you can invoke with

```
$ wsk action invoke hello 
```

It will build and deploy an action

<a name="under-the-hood"></a>
# How it works

When you run an `action create`, you will contact the controller that will gather the image, create a git repository and store the data in it. Then it triggers a Tekton Build.

The Tekton build uses as a source the git repository, builds an image using the provided git repository and stores the image in the registry.

Once done, it will start the image using Knative Serving.


<a name="limitations"></a>
# Current Limitations

Being mostly a proof-of-concept there are many limitations.

- it was tested only on a local Kubernetes cluster built with the provided kubepass script, that was tested so far only on Mac OSX and Ubuntu Linux; next step will be to port on Google GKS and Amazon EKS
- There is no authentication support yet
- The only runtimes are Nodejs 10, Go 1.12 and Python 3
- The only implemented operations so far are action creation and invocation

# Installation

[`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and the [`wsk`](https://github.com/apache/incubator-openwhisk-cli/releases).

## Get a Kubernetes cluster and a Registry

You need a Kubernetes Cluster and a Docker registry. 

Currently the only supported one is a locally built kubernetes cluster using `multipass` that also includes a registry as `registry.k8s:5000`.

### Build a local Kubernetes cluster with multipass

As a prerequisite you need either OSX (tested on High Sierra and Mojave) or Ubuntu Linux (tested on 18.x). 

Since the cluster eats 8 Gb of memory of its own it is unlikely you can run it on a machine with less than 16Gb of RAM.

Before running the script you need to install [`multipass`](https://github.com/CanonicalLtd/multipass/releases) v0.7.1 or later. 

Clone the repository and enter in the setup folder:

```
git clone https://github.com/sciabarracom/openwhisk-knative-operator
cd openwhisk-knative-operator/install
```

Run the `kubepass.sh` script to create the cluster and retrieve the configuration:

```
# optional step to preserve your kubernetes config if any
mv ~/.kube/config ~/.kube/config.saved
# now create the cluster
bash kubepassh.sh create
bash kubepass.sh config
```

As a check, executing the `multipass list` command you should see something like this:

```
Name                    State             IPv4             Image
kube-node2              Running           192.168.64.4     Ubuntu 18.04 LTS
kube-node1              Running           192.168.64.6     Ubuntu 18.04 LTS
kube-node3              Running           192.168.64.5     Ubuntu 18.04 LTS
kube-master             Running           192.168.64.3     Ubuntu 18.04 LTS
```

IPs are assigned by DHCP and may be different. Also note the cluster is not supposed to resist to a reboot so you you have to restart your server you will likely have to destroy the cluster `bash kubepass.sh destroy` and recreate it.

Also check you can connect to the cluster with `kubectl`:

```
$ kubectl get nodes
NAME          STATUS   ROLES    AGE    VERSION
kube-master   Ready    master   121m   v1.14.1
kube-node1    Ready    <none>   119m   v1.14.1
kube-node2    Ready    <none>   118m   v1.14.1
kube-node3    Ready    <none>   120m   v1.14.1
```

## Install Knative-Whisk and dependencies

Once you have a Kubernetes cluster and your `kubectl` can control it, apply the configuration with  and wait until istio, knative-serving and tekton are installed:

```
kubectl apply -f knative-whisk.yaml
POD=$(kubectl -n knative-whisk get po | tail +2 | awk '{print $1}')
kubectl -n knative-whisk logs $POD -f
```

Wait until you see:

```
Web Server starting at http://0.0.0.0:8000
```
Then you can setup  `wsk`

```
IP=$(multipass list | grep kube-node1 | awk '{ print $3}')
wsk property set -u 123:453 --apihost http://$IP:30080
```
Check it works with:

```
$ wsk namespace list
namespaces
local
```




