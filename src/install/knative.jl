SERVING="https://github.com/knative/serving/releases/download/v0.6.0"
APPLY="kubectl apply --selector knative.dev/crd-install=true"

function install_knative_serving()
    cmd = split("$APPLY -f $SERVING/serving.yaml")
    run(`$cmd`)
end
install_knative_serving()