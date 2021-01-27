package spamchk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	plinoURL = "https://plino.herokuapp.com/api/v1/classify/"
)

type PlinoResp struct {
	TextClass  string `json:"email_class"`
	Text       string `json:"email_text"`
	StatusCode int    `json:"status"`
}

func IsStringSpam(str string) bool {
	if strings.Trim(str, " ") == "" {
		fmt.Println("The string can't be blank!")
		return false
	}

	if len(str) < 10 {
		fmt.Println("The string is too short!")
		return false
	}

	reqBody, err := json.Marshal(map[string]string{"email_text": str})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(plinoURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var plinoResp PlinoResp
	json.Unmarshal(body, &plinoResp)

	if plinoResp.TextClass == "ham" {
		fmt.Println("The string is spam-free")
		return false
	}

	if plinoResp.TextClass == "spam" {
		fmt.Println("The string is a spam!")
		return true
	}

	if plinoResp.TextClass == "UnicodeEncodeError" {
		fmt.Println("There was an Encoding issue")
		return false
	}

	fmt.Println("Something went wrong...")
	return false
}

func IsTextFileSpam(fileName string) bool {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileString := string(content)

	fmt.Println(fileString)

	if strings.Trim(fileString, " ") == "" {
		fmt.Println("The file can't be blank!")
		return false
	}

	if len(fileString) < 10 {
		fmt.Println("The file length is too short!")
		return false
	}

	reqBody, err := json.Marshal(map[string]string{"email_text": fileString})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(plinoURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var plinoResp PlinoResp
	json.Unmarshal(body, &plinoResp)

	if plinoResp.TextClass == "ham" {
		fmt.Println("The string is spam-free")
		return false
	}

	if plinoResp.TextClass == "spam" {
		fmt.Println("The string is a spam!")
		return true
	}

	if plinoResp.TextClass == "UnicodeEncodeError" {
		fmt.Println("There was an Encoding issue")
		return false
	}

	fmt.Println("Something went wrong...")
	return false
}
