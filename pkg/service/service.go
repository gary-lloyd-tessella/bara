package service

import (
	"fmt"
	"github.com/gary-lloyd-tessella/bara/pkg/kubectl"
	"github.com/gary-lloyd-tessella/bara/pkg/manifests"
	"github.com/gary-lloyd-tessella/bara/pkg/template"
)

func Execute(templateDir string, configFile string, outputDirectory string) error {
	manifestManager := manifests.NewManifestManager()
	filesList, err := manifestManager.FindFiles(templateDir)
	if err != nil {
		return fmt.Errorf("unable to locate manifest files: %v", err)
	}

	var manifestList []manifests.Manifest
	for _, file := range filesList {
		manifestList = append(manifestList, manifests.Manifest{Path: file})
	}

	manifestManager.CreateDirectory(outputDirectory)

	template.EvaluateTemplates(configFile, outputDirectory, manifestList)
	kubectl.ApplyManifests(outputDirectory, manifestList)

	return nil
}
