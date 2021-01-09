package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCreateTaskCommand(t *testing.T) {

	testcases := []struct {
		Input  []string
		Output string
	}{{
		Input:  []string{"create", "--task-name", "test", "--task-label", "sport", "--task-periodicity", "w", "--task-date", "2020-10-10T00:00:00.000Z"},
		Output: "Task Created Succesfully",
	},
		{
			Input:  []string{"create", "--task-name", "test", "--task-label", "work", "--task-periodicity", "w", "--task-date", "2020-10-10T00:00:00.000Z"},
			Output: "Task Created Succesfully",
		},
		{
			Input:  []string{"create", "--task-name", "test", "--task-label", "sport", "--task-periodicity", "w", "--task-date", "2020-10-10T00:00:00.000Z", "-v"},
			Output: "Task Created Succesfully with: [{test sport 2020-10-17 00:00:00 +0000 UTC w 0} {test sport 2020-10-24 00:00:00 +0000 UTC w 0} {test sport 2020-10-31 00:00:00 +0000 UTC w 0} {test sport 2020-11-07 00:00:00 +0000 UTC w 0} {test sport 2020-11-14 00:00:00 +0000 UTC w 0} {test sport 2020-11-21 00:00:00 +0000 UTC w 0} {test sport 2020-11-28 00:00:00 +0000 UTC w 0} {test sport 2020-12-05 00:00:00 +0000 UTC w 0} {test sport 2020-12-12 00:00:00 +0000 UTC w 0} {test sport 2020-12-19 00:00:00 +0000 UTC w 0} {test sport 2020-12-26 00:00:00 +0000 UTC w 0} {test sport 2021-01-02 00:00:00 +0000 UTC w 0} {test sport 2021-01-09 00:00:00 +0000 UTC w 0} {test sport 2021-01-16 00:00:00 +0000 UTC w 0} {test sport 2021-01-23 00:00:00 +0000 UTC w 0} {test sport 2021-01-30 00:00:00 +0000 UTC w 0} {test sport 2021-02-06 00:00:00 +0000 UTC w 0} {test sport 2021-02-13 00:00:00 +0000 UTC w 0} {test sport 2021-02-20 00:00:00 +0000 UTC w 0} {test sport 2021-02-27 00:00:00 +0000 UTC w 0} {test sport 2021-03-06 00:00:00 +0000 UTC w 0} {test sport 2021-03-13 00:00:00 +0000 UTC w 0} {test sport 2021-03-20 00:00:00 +0000 UTC w 0} {test sport 2021-03-27 00:00:00 +0000 UTC w 0} {test sport 2021-04-03 00:00:00 +0000 UTC w 0} {test sport 2021-04-10 00:00:00 +0000 UTC w 0} {test sport 2021-04-17 00:00:00 +0000 UTC w 0} {test sport 2021-04-24 00:00:00 +0000 UTC w 0} {test sport 2021-05-01 00:00:00 +0000 UTC w 0} {test sport 2021-05-08 00:00:00 +0000 UTC w 0} {test sport 2021-05-15 00:00:00 +0000 UTC w 0} {test sport 2021-05-22 00:00:00 +0000 UTC w 0} {test sport 2021-05-29 00:00:00 +0000 UTC w 0} {test sport 2021-06-05 00:00:00 +0000 UTC w 0} {test sport 2021-06-12 00:00:00 +0000 UTC w 0} {test sport 2021-06-19 00:00:00 +0000 UTC w 0} {test sport 2021-06-26 00:00:00 +0000 UTC w 0} {test sport 2021-07-03 00:00:00 +0000 UTC w 0} {test sport 2021-07-10 00:00:00 +0000 UTC w 0} {test sport 2021-07-17 00:00:00 +0000 UTC w 0} {test sport 2021-07-24 00:00:00 +0000 UTC w 0} {test sport 2021-07-31 00:00:00 +0000 UTC w 0} {test sport 2021-08-07 00:00:00 +0000 UTC w 0} {test sport 2021-08-14 00:00:00 +0000 UTC w 0} {test sport 2021-08-21 00:00:00 +0000 UTC w 0} {test sport 2021-08-28 00:00:00 +0000 UTC w 0} {test sport 2021-09-04 00:00:00 +0000 UTC w 0} {test sport 2021-09-11 00:00:00 +0000 UTC w 0} {test sport 2021-09-18 00:00:00 +0000 UTC w 0} {test sport 2021-09-25 00:00:00 +0000 UTC w 0} {test sport 2021-10-02 00:00:00 +0000 UTC w 0} {test sport 2021-10-09 00:00:00 +0000 UTC w 0} {test sport 2020-10-10 00:00:00 +0000 UTC w 0}]",
		},
		{
			Input:  []string{"create", "--task-name", "test", "--task-label", "sport", "--task-periodicity", "m", "--task-date", "2020-10-10T00:00:00.000Z", "-v"},
			Output: "Task Created Succesfully with: [{test sport 2020-11-10 00:00:00 +0000 UTC m 0} {test sport 2020-12-10 00:00:00 +0000 UTC m 0} {test sport 2021-01-10 00:00:00 +0000 UTC m 0} {test sport 2021-02-10 00:00:00 +0000 UTC m 0} {test sport 2021-03-10 00:00:00 +0000 UTC m 0} {test sport 2021-04-10 00:00:00 +0000 UTC m 0} {test sport 2021-05-10 00:00:00 +0000 UTC m 0} {test sport 2021-06-10 00:00:00 +0000 UTC m 0} {test sport 2021-07-10 00:00:00 +0000 UTC m 0} {test sport 2021-08-10 00:00:00 +0000 UTC m 0} {test sport 2021-09-10 00:00:00 +0000 UTC m 0} {test sport 2021-10-10 00:00:00 +0000 UTC m 0} {test sport 2020-10-10 00:00:00 +0000 UTC m 0}]",
		},
	}

	for _, tc := range testcases {
		t.Run("Test Create", func(t *testing.T) {
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

func TestInvalidCreateTaskCommand(t *testing.T) {
	testcases := []struct {
		Input    []string
		Output   string
		TestCase string
	}{{
		Input:    []string{"create", "--task-label", "dummy"},
		Output:   `Task Label do not match a valid task label`,
		TestCase: `Invalid Task Label`,
	}, {
		Input:    []string{"create", "--task-date", "dummy"},
		Output:   `parsing time "dummy" as "2020-01-01": cannot parse "dummy" as "2"`,
		TestCase: `Invalid Task Date`,
	}}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Test Create %s", tc.TestCase), func(t *testing.T) {
			b := new(bytes.Buffer)
			rootCmd, _ := buildRootCommand("memory")
			rootCmd.SetArgs(tc.Input)
			rootCmd.SetOut(b)
			rootCmd.SetErr(b)
			rootCmd.Execute()
			assert.Equal(t, tc.Output, b.String())
		})

	}
}
