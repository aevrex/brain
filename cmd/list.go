package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCMD = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		vaultPath, err := getVaultPath()
		if err != nil {
			fmt.Println("Error getting vault path:", err)
			return
		}

		walkErr := walkVault(vaultPath, func(path string, d fs.DirEntry) error {
			if d.IsDir() {
				return nil
			}

			if filepath.Ext(path) != ".md" {
				return nil
			}

			relPath, err := filepath.Rel(vaultPath, path)
			if err != nil {
				return err
			}

			fmt.Println(relPath)

			return nil
		})

		if walkErr != nil {
			fmt.Println("Error walking vault:", walkErr)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCMD)
}