package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func walkVault(vaultPath string) error {
	return filepath.WalkDir(vaultPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	})
}