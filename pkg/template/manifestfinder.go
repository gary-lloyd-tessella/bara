package template

import (
	"fmt"
	"github.com/gary-lloyd-tessella/bara/pkg/manifests"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Due to the way file walker works the function needs to be attached to a struct to allow context to be stored
// while the full hierarchy is walked. We don't want to expose this outside this package but want to be able to
// return the Manifests outside the package so keep this definition private.
type manifestFinder struct {
	Manifests []manifests.Manifest
}

// FindTemplates finds all template files in the specified template directory and
// compiles all of the templates into the output directory using the specified configuration
func FindTemplates(templateDir string) []manifests.Manifest {
	manifestFinder := manifestFinder{}
	err := filepath.Walk(templateDir, manifestFinder.processTemplate)
	if err != nil {
		log.Error(fmt.Sprintf("Error processing template in directory: %v\n", err))
		return manifestFinder.Manifests
	}

	return manifestFinder.Manifests
}

func (manifestFinder *manifestFinder) processTemplate(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Error(fmt.Sprintf("Error accessing filePath %q: %v\n", filePath, err))
		return err
	}

	if !info.IsDir() {
		log.Info(fmt.Sprintf("Template to process: %q\n", filePath))
		manifestFinder.Manifests = append(manifestFinder.Manifests, manifests.Manifest{Path: filePath})
	}

	return nil
}
