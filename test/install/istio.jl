using Test
import Base.run

function run(cmd::Cmd, args...; wait=true) 
    @info join(cmd.exec, " ")
end

include("../../src/install/istio.jl")

@test_logs (:info, "helm repo add istio.io $ISTIO_REPO") install_istio()
