include("installer.jl")

node_type="NodePort"

install_istio(node_type)
install_knative_serving()
install_tekton()
