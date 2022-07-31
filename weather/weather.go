package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telebot/config"
)

func searchCity(city string) (*http.Request, error) {
	cfg := config.GetConfig()

	req, err := http.NewRequest(http.MethodGet, cfg.Weather.WeatherApi, nil)
	if err != nil {
		return nil, fmt.Errorf("req error, %v", err)
	}

	query := req.URL.Query()

	query.Add("lang", "ru")
	query.Add("q", city)
	query.Add("units", "metric")
	query.Add("appid", cfg.Weather.WeatherKey)

	req.URL.RawQuery = query.Encode()

	return req, nil
}

func FindWeatherCity(city string) (string, error) {
	req, err := searchCity(city)
	if err != nil {
		return "", fmt.Errorf("req error: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("resp error: %v", err)
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("readAll error: %v", err)
	}

	defer resp.Body.Close()

	var responseData map[string]interface{}

	err = json.Unmarshal(all, &responseData)
	if err != nil {
		return "", fmt.Errorf("unmarshal error: %v", err)
	}

	a := responseData["main"].(map[string]interface{})["temp"].(float64)
	b := responseData["main"].(map[string]interface{})["humidity"].(float64)
	c := responseData["main"].(map[string]interface{})["pressure"].(float64)
	d := responseData["weather"].([]interface{})[0].(map[string]interface{})["description"].(string)

	return fmt.Sprintf("\nТемпература %.2f %s\nДаление %.2f\nВлажность %.2f", a, d, b, c), nil
}
