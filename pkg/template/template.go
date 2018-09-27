package template

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type environment struct {
	config interface{}
	outputDir string
}

func ProcessTemplates(templateDir string, configFile string, outputDir string) {
	env := environment{parseConfig(configFile), outputDir}
	err := filepath.Walk(templateDir, env.ProcessTemplate)
	if err != nil {
		fmt.Printf("Error processing template in directory: %v\n", err)
		return
	}
}

func (env *environment) ProcessTemplate(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing filePath %q: %v\n", filePath, err)
		return err
	}

	if !info.IsDir() {
		fmt.Printf("Template to process: %q\n", filePath)

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