package scanner

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func AnchoreGET(reqURL string) (map[string]interface{}, error) {
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.SetBasicAuth(model.Config.Anchore.UserName, model.Config.Anchore.PassWord)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		xray.ErrMini(perr)
		return nil, perr
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		xray.ErrMini(berr)
		return nil, berr
	}
	responeDate := make(map[string]interface{}, 1)
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate, nil
}

func AnchorePOST(reqURL, reqData string) []map[string]interface{} {
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqData))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(model.Config.Anchore.UserName, model.Config.Anchore.PassWord)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		xray.ErrMini(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		xray.ErrMini(berr)
		return nil
	}
	responeDate := make([]map[string]interface{}, 1)
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}
