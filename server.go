package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type List struct {
	Lists []Respons
}

type Respons struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Locations    string   `json:"locations"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
	ConcertDates string   `json:"concertdates"`
	Relations    string   `json:"relations"`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/home.html")
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var GroupList List
	json.Unmarshal(body, &GroupList.Lists)
	t.Execute(w, GroupList)
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/artiste.html")
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var GroupList List
	json.Unmarshal(body, &GroupList.Lists)
	t.Execute(w, GroupList)
}

func main() {
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorRed := "\033[31m"
	// Setting a file server to hold js & css files there
	// assets := http.FileServer(http.Dir("html/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	fmt.Println(string(colorBlue), "[SERVER_INFO] : Starting local Server...")
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(colorRed), "[SERVER_INFO] : Unable to start server ...")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(colorRed), "[SERVER_INFO] : Unable to start server ...")
	}
	var GroupListe List
	json.Unmarshal(body, &GroupListe.Lists)
	styles := http.FileServer(http.Dir("./static/stylesheets"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/artiste", artistPage)
	fmt.Println(string(colorGreen), "[SERVER_READY] : on http://localhost:8080 ✅")
	fmt.Println(string(colorYellow), "[SERVER_INFO] : To stop the program : Ctrl + c")
	http.ListenAndServe(":8080", nil)
}
