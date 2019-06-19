cd(@__DIR__)
push!(LOAD_PATH, pwd())
using Pkg
pkg"activate ."
pkg"instantiate"
pkg"build"
pkg"add https://github.com/essenciary/Genie.jl"
if length(ARGS) > 0
    Pkg.add(PackageSpec(url=ARGS[1]))
    using KnativeWhisk
end
