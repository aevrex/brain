package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
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
	var bestMatch string
	bestScore := 0

	searchTerm = strings.ToLower(searchTerm)

	err := filepath.WalkDir(vaultPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		fileName := strings.ToLower(d.Name())
		score := scoreMatch(fileName, searchTerm)

		if score > bestScore {
			bestScore = score
			bestMatch = path
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if bestMatch == "" {
		return "", fmt.Errorf("no matching note found for %q", searchTerm)
	}

	return bestMatch, nil
}

func scoreMatch(fileName string, searchTerm string) int {
	if fileName == searchTerm {
		return 100
	}

	if strings.TrimSuffix(fileName, filepath.Ext(fileName)) == searchTerm {
		return 90
	}

	if strings.Contains(fileName, searchTerm) {
		return 70
	}

	score := 0
	searchIndex := 0

	for _, char := range fileName {
		if searchIndex < len(searchTerm) && char == rune(searchTerm[searchIndex]) {
			score += 5
			searchIndex++
		}
	}

	return score
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