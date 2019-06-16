using Logging
using Genie
import Genie.Renderer: json!

KUBECTL_NS=`kubectl get ns -o=jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}'`

function namespace_exists(ns)
    for i in split(read(KUBECTL_NS, String))
        if i == ns
            @info "found namespace $ns"
            return true
        end
    end
    return false
end    

ISTIO_VERSION = "1.1.7"
ISTIO_REPO="https://storage.googleapis.com/istio-release/releases/$ISTIO_VERSION/charts/"

function install_istio(node_type="NodePort")
   if namespace_exists("istio-system")
    return true
   end
   run(`helm repo add istio.io $ISTIO_REPO`)
   run(`kubectl create namespace istio-system`)
   run(`helm fetch istio.io/istio-init --untar`)
   run(`mkdir out`)
   run(`helm template --name istio-init --namespace istio-system --output-dir out --set gateways.istio-ingressgateway.type=$node_type ./istio-init`)
   run(`kubectl apply -f out/istio-init/templates`)
   return true
end

SERVING_VERSION="v0.6.0"
SERVING_URL="https://github.com/knative/serving/releases/download/$SERVING_VERSION"
CRD_ONLY="--selector knative.dev/crd-install=true"

function install_knative_serving()
    if namespace_exists("knative-serving")
        return true
       end
    cmd = split("kubectl apply $CRD_ONLY -f $SERVING_URL/serving.yaml")
    run(`$cmd`)
    cmd = split("kubectl apply -f $SERVING_URL/serving.yaml")
    run(`$cmd`)
    return true
end

TEKTON_URL="https://storage.googleapis.com/tekton-releases/latest/release.yaml"

function install_tekton()
    if namespace_exists("tekton-pipelines")
        return true
    end
    cmd = split("kubectl apply -f $TEKTON_URL")
    run(`$cmd`)
    return true
end

function install_knative(node_type="NodePort") 
    install_istio(node_type)
    install_knative_serving()
    install_tekton()
end
