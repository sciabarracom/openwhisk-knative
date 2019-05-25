using Genie
import Genie.Renderer: json!

function action_list(namespace)

    d1 = Dict(
    "name"        => "hello",
    "exec"        => Dict("binary"=>false),
    "namespace"   => "$namespace/install"
    )
 
    return [d1] |> json!
end