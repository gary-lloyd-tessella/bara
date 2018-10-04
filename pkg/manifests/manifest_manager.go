package manifests

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

//go:generate charlatan ManifestManager

type ManifestManager interface {
	FindFiles(directory string) ([]string, error)
	CreateDirectory(directory string) error
}

type manifestManager struct{}

func NewManifestManager() ManifestManager {
	return &manifestManager{}
}

func (*manifestManager) FindFiles(directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error(fmt.Sprintf("Error accessing filePath %q: %v\n", filePath, err))
			return err
		}

		if !info.IsDir() {
			log.Info(fmt.Sprintf("Template to process: %q\n", filePath))
			files = append(files, filePath)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walk error [%v]\n", err)
	}

	return files, nil
}

func (*manifestManager) CreateDirectory(directory string) error {
	src, err := os.Stat(directory)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directory, 0755)
		if errDir != nil {
			panic(fmt.Sprintf("Unable to create output direcory: %s", directory))
		}
		return nil
	}

	if src.Mode().IsRegular() {
		log.Info(fmt.Sprintf("%s already exist as a file!", directory))
	}

	log.Info("Clearing existing output directory")
	os.Remove(directory)

	return nil
}