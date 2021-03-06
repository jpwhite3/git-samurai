/*
Copyright © 2022 JP White jp@jpw3.me

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/jpwhite3/git-samurai/dbutils"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: SyncCmdRunE,
}

func SyncCmdRunE(cmd *cobra.Command, args []string) error {

	// dbutils.DeleteDatabase()

	repo, err := git.PlainOpen(RepoPath)
	if err != nil {
		return err
	}

	// Retrieve the commit history
	cIter, err := repo.Log(&git.LogOptions{All: true})
	if err != nil {
		return err
	}

	// Store commit data
	err = cIter.ForEach(dbutils.InsertCommit)
	if err != nil {
		return err
	}

	headRef, err := repo.Head()
	if err != nil {
		return err
	}

	// dbutils.RecordBlame()
	c, err := object.GetCommit(repo.Storer, headRef.Hash())
	if err != nil {
		return err
	}
	dbutils.RecordBlame(c)

	cmd.Println("ok")
	return nil
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
