using Genie
import Genie.Router: route
import Genie.Router: @params

route("/api/v1/namespaces/:namespace/actions") do
    namespace =  @params(:namespace)
    action_list(namespace) 
end

route("/") do
    "Hello, world!\n"
end

