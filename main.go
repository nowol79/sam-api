package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "log"
    "fmt"
    "net/http"
)

type AreaCode struct {
    SSCODE    string   `json:"sscode,omitempty"`
    Search   string   `json:"search,omitempty"`
    Click   string   `json:"click,omitempty"`
}

var areaCodes []AreaCode

// Display all from the area codes var
func GetAreaCodes(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(areaCodes)
}

// Display a single data
func GetAreaCode(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range areaCodes {
        if item.Click == params["click"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&AreaCode{})
}

// create a new item
func CreateAreaCode(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fmt.Fprintf(w, "param: %v\n", params["click"])
    var code AreaCode
    _ = json.NewDecoder(r.Body).Decode(&code)
    code.Click = params["click"]
    areaCodes = append(areaCodes, code)
    json.NewEncoder(w).Encode(areaCodes)
}

// Delete an item
func DeleteAreaCode(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range areaCodes {
        if item.Click == params["click"] {
            areaCodes = append(areaCodes[:index], areaCodes[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(areaCodes)
    }
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    corsObj := handlers.AllowedOrigins([]string{"*"})

    areaCodes = append(areaCodes, AreaCode{SSCODE: "tab.m.all", Search: "img_icr", Click: "img_icr*a.flicking"})
    areaCodes = append(areaCodes, AreaCode{SSCODE: "tab.m.all", Search: "itl_pai", Click: "itl_pai.flicking"})

    router.HandleFunc("/areaCode", GetAreaCodes).Methods("GET")
    router.HandleFunc("/areaCode/{click}", GetAreaCode).Methods("GET")
    router.HandleFunc("/areaCode/{click}", CreateAreaCode).Methods("POST")
    router.HandleFunc("/areaCode/{click}", DeleteAreaCode).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))
}
