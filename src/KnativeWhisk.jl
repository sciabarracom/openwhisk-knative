module KnativeWhisk

include("routes.jl")

include("install/istio.jl")
include("install/knative.jl")

end # module
