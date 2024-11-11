package cmd

import (
	"bytes"
	"io/ioutil"
	"quiz/cmd"
	"testing"
)

func Test_ExecuteGetQuestionCommand(t *testing.T) {

	authCmd := cmd.NewAuthCmd()
	authCmdBuff := bytes.NewBufferString("")
	authCmd.SetOut(authCmdBuff)
	authCmd.SetArgs([]string{"--username", "alice", "--password", "rainbow"})
	authCmd.Execute()

	t.Run("should return a quiz question successfully", func(t *testing.T) {
		
		questionCmd := cmd.NewGetQuestionCmd()
		questionCmdBff := bytes.NewBufferString("")
		questionCmd.SetOut(questionCmdBff)
		questionCmd.SetArgs([]string{"1"})
		questionCmd.Execute()
		actual, err := ioutil.ReadAll(questionCmdBff)
		if err != nil {
			t.Fatal(err)
		}
		expected := cmd.ParseQuestion(questionMock)
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})

	t.Run("should return invalid question ID error", func(t *testing.T) {
		questionCmd := cmd.NewGetQuestionCmd()
		b := bytes.NewBufferString("")
		questionCmd.SetOut(b)
		questionCmd.SetArgs([]string{"5"})
		questionCmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := "invalid question ID provided: 5"
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}

var questionMock = cmd.Question{
	ID:          "1",
	Description: "Which OS is more popular?",
	Options: []cmd.Option{
		{
			ID:          "a",
			Description: "MacOS",
		},
		{
			ID:          "b",
			Description: "Windows",
		},
		{
			ID:          "c",
			Description: "Linux",
		},
	},
}