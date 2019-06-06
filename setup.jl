cd(@__DIR__)
push!(LOAD_PATH, pwd())
using Pkg
pkg"activate ."
pkg"instantiate"
pkg"build"
if ! Base.isinteractive()
    pkg"add https://github.com/sciabarracom/openwhisk-knative-operator#devel"
    using KnativeWhisk
end
