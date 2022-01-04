package api

import (
	"api_proyect/parte1/config"
	"api_proyect/parte1/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func DaysReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var date string = ""
	var dayData, finalReport models.Report
	tdc := make(map[string]float32)
	days, _ := strconv.Atoi(r.URL.Query().Get("days"))

	log.Printf("total days: %v", days)
	// iterate through all the days
	for i := 1; i <= days; i++ {
		date = config.START_DATE
		if i > 1 {
			t, err := time.Parse("2006-01-02", config.START_DATE)

			if err != nil {
				log.Print(err)
			}

			t2 := t.AddDate(0, 0, i-1)
			date = t2.Format("2006-01-02")
		}

		url := config.API_URL + date
		log.Print("Getting data from: " + url)
		apiRes, err := http.Get(url)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"Error": "Error consultando API: ` + err.Error() + `"}`))
			return
		}

		data, err := ioutil.ReadAll(apiRes.Body)
		if err != nil {
			log.Print(err)
		}

		// generate report for every single day
		dayData = CalculateSingleDayInfo(data)

		for key, value := range dayData.ComprasPorTDC {
			_, exist := tdc[key]
			if exist {
				tdc[key] += value
			} else {
				tdc[key] = value
			}
		}

		// Consolidate report from all days
		finalReport.Total += dayData.Total
		finalReport.ComprasPorTDC = tdc
		finalReport.NoCompraron += dayData.NoCompraron
		finalReport.CompraMasAlta += dayData.CompraMasAlta
	}

	// set response
	res := models.DefaultResponse{
		Success:   true,
		StartDate: config.START_DATE,
		EndDate:   date,
		Data:      finalReport,
	}
	json, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
