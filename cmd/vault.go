package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	VaultPath string `json:"vaultPath"`
}

func getVaultPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	exeDir := filepath.Dir(exePath)
	configPath := filepath.Join(exeDir, "vault.json")

	file, err := os.Open(configPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var config Config

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return "", err
	}

	if config.VaultPath == "" {
		return "", fmt.Errorf("vaultPath is missing from vault.json")
	}

	return config.VaultPath, nil
}