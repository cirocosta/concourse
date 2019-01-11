package k8s

import (
	"bufio"
	"bytes"
	"encoding/json"
	"regexp"
	"time"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	. "github.com/concourse/concourse/topgun"
	. "github.com/onsi/gomega"
)

type K8s struct {
	Namespace string
}

func (k *K8s) GetWebInstances() []Instance {
	return nil
}

func GetPods(namespace string, flags ...string) []Pod {
	var (
		pods struct {
			Items []Pod `json:"items`
		}
		args = append([]string{"get", "pods",
			"--namespace=" + namespace,
			"--output=json",
			"--no-headers"}, flags...)
		session = Start(nil, "kubectl", args...)
	)

	Wait(session)

	err := json.Unmarshal(session.Out.Contents(), &pods)
	Expect(err).ToNot(HaveOccurred())

	return pods.Items
}

func WaitAllPodsInNamespaceToBeReady(namespace string) {
	Eventually(func() bool {
		expectedPods := GetPods(namespace)
		actualPods := GetPods(namespace, "--field-selector=status.phase=Running")

		if len(expectedPods) != len(actualPods) {
			return false
		}

		podsReady := 0
		for _, pod := range actualPods {
			if IsPodReady(pod) {
				podsReady++
			}
		}

		return podsReady == len(expectedPods)
	}, 5*time.Minute, 10*time.Second).Should(BeTrue(), "expected all pods to be ready")
}

func DeletePods(namespace string, flags ...string) []string {
	var (
		podNames []string
		args     = append([]string{"delete", "pod",
			"--namespace=" + namespace,
		}, flags...)
		session = Start(nil, "kubectl", args...)
	)

	Wait(session)

	scanner := bufio.NewScanner(bytes.NewBuffer(session.Out.Contents()))
	for scanner.Scan() {
		podNames = append(podNames, scanner.Text())
	}

	return podNames
}

func StartPortForwarding(namespace, service, port string) (*gexec.Session, string) {
	session := Start(nil, "kubectl", "port-forward",
		"--namespace="+namespace,
		"service/"+service,
		":"+port)
	Eventually(session.Out).Should(gbytes.Say("Forwarding"))

	address := regexp.MustCompile(`127\.0\.0\.1:[0-9]+`).
		FindStringSubmatch(string(session.Out.Contents()))

	Expect(address).NotTo(BeEmpty())

	return session, "http://" + address[0]
}

func IsPodReady(p Pod) bool {
	total := len(p.Status.ContainerStatuses)
	actual := 0

	for _, containerStatus := range p.Status.ContainerStatuses {
		if containerStatus.Ready {
			actual++
		}
	}

	return total == actual
}
