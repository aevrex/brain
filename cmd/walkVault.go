package cmd

import (
	"io/fs"
	"path/filepath"
)

func walkVault(vaultPath string, visit func(path string, d fs.DirEntry) error) error {
	return filepath.WalkDir(vaultPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		return visit(path, d)
	})
}