package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telebot/config"
)

func searchTrack(trackName string) (*http.Request, error) {
	cfg := config.GetConfig()
	//logger := logger.NewBuiltinLogger()

	req, err := http.NewRequest(http.MethodGet, cfg.YouTube.YoutubeUrl, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("part", "snippet")
	query.Add("maxResults", "1")
	query.Add("q", trackName)
	query.Add("key", cfg.YouTube.ApiToken)

	req.URL.RawQuery = query.Encode()

	return req, err
}

func FindByTrack(trackName string) (string, error) {
	req, err := searchTrack(trackName)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return "", err
	}

	a := responseData["items"].([]interface{})[0].(map[string]interface{})["id"].(map[string]interface{})["videoId"].(string)

	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", a), err
}
