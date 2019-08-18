APIHOST=$(multipass list | grep kube-node1 | awk '{ print $3}')
wsk -u 123:456 --apihost http://$APIHOST:30808 "$@"
