package main

import (
	"fmt"
	"github.com/fatih/color"
	flag "gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var (
	debug       = flag.Flag("debug", "Enable debug mode.").Short('d').Bool()
	templateDir = flag.Flag("template", "Kubernetes deployment template file to use").Short('t').Required().String()
	configFile  = flag.Flag("config", "Environment configuration file to use").Short('c').Required().String()
	dryrun      = flag.Flag("dryrun", "Build templates without applying to cluster").Short('r').Bool()
)

func main() {
	flag.Version("0.0.1")
	flag.CommandLine.HelpFlag.Short('h')
	flag.Parse()
	color.Blue("Using template: %s\n", *templateDir)
	color.Blue("Using config: %s\n", *configFile)
	if *dryrun {
		color.Yellow("Dry run - Templates will no be applied")
	}

	processTemplates(*templateDir, *configFile)
}

type environment struct {
	config interface{}
}

func parseConfig(configFile string) interface{} {
	var config interface{}

	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}

func processTemplates(templateDir string, configFile string) {
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

