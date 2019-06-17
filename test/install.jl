using Test
import Base.run
import Logging

include("../src/install.jl")

# disable acutual run - just check log
function run(cmd::Cmd, args...; wait = true) 
end

# mock read with provided values
mock_read = [""]

function read(cmd::Cmd, String)
    global mock_read
    r = mock_read[1]
    mock_read = cat(mock_read[2:end], r; dims=1)
    return r
end

# tests

mock_read = ["default\n"]

@test_logs( 
    (:info, "kubectl get ns -o=jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"),
    (:info, "helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.1.7/charts/"),
    (:info, "kubectl create namespace istio-system"),
    (:info, "helm fetch istio.io/istio-init --untar"),
    (:info, "mkdir out"),
    (:info, "helm template --name istio-init --namespace istio-system --output-dir out --set gateways.istio-ingressgateway.type=NodePort ./istio-init"),
    (:info, "kubectl apply -f out/istio-init/templates"), 
    install_istio()
    )

@test_logs( 
    (:info, "kubectl get ns -o=jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"),
    (:info, "kubectl apply --selector knative.dev/crd-install=true -f https://github.com/knative/serving/releases/download/v0.6.0/serving.yaml"),
    (:info, "kubectl apply -f https://github.com/knative/serving/releases/download/v0.6.0/serving.yaml"), 
    install_knative_serving()
    )

@test_logs( 
    (:info, "kubectl get ns -o=jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"),
    (:info, "kubectl apply -f https://storage.googleapis.com/tekton-releases/latest/release.yaml"),
    install_tekton()
    )

mock_read = ["default\nistio-system\ntekton-pipelines\nknative-serving\n"]

@test_logs( 
    (:info, "kubectl get ns -o=jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"),
    (:info, "found namespace istio-system"),
    (:info, "kubectl get ns -o=jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"),
    (:info, "found namespace knative-serving"),
    (:info, "kubectl get ns -o=jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}"),
    (:info, "found namespace tekton-pipelines"),
    install_knative())
