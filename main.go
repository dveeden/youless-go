package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type YoulessResponse struct {
	Count string `json:"cnt"`
	Power int    `json:"pwr"`
}

func metricHandler(w http.ResponseWriter, req *http.Request) {
	url := "http://192.168.1.20/a?f=j"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var dat YoulessResponse
	err = json.NewDecoder(resp.Body).Decode(&dat)
	if err != nil {
		panic(err)
	}

	count, err := strconv.ParseFloat(strings.Replace(dat.Count, ",", ".", 1), 64)
	if err != nil {
		panic(err)
	}
	metricData := fmt.Sprintf("youless_total %f\n", count)
	metricData += fmt.Sprintf("youless_power %d\n", dat.Power)
	io.WriteString(w, metricData)
}

func main() {
	http.HandleFunc("/metrics", metricHandler)
	log.Fatal(http.ListenAndServe(":8002", nil))
}
