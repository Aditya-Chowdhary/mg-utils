/*
Copyright © 2024 Aditya Chowdhary
*/
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap two migration files",
	Long: `Swap the position of two migration files, leaving the remaining untouched. For example:
	swap -i 5 2 -> Swap the migration files at position 5 and 2.
	swap 000003_migrate.up.sql 000010_migrate.up.sql -> Swap the migration files at position 3 and 10.

	Mixing integer and files names is not allowed. Eg. 
	swap -i 000003_migrate.up.sql 10 -> Not allowed
	`,
	Run: swapCommand,
}

func init() {
	rootCmd.AddCommand(swapCmd)
}

func swapCommand(cmd *cobra.Command, args []string) {
	dir, err := cmd.Flags().GetString("directory")
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error while parsing flag `directory`")
		return
	}

	in, err := cmd.Flags().GetBool("integer")
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error while parsing flag `integer`")
		return
	}

	if !checkSwapArguments(in, args) {
		fmt.Fprintln(cmd.ErrOrStderr(), "Incorrect Arguments")
		return
	}

	if err = checkDirectory(dir); err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), err)
		return
	}

	fmt.Fprintln(cmd.OutOrStderr(), "Success")
}

func checkDirectory(dir string) error {

	directory, err := os.Open(dir)
	if errors.Is(err, fs.ErrNotExist) {
		errMsg := fmt.Sprintf("The given directory [%s] does not exist\n", dir)
		return errors.New(errMsg)
	} else if err != nil {
		errMsg := fmt.Sprintf("Error while opening directory %s: %s\n", dir, err)
		return errors.New(errMsg)
	}

	dirinfo, err := directory.Stat()
	if err != nil {
		errMsg := fmt.Sprintf("Error while reading directory info %s: %s\n", dir, err)
		return errors.New(errMsg)
	}

	if !dirinfo.IsDir() {
		errMsg := fmt.Sprintf("%s is not a valid directory\n", dir)
		return errors.New(errMsg)
	}

	files, err := directory.Readdirnames(0)
	if err != nil {
		errMsg := fmt.Sprintf("Error while reading file names in %s: %s", dir, err)
		return errors.New(errMsg)
	}

	return nil
}

func checkSwapArguments(in bool, args []string) bool {
	if in {
		if !checkInt(args[0]) || !checkInt(args[1]) {
			return false
		}
	}

	return true
}
