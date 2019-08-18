APIHOST=$(multipass list | grep kube-master | awk '{ print $3}')
curl -u registry:password -k https://$APIHOST/v2/_catalog
