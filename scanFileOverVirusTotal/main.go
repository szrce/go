package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func sendRequest(hash string) {
	//send request
	url := "https://www.virustotal.com/api/v3/files/275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", "API_KEY")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func main() {

	if len(os.Args) > 1 {
		directory := os.Args[1]
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
				sendRequest(fmt.Sprintf("%x", h.Sum(nil)))
			}
		}

	} else {
		fmt.Println("directory not provided")
	}
}
