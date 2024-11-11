package cmd

import (
	"bytes"
	"io/ioutil"
	"quiz/cmd"
	"testing"
)

func Test_ExecuteSelectOptionsCommand(t *testing.T) {
	t.Run("should select an option successfully", func(t *testing.T) {
		selectOptionsCmd := cmd.NewSelectOptionsCommand()
		b := bytes.NewBufferString("")
		selectOptionsCmd.SetOut(b)
		selectOptionsCmd.SetArgs([]string{"1:b"})
		selectOptionsCmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := cmd.ParseSelectedOptions([]cmd.Answer{answersMock[0]})
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
	t.Run("should select all quiz options successfully", func(t *testing.T) {
		selectOptionsCmd := cmd.NewSelectOptionsCommand()
		b := bytes.NewBufferString("")
		selectOptionsCmd.SetOut(b)
		selectOptionsCmd.SetArgs([]string{"1:b,2:b,3:a"})
		selectOptionsCmd.Execute()
		actual, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		expected := cmd.ParseSelectedOptions(answersMock)
		if string(actual) != expected {
			t.Fatalf("expected \"%s\" got \"%s\"", expected, string(actual))
		}
	})
}

var answersMock = []cmd.Answer{
	{
		QuestionID: "1",
		OptionID: "b",
	},
	{
		QuestionID: "2",
		OptionID: "b",
	},
	{
		QuestionID: "3",
		OptionID: "a",
	},
}