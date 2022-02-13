package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(name string) UserInfo {
	log.Println("getUser : ", name)
	var UserInfo UserInfo

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+name, nil)
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
	err = decoder.Decode(&UserInfo)
	if err != nil {
		return UserInfo
	}
	return UserInfo
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("getUserHandler : ", name)

	user := getUser(name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
