package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

var baseUrl string

// YoulessResponseE is the response of the /e endpoint
type YoulessResponseE struct {
	Time           int     `json:"tm"`
	Power          int     `json:"pwr"`
	Netto          float64 `json:"net"`
	TimeS0         int     `json:"ts0"`
	CounterS0      float64 `json:"cs0"`
	PowerS0        int     `json:"ps0"`
	P1             float64 `json:"p1"`
	P2             float64 `json:"p2"`
	N1             float64 `json:"n1"`
	N2             float64 `json:"n2"`
	Gas            float64 `json:"gas"`
	GasTimeStamp   int     `json:"gts"`
	Water          float64 `json:"wtr"`
	WaterTimeStamp int     `json:"wts"`
}

// YoulessResponseF is the response of the /f endpoint
type YoulessResponseF struct {
	Tarif int `json:"tr"`
	// int      `json:"pa"`
	// int      `json:"pp"`
	// int      `json:"pts"`
	Current1 float64 `json:"i1"`
	Current2 float64 `json:"i2"`
	Current3 float64 `json:"i3"`
	Voltage1 float64 `json:"v1"`
	Voltage2 float64 `json:"v2"`
	Voltage3 float64 `json:"v3"`
	Power1   int     `json:"l1"`
	Power2   int     `json:"l2"`
	Power3   int     `json:"l3"`
}

func getE(baseUrl string) (string, error) {
	url := baseUrl + "/e"
	slog.Info("getE", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var yr []YoulessResponseE
	err = json.NewDecoder(resp.Body).Decode(&yr)
	if err != nil {
		return "", err
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
	metricData += fmt.Sprintf("youless_n1 %f\n", dat.N1)
	metricData += fmt.Sprintf("youless_n2 %f\n", dat.N2)
	metricData += fmt.Sprintf("youless_gas %f\n", dat.Gas)
	metricData += fmt.Sprintf("youless_gas_timestamp %d\n", dat.GasTimeStamp)
	metricData += fmt.Sprintf("youless_water %f\n", dat.Water)
	metricData += fmt.Sprintf("youless_water_timestamp %d\n", dat.WaterTimeStamp)
	return metricData, nil
}

func getF(baseUrl string) (string, error) {
	url := baseUrl + "/f"
	slog.Info("getF", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var dat YoulessResponseF
	err = json.NewDecoder(resp.Body).Decode(&dat)
	if err != nil {
		return "", err
	}

	metricData := fmt.Sprintf("youless_tarif %d\n", dat.Tarif)
	metricData += fmt.Sprintf("youless_current1 %f\n", dat.Current1)
	metricData += fmt.Sprintf("youless_current2 %f\n", dat.Current2)
	metricData += fmt.Sprintf("youless_current3 %f\n", dat.Current3)
	metricData += fmt.Sprintf("youless_voltage1 %f\n", dat.Voltage1)
	metricData += fmt.Sprintf("youless_voltage2 %f\n", dat.Voltage2)
	metricData += fmt.Sprintf("youless_voltage3 %f\n", dat.Voltage3)
	metricData += fmt.Sprintf("youless_power1 %d\n", dat.Power1)
	metricData += fmt.Sprintf("youless_power2 %d\n", dat.Power2)
	metricData += fmt.Sprintf("youless_power3 %d\n", dat.Power3)
	return metricData, nil
}

func metricHandler(w http.ResponseWriter, req *http.Request) {
	metricData, err := getE(baseUrl)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, metricData)

	metricData, err = getF(baseUrl)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, metricData)
}

func main() {
	flag.StringVar(&baseUrl, "url", "http://192.168.1.20", "URL base for YouLess")
	listen := flag.String("listen", ":8002", "listen address")
	flag.Parse()
	http.HandleFunc("/metrics", metricHandler)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
