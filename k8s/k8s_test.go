package k8s

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	fake "k8s.io/client-go/dynamic/fake"
)

func TestConnect(t *testing.T) {

	//out of cluster
	var k KubernetesAPI = KubernetesAPI{Client: nil}

	k.Client = k.Connect()

	if k.Client == nil {
		t.Errorf("Clientset should not be nil")
	}
}

func TestGetDeployments(t *testing.T) {

	var k KubernetesAPI = KubernetesAPI{Client: fake.NewSimpleDynamicClient(runtime.NewScheme())}

	//no deployments
	result := k.GetDeployments("default")
	if result != nil {
		t.Errorf("result should be nil")
	}

	//one deployment - no replicas
	deploy := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"namespace": "default",
				"name":      "deployment",
			},
		},
	}

	k.Client = fake.NewSimpleDynamicClient(runtime.NewScheme(), deploy)

	result = k.GetDeployments("default")
	if result != nil {
		t.Errorf("result should be nil")
	}

	//error
	//TODO
}

func TestGetHorizontalPodAutoscalers(t *testing.T) {
	var k KubernetesAPI = KubernetesAPI{Client: fake.NewSimpleDynamicClient(runtime.NewScheme())}

	//no hpas
	result := k.GetHorizontalPodAutoscalers("default")
	if result != nil {
		t.Errorf("result should be nil")
	}

	//one hpa - no name
	hpa := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "autoscaling/v1",
			"kind":       "HorizontalPodAutoscaler",
			"metadata": map[string]interface{}{
				"namespace": "default",
			},
		},
	}

	//one hpa - no minReplicas
	hpa = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "autoscaling/v1",
			"kind":       "HorizontalPodAutoscaler",
			"metadata": map[string]interface{}{
				"namespace": "default",
				"name":      "hpa",
			},
		},
	}
	k.Client = fake.NewSimpleDynamicClient(runtime.NewScheme(), hpa)

	result = k.GetHorizontalPodAutoscalers("default")
	if result != nil {
		t.Errorf("result should be nil")
	}

	//one hpa - no scaleTargetRef
	var replicas int64
	replicas = 1
	hpa = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "autoscaling/v1",
			"kind":       "HorizontalPodAutoscaler",
			"metadata": map[string]interface{}{
				"namespace": "default",
				"name":      "hpa",
			},
			"spec": map[string]interface{}{
				"minReplicas": replicas,
			},
		},
	}
	k.Client = fake.NewSimpleDynamicClient(runtime.NewScheme(), hpa)

	result = k.GetHorizontalPodAutoscalers("default")
	if result != nil {
		t.Errorf("result should be nil")
	}

	//one hpa
	hpa = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "autoscaling/v1",
			"kind":       "HorizontalPodAutoscaler",
			"metadata": map[string]interface{}{
				"namespace": "default",
				"name":      "hpa",
			},
			"spec": map[string]interface{}{
				"minReplicas": replicas,
				"scaleTargetRef": map[string]interface{}{
					"name": "default",
				},
			},
		},
	}
	k.Client = fake.NewSimpleDynamicClient(runtime.NewScheme(), hpa)

	result = k.GetHorizontalPodAutoscalers("default")
	if result == nil {
		t.Errorf("result should not be nil")
	}
}

func TestScaleDeployments(t *testing.T) {
	// ScaleDeployments(deployments []unstructured.Unstructured, hpas []Hpa, scaleAction string, dryRun bool) {
}
