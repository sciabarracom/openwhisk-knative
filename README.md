# OpenWhisk Knative

Warning: this is still work in progress.

This is a proof-of-concept implementation of OpenWhisk on top of Knative. It tries to be compatible with the current user experience of OpenWhisk users implementing it on top of Knative.

It actually uses Knative Serving version 0.6 with Istio and Tekton Pipelines. An installer of the prerequisites is included in the kit.

In the current state it has many [limitations](#limitations) See below.

Once you have installed both the Kubernetes cluster and the knative-whisk operator and configures your wsk command, you can do:

``` bash
$ wsk action create hello hello.js
```

It will build and create an action that you can invoke with

```
$ wsk action invoke hello
```

<a name="under-the-hood"></a>
## How it works

When you run an `action create`, you will contact the controller that will gather the image, create a git repository and store the data in it. Then it triggers a Tekton Build.

The Tekton build uses as a source the git repository, builds an image using the provided git repository and stores the image in the registry.

Once done, it will start the image using Knative Serving.

<a name="limitations"></a>
## Current Limitations

Being mostly a proof-of-concept there are many limitations.

- it was tested only on a local Kubernetes cluster built with the provided kubepass script, that was tested so far only on Mac OSX and Ubuntu Linux; next step will be to port on Google GKS and Amazon EKS
- There is no authentication support yet
- The only runtimes are Nodejs 10, Go 1.12 and Python 3
- The only implemented operations so far are action creation and invocation

# Installation

Let's see the installation:

- installing prerequisites
- installing a kubernetes cluster
- installing knative
- installint the knative-whisk operator

## Prerequisites

Install first the command line tools  [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and [`wsk`](https://github.com/apache/incubator-openwhisk-cli/releases).

If you want to install a local development cluster using the provided script below, you need to install [`multipass`](https://github.com/CanonicalLtd/multipass/releases) v0.7.1 or later. 

## Get a Kubernetes cluster and a Registry

You need a Kubernetes Cluster and a Docker registry. 

Currently the only supported one is a locally built kubernetes cluster using `multipass` that also includes a registry as `registry.k8s`.

### Build a local Kubernetes cluster with multipass

The provided `kubepass.sh` is able to create a Kubernetes cluster using multipass.

You need either OSX (tested on High Sierra and Mojave) or Ubuntu Linux (tested on 18.x). It can work also on Windows but is not tested.

The default is to build a cluster with 1 master and 2 workers with 2Gb of memory each. So you need at least 16Gb for your machine. THe master consumes 2 CPU while the workers 1 VCP each to you need at least 4 VCPU available. Note that if you have 4 cores, then you have 8 VCPU available.

On Mac, if you use Docker for Desktop, it consumes by default 4 VCPU so you need to turn it off or reduce the usage to 2 VCPU otherwise you will experience failures in creation of multipass instances.

It also uses 15GB for each virtual machine so you need at least 60GB of disk space available.

Once you checked the prerequisites, clone the repository and enter in the setup folder:

```
git clone https://github.com/sciabarracom/openwhisk-knative-operator
cd openwhisk-knative-operator/install
```

Run the `kubepass.sh` script to create the cluster 

```
bash kubepassh.sh create
```

It will take a while to complete (up to 10 minutes). Once you have finished you can retrieve the Kubernetes configuration file with:

```
# optional step to preserve your kubernetes config if any
mv ~/.kube/config ~/.kube/config.saved
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

IPs are assigned by DHCP and may be different than in the example. 

Also check you can connect to the cluster with `kubectl`:

```
$ kubectl get nodes
NAME          STATUS   ROLES    AGE    VERSION
kube-master   Ready    master   121m   v1.14.1
kube-node1    Ready    <none>   119m   v1.14.1
kube-node2    Ready    <none>   118m   v1.14.1
kube-node3    Ready    <none>   120m   v1.14.1
```

### Notes on managing multipass vms

Multipass is used for development purposes only.

The cluster is not supposed to survive to a reboot so if you have to restart your server you will likely have to destroy the cluster `bash kubepass.sh destroy` and recreate it.

On OSX If your computer goes to sleep, multipass virtual machines goes to sleep too. When you wake it they may be unresponsive. You can do a quick check of the state with `kubectl get nodes`. Generally a `multipass start kube-master` is enough to wake up the cluster but it can take a bit.

## Install Knative

Once you have a Kubernetes cluster you need to install Istio, Knative-Serving and Tekton. 

You can do it quickly using the provided `knative-install.yaml` batch job under `install`.

```
kubectl apply -f knative-install.yaml
kubectl wait --timeout=10m --for=condition=complete job/knative-install
```

After the installation you should see the namespaces `istio-system`, `knative-serving` and `tekton-pipelines` when listing with `kubectl get ns`.

## Install Knative-Whisk

Once you have a Kubernetes cluster  with Knative and your `kubectl` can control it, apply the configuration:

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
