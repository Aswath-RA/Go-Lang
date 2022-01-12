package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Mobile structure
type Mobile struct {
	MobileId     string        `json:"mobileid"`
	MobileName   string        `json:"mobilename"`
	MobilePrice  int           `json:"price"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

//Manufacturer structure
type Manufacturer struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// DB
var mobiles []Mobile

func (c *Mobile) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.MobileName == ""
}
func main() {
	fmt.Println("Welcome to Mobile shop - Api")
	r := mux.NewRouter()

	//seeding
	mobiles = append(mobiles, Mobile{MobileId: "2", MobileName: "Moto E4", MobilePrice: 8999, Manufacturer: &Manufacturer{Fullname: "Motorola ltd", Website: "www.motorola.com"}})
	mobiles = append(mobiles, Mobile{MobileId: "4", MobileName: "Realme XT", MobilePrice: 16999, Manufacturer: &Manufacturer{Fullname: "Realme ltd", Website: "www.realme.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/mobiles", getAllmobiles).Methods("GET")
	r.HandleFunc("/mobile/{id}", getOnemobile).Methods("GET")
	r.HandleFunc("/mobile", createOnemobile).Methods("POST")
	r.HandleFunc("/mobile/{id}", updateOnemobile).Methods("PUT")
	r.HandleFunc("/mobile/{id}", deleteOnemobile).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Mobile shop</h1>"))
}

func getAllmobiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all mobiles")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(mobiles)
}

func getOnemobile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Mobile")
	w.Header().Set("Content-Type", "applicatioan/json")

	//getting id from request
	params := mux.Vars(r)

	for _, mobile := range mobiles {
		if mobile.MobileId == params["id"] {
			json.NewEncoder(w).Encode(mobile)
			return
		}
	}
	json.NewEncoder(w).Encode("No Mobile found with given id")
	return
}

func createOnemobile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Mobile")
	w.Header().Set("Content-Type", "applicatioan/json")

	// If body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var mobile Mobile
	_ = json.NewDecoder(r.Body).Decode(&mobile)
	if mobile.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//seeding to get random value
	rand.Seed(time.Now().UnixNano())
	mobile.MobileId = strconv.Itoa(rand.Intn(100))
	mobiles = append(mobiles, mobile)
	json.NewEncoder(w).Encode(mobile)
	return

}

func updateOnemobile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Mobile")
	w.Header().Set("Content-Type", "applicatioan/json")

	//grabing id from request
	params := mux.Vars(r)

	for index, mobile := range mobiles {
		if mobile.MobileId == params["id"] {
			mobiles = append(mobiles[:index], mobiles[index+1:]...)
			var mobile Mobile
			_ = json.NewDecoder(r.Body).Decode(&mobile)
			mobile.MobileId = params["id"]
			mobiles = append(mobiles, mobile)
			json.NewEncoder(w).Encode(mobile)
			return
		}
	}

}

func deleteOnemobile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	for index, mobile := range mobiles {
		if mobile.MobileId == params["id"] {
			mobiles = append(mobiles[:index], mobiles[index+1:]...)

			break
		}
	}
}
