using Genie

include("actions/list.jl")
include("route.jl")

function start()
    Genie.config.run_as_server = ! Base.isinteractive()
    Genie.startup()
end
