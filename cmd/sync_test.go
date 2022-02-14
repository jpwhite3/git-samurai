package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestSyncCmd(t *testing.T) {
	sync := &cobra.Command{Use: "sync", RunE: SyncCmdRunE}
	RootCmdFlags(sync)
	err := sync.Execute()
	assert.Error(t, err)
}

func TestSyncCmdFlags(t *testing.T) {

	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: []string{"-r arbitrary"},
			err:  errors.New("repository does not exist"),
		},
		{
			args: []string{"--repo-path=arbitrary"},
			err:  errors.New("repository does not exist"),
		},
	}

	sync := &cobra.Command{Use: "sync", RunE: SyncCmdRunE}
	RootCmdFlags(sync)

	for _, tc := range tt {
		out, err := execute(t, sync, tc.args...)
		assert.Equal(t, tc.err, err)
		if tc.err == nil {
			assert.Equal(t, tc.out, out)
		}
	}
}
