package k8s

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/plugin/pkg/client/auth" //_ blank import for all auth packages
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

//K interace for k8s package
type K interface {
	Connect() dynamic.Interface
	GetDeployments(client dynamic.Interface) []unstructured.Unstructured
	GetHorizontalPodAutoscalers(client dynamic.Interface) []H
	ScaleDeployments(client dynamic.Interface, deployments []unstructured.Unstructured)
}

const (
	kubeSystem = "kube-system"
	//ScaleUp const for scaling up action
	ScaleUp = "UP"
	//ScaleDown const for scaling up action
	ScaleDown = "DOWN"
)

//H is the hpa struct
type H struct {
	name             string
	targetDeployment string
	minReplicas      int64
}

//Connect returns new kubernetes client
func Connect() dynamic.Interface {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return client
}

//GetDeployments gets all deployments from all namespaces except kube-system
func GetDeployments(client dynamic.Interface) []unstructured.Unstructured {

	var result []unstructured.Unstructured
	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	list, err := client.Resource(deploymentRes).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		if d.GetNamespace() != kubeSystem {
			_, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
			if err != nil || !found {
				fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
				continue
			}

			result = append(result, d)
		}
	}

	return result
}

//GetHorizontalPodAutoscalers gets all hpas from all namespaces except kube-system
func GetHorizontalPodAutoscalers(client dynamic.Interface) []H {

	var result []H
	hpaRes := schema.GroupVersionResource{Group: "autoscaling", Version: "v1", Resource: "horizontalpodautoscalers"}

	list, err := client.Resource(hpaRes).Namespace("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, h := range list.Items {
		if h.GetNamespace() != kubeSystem {
			name, found, err := unstructured.NestedString(h.Object, "metadata", "name")
			if err != nil || !found {
				fmt.Printf("name not found for hpa %s: error=%s", h.GetName(), err)
				continue
			}
			minReplicas, found, err := unstructured.NestedInt64(h.Object, "spec", "minReplicas")
			if err != nil || !found {
				fmt.Printf("minReplicas not found for hpa %s: error=%s", h.GetName(), err)
				continue
			}
			targetDeployment, found, err := unstructured.NestedString(h.Object, "spec", "scaleTargetRef", "name")
			if err != nil || !found {
				fmt.Printf("scaleTargetRef name not found for hpa %s: error=%s", h.GetName(), err)
				continue
			}

			var hpa = new(H)
			hpa.name = name
			hpa.minReplicas = minReplicas
			hpa.targetDeployment = targetDeployment

			result = append(result, *hpa)
		}
	}

	return result
}

//ScaleDeployments scales all deployments from all namespaces except kube-system either down to 0, or up to the minimum replicas if available(or 1)
func ScaleDeployments(client dynamic.Interface, deployments []unstructured.Unstructured, hpas []H, scaleAction string) {

	var updateRequired bool = false
	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	var newReplicas int64 = 0

	for _, d := range deployments {

		result, getErr := client.Resource(deploymentRes).Namespace(d.GetNamespace()).Get(context.TODO(), d.GetName(), metav1.GetOptions{})

		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		if replicas, _, _ := unstructured.NestedInt64(d.Object, "spec", "replicas"); replicas > 0 && scaleAction == ScaleDown {
			if err := unstructured.SetNestedField(result.Object, int64(0), "spec", "replicas"); err != nil {
				panic(fmt.Errorf("failed to set replica value: %v", err))
			}
			updateRequired = true
			newReplicas = 0
		} else if replicas, _, _ := unstructured.NestedInt64(d.Object, "spec", "replicas"); replicas == 0 && scaleAction == ScaleUp {

			//check for hpa and assing minReplicas otherwise replicas will be 1
			var minReplicas int64 = 1

			for _, h := range hpas {
				if h.targetDeployment == d.GetName() {
					minReplicas = h.minReplicas
					newReplicas = minReplicas
					break
				}
			}
			if err := unstructured.SetNestedField(result.Object, int64(minReplicas), "spec", "replicas"); err != nil {
				panic(fmt.Errorf("failed to set replica value: %v", err))
			}
			updateRequired = true
		}

		if updateRequired {
			updateRequired = false
			replicas, _, _ := unstructured.NestedInt64(d.Object, "spec", "replicas")
			_, updateErr := client.Resource(deploymentRes).Namespace(d.GetNamespace()).Update(context.TODO(), result, metav1.UpdateOptions{})
			fmt.Printf(" * %s (%d replicas) -> (%d replicas)\n", d.GetName(), replicas, newReplicas)
			if updateErr != nil {
				panic(fmt.Errorf("failed to set replica value: %v", updateErr))
			}
		}
	}
}
