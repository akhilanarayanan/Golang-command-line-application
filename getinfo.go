package main

import (
	//"context"
	"flag"
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"time"
	"bufio"
  "os"
	"log"
	"encoding/csv"
	"strconv"
)

type queryRes struct {
	url string
	HTTPStatus int
	startTime int64
	endTime int64
	resBytesSize int
	requestId int
}

func getInfo(url string, id int) (res queryRes, err error) {
	res.requestId = id
	res.startTime = time.Now().UnixNano()
	response, err := http.Get(url)
	// Not sure when the end request time should be measured. Is it just after
	// the get request or after reading the body?
	res.endTime = time.Now().UnixNano()
	
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	
	res.url = url
	res.HTTPStatus = response.StatusCode
	res.resBytesSize = len(body)
	return
}

func main() {
	url := flag.String("url", "", "")
	inputFile := flag.String("file", "", "")
	flag.Parse()
	
	writeFile, createErr := os.Create("large_output.csv")
	if createErr != nil {
		log.Println("failed creating file:", createErr)
		return
	}
	defer writeFile.Close()
	fileWriter := csv.NewWriter(writeFile)
	defer fileWriter.Flush()
	
	if (*url != "") {
		res, err := getInfo(*url, 0)
		if err != nil {
			fmt.Println("url getinfo error:", err)
			return
		}
		fmt.Println(res)
	}
	if (*inputFile != "") {
		readFile, readErr := os.Open(*inputFile)
		if readErr != nil {
			log.Fatalln(readErr)
			return
		}
		
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		id := 0
		for fileScanner.Scan() {
			res, err := getInfo("http://" + fileScanner.Text(), id)
			if err != nil {
				log.Println("file getinfo error:", err)
			} else {
				record := []string{
					res.url, 
					strconv.Itoa(res.HTTPStatus), 
					strconv.FormatInt(res.startTime, 10),
					strconv.FormatInt(res.endTime, 10),
					strconv.Itoa(res.resBytesSize),
					strconv.Itoa(res.requestId)}
				writeErr := fileWriter.Write(record)
				if (writeErr != nil) {
					log.Println("file writing error:", writeErr)
					return
				}
			}
			id++
		}
		readFile.Close()
	}
}