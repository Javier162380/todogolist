package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidListTaskCommand(t *testing.T) {

	testcases := []struct {
		Input    []string
		Output   string
		TestName string
	}{{
		Input:    []string{"list", "-L", "2"},
		Output:   "TaskID  \tTaskName            \tTaskLabel      \tTaskDate                 \t\tPeriodicity \tTaskTimeInvested    \t\n1234    \ttest1               \twork           \t2020-01-01 10:10:10.00000\t\tw           \t600                 \t\n1235    \ttest2               \tsport          \t2020-02-10 10:10:10.00000\t\t            \t10                  \t\n1236    \ttest3               \ttech           \t2020-08-01 10:10:10.00000\t\tw           \t10                  \t\n1234    \ttest1               \twork           \t2020-01-01 10:10:10.00000\t\tw           \t600                 \t\n1235    \ttest2               \tsport          \t2020-02-10 10:10:10.00000\t\t            \t10                  \t\n1236    \ttest3               \ttech           \t2020-08-01 10:10:10.00000\t\tw           \t10                  \t\n",
		TestName: "Valid List command with Limit",
	}, {
		Input:    []string{"list"},
		Output:   "TaskID  \tTaskName            \tTaskLabel      \tTaskDate                 \t\tPeriodicity \tTaskTimeInvested    \t\n1234    \ttest1               \twork           \t2020-01-01 10:10:10.00000\t\tw           \t600                 \t\n1235    \ttest2               \tsport          \t2020-02-10 10:10:10.00000\t\t            \t10                  \t\n1236    \ttest3               \ttech           \t2020-08-01 10:10:10.00000\t\tw           \t10                  \t\n1234    \ttest1               \twork           \t2020-01-01 10:10:10.00000\t\tw           \t600                 \t\n1235    \ttest2               \tsport          \t2020-02-10 10:10:10.00000\t\t            \t10                  \t\n1236    \ttest3               \ttech           \t2020-08-01 10:10:10.00000\t\tw           \t10                  \t\n",
		TestName: "Valid List command without limit",
	}}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Test List %s", tc.TestName), func(t *testing.T) {
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
