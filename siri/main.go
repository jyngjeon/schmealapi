package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gorilla/mux"
)

//var for siri (to read Korean only)
var sch [4]string

//GetMeal parses the meal info from school website
func GetMeal() []soup.Root {
	resp, err := soup.Get("http://hana.hs.kr/life/meal.asp")
	fmt.Println("http transport error is:", err)

	doc := soup.HTMLParse(resp)
	meal := doc.Find("table", "class", "today_meal").FindAll("td")

	return meal
}

//WriteMeal makes the struct
func WriteMeal() [4]string {
	meal := GetMeal()
	if len(meal) > 1 {
		brkf := strings.Split(meal[0].Text(), ",")
		lnch := strings.Split(meal[1].Text(), ",")
		dinr := strings.Split(meal[2].Text(), ",")
		snck := strings.Split(meal[3].Text(), ",")

		sch[0] = strings.Join(brkf, ",")
		sch[1] = strings.Join(lnch, ",")
		sch[2] = strings.Join(dinr, ",")
		sch[3] = strings.Join(snck, ",")
	} else {
		sch[0] = meal[0].Text()
	}
	return sch
}

//SendMeal for api work
func SendBrkf(w http.ResponseWriter, r *http.Request) {
	schmeal := WriteMeal()
	json.NewEncoder(w).Encode(schmeal[0])
}

func SendLnch(w http.ResponseWriter, r *http.Request) {
	schmeal := WriteMeal()
	json.NewEncoder(w).Encode(schmeal[1])
}

func SendDinr(w http.ResponseWriter, r *http.Request) {
	schmeal := WriteMeal()
	json.NewEncoder(w).Encode(schmeal[2])
}

func SendSnck(w http.ResponseWriter, r *http.Request) {
	schmeal := WriteMeal()
	json.NewEncoder(w).Encode(schmeal[3])
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/schmeal/brkf", SendBrkf).Methods("GET")
	router.HandleFunc("/schmeal/lnch", SendLnch).Methods("GET")
	router.HandleFunc("/schmeal/dinr", SendDinr).Methods("GET")
	router.HandleFunc("/schmeal/snck", SendSnck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
