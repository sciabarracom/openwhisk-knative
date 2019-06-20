module KnativeWhisk
using Genie
export install_knative, start_whisk, install_istio, install_knative_serving, install_tekton

include("install.jl")

include("routes.jl")

end # module
