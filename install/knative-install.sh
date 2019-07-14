#!/bin/bash
if ! which kubectl
then echo "install kubectl please" ; exit 1
fi
if ! kubectl get nodes
then echo "configure access to a kubernetes cluster please" ; exit 1
fi 
kubectl apply -f knative-install.yaml
echo "Waiting for knative installation to complete..."
kubectl wait --timeout=10m --for=condition=complete job/knative-install
echo "Check you have istio-system, knative-serving and tekton-pipelines:"
kubectl get ns
