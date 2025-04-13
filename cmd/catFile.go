// Copyright © 2025 mook-jp <mook24.jp@gmail.com>
package cmd

import (
	"errors"
	"fmt"

	"github.com/mook-jp/mygit/internal/objectstore"
	"github.com/spf13/cobra"
)

var (
	printFlag bool
	typeFlag  bool
	sizeFlag  bool
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content of repository objects",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("you must provide an object hash")
		}
		hash := args[0]

		modeCount := 0
		if printFlag {
			modeCount++
		}
		if typeFlag {
			modeCount++
		}
		if sizeFlag {
			modeCount++
		}
		if modeCount != 1 {
			return errors.New("specify exactly one of -p, -t, or -s")
		}

		if printFlag {
			content, err := objectstore.ReadObjectContent(args[0])
			if err != nil {
				return err
			}
			fmt.Fprintln(OutWriter, string(content))
		}
		if typeFlag {
			objectType, err := objectstore.ReadObjectType(hash)
			if err != nil {
				return err
			}
			fmt.Fprintln(OutWriter, objectType)
		}
		if sizeFlag {
			size, err := objectstore.ReadObjectSize(hash)
			if err != nil {
				return err
			}
			fmt.Fprintln(OutWriter, size)
		}
		return nil
	},
}

func init() {
	catFileCmd.Flags().BoolVarP(&printFlag, "pretty", "p", false, "pretty-print the contents of object")
	catFileCmd.Flags().BoolVarP(&typeFlag, "type", "t", false, "Show object type")
	catFileCmd.Flags().BoolVarP(&sizeFlag, "size", "s", false, "Show object size")
	rootCmd.AddCommand(catFileCmd)
}
