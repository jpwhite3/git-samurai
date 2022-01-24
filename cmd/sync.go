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
	"errors"
	"fmt"

	"time"

	"github.com/spf13/cobra"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")

		r, err := git.PlainOpen(".")

		CheckIfError(err)

		Info("git log")

		// ... retrieves the branch pointed by HEAD
		ref, err := r.Head()
		CheckIfError(err)

		// ... retrieves the commit history
		since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
		until := time.Date(2022, 7, 30, 0, 0, 0, 0, time.UTC)
		cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
		CheckIfError(err)

		// ... just iterates over the commits, printing it
		err = cIter.ForEach(func(c *object.Commit) error {
			fmt.Println(c)

			return nil
		})
		CheckIfError(err)
	},
}

func SyncCmdRunE(cmd *cobra.Command, args []string) error {
	option, err := cmd.Flags().GetBool("option")

	if err != nil {
		return err
	}

	if option {
		cmd.Println("ok")
		return nil
	}

	return errors.New("not ok")
}

func SyncCmdFlags(cmd *cobra.Command) {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.Flags().BoolP("option", "o", false, "Help message for option")
}

func init() {
	rootCmd.AddCommand(syncCmd)
	SyncCmdFlags(rootCmd)
}
