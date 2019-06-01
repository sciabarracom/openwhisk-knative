using Genie
import Genie.Router: route
import Genie.Router: @params

include("actions/list.jl")
route("/api/v1/namespaces/:namespace/actions") do
    namespace =  @params(:namespace)
    action_list(namespace) 
end

route("/") do
    "OpenWhiskKnative 0.0.1\n"
end

Genie.config.run_as_server = ! Base.isinteractive()
Genie.config.server_host = "0.0.0.0"
Genie.startup()
