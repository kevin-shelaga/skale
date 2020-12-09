package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRootCmdExecute(t *testing.T) {
	cmd := rootCmd
	cmd.Execute()
}

func TestHelpRootCmdExecute(t *testing.T) {
	cmd := rootCmd
	cmd.SetArgs([]string{"--helpss"})

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "" {
		t.Fatalf("expected \"%s\" got \"%s\"", "", string(out))
	}
}

func TestDownCmdExecute(t *testing.T) {
	cmd := downCmd
	cmd.Execute()
}

func TestUpCmdExecute(t *testing.T) {
	cmd := downCmd
	cmd.Execute()
}
