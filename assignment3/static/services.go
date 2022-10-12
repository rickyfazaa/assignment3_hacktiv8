package static

import (
	"assignment3/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

func AutoReload() {
	for {
		var (
			water int = RandomNumberWater()
			wind  int = RandomNumberWind()
		)

		number := models.StatusWeather{}
		number.Status.Water = water
		number.Status.Wind = wind

		jsonData, err := json.Marshal(number)

		if err != nil {
			log.Fatal(err.Error())
		}
		err = ioutil.WriteFile("weather.json", jsonData, 0644)

		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

func ReloadWeb(w http.ResponseWriter, r *http.Request) {
	fileData, err := ioutil.ReadFile("weather.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	var statusData models.StatusWeather

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal(err.Error())
	}

	waterValue := statusData.Status.Water
	windValue := statusData.Status.Wind

	var waterStatus string
	var windStatus string

	waterStatus = WaterCondition(waterValue)
	windStatus = WindCondition(windValue)

	data := map[string]interface{}{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
		"waterHeight": waterValue,
		"windSpeed":   windValue,
	}

	tpl, err := template.ParseFiles("./asset/layout.html")

	if err != nil {
		log.Fatal(err.Error())
	}

	tpl.Execute(w, data)

}
