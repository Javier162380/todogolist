package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDeleteTaskCommand(t *testing.T) {
	testcases := []struct {
		Input    []string
		Output   string
		TestName string
	}{{
		Input:    []string{"delete", "--task-name", "test"},
		Output:   "Task Deleted Succesfully",
		TestName: "TaskName parameter",
	},
		{
			Input:    []string{"delete"},
			Output:   "Empty task nothing is deleted from the repository",
			TestName: "TaskName missing parameter",
		}}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Test Delete %s", tc.TestName), func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd, err := buildRootCommand("memory")
			rootCmd.SetArgs(tc.Input)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			err = rootCmd.Execute()
			assert.NoError(t, err)
			assert.Equal(t, tc.Output, b.String())
		})

	}

}
