package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*

Example from /e

[
  {
    "tm": 1737031587,
    "net": 7692.057,
    "pwr": 1710,
    "ts0": 1737030600,
    "cs0": 0.000,
    "ps0": 0,
    "p1": 6452.689,
    "p2": 6995.230,
    "n1": 1855.186,
    "n2": 3900.676,
    "gas": 4386.952,
    "gts": 2501161345,
    "wtr": 0.000,
    "wts": 0
  }
]
*/

type YoulessResponse struct {
	Time         int     `json:"tm"`
	Power        int     `json:"pwr"`
	Netto        float64 `json:"net"`
	TimeS0       int     `json:"ts0"`
	CounterS0    float64 `json:"cs0"`
	PowerS0      int     `json:"ps0"`
	P1           float64 `json:"p1"`
	P2           float64 `json:"p2"`
	Gas          float64 `json:"gas"`
	GasTimeStamp int     `json:"gts"`
	Wtr          float64 `json:"wtr"`
	Wts          int     `json:"wts"`
}

func metricHandler(w http.ResponseWriter, req *http.Request) {
	url := "http://192.168.1.20/e"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var yr []YoulessResponse
	err = json.NewDecoder(resp.Body).Decode(&yr)
	if err != nil {
		panic(err)
	}
	dat := yr[0]

	metricData := fmt.Sprintf("youless_time %d\n", dat.Time)
	metricData += fmt.Sprintf("youless_power %d\n", dat.Power)
	metricData += fmt.Sprintf("youless_netto %f\n", dat.Netto)
	metricData += fmt.Sprintf("youless_times0 %d\n", dat.TimeS0)
	metricData += fmt.Sprintf("youless_total %f\n", dat.CounterS0)
	metricData += fmt.Sprintf("youless_powers0 %d\n", dat.PowerS0)
	metricData += fmt.Sprintf("youless_p1 %f\n", dat.P1)
	metricData += fmt.Sprintf("youless_p2 %f\n", dat.P2)
	metricData += fmt.Sprintf("youless_gas %f\n", dat.Gas)
	metricData += fmt.Sprintf("youless_gas_timestamp %d\n", dat.GasTimeStamp)
	metricData += fmt.Sprintf("youless_wtr %f\n", dat.Wtr)
	metricData += fmt.Sprintf("youless_wts %d\n", dat.Wts)

	io.WriteString(w, metricData)
}

func main() {
	http.HandleFunc("/metrics", metricHandler)
	log.Fatal(http.ListenAndServe(":8002", nil))
}
