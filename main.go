package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataFormat struct {
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

func fetchData() (covidData map[string][]dataFormat) {
	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodyJSON := make(map[string][]dataFormat)
	jsonErr := json.Unmarshal(body, &bodyJSON)

	if jsonErr != nil {
		panic(jsonErr)
	}

	return bodyJSON
}

func countAgeGroup(covidData map[string][]dataFormat) (ageGroup map[string]int) {
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
		// switch {
		// case *value.Age >= 0 && *value.Age <= 30:
		// 	ageMap["0-30"] += 1
		// case *value.Age >= 31 && *value.Age <= 60:
		// 	ageMap["31-60"] += 1
		// case *value.Age >= 61:
		// 	ageMap["61+"] += 1
		// default:
		// 	ageMap["N/A"] += 1
		// }
	}

	return ageMap
}

func main() {
	r := gin.Default()

	r.GET("/covid/summary", func(c *gin.Context) {

		covidData := fetchData()

		ageGroup := countAgeGroup(covidData)

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"Age":     ageGroup,
			"Data":    covidData["Data"],
		})
	})

	r.Run(":9090")
}

/*
 */
