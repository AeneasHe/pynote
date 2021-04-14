package server

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestWalk(t *testing.T) {

	c := testcli.Command("greetings", "--name", "John")
	c.Run()
	if !c.Success() {
		t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
	}

	if !c.StdoutContains("Hello John!") {
		t.Fatalf("Expected %q to contain %q", c.Stdout(), "Hello John!")
	}

}
