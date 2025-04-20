/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mook-jp/ggit/internal/repository"
	"github.com/mook-jp/ggit/internal/tree"
	"github.com/spf13/cobra"
)

// writeTreeCmd represents the writeTree command
var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "Create a tree object from the working directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := repository.FindRepoRoot(".")
		if err != nil {
			return err
		}
		hash, err := tree.Write(repoRoot)
		if err != nil {
			return err
		}
		fmt.Fprintln(OutWriter, hash)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeTreeCmd)
}
