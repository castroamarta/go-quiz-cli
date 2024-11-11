package cmd

import (
	"bytes"
	"io/ioutil"
	"quiz/cmd"
	"testing"
)

func Test_ExecuteAuthCommand(t *testing.T) {
	t.Run("should return authentication successfull message", func(t *testing.T) {
		authCmd := cmd.NewAuthCmd()
		b := bytes.NewBufferString("")
		authCmd.SetOut(b)
		authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
		authCmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := "user authenticated successfully"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})

	t.Run("should return authentication failed message", func(t *testing.T) {
		authCmd := cmd.NewAuthCmd()
		b := bytes.NewBufferString("")
		authCmd.SetOut(b)
		authCmd.SetArgs([]string{"--username", "alice", "--password", "seasalt"})
		authCmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := "user authentication failed"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}