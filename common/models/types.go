package models

type Response struct {
	Status   int      `json:"status"`
	CityInfo CityInfo `json:"cityInfo"`
	Data     Data     `json:"data"`
	Date     string   `json:"date"`
	Message  string   `json:"message"`
	Count    int      `json:"count"`
}

type Data struct {
	Moisture  string `json:"shidu"`
	Quality   string `json:"quality"`
	Cold      string `json:"ganmao"`
	Yesterday Day    `json:"yesterday"`
	Forecast  []Day  `json:"forecast"`
}

type CityInfo struct {
	CityName string `json:"city"`
	CityCode string `json:"cityKey"`
	Parent   string `json:"parent"`
}

type City struct {
	Name string `json:"city_name"`
	Code string `json:"city_code"`
}

type Day struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float32 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}
