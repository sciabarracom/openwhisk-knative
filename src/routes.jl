using Genie
using Genie.Router
using Genie.Renderer

route("/api/v1/namespaces") do
    return ["local"] |> json!
end

route("/api/v1/namespaces/:namespace/actions") do
    namespace =  @params(:namespace)
    dict = Dict(
        "name"        => "knative-install",
        "exec"        => Dict("binary"=>false),
        "namespace"   => "$namespace"
        )
    return [dict] |> json!
end

route("/api/v1/namespaces") do
    return ["local"] |> json!
end

route("/") do
    Dict("name" => "KnativeWhisk", "version" => "0.0.1") |> json!
end

