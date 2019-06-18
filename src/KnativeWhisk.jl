module KnativeWhisk
using Genie

include("install.jl")
install_knative()

include("routes.jl")
Genie.config.run_as_server = true
Genie.config.server_host = "0.0.0.0"
Genie.startup()

end # module
