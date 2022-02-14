package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestRootCmd(t *testing.T) {
	root := &cobra.Command{Use: "root", RunE: RootCmdRunE}
	RootCmdFlags(root)
	err := root.Execute()
	assert.Error(t, err)

}

func TestRootCmdToggles(t *testing.T) {

	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: nil,
			err:  errors.New("not implemented"),
		},
		{
			args: []string{"-r"},
			err:  errors.New("flag needs an argument: 'r' in -r"),
		},
		{
			args: []string{"--repo-path"},
			err:  errors.New("flag needs an argument: --repo-path"),
		},
		{
			args: []string{"--arbitrary"},
			err:  errors.New("unknown flag: --arbitrary"),
		},
	}

	root := &cobra.Command{Use: "root", RunE: RootCmdRunE}
	RootCmdFlags(root)

	for _, tc := range tt {
		out, err := execute(t, root, tc.args...)
		assert.Equal(t, tc.err, err)
		if tc.err == nil {
			assert.Equal(t, tc.out, out)
		}
	}
}
