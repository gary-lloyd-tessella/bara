package template

import (
	"github.com/gary-lloyd-tessella/bara/pkg/manifests"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"text/template"
)

// EvaluateTemplates iterates over the supplied templates evaluating them into the output directory
func EvaluateTemplates(configFile string, outputDir string, manifests []manifests.Manifest) error {
	config := parseConfig(configFile)

	for _, manifest := range manifests {
		evaluateTemplate(config, outputDir, manifest)
	}

	return nil
}

func evaluateTemplate(config interface{}, outputDir string, manifest manifests.Manifest) error {
	t, err := template.ParseFiles(manifest.Path)
	if err != nil {
		log.Error(err)
		return err
	}

	outputPath := outputDir + "/" + manifest.Path
	os.MkdirAll(path.Dir(outputPath), 0777)
	file, err := os.Create(outputPath)

	if err != nil {
		log.Error(err)
		return err
	}

	err = t.Execute(file, config)
	if err != nil {
		return err
	}

	return nil
}
