package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/sigmon"
)

func UnmarshallDataFromRequest(r *http.Request, object interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}

	return nil
}

func ExitOnFailure(logger lager.Logger, err error) {
	if err != nil {
		logger.Error("fatal-error-aborting", err)
		os.Exit(1)
	}
}

func UntilTerminated(logger lager.Logger, process ifrit.Process) {
	err := <-process.Wait()
	ExitOnFailure(logger, err)
}

func ProcessRunnerFor(servers grouper.Members) ifrit.Runner {
	return sigmon.New(grouper.NewOrdered(os.Interrupt, servers))
}
