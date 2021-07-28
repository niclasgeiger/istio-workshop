package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Println("Starting Server and Listen on Port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(resp http.ResponseWriter, _ *http.Request) {
	t := TimeResponse{Time: time.Now().Add(time.Hour * 2).Format("02.01.2006 - 15:04:05")}
	body, err := json.Marshal(t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := resp.Write(body); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
