package main

import (
	"github.com/fatih/color"
	flag "gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug    = flag.Flag("debug", "Enable debug mode.").Short('d').Bool()
	template = flag.Flag("template", "Kubernetes deployment template file to use").Short('t').Required().String()
	config   = flag.Flag("config", "Environment configuration file to use").Short('c').Required().String()
	dryrun   = flag.Flag("dryrun", "Build templates without applying to cluster").Short('r').Bool()
)

func main() {
	flag.Version("0.0.1")
	flag.CommandLine.HelpFlag.Short('h')
	flag.Parse()
	color.Blue("Using template: %s\n", *template)
	color.Blue("Using config: %s\n", *config)
	if *dryrun {
		color.Yellow("Dry run - Templates will no be applied")
	}
}