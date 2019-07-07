module KnativeWhisk

include("install.jl")
include("routes.jl")

using Genie

start() = Genie.startup(8000, "0.0.0.0"; async=false)

export start
export install_knative

end # module

