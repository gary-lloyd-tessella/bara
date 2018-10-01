package main

import (
	"fmt"
	"github.com/gary-lloyd-tessella/bara/pkg/kubectl"
	"github.com/gary-lloyd-tessella/bara/pkg/template"
	log "github.com/sirupsen/logrus"
	flag "gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	debug       = flag.Flag("debug", "Enable debug mode.").Short('d').Bool()
	templateDir = flag.Flag("template", "Kubernetes deployment template file to use").Short('t').Required().String()
	configFile  = flag.Flag("config", "Environment configuration file to use").Short('c').Required().String()
	dryrun      = flag.Flag("dryrun", "Build templates without applying to cluster").Short('r').Bool()
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

const outputDirectory string = ".bara"

func main() {
	flag.Version("0.0.1")
	flag.CommandLine.HelpFlag.Short('h')
	flag.Parse()
	log.Info(fmt.Sprintf("Using template directory: %s", *templateDir))
	log.Info(fmt.Sprintf("Using config: %s\n", *configFile))

	if *dryrun {
		log.Info("Dry run - Templates will no be applied")
	}

	createOutputDirectory(outputDirectory)
	template.ProcessTemplates(*templateDir, *configFile, outputDirectory)
	kubectl.ApplyManifests(outputDirectory, *templateDir)
}

func createOutputDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(fmt.Sprintf("Unable to create output direcory: %s", dirName))
		}
		return true
	}

	log.Info("Clearing existing output directory")
	os.Remove(dirName)

	if src.Mode().IsRegular() {
		log.Info(fmt.Sprintf("%s already exist as a file!", dirName))
		return false
	}

	return false
}
