package main

import (
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"os"
)

type Malicious struct {
	Data struct {
		Attributes struct {
			AntiyInfo         string   `json:"antiy_info"`
			Names             []string `json:"names"`
			Magika            string   `json:"magika"`
			UniqueSources     int      `json:"unique_sources"`
			LastAnalysisStats struct {
				Malicious        int `json:"malicious"`
				Suspicious       int `json:"suspicious"`
				Undetected       int `json:"undetected"`
				Harmless         int `json:"harmless"`
				Timeout          int `json:"timeout"`
				ConfirmedTimeout int `json:"confirmed-timeout"`
				Failure          int `json:"failure"`
				TypeUnsupported  int `json:"type-unsupported"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

func sendRequest(hash string, apiKey string) {
	//send request
	var malicus Malicious

	url := "https://www.virustotal.com/api/v3/files/" + hash

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal([]byte(string(body)), &malicus)
	color.Red("malicious: %d, ", malicus.Data.Attributes.LastAnalysisStats.Malicious)

	color.White(string(body))
}

func main() {

	directory := flag.String("d", "x", "a directory")
	apiKey := flag.String("k", "1", "your api key")
	flag.Parse()

	if *directory == "x" && *apiKey == "1" {
		flag.PrintDefaults()
		return
	}

	color.Red("your directory we scan this directory:" + *directory)
	color.Red("your api key:" + *apiKey)

	if _, err := os.Stat(*directory); os.IsNotExist(err) {
		fmt.Println("directory does not exist")
		return
	}

	files, err := os.ReadDir(*directory)
	if err != nil {
		fmt.Println("error reading directory")
		return
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directoryy: ", file.Name())

		} else {

			f, err := os.Open(*directory + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			h := md5.New()
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fmt.Printf("%x", h.Sum(nil))

			fmt.Println("File: ", file.Name())
			sendRequest(fmt.Sprintf("%x", h.Sum(nil)), *apiKey)
		}
	}

	/*
		} else {
			fmt.Println("directory not provided")
		}*/
}
