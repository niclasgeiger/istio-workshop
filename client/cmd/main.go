package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	file, err := ioutil.ReadFile("./assets/index.html.tpl")
	if err != nil {
		panic(err)
	}

	tmpl, err = template.New("name").Parse(string(file))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", page)
	log.Println("Starting Server and Listen on Port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func page(resp http.ResponseWriter, _ *http.Request) {
	time, err := getTime()
	if err != nil {
		_, _ = resp.Write([]byte("could not get time!"))
		return
	}

	body := new(bytes.Buffer)
	if err := tmpl.Execute(body, time); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := resp.Write(body.Bytes()); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type TimeResponse struct {
	Time string `json:"time"`
}

func getTime() (string, error) {
	resp, err := http.Get("http://server/time")
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var timeResponse TimeResponse
	if err := json.Unmarshal(body, &timeResponse); err != nil {
		return "", err
	}
	return timeResponse.Time, nil
}
