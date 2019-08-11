#!/bin/bash
CONF="$(dirname $0)/conf/knative-install.yaml"
if ! which kubectl
then echo "install kubectl please" ; exit 1
fi
if ! kubectl get nodes
then echo "configure access to a kubernetes cluster please" ; exit 1
fi 
kubectl apply -f "$CONF"
kubectl logs -f $(kubectl get po | grep knative-install | awk '{ print $1}') &
kubectl wait --timeout=10m --for=condition=complete job/knative-install
echo "*** Namespaces:"
kubectl get ns
echo "*** Check you have istio-system, knative-serving and tekton-pipelines"
