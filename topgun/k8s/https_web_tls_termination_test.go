package k8s_test

import (
	"io/ioutil"
	"net"
	"os"
	"time"

	"github.com/onsi/gomega/gbytes"
	"github.com/square/certstrap/pkix"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Web HTTP or HTTPS(TLS) termination at web node", func() {
	var (
		serverCertBytes []byte
		serverKeyBytes  []byte
		caCertFile      *os.File
	)

	BeforeEach(func() {
		var err error

		CACert, serverKey, serverCert := generateKeyPairWithCA()
		CACertBytes, err := CACert.Export()
		Expect(err).NotTo(HaveOccurred())

		caCertFile, err = ioutil.TempFile("", "ca")
		caCertFile.Write(CACertBytes)
		caCertFile.Close()

		serverKeyBytes, err = serverKey.ExportPrivate()
		Expect(err).NotTo(HaveOccurred())

		serverCertBytes, err = serverCert.Export()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.Remove(caCertFile.Name())
	})

	Context("when configured correctly", func() {

		var (
			atc Endpoint

			chartConfig []string
			proxyPort   string
		)

		JustBeforeEach(func() {
			setReleaseNameAndNamespace("wtt")
			deployConcourseChart(releaseName,
				chartConfig...)

			waitAllPodsInNamespaceToBeReady(namespace)

			atc = endpointFactory.NewServiceEndpoint(
				namespace,
				releaseName+"-web",
				proxyPort,
			)
		})

		AfterEach(func() {
			atc.Close()
			cleanup(releaseName, namespace)
		})

		Context("with tls termination at web", func() {

			BeforeEach(func() {
				chartConfig = generateChartConfig(
					"--set=concourse.web.externalUrl=https://test.com",
					"--set=concourse.web.tls.enabled=true",
					"--set=secrets.webTlsCert="+string(serverCertBytes),
					"--set=secrets.webTlsKey="+string(serverKeyBytes),
				)

				proxyPort = "443"
			})

			It("fly login succeeds when using the correct CA and host", func() {
				fly.Run("login", "-u", "test", "-p", "test",
					"--ca-cert", caCertFile.Name(),
					"-c", "https://"+atc.Address(),
				)
			})

			It("fly login fails when NOT using the correct CA", func() {
				sess := fly.Start("login", "-u", "test", "-p", "test",
					"--ca-cert", "k8s/certs/wrong-ca.crt",
					"-c", "https://"+atc.Address(),
				)
				<-sess.Exited

				Expect(sess.ExitCode()).ToNot(Equal(0))
				Expect(sess.Err).To(gbytes.Say(`x509: certificate signed by unknown authority`))
			})
		})

	})

	Context("When NOT configured correctly", func() {

		BeforeEach(func() {
			setReleaseNameAndNamespace("wtt")
		})

		It("fails if tls is enabled but externalURL is NOT set", func() {
			expectedErr := "Must specify HTTPS external URL when concourse.web.tls.enabled is true"

			chartConfig := generateChartConfig(
				"--set=concourse.web.tls.enabled=true",
				"--set=secrets.webTlsCert="+string(serverCertBytes),
				"--set=secrets.webTlsKey="+string(serverKeyBytes),
			)

			deployFailingConcourseChart(releaseName, expectedErr,
				chartConfig...,
			)
		})

		It("fails when tls is enabled but ssl cert and ssl key are NOT set", func() {
			expectedErr := "secrets.webTlsCert is required because secrets.create is true and concourse.web.tls.enabled is true"

			chartConfig := generateChartConfig("--set=concourse.web.externalUrl=https://test.com",
				"--set=concourse.web.tls.enabled=true")

			deployFailingConcourseChart(releaseName, expectedErr,
				chartConfig...,
			)
		})
	})

})

func generateChartConfig(args ...string) []string {
	return append(args,
		"--set=worker.replicas=1",
		"--set=concourse.worker.baggageclaim.driver=detect",
		"--set=concourse.web.tls.bindPort=443",
	)
}
func generateKeyPairWithCA() (*pkix.Certificate, *pkix.Key, *pkix.Certificate) {
	CAKey, err := pkix.CreateRSAKey(1024)
	Expect(err).NotTo(HaveOccurred())

	CACert, err := pkix.CreateCertificateAuthority(CAKey, "", time.Now().Add(time.Hour), "Pivotal", "", "", "", "CA")
	Expect(err).NotTo(HaveOccurred())

	serverKey, err := pkix.CreateRSAKey(1024)
	Expect(err).NotTo(HaveOccurred())

	certificateSigningRequest, err := pkix.CreateCertificateSigningRequest(serverKey, "", []net.IP{net.IPv4(127, 0, 0, 1)},
		nil, "Pivotal", "", "", "", "127.0.0.1")
	Expect(err).NotTo(HaveOccurred())

	serverCert, err := pkix.CreateCertificateHost(CACert, CAKey, certificateSigningRequest, time.Now().Add(time.Hour))
	Expect(err).NotTo(HaveOccurred())

	return CACert, serverKey, serverCert
}
