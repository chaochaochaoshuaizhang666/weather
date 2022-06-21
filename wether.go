package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Weather struct {
	CityInfo struct {
		City       string `json:"city"`
		CityID     string `json:"cityId"`
		Parent     string `json:"parent"`
		UpdateTime string `json:"updateTime"`
	} `json:"cityInfo"`
	Data struct {
		Forecast []struct {
			Aqi     int64  `json:"aqi"`
			Date    string `json:"date"`
			Fl      string `json:"fl"`
			Fx      string `json:"fx"`
			High    string `json:"high"`
			Low     string `json:"low"`
			Notice  string `json:"notice"`
			Sunrise string `json:"sunrise"`
			Sunset  string `json:"sunset"`
			Type    string `json:"type"`
			Week    string `json:"week"`
			Ymd     string `json:"ymd"`
		} `json:"forecast"`
		Ganmao    string  `json:"ganmao"`
		Pm10      float32 `json:"pm10"`
		Pm25      float32 `json:"pm25"`
		Quality   string  `json:"quality"`
		Shidu     string  `json:"shidu"`
		Wendu     string  `json:"wendu"`
		Yesterday struct {
			Aqi     int64  `json:"aqi"`
			Date    string `json:"date"`
			Fl      string `json:"fl"`
			Fx      string `json:"fx"`
			High    string `json:"high"`
			Low     string `json:"low"`
			Notice  string `json:"notice"`
			Sunrise string `json:"sunrise"`
			Sunset  string `json:"sunset"`
			Type    string `json:"type"`
			Week    string `json:"week"`
			Ymd     string `json:"ymd"`
		} `json:"yesterday"`
	} `json:"data"`
	Date    string `json:"date"`
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Time    string `json:"time"`
}

func main() {
	res, err := http.Get("http://t.weather.sojson.com/api/weather/city/101030100")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var wt = new(Weather)
	err = json.Unmarshal(result, wt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(`
    日期：%s %s
    天气：%s
    温度：%s - %s
    风力风向：%s %s
`,
		wt.Data.Forecast[0].Ymd,
		wt.Data.Forecast[0].Week,
		wt.Data.Forecast[0].Type,
		wt.Data.Forecast[0].Low,
		wt.Data.Forecast[0].High,
		wt.Data.Forecast[0].Fl,
		wt.Data.Forecast[0].Fx,
	)
}
