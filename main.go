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

//SchoolMeal is struct of schmeal for output(GET)
type SchoolMeal struct {
	Brkf []string
	Lnch []string
	Dinr []string
	Snck []string
	err  string
}

var schmeal []SchoolMeal

//GetMeal parses the meal info from school website
func GetMeal() []soup.Root {
	resp, err := soup.Get("http://hana.hs.kr/life/meal.asp?yy=2018&mm=11")
	fmt.Println("http transport error is:", err)

	doc := soup.HTMLParse(resp)
	meal := doc.Find("table", "class", "today_meal").FindAll("td")

	return meal
}

//WriteMeal makes the struct
func WriteMeal() SchoolMeal {
	meal := GetMeal()
	sch := SchoolMeal{}
	if len(meal) > 1 {
		brkf := strings.Split(meal[0].Text(), ",")
		lnch := strings.Split(meal[1].Text(), ",")
		dinr := strings.Split(meal[2].Text(), ",")
		snck := strings.Split(meal[3].Text(), ",")

		sch.Brkf = brkf
		sch.Lnch = lnch
		sch.Dinr = dinr
		sch.Snck = snck

		fmt.Println(brkf)
	} else {
		sch.err = meal[0].Text()
	}
	return sch
}

//SendMeal is for api work
func SendMeal(w http.ResponseWriter, r *http.Request) {
	schmeal := WriteMeal()
	json.NewEncoder(w).Encode(schmeal)
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/schmeal", SendMeal).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//idk what is happening...
//https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
