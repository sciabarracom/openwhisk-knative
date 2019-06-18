cd(@__DIR__)
push!(LOAD_PATH, pwd())
using Pkg
pkg"activate ."
pkg"instantiate"
pkg"build"
if length(ARGS) > 0
    Pkg.add(PackageSpec(url=ARGS[1]))
    using KnativeWhisk
end
