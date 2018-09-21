package template

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"text/template"
)

type environment struct {
	config interface{}
}

func ProcessTemplates(templateDir string, configFile string) {
	env := environment{parseConfig(configFile)}
	err := filepath.Walk(templateDir, env.ProcessTemplate)
	if err != nil {
		fmt.Printf("Error processing template in directory: %v\n", err)
		return
	}
}

func (env *environment) ProcessTemplate(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}

	if !info.IsDir() {
		fmt.Printf("Template to process: %q\n", path)

		t, err := template.ParseFiles(path)
		if err != nil {
			log.Print(err)
			return err
		}

		//config := parseConfig(*configFile)
		err = t.Execute(os.Stdout, env.config)
		if err != nil {
			return err
		}
	}

	return nil
}