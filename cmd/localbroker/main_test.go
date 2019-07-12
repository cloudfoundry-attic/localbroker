package main_test

import (
	"io"
	"net/http"
	"os/exec"
	"strconv"

	"encoding/json"
	"io/ioutil"

	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"
)

var _ = Describe("Localbroker Main", func() {
	var process ifrit.Process
	var volmanRunner *ginkgomon.Runner

	JustBeforeEach(func() {
		process = ifrit.Invoke(volmanRunner)
	})

	AfterEach(func() {
		ginkgomon.Kill(process)
	})

	Context("Missing required args", func() {
		BeforeEach(func() {
			var args []string
			volmanRunner = ginkgomon.New(ginkgomon.Config{
				Name:       "localbroker",
				Command:    exec.Command(binaryPath, args...),
				StartCheck: "ERROR: Required parameter dataDir not defined.",
			})

		})

		It("shows usage", func() {
			var err error
			Eventually(process.Wait(), 1*time.Minute).Should(Receive(&err))
		})
	})

	Context("Has required args", func() {
		var (
			args               []string
			listenAddr         string
			tempDir            string
			username, password string
		)

		BeforeEach(func() {
			listenAddr = "0.0.0.0:" + strconv.Itoa(8999+GinkgoParallelNode())
			username = "admin"
			password = "password"
			tempDir = os.TempDir()

			args = append(args, "-listenAddr", listenAddr)
			args = append(args, "-username", username)
			args = append(args, "-password", password)
			args = append(args, "-dataDir", tempDir)

			volmanRunner = ginkgomon.New(ginkgomon.Config{
				Name:       "localbroker",
				Command:    exec.Command(binaryPath, args...),
				StartCheck: "started",
			})
		})

		httpDoWithAuth := func(method, endpoint string, body io.ReadCloser) (*http.Response, error) {
			req, err := http.NewRequest(method, "http://"+listenAddr+endpoint, body)
			Expect(err).NotTo(HaveOccurred())

			req.SetBasicAuth(username, password)
			return http.DefaultClient.Do(req)
		}

		It("should listen on the given address", func() {
			resp, err := httpDoWithAuth("GET", "/v2/catalog", nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(200))
		})

		Context("given arguments", func() {
			BeforeEach(func() {
				args = append(args, "-serviceName", "something")
				args = append(args, "-serviceId", "someguid")
				args = append(args, "-planName", "some name")
				args = append(args, "-planId", "some other guid")
				args = append(args, "-planDesc", "a description")

				volmanRunner = ginkgomon.New(ginkgomon.Config{
					Name:       "localbroker",
					Command:    exec.Command(binaryPath, args...),
					StartCheck: "started",
				})
			})

			It("should pass arguments though to catalog", func() {
				resp, err := httpDoWithAuth("GET", "/v2/catalog", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))

				bytes, err := ioutil.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())

				var catalog brokerapi.CatalogResponse
				err = json.Unmarshal(bytes, &catalog)
				Expect(err).NotTo(HaveOccurred())

				Expect(catalog.Services[0].Name).To(Equal("something"))
				Expect(catalog.Services[0].ID).To(Equal("someguid"))
				Expect(catalog.Services[0].Plans[0].ID).To(Equal("some other guid"))
				Expect(catalog.Services[0].Plans[0].Name).To(Equal("some name"))
				Expect(catalog.Services[0].Plans[0].Description).To(Equal("a description"))
			})
		})
	})
})
