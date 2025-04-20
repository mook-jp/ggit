/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/mook-jp/ggit/internal/repository"
	"github.com/spf13/cobra"
)

var (
	showGitDir   bool
	showTopLevel bool
)

// revParseCmd represents the revParse command
var revParseCmd = &cobra.Command{
	Use:   "rev-parse",
	Short: "A brief description of your command",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := repository.FindRepoRoot(".")
		if err != nil {
			return fmt.Errorf("not a ggit repository (or any parent dir): .mygit not found")
		}

		// --show-toplevel オプション: ルートディレクトリのパス
		if showTopLevel {
			fmt.Fprintln(OutWriter, repoRoot)
			return nil
		}

		// --git-dir オプション: .mygit ディレクトリのパス
		if showGitDir {
			goto SHOWGITDIR
		}

	SHOWGITDIR:
		fmt.Fprintln(OutWriter, filepath.Join(repoRoot, ".mygit"))
		return nil
	},
}

func init() {
	revParseCmd.Flags().BoolVar(&showGitDir, "git-dir", false, "show path to the .mygit directory")
	revParseCmd.Flags().BoolVar(&showTopLevel, "show-toplevel", false, "show path to the repository root")
	rootCmd.AddCommand(revParseCmd)
}
