package k8s

import (
	. "github.com/concourse/concourse/topgun"
)

func HelmDeploy(releaseName, namespace, chartDir string, args ...string) {
	helmArgs := []string{
		"upgrade",
		"--install",
		"--force",
		"--wait",
		"--namespace", namespace,
	}

	helmArgs = append(helmArgs, args...)
	helmArgs = append(helmArgs, releaseName, chartDir)

	Wait(Start(nil, "helm", helmArgs...))
}

func HelmDestroy(releaseName string) {
	helmArgs := []string{
		"delete",
		"--purge",
		releaseName,
	}

	Wait(Start(nil, "helm", helmArgs...))
}
