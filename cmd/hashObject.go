/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mook24/mygit/internal/objectstore"
	"github.com/spf13/cobra"
)

var write bool

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object <file>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hash, err := objectstore.HashObject(args[0], write, OutWriter)
		if err != nil {
			return err
		}
		fmt.Fprintln(OutWriter, hash)
		return nil
	},
}

func init() {
	hashObjectCmd.Flags().BoolVarP(&write, "write", "w", false, "Actually write the object into the object database")
	rootCmd.AddCommand(hashObjectCmd)
}
