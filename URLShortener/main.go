package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
)

type urlData struct {
	ID       string `json:"id"`
	OrigUrl  string `json:"origUrl"`
	ShortUrl string `json:"shortUrl"`
}

var urls []urlData

func generateShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var url urlData
	_ = json.NewDecoder(r.Body).Decode(&url)

	for _, item := range urls {
		fmt.Println(item.OrigUrl)
		fmt.Println(url.OrigUrl)
		if item.OrigUrl == url.OrigUrl {
			fmt.Println("already generated")
			url.ShortUrl = item.ShortUrl
			json.NewEncoder(w).Encode(url.ShortUrl)
			return
		}
	}
	url.ID = generateHashID()
	url.ShortUrl = "http://localhost:8000/" + url.ID
	fmt.Println(url.ShortUrl)
	urls = append(urls, url)
	json.NewEncoder(w).Encode(url.ShortUrl)

}

func getAllUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}

func generateHashID() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	//fmt.Println(now.Unix())
	e, _ := h.Encode([]int{int(now.Unix())})
	//fmt.Println(e)
	return e

}

func main() {

	router := mux.NewRouter()

	//urls = append(urls, urlData{ID: "1", OrigUrl: "www.facebook.com", ShortUrl: "www.fb.com"})

	router.HandleFunc("/generate", generateShortUrl).Methods("POST")
	router.HandleFunc("/getall", getAllUrls).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
