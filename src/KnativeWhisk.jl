module KnativeWhisk

include("install.jl")
include("routes.jl")

if haskey(ENV, "INSTALL_KNATIVE")    
    install_knative()
end

if haskey(ENV, "START_KNATIVE_WHISK")
    using Genie
    Genie.startup(8000, "0.0.0.0"; async=false)
end

end # module
