package helpers

import "testing"

func TestProcessFlagsPositive(t *testing.T) {

	flagsTofind := []string{"d", "n", "default"}
	result := ProcessFlags(flagsTofind, "n")

	if result[0] != "default" {
		t.Errorf("default not found!")
	}
}

func TestProcessFlagsNegative(t *testing.T) {

	flagsTofind := []string{"d", "n", "default"}
	result := ProcessFlags(flagsTofind, "z")

	if len(result) > 0 {
		t.Errorf("Length of result should be 0!")
	}
}

func BenchmarkProcessFlagsPositive(t *testing.B) {

	flagsTofind := []string{"d", "n", "default"}
	result := ProcessFlags(flagsTofind, "n")

	if result[0] != "default" {
		t.Errorf("default not found!")
	}
}

func BenchmarkProcessFlagsNegative(t *testing.B) {

	flagsTofind := []string{"d", "n", "default"}
	result := ProcessFlags(flagsTofind, "z")

	if len(result) > 0 {
		t.Errorf("Length of result should be 0!")
	}
}

func TestIsDryRunPositive(t *testing.T) {

	flagsTofind := []string{"d", "n", "default"}
	result := IsDryRun(flagsTofind)

	if result != true {
		t.Errorf("Dry run not found!")
	}
}

func TestIsDryRunNegative(t *testing.T) {

	flagsTofind := []string{"n", "default"}
	result := IsDryRun(flagsTofind)

	if result != false {
		t.Errorf("Dry run found!")
	}
}

func BenchmarkIsDryRunPositive(t *testing.B) {

	flagsTofind := []string{"d", "n", "default"}
	result := IsDryRun(flagsTofind)

	if result != true {
		t.Errorf("Dry run not found!")
	}
}

func BenchmarkIsDryRunNegative(t *testing.B) {

	flagsTofind := []string{"n", "default"}
	result := IsDryRun(flagsTofind)

	if result != false {
		t.Errorf("Dry run found!")
	}
}
