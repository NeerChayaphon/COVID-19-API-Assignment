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
	No             int    `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
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

func main() {
	r := gin.Default()

	r.GET("/covid/summary", func(c *gin.Context) {

		covidData := fetchData()

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"Data":    covidData["Data"],
		})
	})

	r.Run(":9090")
}

/*
 */
