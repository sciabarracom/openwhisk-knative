module KnativeWhisk
using Genie
export install_knative, start_whisk

include("install.jl")

include("routes.jl")

end # module
