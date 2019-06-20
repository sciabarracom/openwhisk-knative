using Logging
using Genie
import Genie.Renderer: json!

include("util.jl")

function namespace_exists(ns)
    sel = """{range .items[*]}{.metadata.name}{"\\n"}{end}'"""
    cmd = `kubectl get ns -o=jsonpath=$sel`
    for i in split(read_log(cmd))
        if i == ns
            @info "found namespace $ns"
            return true
        end
    end
    return false
end    

function check_all_pod_phase(namespace, state="Running")
    sel = """{range .items[*]}{.status.phase}{"\\n"}{end}"""
    cmd = `kubectl -n $namespace get pod -o jsonpath=$sel`
    for line in split(read_log(cmd))
        @debug "$line\n$state\n"
        if line != state
            return false
        end
    end
    return true
end

function count_crd(sub)
    cmd = `kubectl get crd -ojsonpath='{.items[*].metadata.name}'`
    crd = read_log(cmd, String)
    filtered = filter((s) -> occursin(sub, s), split(crd))
    return length(filtered)
end

ISTIO_VERSION = "1.1.7"
ISTIO_REPO="https://storage.googleapis.com/istio-release/releases/$ISTIO_VERSION/charts/"

ISTIO_NAMESPACE = """
        apiVersion: v1
        kind: Namespace
        metadata:
          name: istio-system
          labels:
            istio-injection: disabled       
        """

ISTIO_OPTIONS(node_type) = replace("""--namespace=istio-system
        --set prometheus.enabled=false
        --set mixer.enabled=false
        --set mixer.policy.enabled=false
        --set mixer.telemetry.enabled=false
        --set pilot.sidecar=false
        --set galley.enabled=false
        --set global.useMCP=false
        --set security.enabled=false
        --set global.disablePolicyChecks=true
        --set sidecarInjectorWebhook.enabled=false
        --set global.proxy.autoInject=disabled
        --set global.omitSidecarInjectorConfigMap=true
        --set gateways.istio-ingressgateway.autoscaleMin=1
        --set gateways.istio-ingressgateway.autoscaleMax=1
        --set pilot.traceSampling=100
        --set gateways.istio-ingressgateway.type=$node_type
        """, '\n' => ' ')


function install_istio(node_type="NodePort")
    # check namespace
    if namespace_exists("istio-system")
     return true
    end
    opts = split(ISTIO_OPTIONS(node_type))
    run_log(`helm repo add istio.io $ISTIO_REPO`)
    run_log(`helm fetch istio.io/istio-init --untar`)
    run(pipeline(IOBuffer(ISTIO_NAMESPACE), `kubectl apply -f -`))
    run(pipeline(`helm template $opts ./istio-init`, 
        `kubectl apply -f -`))
 
    retry(100, 10) do
        return count_crd("istio.io") == 53
    end

    retry(100,10) do
      check_all_pod_phase("istio-system", "Succeeded")
    end
end

function kind_exists(kind)
    try
        run(`kubectl get $kind`)
        return true
    catch
        return false
    end
end

function install_knative_serving()
    if namespace_exists("knative-serving")
        return true
    end

    SERVING_VERSION="v0.6.0"
    SERVING_URL="https://github.com/knative/serving/releases/download/$SERVING_VERSION"
    CRD_ONLY="--selector knative.dev/crd-install=true"

    cmd = split("kubectl apply $CRD_ONLY -f $SERVING_URL/serving.yaml")
    run_log(`$cmd`)
   
    retry(100, 10) do
        return count_crd("knative.dev") == 9
    end

    retry(100, 10) do
        kind_exists("Image")
    end
    
    cmd = split("kubectl apply -f $SERVING_URL/serving.yaml")
    run_log(`$cmd`)

    retry(100,10) do
       check_all_pod_phase("knative-serving", "Running") 
    end
end

function install_tekton()
    if namespace_exists("tekton-pipelines")
        return true
    end
    TEKTON_URL="https://storage.googleapis.com/tekton-releases/latest/release.yaml"

    cmd = split("kubectl apply -f $TEKTON_URL")
    run_log(`$cmd`)

    retry(100,10) do
        check_all_pod_phase("tekton-pipelines", "Running")
    end
end
 
function install_knative(node_type="NodePort") 
    install_istio(node_type)
    install_knative_serving()
    install_tekton()
end

