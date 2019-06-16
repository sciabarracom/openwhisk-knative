using Genie
module KnativeWhisk

include("routes.jl")

Genie.config.run_as_server = true
Genie.config.server_host = "0.0.0.0"
Genie.startup()

end # module
