module KnativeWhisk

include("install.jl")
include("routes.jl")

using Genie
Genie.config.run_as_server = true
Genie.config.server_host = "0.0.0.0"

export startup
export install_knative

end # module

