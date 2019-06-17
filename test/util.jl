using Test

include("../src/util.jl")

@test_logs( 
    (:info, "not ready yet, retring in 1 second(s)"),
    (:info, "not ready yet, retring in 1 second(s)"),
begin
    n = 0
    retry(3) do
        n = n+1  
        return n == 3
    end
end)

@test_logs( 
    (:info, "not ready yet, retring in 1 second(s)"),
    (:info, "not ready yet, retring in 1 second(s)"),
    (:warn, "max retries (2) reached, exiting"),
retry(2) do
    return false
end)
