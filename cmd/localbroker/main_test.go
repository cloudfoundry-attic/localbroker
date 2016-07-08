package main_test

import (
	"net/http"
	"os/exec"
	"strconv"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"
	"io/ioutil"
	"encoding/json"
	"github.com/cloudfoundry-incubator/localbroker/model"
)

var _ = Describe("Localbroker Main", func() {
	var (
		args       []string
		listenAddr string
		process    ifrit.Process
		testLogger lager.Logger
	)

	BeforeEach(func() {
		listenAddr = "0.0.0.0:" + strconv.Itoa(8999+GinkgoParallelNode())
		args = append(args, "-listenAddr", listenAddr)

		testLogger = lagertest.NewTestLogger("test")
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

	It("should listen on the given address", func() {
		resp, err := http.Get("http://" + listenAddr + "/v2/catalog")
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
			resp, err := http.Get("http://" + listenAddr + "/v2/catalog")
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(200))
			bytes, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			var catalog model.Catalog
			err = json.Unmarshal(bytes, &catalog)
			Expect(err).NotTo(HaveOccurred())

			Expect(catalog.Services[0].Name).To(Equal("something"))
			Expect(catalog.Services[0].Id).To(Equal("someguid"))
			Expect(catalog.Services[0].Plans[0].Id).To(Equal("some other guid"))
			Expect(catalog.Services[0].Plans[0].Name).To(Equal("some name"))
			Expect(catalog.Services[0].Plans[0].Description).To(Equal("a description"))
		})
	})
})
