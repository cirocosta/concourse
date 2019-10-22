package k8s_test

import (
	"path"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("External PostgreSQL", func() {
	var (
		pgReleaseName string
		atc           Endpoint
	)

	BeforeEach(func() {
		setReleaseNameAndNamespace("ep")
		pgReleaseName = releaseName + "-pg"

		helmDeploy(pgReleaseName,
			namespace,
			path.Join(Environment.ChartsDir, "stable/postgresql"),
			"--set=livenessProbe.initialDelaySeconds=3",
			"--set=livenessProbe.periodSeconds=3",
			"--set=persistence.enabled=false",
			"--set=postgresqlDatabase=pg-database",
			"--set=postgresqlPassword=pg-password",
			"--set=postgresqlUsername=pg-user",
			"--set=readinessProbe.initialDelaySeconds=3",
			"--set=readinessProbe.periodSeconds=3",
		)

		deployConcourseChart(releaseName,
			"--set=concourse.web.postgres.database=pg-database",
			"--set=concourse.web.postgres.host="+pgReleaseName+"-postgresql",
			"--set=concourse.worker.ephemeral=true",
			"--set=postgresql.enabled=false",
			"--set=secrets.postgresPassword=pg-password",
			"--set=secrets.postgresUser=pg-user",
			"--set=worker.replicas=0",
		)

		waitAllPodsInNamespaceToBeReady(namespace)

		atc = endpointFactory.NewServiceEndpoint(namespace, releaseName+"-web", "8080")
	})

	AfterEach(func() {
		helmDestroy(pgReleaseName)

		atc.Close()
		cleanup(releaseName, namespace, nil) // TODO
	})

	FIt("can have pipelines set", func() {
		By("Logging in")
		fly.Login("test", "test", "http://"+atc.Address())

		By("Setting and triggering a dumb pipeline")
		fly.Run("set-pipeline", "-n", "-c", "pipelines/get-task.yml", "-p", "pipeline")
	})
})
