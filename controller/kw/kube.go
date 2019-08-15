package kw

import (
	homedir "github.com/mitchellh/go-homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Kube is a kubernetes clieent
type Kube struct {
	Config *rest.Config
	Client *kubernetes.Clientset
}

// NewKube create a new Kubernetes client
func NewKube() *Kube {
	kubeConfig, err := rest.InClusterConfig()
	if err != nil {
		var config string
		config, err = homedir.Expand("~/.kube/config")
		if err == nil {
			kubeConfig, err = clientcmd.BuildConfigFromFlags("", config)
		}
	}
	FatalIf(err)
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	FatalIf(err)
	return &Kube{
		Config: kubeConfig,
		Client: kubeClient,
	}
}

// ListNamespaces lists namespaces
func (k *Kube) ListNamespaces() []string {
	v1 := k.Client.CoreV1()
	res := []string{}
	nsl, err := v1.Namespaces().List(metav1.ListOptions{})
	LogIf(err)
	if err != nil {
		return res
	}
	for _, ns := range nsl.Items {
		res = append(res, ns.Name)
	}
	return res
}


func