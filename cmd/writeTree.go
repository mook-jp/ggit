/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mook-jp/ggit/internal/tree"
	"github.com/spf13/cobra"
)

// writeTreeCmd represents the writeTree command
var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "Create a tree object from the working directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		hash, err := tree.Write(".")
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
