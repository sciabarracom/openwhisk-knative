using Test
import Base.run

function run(cmd::Cmd, args...; wait=true) 
    @info join(cmd.exec, " ")
end

include("../../src/install/istio.jl")

@test_logs (:info, "Info: helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.1.7/charts/")
(:info, "Info: kubectl create namespace istio-system")
(:info, "Info: helm fetch istio.io/istio-init --untar")
(:info, "Info: mkdir out")
(:info, "Info: helm template --name istio-init --namespace istio-system --output-dir out --set gateways.istio-ingressgateway.type=NodePort ./istio-init")
(:info, "Info: kubectl apply -f out/istio-init/templates")
install_istio()


install_istio()
