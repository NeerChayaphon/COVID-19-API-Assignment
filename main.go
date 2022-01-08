package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DataFormat - The format of the covid 19 data
type DataFormat struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             *int64 `json:"No"` // Use int64 as pointer because it can work with nil
	Age            *int64 `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     *int64 `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine *int64 `json:"StatQuarantine"`
}

// FetchData - Function to get the data from the global API
// The data will be store in a Map variable that have a string as a key and slice of DataFormat struct is value
func FetchData() (covidData map[string][]DataFormat) {
	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json") // get data
	if err != nil {
		panic(err) // error
	}

	body, err := ioutil.ReadAll(resp.Body) // get information as slice of byte
	if err != nil {
		panic(err)
	}

	bodyJSON := make(map[string][]DataFormat)  // create map to store data.
	jsonErr := json.Unmarshal(body, &bodyJSON) // convert JSON to map, The map should have "Data" as a key as have Slice of DataFormat struct as value

	if jsonErr != nil { // check for error
		panic(jsonErr)
	}

	return bodyJSON
}

// CountAgeGroup - A function that count each age group
func CountAgeGroup(covidData map[string][]DataFormat) (ageGroup map[string]int) {
	ageMap := make(map[string]int) // map to store each each group
	ageMap["0-30"] = 0
	ageMap["31-60"] = 0
	ageMap["61+"] = 0
	ageMap["N/A"] = 0

	data := covidData["Data"] // get the data

	for _, value := range data { // loop all the data in an array

		if value.Age == nil { // Age is null
			ageMap["N/A"] += 1
		} else {
			if *value.Age >= 0 && *value.Age <= 30 { // count age from 0-30
				ageMap["0-30"] += 1
			} else if *value.Age >= 31 && *value.Age <= 60 { // count age from 31-60
				ageMap["31-60"] += 1
			} else if *value.Age >= 61 { // count age from 61+
				ageMap["61+"] += 1
			}
		}
	}

	return ageMap
}

// CountProvince - A function that find and count province
func CountProvince(covidData map[string][]DataFormat) (province map[string]int) {
	provinceMap := make(map[string]int) // create a map to store Provinces
	data := covidData["Data"]           // get the data

	for _, value := range data { // loop all the data in the array
		if len(value.Province) != 0 { // check for the Province value is not empty
			if _, ok := provinceMap[value.Province]; !ok { // check that the Province is exist in the map or not
				provinceMap[value.Province] = 1 // if not then set to 1
			} else {
				provinceMap[value.Province] += 1 // already exist, increment by 1
			}
		} else {
			if _, ok := provinceMap["N/A"]; !ok { // check that the Province is empty
				provinceMap["N/A"] = 1
			} else {
				provinceMap["N/A"] += 1
			}
		}
	}

	return provinceMap
}

// main - This is the main function of the API
func main() {
	r := gin.Default()

	r.GET("/covid/summary", func(c *gin.Context) { // /covid/summary route

		covidData := FetchData() // get the data from the global API

		ageGroup := CountAgeGroup(covidData) // find and count age group

		province := CountProvince(covidData) // find and count province

		c.JSON(http.StatusOK, gin.H{ // return JSON responce
			"Province": province,
			"AgeGroup": ageGroup,
		})
	})

	r.Run(":9090") // running on port 9090
}
