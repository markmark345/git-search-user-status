package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRateLimit(name string) RateLimit {
	var RateLimit RateLimit

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/rate_limit", nil)
	if err != nil {
		log.Println("Error : ", err)
	}
	auth := basicAuth("markmark345", "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")
	log.Println("basic auth : ", auth)
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error : ", err)
	}

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&RateLimit)
	if err != nil {
		return RateLimit
	}
	return RateLimit
}

func getUserRateLimit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("getUserHandler : ", name)

	rate := getRateLimit(name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rate)
}
