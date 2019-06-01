ISTIO_VERSION = "1.1.7"
ISTIO_REPO="https://storage.googleapis.com/istio-release/releases/$ISTIO_VERSION/charts/"
ISTIO_OPTS="
ISTIO_ARGS(node_type)="

function install_istio(node_type="NodePort")
   run(`helm repo add istio.io $ISTIO_REPO`)
   run(`kubectl create namespace istio-system`)
   run(`helm fetch istio.io/istio-init --untar`)
   run(`mkdir out`)
   run(`helm template --name istio-init --namespace istio-system --output-dir out --set gateways.istio-ingressgateway.type=$node_type ./istio-init`)
   run(`kubectl apply -f out/istio-init/templates`)
   return true
end
