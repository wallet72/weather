package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	type metric struct {
		temp        string
		heatIndex   string
		dewpt       string
		windChill   string
		windSpeed   string
		windGust    string
		pressure    string
		precipRate  string
		precipTotal string
		elev        string
	}

	type siteData struct {
		stationID         string
		obsTimeUtc        string
		obsTimeLocal      string
		neighborhood      string
		softwareType      string
		country           string
		solarRadiation    string
		lon               string
		realtimeFrequency string
		epoch             string
		lat               string
		uv                string
		winddir           string
		humidity          string
		qcStatus          string
		metric            metric
	}

	type topLevel struct {
		observations siteData
	}

	response, err := http.Get("https://api.weather.com/v2/pws/observations/current?stationId=ISYDNE2691&format=json&units=m&apiKey=c4ea432face84316aa432face80316c7")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var result topLevel
	json.Unmarshal([]byte(responseData), &result)

	//fmt.Println("now for the JSON...")
	//var jsonData bytes.Buffer
	//jsonErr := json.Indent(&jsonData, responseData, "", "\t")
	//if jsonErr != nil {
	//	log.Fatal(jsonErr)
	//
	//}
	//fmt.Println(string(jsonData.Bytes()))

	fmt.Println("result = ", result)

	fmt.Println("stationID = ", result.observations.stationID)

}
