package k8s_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/caarlos0/env"

	. "github.com/concourse/concourse/topgun"
	. "github.com/concourse/concourse/topgun/k8s"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestK8s(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8s Suite")
}

type environment struct {
	ChartsDir            string `env:"CHARTS_DIR,required"`
	ConcourseChartDir    string `env:"CONCOURSE_CHART_DIR"`
	ConcourseImageDigest string `env:"CONCOURSE_IMAGE_DIGEST"`
	ConcourseImageName   string `env:"CONCOURSE_IMAGE_NAME,required"`
	ConcourseImageTag    string `env:"CONCOURSE_IMAGE_TAG"`
	FlyPath              string `env:"FLY_PATH"`
}

var (
	Environment environment
	fly         Fly
)

var _ = SynchronizedBeforeSuite(func() []byte {
	var parsedEnv environment

	err := env.Parse(&parsedEnv)
	Expect(err).ToNot(HaveOccurred())

	if parsedEnv.FlyPath == "" {
		parsedEnv.FlyPath = BuildBinary()
	}

	if parsedEnv.ConcourseChartDir == "" {
		parsedEnv.ConcourseChartDir = path.Join(
			parsedEnv.ChartsDir, "stable/concourse")
	}

	Run(nil, "kubectl", "config", "current-context")
	Run(nil, "helm", "init", "--client-only")
	Run(nil, "helm", "dependency", "update", parsedEnv.ConcourseChartDir)

	envBytes, err := json.Marshal(parsedEnv)
	Expect(err).ToNot(HaveOccurred())

	return envBytes
}, func(data []byte) {
	err := json.Unmarshal(data, &Environment)
	Expect(err).ToNot(HaveOccurred())
})

var _ = BeforeEach(func() {
	tmp, err := ioutil.TempDir("", "topgun-tmp")
	Expect(err).ToNot(HaveOccurred())

	fly = Fly{
		Bin:    Environment.FlyPath,
		Target: "concourse-topgun-k8s-" + strconv.Itoa(GinkgoParallelNode()),
		Home:   filepath.Join(tmp, "fly-home-"+strconv.Itoa(GinkgoParallelNode())),
	}

	err = os.Mkdir(fly.Home, 0755)
	Expect(err).ToNot(HaveOccurred())
})

func deployConcourseChart(releaseName string, args ...string) {
	helmArgs := []string{
		"--set=concourse.web.kubernetes.keepNamespaces=false",
		"--set=postgresql.persistence.enabled=false",
		"--set=image=" + Environment.ConcourseImageName}

	if Environment.ConcourseImageDigest != "" {
		helmArgs = append(helmArgs, "--set=imageTag="+Environment.ConcourseImageTag)
	}

	if Environment.ConcourseImageDigest != "" {
		helmArgs = append(helmArgs, "--set=imageDigest="+Environment.ConcourseImageDigest)
	}

	helmArgs = append(helmArgs, args...)
	HelmDeploy(releaseName, releaseName, Environment.ConcourseChartDir, helmArgs...)
}

func getRunningWorkers(workers []Worker) (running []Worker) {
	for _, w := range workers {
		if w.State == "running" {
			running = append(running, w)
		}
	}
	return
}
