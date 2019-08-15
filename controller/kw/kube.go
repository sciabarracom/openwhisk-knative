package kw

import (
	"fmt"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
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

func parseK8sYaml(fileAsString string) []runtime.Object {

	//acceptedK8sTypes := regexp.MustCompile(`(Role|ClusterRole|RoleBinding|ClusterRoleBinding|ServiceAccount)`)
	sepYamlfiles := strings.Split(fileAsString, "---")
	retVal := make([]runtime.Object, 0, len(sepYamlfiles))
	for _, f := range sepYamlfiles {
		if f == "\n" || f == "" {
			// ignore empty cases
			continue
		}

		decode := scheme.Codecs.UniversalDeserializer().Decode
		obj, groupVersionKind, err := decode([]byte(f), nil, nil)

		if err != nil {
			log.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
			continue
		}

		/*if !acceptedK8sTypes.MatchString(groupVersionKind.Kind) {
			log.Printf("The custom-roles configMap contained K8s object types which are not supported! Skipping object with type: %s", groupVersionKind.Kind)
		} else {
		}*/
		log.Println(groupVersionKind)
		retVal = append(retVal, obj)
	}
	return retVal
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
