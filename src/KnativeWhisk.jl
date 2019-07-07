module KnativeWhisk

include("routes.jl")

using Genie
Genie.startup(8000, "0.0.0.0"; async=false)

end # module
