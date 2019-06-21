module KnativeWhisk

include("install.jl")
include("routes.jl")
    
install_knative()

using Genie
Genie.config.run_as_server = true
Genie.config.server_host = "0.0.0.0"
Genie.startup()

end # module
