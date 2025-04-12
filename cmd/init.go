/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mook24/mygit/internal/initrepo"
	"github.com/spf13/cobra"
)

var initialBranch string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Git-like repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := initrepo.Options{
			BaseDir:       ".",
			InitialBranch: initialBranch,
			Stdout:        OutWriter,
			Stderr:        ErrWriter,
		}
		return initrepo.InitRepo(opts)
	},
}

func init() {
	initCmd.Flags().StringVar(&initialBranch, "initial-branch", "main", "initial branch name")
	rootCmd.AddCommand(initCmd)
}
