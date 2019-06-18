cd(@__DIR__)
push!(LOAD_PATH, pwd())
using Pkg
pkg"activate ."
pkg"instantiate"
pkg"build"
if ! Base.isinteractive()
    branch = ENV["BRANCH"]
    Pkg.add("https://github.com/sciabarracom/openwhisk-knative-operator#$branch")
    using KnativeWhisk
end
