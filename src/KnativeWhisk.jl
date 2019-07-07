module KnativeWhisk

if haskey(ENV, "INSTALL_KNATIVE_ONLY")    
    include("install.jl")
    install_knative()
else
    include("routes.jl")
    using Genie
    Genie.startup(8000, "0.0.0.0"; async=false)
end

end # module
