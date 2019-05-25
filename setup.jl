cd(@__DIR__)
using Pkg
Pkg.activate(".")
pkg"add https://github.com/essenciary/Genie.jl"
pkg"add https://github.com/sciabarracom/openwhisk-knative-operator#devel"
Pkg.test("OpenWhiskKnative")
Pkg.build()

