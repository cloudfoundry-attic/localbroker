package main

import (
	"flag"

	"code.cloudfoundry.org/cflager"
	"code.cloudfoundry.org/debugserver"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-incubator/localbroker/handlers"
	"github.com/cloudfoundry-incubator/localbroker/service"
	"github.com/cloudfoundry-incubator/localbroker/utils"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
)

var atAddress = flag.String(
	"listenAddr",
	"0.0.0.0:8999",
	"host:port to serve service broker API",
)
var configPath = flag.String(
	"configPath",
	"/tmp/localbroker",
	"config directory to store book-keeping info",
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

func main() {
	parseCommandLine()
	logger, logSink := cflager.New("localbroker")
	defer logger.Info("ends")

	servers, err := createServer(logger)
	utils.ExitOnFailure(logger, err)

	if dbgAddr := debugserver.DebugAddress(flag.CommandLine); dbgAddr != "" {
		servers = append(grouper.Members{
			{"debug-server", debugserver.Runner(dbgAddr, logSink)},
		}, servers...)
	}

	process := ifrit.Invoke(utils.ProcessRunnerFor(servers))
	logger.Info("started")
	utils.UntilTerminated(logger, process)
}

func parseCommandLine() {
	cflager.AddFlags(flag.CommandLine)
	debugserver.AddFlags(flag.CommandLine)
	flag.Parse()
}

func createServer(logger lager.Logger) (grouper.Members, error) {
	controller := service.NewController(*serviceName, *serviceId, *planId, *planName, *planDesc, *configPath, nil, nil)
	handler, err := handlers.NewHandler(logger, controller)
	utils.ExitOnFailure(logger, err)

	return grouper.Members{
		{"http-server", http_server.New(*atAddress, handler)},
	}, nil
}
