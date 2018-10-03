package kubectl

import (
	"fmt"
	"github.com/gary-lloyd-tessella/bara/pkg/manifests"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

// ApplyManifests iterates over manifests applying them to the cluster
func ApplyManifests(outputDir string, manifests []manifests.Manifest) error {
	kubectlPath, _ := exec.LookPath("kubectl")
	log.Info(fmt.Sprintf("Using kubectl from path: %s", kubectlPath))

	for _, manifest := range manifests {
		applyManifest(kubectlPath, outputDir, manifest)
	}

	return nil
}

func applyManifest(kubectlPath string, outputDirectory string, manifest manifests.Manifest) error {
	log.Info(fmt.Sprintf("Applying Template: %q\n", manifest.Path))

	cmd := exec.Command(kubectlPath, "apply", "-f", outputDirectory+"/"+manifest.Path)
	out, err := cmd.Output()

	if err != nil {
		// Log the error and continue as we want to apply all valid manifests
		fmt.Println(string(err.(*exec.ExitError).Stderr))
		return err
	}
	log.Info(fmt.Sprintf("Response: %s", string(out)))
	return nil
}
