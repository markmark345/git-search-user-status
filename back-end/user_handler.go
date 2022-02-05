package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("getUserHandler : ", name)
	resp, err := http.Get("https://api.github.com/users/" + name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var u UserInfo
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&u)
	if err != nil {
		log.Println("Error : ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewErrorResponse(err.Error()))
		return
	}
	log.Println("Response status:", resp.Status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
