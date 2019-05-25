cd(@__DIR__)
using Pkg
pkg"activate ."

using Genie

include("src/actions/list.jl")
include("src/route.jl")

if haskey(ENV, "AS_SERVER")
    Genie.config.run_as_server = true
end
Genie.startup()
