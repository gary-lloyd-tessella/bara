package kubectl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

const KUBECTL string = "kubectl"

type environment struct {
	kubectlPath string
}

func ApplyManifests(outputDirectory string, templateDir string) {
	kubectlPath, _ := exec.LookPath(KUBECTL)
	log.Info(fmt.Sprintf("Using kubectl from path: %s", kubectlPath))

	templateDirToWalk := outputDirectory + "/" + templateDir
	log.Info(fmt.Sprintf("Executing templates in directory: %s", templateDirToWalk))

	env := environment{kubectlPath}
	filepath.Walk(templateDirToWalk, env.ApplyManifest)
}

func (env *environment) ApplyManifest(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Error(fmt.Sprintf("Error accessing filePath %q: %v\n", filePath, err))
		return err
	}

	if !info.IsDir() {
		log.Info(fmt.Sprintf("Applying Template: %q\n", filePath))

		cmd := exec.Command(env.kubectlPath, "apply", "-f", filePath)
		out, err := cmd.Output()

		if err != nil {
			// Log the error and continue as we want to apply all valid manifests
			fmt.Println(string(err.(*exec.ExitError).Stderr))
		}
		log.Info(fmt.Sprintf("Response: %s", string(out)))
	}

	return nil
}