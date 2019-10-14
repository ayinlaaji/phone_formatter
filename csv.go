package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var delim string = " ::: "

//Person struct
type Person struct {
	PhoneNumber string `json:"phoneNumber"`
}

func RemoveChar(text string, char string) string {
	splits := strings.Split(text, char)
	return strings.Join(splits, "")
}

func CheckPlus(text string) bool {
	return string(text[0]) == "+"
}

func CountryZip(country string) (string, error) {
	switch country {
	case "NG":
		return "+234", nil
	}
	return "", errors.New("Country does not exist")
}

func CountryNOLen(country string) (int, error) {
	switch country {
	case "NG":
		return 11, nil
	}
	return 0, errors.New("Country does not exist")
}

func AddZip(text string, country string) (string, error) {
	noLen, _ := CountryNOLen(country)
	if noLen != len(text) {
		return "", errors.New("Invalid length")
	}
	zip, _ := CountryZip(country)
	return (zip + text[1:]), nil
}

func RemoveDelimiter(numbs string, delimiter string) []string {

	strs := strings.Split(numbs, delimiter)
	var arr []string
	for _, item := range strs {
		arr = append(arr, strings.TrimSpace(item))
	}
	return arr
}

func ContactFormat(text string) string {
	if strings.TrimSpace(text) == "" {
		return ""
	}

	chars := []string{" ", "-", "(", ")"}

	for _, char := range chars {
		text = RemoveChar(text, char)
	}

	firstItem := string(text[0])

	if firstItem == "+" {
		return text
	}

	if firstItem != "0" {
		text = "0" + text
	}

	txt, _ := AddZip(text, "NG")
	return txt

}

func main() {

	var fileName string
	flag.StringVar(&fileName, "fileName", "contacts.csv", "Specify csv file name")
	flag.Parse()

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error occured", err)
	}

	r := csv.NewReader(file)

	var allData [][]string

	phoneCell := 39

	title, err := r.Read()
	nTb := title[:phoneCell+1]
	nTa := title[phoneCell+5:]
	header := append(nTb, nTa...)
	allData = append(allData, header)

	for {
		rec, err := r.Read()

		if err == io.EOF {
			break
		}

		phone := rec[phoneCell]

		if rec[phoneCell+2] != "" {
			phone += delim + rec[phoneCell+2]
		}

		if rec[phoneCell+4] != "" {
			phone += delim + rec[phoneCell+4]
		}

		var cPhone string = ""

		if phone != "" {

			for ind, item := range strings.Split(phone, ":::") {

				if ind > 0 {
					cPhone += delim
				}

				cPhone += ContactFormat(item)
			}

		}
		rec[phoneCell] = cPhone
		bfAr := rec[:phoneCell+1]
		afAr := rec[phoneCell+5:]

		//fmt.Println(bfAr)
		allAr := append(bfAr, afAr...)
		//fmt.Println(cPhone)

		allData = append(allData, allAr)

	}

	newFile, err := os.Create("clean_contacts.csv")
	defer file.Close()

	w := csv.NewWriter(newFile)
	defer w.Flush()

	w.WriteAll(allData)
}
