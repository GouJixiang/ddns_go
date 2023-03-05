package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main2() {

	httpPostJson()

}

func httpPostJson() {
	jsonStr := []byte(`{ "username": "auto", "password": "auto123123" }`)
	url := "https://dns.aliyuncs.com"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(statuscode)
	fmt.Println(hea)

}
