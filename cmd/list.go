package cmd

import (
	"fmt"
	"os"

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

		fmt.Println("Vault path:", vaultPath)

		entries, err := os.ReadDir(vaultPath)
		if err != nil {
			fmt.Println("Error reading vault:", err)
			return
		}

		for _, entry := range entries {
			fmt.Printf("Name: %s | IsDir: %t\n", entry.Name(), entry.IsDir())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCMD)
}