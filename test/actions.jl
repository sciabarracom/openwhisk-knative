using Test 

using Test

include("../src/actions/list.jl")

@testset "actions" begin
@test String(action_list("hello").body) == "[{\"name\":\"hello\",\"exec\":{\"binary\":false},\"namespace\":\"hello/install\"}]"
end
