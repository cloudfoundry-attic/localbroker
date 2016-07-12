package main

import (
	"flag"

	"code.cloudfoundry.org/cflager"
	"code.cloudfoundry.org/debugserver"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/localbroker/localbroker"
	"github.com/cloudfoundry-incubator/localbroker/utils"
	"github.com/pivotal-cf/brokerapi"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
)

var atAddress = flag.String(
	"listenAddr",
	"0.0.0.0:8999",
	"host:port to serve service broker API",
)
var serviceName = flag.String(
	"serviceName",
	"localvolume",
	"name of the service to register with cloud controller",
)
var serviceId = flag.String(
	"serviceId",
	"service-guid",
	"ID of the service to register with cloud controller",
)
var planName = flag.String(
	"planName",
	"free",
	"name of the service plan to register with cloud controller",
)
var planId = flag.String(
	"planId",
	"free-plan-guid",
	"ID of the service plan to register with cloud controller",
)
var planDesc = flag.String(
	"planDesc",
	"free local filesystem",
	"description of the service plan to register with cloud controller",
)
var username = flag.String(
	"username",
	"admin",
	"basic auth username to verify on incoming requests",
)
var password = flag.String(
	"password",
	"admin",
	"basic auth password to verify on incoming requests",
)

func main() {
	parseCommandLine()

	logger, logSink := cflager.New("localbroker")
	logger.Info("starting")
	defer logger.Info("ends")

	server := createServer(logger)

	if dbgAddr := debugserver.DebugAddress(flag.CommandLine); dbgAddr != "" {
		server = utils.ProcessRunnerFor(grouper.Members{
			{"debug-server", debugserver.Runner(dbgAddr, logSink)},
			{"broker-api", server},
		})
	}

	process := ifrit.Invoke(server)
	logger.Info("started")
	utils.UntilTerminated(logger, process)
}

func parseCommandLine() {
	cflager.AddFlags(flag.CommandLine)
	debugserver.AddFlags(flag.CommandLine)
	flag.Parse()
}

func createServer(logger lager.Logger) ifrit.Runner {
	serviceBroker := localbroker.New(logger, *serviceName, *serviceId, *planName, *planId, *planDesc)

	credentials := brokerapi.BrokerCredentials{Username: *username, Password: *password}
	handler := brokerapi.New(serviceBroker, logger, credentials)

	return http_server.New(*atAddress, handler)
}
