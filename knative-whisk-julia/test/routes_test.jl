using Test 
using Genie
using HTTP
using JSON

include("../src/routes.jl")

Genie.startup()

function get_json(path, host="http://localhost:8000/api/v1/namespaces/local")
    url = "$host$path"
    headers = [("Content-Type", "application/json; charset=utf-8")]
    res = HTTP.request("GET", url, headers)
    body = String(res.body)
    JSON.Parser.parse(body)
end

function check_dict(dict::Dict, json::String)
    other = JSON.Parser.parse(json)
    for key in keys(other)
        if dict[key] != other[key]
            return false
        end
    end
    return true
end

@test check_dict(get_json("/", "http://localhost:8000"), 
"""{"name":"KnativeWhisk","version":"0.0.1"}""")

@test get_json("/api/v1/namespaces", "http://localhost:8000") == ["local"]

@test check_dict(get_json("/actions")[1],
"""{"name":"knative-install", "namespace":"local"}""")


