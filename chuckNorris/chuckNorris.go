package chuckNorris

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telebot/config"
)

func reqChuckNorrisApi() (*http.Request, error) {
	cfg := config.GetConfig()

	request, err := http.NewRequest(http.MethodGet, cfg.ChuckNorris.ChuckUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("bad req %v", err)
	}

	return request, nil
}

func GetJoke() (string, error) {
	req, err := reqChuckNorrisApi()
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println(resp)

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseData map[string]interface{}
	err = json.Unmarshal(all, &responseData)
	if err != nil {
		return "", err
	}

	result := responseData["value"].(map[string]interface{})["joke"].(string)

	return result, nil
}
