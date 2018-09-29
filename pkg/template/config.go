package template

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

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
