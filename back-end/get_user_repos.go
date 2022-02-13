package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRepos(name string) []Repo {
	log.Println("getRepos : ", name)
	var repos []Repo

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+name+"/repos", nil)
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
	//resp, err := http.Get("https://api.github.com/users/" + name + "/repos")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&repos)
	if err != nil {
		return []Repo{}
	}
	return repos
}

func getUserReposHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("getUserReposHandler : ", name)

	repos := getRepos(name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repos)
}
