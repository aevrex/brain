package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sahilm/fuzzy"

)

var openCmd = &cobra.Command{
	Use:   "open [search]",
	Short: "Open the closest matching note in Notepad",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultPath, err := getVaultPath()
		if err != nil {
			fmt.Println("Error getting vault path:", err)
			return
		}

		searchTerm := args[0]

		match, err := findClosestNote(vaultPath, searchTerm)
		if err != nil {
			fmt.Println("Error finding note:", err)
			return
		}

		fmt.Println("Opening:", match)

		err = openInNotepad(match)
		if err != nil {
			fmt.Println("Error opening note:", err)
			return
		}
	},
}

func findClosestNote(vaultPath string, searchTerm string) (string, error) {
	var notes []string

	err := walkVault(vaultPath, func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		notes = append(notes, path)
		return nil
	})

	if err != nil {
		return "", err
	}

	matches := fuzzy.Find(searchTerm, notes)
	if len(matches) == 0 {
		return "", fmt.Errorf("no matching note found for %q", searchTerm)
	}

	return notes[matches[0].Index], nil
}

func openInNotepad(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	return exec.Command("notepad.exe", path).Start()
}

func init() {
	rootCmd.AddCommand(openCmd)
}