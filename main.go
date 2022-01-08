package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataFormat struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             *int64 `json:"No"`
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

func FetchData() (covidData map[string][]DataFormat) {
	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	bodyJSON := make(map[string][]DataFormat)
	jsonErr := json.Unmarshal(body, &bodyJSON)

	if jsonErr != nil {
		panic(jsonErr)
	}

	return bodyJSON
}

func CountAgeGroup(covidData map[string][]DataFormat) (ageGroup map[string]int) {
	ageMap := make(map[string]int)
	ageMap["0-30"] = 0
	ageMap["31-60"] = 0
	ageMap["61+"] = 0
	ageMap["N/A"] = 0

	data := covidData["Data"]

	for _, value := range data {

		if value.Age == nil {
			ageMap["N/A"] += 1
		} else {
			if *value.Age >= 0 && *value.Age <= 30 {
				ageMap["0-30"] += 1
			} else if *value.Age >= 31 && *value.Age <= 60 {
				ageMap["31-60"] += 1
			} else if *value.Age >= 61 {
				ageMap["61+"] += 1
			}
		}
	}

	return ageMap
}

func CountProvince(covidData map[string][]DataFormat) (province map[string]int) {
	provinceMap := make(map[string]int)
	data := covidData["Data"]

	for _, value := range data {
		if len(value.Province) != 0 {
			if _, ok := provinceMap[value.Province]; !ok {
				provinceMap[value.Province] = 1
			} else {
				provinceMap[value.Province] += 1
			}
		} else {
			if _, ok := provinceMap["N/A"]; !ok {
				provinceMap["N/A"] = 1
			} else {
				provinceMap["N/A"] += 1
			}
		}
	}

	return provinceMap
}

func main() {
	r := gin.Default()

	r.GET("/covid/summary", func(c *gin.Context) {

		covidData := FetchData()

		ageGroup := CountAgeGroup(covidData)

		province := CountProvince(covidData)

		c.JSON(http.StatusOK, gin.H{
			"Province": province,
			"AgeGroup": ageGroup,
		})
	})

	r.Run(":9090")
}
