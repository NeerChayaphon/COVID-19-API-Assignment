// Create by Chayaphon Bunyakan, Email: bun.chayaphon@gmail.com
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFetchingData - Use to test the getting data feature from globle API
func TestFetchingData(t *testing.T) {
	covidData := FetchData()
	if _, ok := covidData["Data"]; !ok {
		t.Error("Expected to have information of the Covid 19")
	}
}

// TestCountAgeGroup - Test the result of the counting of each age group from the sample data
func TestCountAgeGroup(t *testing.T) {
	covidData := createSampleData()
	ageMap := CountAgeGroup(covidData)
	if ageMap["31-60"] != 1 || ageMap["61+"] != 1 || ageMap["N/A"] != 1 {
		t.Error("Counting age groups are in correct. The result should 1 for '31-60', '61+' and 'N/A'")
	}
}

// TestCountProvince - Test the result of the counting of each Province from the sample data
func TestCountProvince(t *testing.T) {
	covidData := createSampleData()
	province := CountProvince(covidData)
	if province["Phrae"] != 1 || province["Roi Et"] != 2 {
		t.Error("Counting Province is in correct. The result should 1 for 'Phrae' and 2 for 'Roi Et'")
	}
}

// TestHttpRequest - Testing the HTTP request result in term of Status Code
func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"expected service response\"}")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Error("Status Code Not OK")
	}
}

// createSampleData - A function to create sample data that use for testing. The format is the same as FetchData()
func createSampleData() (covidData map[string][]DataFormat) {

	// ******* This is a first example data *******
	firstData := DataFormat{
		ConfirmDate:    "2021-05-04",
		No:             nil,
		Age:            new(int64),
		Gender:         "หญิง",
		GenderEn:       "Female",
		Nation:         "",
		NationEn:       "China",
		Province:       "Phrae",
		ProvinceId:     new(int64),
		District:       "",
		ProvinceEn:     "Phrae",
		StatQuarantine: new(int64),
	}
	*firstData.Age = 51
	*firstData.ProvinceId = 46
	*firstData.StatQuarantine = 5

	// ******* This is a second example data *******
	secondData := DataFormat{
		ConfirmDate:    "2021-05-01",
		No:             nil,
		Age:            new(int64),
		Gender:         "",
		GenderEn:       "",
		Nation:         "",
		NationEn:       "India",
		Province:       "Roi Et",
		ProvinceId:     new(int64),
		District:       "",
		ProvinceEn:     "Roi Et",
		StatQuarantine: new(int64),
	}
	*secondData.Age = 79
	*secondData.ProvinceId = 53
	*secondData.StatQuarantine = 1

	// ******* This is a third example data *******
	thirdData := DataFormat{
		ConfirmDate:    "2021-05-01",
		No:             nil,
		Age:            nil,
		Gender:         "หญิง",
		GenderEn:       "Female",
		Nation:         "",
		NationEn:       "China",
		Province:       "Roi Et",
		ProvinceId:     new(int64),
		District:       "",
		ProvinceEn:     "Roi Et",
		StatQuarantine: nil,
	}
	*thirdData.ProvinceId = 53

	sampleData := make(map[string][]DataFormat)
	sampleData["Data"] = make([]DataFormat, 0)
	sampleData["Data"] = append(sampleData["Data"], firstData)
	sampleData["Data"] = append(sampleData["Data"], secondData)
	sampleData["Data"] = append(sampleData["Data"], thirdData)

	return sampleData
}
