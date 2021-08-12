package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PolicyList data structure
type PolicyList struct {
}

//Help doc helper for cli
func (*PolicyList) Help() string {
	return "Lists all Custom Policies"
}

//Run  exec for cli
func (*PolicyList) Run(args []string) int {
	url := "https://www.bridgecrew.cloud/api/v1/policies/table/data"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", "e1debacc-fb6d-5230-89f4-ec76f383d092")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var result map[string]interface{}

	//json.NewDecoder([]byte(body)).Decode(&result)
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		panic(err)
	}

	data := result["data"].([]interface{})
	//  filters := result["filters"].([]interface{})
	log.Print(data[0])
	//log.Print(filters)
	return 0
}

//Synopsis of usage
func (h *PolicyList) Synopsis() string {
	return h.Help()
}
