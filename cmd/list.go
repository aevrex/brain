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

		err = filepath.WalkDir(vaultPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			fmt.Println(path)
			return nil
		})

		if err != nil {
			fmt.Println("Error walking vault:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCMD)
}