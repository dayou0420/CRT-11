package openweathermap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	ICON        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Geocoding struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func timeNow() string {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	t := time.Now().UTC().In(jst).String()
	return t
}

func GetWeatherData() (str string, err error) {
	url := os.Getenv("WEATHERDATA")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Data

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var s string
	slice := GetGeocoding()
	for _, v := range slice {
		s += strconv.FormatFloat(v, 'f', 2, 64) + " "
	}

	for _, item := range data.Weather {
		main := strings.ToUpper(item.Main)
		desc := strings.ToUpper(item.Description)
		str = "[" +
			timeNow() +
			": " +
			s + main +
			" " +
			desc +
			" " +
			strconv.FormatFloat(data.Main.Temp, 'f', 2, 64) +
			" " +
			strconv.FormatFloat(float64(data.Main.Humidity), 'f', 2, 64) +
			"]" +
			"[" +
			"https://crt-11.onrender.com/ts" +
			"]"
	}

	return str, err
}

func GetGeocoding() (sli []float64) {
	url := os.Getenv("GEOCODING")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Geocoding

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, item := range data {
		slices := []float64{item.Lat, item.Lon}
		sli = append(sli, slices...)
	}

	return sli
}
