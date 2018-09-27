package kubectl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

const KUBECTL string = "kubectl"

func ApplyManifests(outputDirectory string, templateDir string) {
	kubectlPath, _ := exec.LookPath(KUBECTL)
	log.Info(fmt.Sprintf("Using kubectl from path: %s", kubectlPath))

	templateDirToExecute := outputDirectory + "/" + templateDir
	log.Info(fmt.Sprintf("Executing templates in directory: %s", templateDirToExecute))

	cmd := exec.Command(kubectlPath, "apply", "-f", templateDirToExecute)
	out, err := cmd.Output()

	if err != nil {
		// Log the error and continue as we want to apply all valid manifests
		fmt.Println(string(err.(*exec.ExitError).Stderr))
	}
	log.Info(string(out))
}