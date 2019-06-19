cd(@__DIR__)
push!(LOAD_PATH, pwd())
using Pkg
pkg"activate ."
pkg"instantiate"
pkg"build"
pkg"add https://github.com/essenciary/Genie.jl"
if length(ARGS) > 1
    Pkg.add(PackageSpec(url=ARGS[1], rev=ARGS[2]))
    using KnativeWhisk
end
