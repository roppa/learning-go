package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type healthReport struct {
	Name  string           `json:"name"`
	Stats runtime.MemStats `json:"stats"`
}

func sendHealthReport() {
	url := "http://localhost:8888/health-check"
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	service, _ := json.Marshal(&healthReport{Name: "serviceA", Stats: stats})

	fmt.Println(url)
	fmt.Println(string(service))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(service))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	for range time.Tick(time.Second) {
		sendHealthReport()
	}
}
