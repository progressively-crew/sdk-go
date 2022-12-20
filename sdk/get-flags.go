package progressively

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetFlags(apiUrl string) map[string]interface{} {
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return make(map[string]interface{})
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return make(map[string]interface{})
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return make(map[string]interface{})
	}

	var flags map[string]interface{}
	errMarshal := json.Unmarshal(resBody, &flags)

	if errMarshal != nil {
		return make(map[string]interface{})
	}

	return flags
}
