package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var listCMD = &cobra.Command{
	Use:   "list",
	Run: func(cmd *cobra.Command, args []string) {
		vaultPath, _ := getVaultPath()

		entries, _ := os.ReadDir(vaultPath)

		for _, entry := range entries {
			fmt.Printf("Name: %s | IsDir: %t\n", entry.Name(), entry.IsDir())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCMD)
}
