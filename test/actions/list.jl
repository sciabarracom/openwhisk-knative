using Test 

include("../../src/actions/list.jl")

@test String(action_list("hello").body) == "[{\"name\":\"hello\",\"exec\":{\"binary\":false},\"namespace\":\"hello/install\"}]"
