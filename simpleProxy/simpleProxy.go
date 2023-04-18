package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {

	ls, err := net.Listen("tcp4", ":9020")

	fmt.Println("listen:9020")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println("connect failed", err)
		}

		go handler(conn)
	}

}
func checkdomain(domain string) bool {
	readFile, err := os.Open("sitelist")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == domain {
			return true
		}
	}
	readFile.Close()
	return false
}

func handler(conn net.Conn) {
	fmt.Printf("======>connection coming this  %s \n\n", conn.RemoteAddr().String())
	for {

		buf := make([]byte, 1024)
		_, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("\nconnection broken %s \n", err)
			conn.Close()
			break
		}

		requestStr := string(buf)
		requestParts := strings.Split(requestStr, " ")
		requrl, err := url.Parse(requestParts[1])
		fmt.Printf("%s", requrl)

		if checkdomain(requrl.Host) {
			conn.Write([]byte("<html><body><div style=\"background-color:red;\">blocked</div></body></html>"))
			conn.Close()
		} else {
			resp, err := http.Get("http://example.com/")
			if err != nil {
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			conn.Write(body)
			conn.Close()

		}

	}

}

/*
	//fmt.Printf("%s", buf)

	//conn.Close()

	//conn.Write([]byte("<html><body><div style=\"background-color:red;\">test</div></body></html>"))


		resp, err := http.Get("http://example.com/")
		if err != nil {
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		//fmt.Print(body)
		//conn.Write([]byte("<html><body><div style=\"background-color:red;\">test</div></body></html>"))
		conn.Write(body)
		//conn.Close()
		//_, err = conn.Write(buf)
*/
