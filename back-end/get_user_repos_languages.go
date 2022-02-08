package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getRepoLang(name, repo string) string {
	resp, err := http.Get("https://api.github.com/repos/" + name + "/" + repo + "/languages")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	return sb
}

func getUserLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("getUserLanguagesHandler : ", name)

	repos := getRepos(name)

	var languages []string
	for _, repo := range repos {
		s := getRepoLang(name, repo.Name)
		languages = append(languages, s)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(languages)
	ret := "[" + strings.Join(languages[:], ", ") + "]"
	w.Write([]byte(ret))
}
