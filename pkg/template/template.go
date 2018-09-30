package template

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type environment struct {
	config interface{}
	outputDir string
}

// ProcessTemplates finds all template files in the specified template directory and
// compiles all of the templates into the output directory using the specified configuration
func ProcessTemplates(templateDir string, configFile string, outputDir string) {
	env := environment{parseConfig(configFile), outputDir}
	err := filepath.Walk(templateDir, env.processTemplate)
	if err != nil {
		log.Error(fmt.Sprintf("Error processing template in directory: %v\n", err))
		return
	}
}

func (env *environment) processTemplate(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Error(fmt.Sprintf("Error accessing filePath %q: %v\n", filePath, err))
		return err
	}

	if !info.IsDir() {
		log.Info(fmt.Sprintf("Template to process: %q\n", filePath))

		t, err := template.ParseFiles(filePath)
		if err != nil {
			log.Print(err)
			return err
		}

		outputPath := env.outputDir + "/" + filePath
		os.MkdirAll(path.Dir(outputPath), 0777)
		file, err := os.Create(outputPath)

		if err != nil {
			log.Print(err)
			return err
		}

		err = t.Execute(file, env.config)
		if err != nil {
			return err
		}
	}

	return nil
}