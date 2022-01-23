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

func TestSyncCmdToggles(t *testing.T) {

	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: nil,
			err:  errors.New("not ok"),
		},
		{
			args: []string{"-o"},
			err:  nil,
			out:  "ok",
		},
		{
			args: []string{"--option"},
			err:  nil,
			out:  "ok",
		},
	}

	root := &cobra.Command{Use: "sync", RunE: SyncCmdRunE}
	SyncCmdFlags(root)

	for _, tc := range tt {
		out, err := execute(t, root, tc.args...)

		assert.Equal(t, tc.err, err)

		if tc.err == nil {
			assert.Equal(t, tc.out, out)
		}
	}
}
