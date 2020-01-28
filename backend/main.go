package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Answer struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"questionId"`
	IsSelected bool   `json:"isSelected"`
	Text       string `json:"text"`
}

type Question struct {
	Id      int64    `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Questions []Question

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: (GET) \"/\" (homePage)")
}

func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions := Questions{
		Question{
			Id:   0,
			Text: "What is the color of sky?",
			Answers: []Answer{
				{
					Id:         0,
					QuestionId: 0,
					IsSelected: false,
					Text:       "Green",
				},
				{
					Id:         1,
					QuestionId: 0,
					IsSelected: false,
					Text:       "Red",
				},
				{
					Id:         2,
					QuestionId: 0,
					IsSelected: false,
					Text:       "Yellow",
				},
				{
					Id:         3,
					QuestionId: 0,
					IsSelected: false,
					Text:       "Blue",
				},
			},
		},
		Question{
			Id:   1,
			Text: "What is the color of grass?",
			Answers: []Answer{
				{
					Id:         4,
					QuestionId: 1,
					IsSelected: false,
					Text:       "Green",
				},
				{
					Id:         5,
					QuestionId: 1,
					IsSelected: false,
					Text:       "Red",
				},
				{
					Id:         6,
					QuestionId: 1,
					IsSelected: false,
					Text:       "Yellow",
				},
				{
					Id:         7,
					QuestionId: 1,
					IsSelected: false,
					Text:       "Blue",
				},
			},
		},
		Question{
			Id:   2,
			Text: "When phone was invented?",
			Answers: []Answer{
				{
					Id:         8,
					QuestionId: 2,
					IsSelected: false,
					Text:       "1854",
				},
				{
					Id:         9,
					QuestionId: 2,
					IsSelected: false,
					Text:       "1853",
				},
				{
					Id:         10,
					QuestionId: 2,
					IsSelected: false,
					Text:       "1855",
				},
				{
					Id:         11,
					QuestionId: 2,
					IsSelected: false,
					Text:       "1878",
				},
			},
		},
		Question{
			Id:   3,
			Text: "When steam engine was invented?",
			Answers: []Answer{
				{
					Id:         12,
					QuestionId: 3,
					IsSelected: false,
					Text:       "1711",
				},
				{
					Id:         13,
					QuestionId: 3,
					IsSelected: false,
					Text:       "1733",
				},
				{
					Id:         14,
					QuestionId: 3,
					IsSelected: false,
					Text:       "1745",
				},
				{
					Id:         15,
					QuestionId: 3,
					IsSelected: false,
					Text:       "1819",
				},
			},
		},
		Question{
			Id:   4,
			Text: "When car was invented?",
			Answers: []Answer{
				{
					Id:         16,
					QuestionId: 4,
					IsSelected: false,
					Text:       "1886",
				},
				{
					Id:         17,
					QuestionId: 4,
					IsSelected: false,
					Text:       "1876",
				},
				{
					Id:         18,
					QuestionId: 4,
					IsSelected: false,
					Text:       "1898",
				},
				{
					Id:         19,
					QuestionId: 4,
					IsSelected: false,
					Text:       "1885",
				},
			},
		},
	}
	fmt.Println("Endpoint hit: (GET) \"/questions/all\"")
	enableCors(&w)
	json.NewEncoder(w).Encode(questions)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/question/all", getAllQuestions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	handleRequests()
}
