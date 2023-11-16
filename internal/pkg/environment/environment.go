package environment

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadEnv() error {
	envFilePath, err := findEnvFile()
	if err != nil {
		return errors.New(".env file not found")
	}

	err = godotenv.Load(envFilePath)
	if err != nil {
		return fmt.Errorf("Error loading .env file: %+v", err)
	}
	return nil
}

func findEnvFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, ".env")); err != nil {
			parentDir := filepath.Dir(currentDir)

			if parentDir == currentDir {
				break
			}
			currentDir = parentDir
			continue
		}
		return filepath.Join(currentDir, ".env"), nil
	}
	return "", fmt.Errorf(".env file not found")
}

func GetAddr() string {
	return os.Getenv("GRPC_ADDR")
}
