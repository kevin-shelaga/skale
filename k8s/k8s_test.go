package k8s

import "testing"

func TestConnect(t *testing.T) {

	Connect()

	if Client == nil {
		t.Errorf("Client should not be nil!")
	}
}

func TestGetDeployments(t *testing.T) {

	Connect()

	if Client == nil {
		t.Errorf("Client should not be nil!")
	}

	deployments := GetDeployments("default")

	if deployments == nil {
		t.Errorf("Deployments should not be nil!")
	}
}

func TestGetHorizontalPodAutoscalers(t *testing.T) {

	Connect()

	if Client == nil {
		t.Errorf("Client should not be nil!")
	}

	hpas := GetHorizontalPodAutoscalers("default")

	if hpas == nil {
		t.Errorf("Hpas should not be nil!")
	}
}

func TestScaleDeploymentsDown(t *testing.T) {

	Connect()

	if Client == nil {
		t.Errorf("Client should not be nil!")
	}

	deployments := GetDeployments("default")

	if deployments == nil {
		t.Errorf("Deployments should not be nil!")
	}

	ScaleDeployments(deployments, nil, ScaleDown, false)
}

func TestScaleDeploymentsUp(t *testing.T) {

	Connect()

	if Client == nil {
		t.Errorf("Client should not be nil!")
	}

	deployments := GetDeployments("default")

	if deployments == nil {
		t.Errorf("Deployments should not be nil!")
	}

	hpas := GetHorizontalPodAutoscalers("default")

	if hpas == nil {
		t.Errorf("Hpas should not be nil!")
	}

	ScaleDeployments(deployments, hpas, ScaleUp, false)
}
