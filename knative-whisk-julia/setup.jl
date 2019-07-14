cd(@__DIR__)
using Pkg
pkg"add https://github.com/essenciary/Genie.jl"
if length(ARGS) > 1
    Pkg.add(PackageSpec(url=ARGS[1], rev=ARGS[2]))
else
    push!(LOAD_PATH, "$(pwd())/src")
end
pkg"build"
