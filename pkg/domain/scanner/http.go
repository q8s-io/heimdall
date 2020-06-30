package scanner

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/q8s-io/heimdall/pkg/models"
)

func AnchoreGET(reqURL string) map[string]interface{} {
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.SetBasicAuth(models.Config.Anchore.UserName, models.Config.Anchore.PassWord)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		log.Println(berr)
	}
	responeDate := make(map[string]interface{}, 1)
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}

func AnchorePOST(reqURL, reqData string) []map[string]interface{} {
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqData))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(models.Config.Anchore.UserName, models.Config.Anchore.PassWord)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, perr := c.Do(req)
	if perr != nil {
		log.Println(perr)
		return nil
	}
	resBody, berr := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if berr != nil {
		log.Println(berr)
		return nil
	}
	responeDate := make([]map[string]interface{}, 1)
	_ = json.Unmarshal(resBody, &responeDate)
	return responeDate
}
