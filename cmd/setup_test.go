package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupCommand(t *testing.T) {
	testcases := []struct {
		Input  []string
		Output string
	}{{
		Input:  []string{"setup"},
		Output: "DB Setup Succesfully",
	},
		{
			Input:  []string{"setup", "-d"},
			Output: "BEGIN;DB Setup Succesfully;COMMIT",
		},
		{
			Input:  []string{"setup", "--dry-run"},
			Output: "BEGIN;DB Setup Succesfully;COMMIT",
		}}

	for _, testcase := range testcases {
		t.Run("Test Setup", func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd, err := buildRootCommand("memory")
			rootCmd.SetArgs(testcase.Input)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			err = rootCmd.Execute()

			assert.NoError(t, err)
			assert.Equal(t, testcase.Output, b.String())

		})

	}

}
