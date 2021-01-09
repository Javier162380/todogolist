package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandsUses(t *testing.T) {
	testcases := []string{
		"create",
		"get",
		"setup",
		"delete",
		"list",
	}

	for _, tc := range testcases {
		t.Run(tc, func(t *testing.T) {
			osargs := strings.Split(tc, " ")
			rootCmd, err := buildRootCommand("memory")
			cmd, _, err := rootCmd.Find(osargs)
			assert.NoError(t, err)
			assert.Equal(t, osargs[len(osargs)-1], cmd.Name())
		})
	}
}
