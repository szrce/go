package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func sendRequest(hash string, apiKey string) {
	//send request
	url := "https://www.virustotal.com/api/v3/files/" + hash

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func main() {

	if len(os.Args) > 2 {
		directory := os.Args[1]
		apiKey := os.Args[2]

		if _, err := os.Stat(directory); os.IsNotExist(err) {
			fmt.Println("directory does not exist")
			return
		}

		files, err := os.ReadDir(directory)
		if err != nil {
			fmt.Println("error reading directory")
			return
		}
		for _, file := range files {
			if file.IsDir() {
				fmt.Println("Directoryy: ", file.Name())

			} else {

				f, err := os.Open(directory + "/" + file.Name())
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
				sendRequest(fmt.Sprintf("%x", h.Sum(nil)), apiKey)
			}
		}

	} else {
		fmt.Println("directory not provided")
	}
}
