package cmd

import (
	"fmt"

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

		walkErr := walkVault(vaultPath)

		if walkErr != nil {
			fmt.Println("Error walking vault:", walkErr)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCMD)
}