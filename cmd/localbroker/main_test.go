package main_test

import (
	"io"
	"net/http"
	"os/exec"
	"strconv"

	"encoding/json"
	"io/ioutil"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"
)

var _ = Describe("Localbroker Main", func() {
	var (
		args               []string
		listenAddr         string
		username, password string

		process    ifrit.Process
		testLogger lager.Logger
	)

	BeforeEach(func() {
		testLogger = lagertest.NewTestLogger("test-broker")

		listenAddr = "0.0.0.0:" + strconv.Itoa(8999+GinkgoParallelNode())
		username = "admin"
		password = "password"

		args = append(args, "-listenAddr", listenAddr)
		args = append(args, "-username", username)
		args = append(args, "-password", password)
	})

	JustBeforeEach(func() {
		volmanRunner := ginkgomon.New(ginkgomon.Config{
			Name:       "localbroker",
			Command:    exec.Command(binaryPath, args...),
			StartCheck: "started",
		})
		process = ginkgomon.Invoke(volmanRunner)
	})

	AfterEach(func() {
		ginkgomon.Kill(process)
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
