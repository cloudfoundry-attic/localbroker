package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
