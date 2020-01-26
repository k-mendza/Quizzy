package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Question struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

type Questions []Question

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: \"/\" (homePage)")
}

func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions := Questions{
		Question{
			Id:   0,
			Text: "What is the color of sky?",
		},
		Question{
			Id:   1,
			Text: "What is the color of grass?",
		},
		Question{
			Id:   2,
			Text: "When phone was invented?",
		},
		Question{
			Id:   3,
			Text: "When steam engine was invented?",
		},
		Question{
			Id:   4,
			Text: "When car was invented?",
		},
	}
	fmt.Println("Endpoint questions (GET) hit")
	json.NewEncoder(w).Encode(questions)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/question/all", getAllQuestions)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
