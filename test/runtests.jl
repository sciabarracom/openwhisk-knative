using Test
using Logging

@testset "Install" begin 
  include("install.jl")
end

@testset "Routes" begin 
    include("routes.jl")
end
