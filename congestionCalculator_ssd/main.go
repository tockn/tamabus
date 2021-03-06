package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const URL = "http://13.231.211.87/api/bus/image"

type BusImage struct {
	BusID  int64  `json:"bus_id"`
	Base64 string `json:"base64"`
}

func main() {

	for {
		req, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			log.Println(err)
		}

		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		bs := make([]*BusImage, 0, 4)

		if err := json.Unmarshal(body, &bs); err != nil {
			log.Println(err)
		}

		for _, b := range bs {
			_, err := decode(b.Base64, b.BusID)
			if err != nil {
				log.Println(err)
			}
			log.Println("busID:", b.BusID)
		}
		time.Sleep(8 * time.Second)
	}
}

func decode(body string, busID int64) (string, error) {
	data, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return "", err
	}

	fullPath := fmt.Sprintf("./images/raw/%d.png", 1) // ほんとはbusID
	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	return "", err
}
