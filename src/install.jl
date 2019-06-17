using Logging
using Genie
import Genie.Renderer: json!

function run_log(cmd::Cmd) 
    @info join(cmd.exec, " ")
    run(cmd)
end

function read_log(cmd::Cmd, args)
    @info join(cmd.exec, " ")
    return read(cmd, String)
end

KUBECTL_NS=`kubectl get ns -o=jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}'`

function namespace_exists(ns)
    for i in split(read_log(KUBECTL_NS))
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
   run_log(`helm repo add istio.io $ISTIO_REPO`)
   run_log(`kubectl create namespace istio-system`)
   run_log(`helm fetch istio.io/istio-init --untar`)
   run_log(`mkdir out`)
   run_log(`helm template --name istio-init --namespace istio-system --output-dir out --set gateways.istio-ingressgateway.type=$node_type ./istio-init`)
   run_log(`kubectl apply -f out/istio-init/templates`)
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
    run_log(`$cmd`)
    cmd = split("kubectl apply -f $SERVING_URL/serving.yaml")
    run_log(`$cmd`)
    return true
end

TEKTON_URL="https://storage.googleapis.com/tekton-releases/latest/release.yaml"

function install_tekton()
    if namespace_exists("tekton-pipelines")
        return true
    end
    cmd = split("kubectl apply -f $TEKTON_URL")
    run_log(`$cmd`)
    return true
end

function install_knative(node_type="NodePort") 
    install_istio(node_type)
    install_knative_serving()
    install_tekton()
end

function check_all_pod_phase(namespace, state="Running")
    selector = """{range .items[*]}{.status.phase}{"\\n"}{end}"""
    cmd = `kubectl -n $namespace get pod -o jsonpath=$selector`
    for line in split(read_log(cmd))
        if line != state
            return false
        end
    end
    return true
end


