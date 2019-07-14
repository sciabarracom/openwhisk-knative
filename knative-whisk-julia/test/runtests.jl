using Test
using Logging

global_logger(NullLogger())

@testset "Util" begin 
  include("util_test.jl")
end

@testset "Install" begin 
  #include("install_test.jl")
end

@testset "Routes" begin 
    include("routes_test.jl")
end
