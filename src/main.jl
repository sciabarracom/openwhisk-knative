using Genie

include("actions/list.jl")
include("route.jl")

Genie.config.run_as_server = ! Base.isinteractive()
Genie.config.server_host = "0.0.0.0"
Genie.startup()
