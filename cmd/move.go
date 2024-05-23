/*
Copyright © 2024 Aditya Chowdhary
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a migration file from one sequence number to another",
	Long: `Move a migration file from one sequence number to another, displacing the files in between. For example: 
	move -i 5 2 -> Move a migration file at position 5 to position 2, increasing the position of all the files in between
	move 000003_migrate.up.sql 10 -> Move the given migration file to position 10, decreasing the position of all files in between 
`,
	Run:  moveCommand,
	Args: cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func moveCommand(cmd *cobra.Command, args []string) {
	in, err := cmd.Flags().GetBool("integer")
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Error while parsing flag")
		return
	}

	if !checkMoveArguments(in, args) {
		fmt.Fprintln(cmd.ErrOrStderr(), "Incorrect Arguments")
		return
	}

	fmt.Fprintln(cmd.ErrOrStderr(), "Success")
}

func checkMoveArguments(in bool, args []string) bool {
	if !checkInt(args[1]) {
		return false
	}

	if in {
		if !checkInt(args[0]) {
			return false
		}
	}

	return true
}

func checkInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
