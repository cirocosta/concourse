package k8s_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/concourse/concourse/topgun/generic"
	"github.com/onsi/gomega/gexec"

	. "github.com/concourse/concourse/topgun"
	. "github.com/concourse/concourse/topgun/k8s"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Worker Rebalancing", func() {
	var (
		releaseName  string
		namespace    string
		proxySession *gexec.Session
		atcEndpoint  string
	)

	BeforeEach(func() {
		releaseName = fmt.Sprintf("topgun-wr-%d-%d", GinkgoRandomSeed(), GinkgoParallelNode())
		namespace = releaseName

		deployConcourseChart(releaseName,
			"--set=concourse.worker.ephemeral=true",
			"--set=worker.replicas=1",
			"--set=web.replicas=2",
			"--set=concourse.worker.rebalanceInterval=5s",
			"--set=concourse.worker.baggageclaim.driver=detect")

		WaitAllPodsInNamespaceToBeReady(namespace)

		By("Creating the web proxy")
		proxySession, atcEndpoint = StartPortForwarding(namespace, releaseName+"-web", "8080")

		By("Logging in")
		fly.Login("test", "test", atcEndpoint)

		By("waiting for a running worker")
		Eventually(func() []Worker {
			return getRunningWorkers(fly.GetWorkers())
		}, 2*time.Minute, 10*time.Second).
			ShouldNot(HaveLen(0))
	})

	AfterEach(func() {
		HelmDestroy(releaseName)
		Wait(Start(nil, "kubectl", "delete", "namespace", namespace, "--wait=false"))
		Wait(proxySession.Interrupt())
	})

	generic.WorkerRebalancing(K8s{
		Namespace:            namespace,
		ConcourseReleaseName: releaseName,
	}, fly)

})
