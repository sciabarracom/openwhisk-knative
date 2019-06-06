using Test
using Logging

disable_logging(Base.CoreLogging.Info)

@testset "KnativeWhisk" begin

include("actions/list.jl")
include("install/istio.jl")

end
