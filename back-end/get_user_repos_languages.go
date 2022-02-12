package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	fmt.Println(repos)

	var x map[string]RepoLanguages
	x = make(map[string]RepoLanguages)

	var str string
	var languages []string
	for _, repo := range repos {
		s := getRepoLang(name, repo.Name)
		json.Unmarshal([]byte(s), &str)
		languages = append(languages, s)
		//fmt.Println(x)
		fmt.Println("******Begin*******")
		fmt.Println(str)
		t := strings.Replace(s, "{", "", -1)
		fmt.Println("t1 : ", t)
		t = strings.Replace(t, "}", "", -1)
		fmt.Println("t2 : ", t)
		t = strings.Replace(t, "\"", "", -1)
		fmt.Println("t3 : ", t)
		fmt.Println("******End*******")
		if len(t) > 0 {
			ss := strings.Split(t, ",")
			fmt.Println("ss : ", ss)
			if len(ss) > 0 {
				for _, tok := range ss {
					var repoLanguages RepoLanguages
					ls := strings.Split(tok, ":")
					num, _ := strconv.Atoi(ls[1])
					if _, ok := x[ls[0]]; ok {
						repoLanguages.Language = ls[0]
						repoLanguages.NumberOfLine = x[ls[0]].NumberOfLine + num
						repoLanguages.Repos = x[ls[0]].Repos + 1

					} else {
						repoLanguages.Language = ls[0]
						repoLanguages.NumberOfLine = num
						repoLanguages.Repos = 1
					}
					x[ls[0]] = repoLanguages
				}
				for k, v := range x {
					fmt.Println("Key : ", k, " number of lines : ", v.NumberOfLine, " repo use : ", v.Repos)
				}
			}
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(languages)
	ret := "[" + strings.Join(languages[:], ", ") + "]"
	w.Write([]byte(ret))
}
