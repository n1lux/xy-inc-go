package main

import(
	"fmt"
	"log"
	"math"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	
)

type Poi struct{	
	Name string `json: "Name"`
	X int `json: "X"`
	Y int `json: "Y"`
}

type Pois []Poi

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the index")
	fmt.Println("Endpoint Hit: index")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", index).Methods("GET")
	myRouter.HandleFunc("/api/pois", listPoisHandler).Methods("GET")
	myRouter.HandleFunc("/api/pois", createPoiHandler).Methods("POST")
	myRouter.HandleFunc("/api/pois/search", searchPoisHandler).Methods("GET")		
	log.Fatal(http.ListenAndServe(":8080", myRouter))	
}

func listPoisHandler(res http.ResponseWriter, req *http.Request) {
	var pois []Poi
	db := InitDb()

	if err := db.Find(&pois).Error; err != nil {		
	    fmt.Println(err)
	} else {
	    json.NewEncoder(res).Encode(pois)
	}
	
	fmt.Println("Endpoint Hit: returnAllPois")
	
}

func createPoiHandler(res http.ResponseWriter, req *http.Request){
	//db := InitDb()
	//defer db.Close()
	decoder := json.NewDecoder(req.Body)
	var poi Poi

	err := decoder.Decode(&poi)

	if err != nil{
		panic(err)
	}

	db := InitDb()
	db.Create(&poi)

	json.NewEncoder(res).Encode(poi)
	
}

func searchPoisHandler(res http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	var pois []Poi
	var pois_return []Poi
	db := InitDb()
	poix, _ := strconv.Atoi(vars.Get("x"))
	poiy, _ := strconv.Atoi(vars.Get("y"))
	radius, _ := strconv.Atoi(vars.Get("d-max"))

	if err := db.Find(&pois).Error; err != nil {		
	    fmt.Println(err)
	} else {
		for _, poi := range pois{
			sqrt_r := math.Sqrt(math.Pow(float64(poi.X - poix), float64(2)) + math.Pow(float64(poi.Y - poiy), float64(2)))
			if ( int(sqrt_r) <= radius ){
				pois_return = append(pois_return, poi)
			}
		}

	    json.NewEncoder(res).Encode(pois_return)
	}
		
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprint(w, "Key: " + key)
	
}

func InitDb() *gorm.DB {
	//Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	//Error
	if err != nil{
		panic(err)
	}

	db.AutoMigrate(&Poi{})

	//Creating the table
	if !db.HasTable(&Poi{}){
		db.CreateTable(&Poi{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(Poi{})
	}

	return db	
}

func xy_mock() {
	db := InitDb()
	fmt.Println("Erase all data")
	db.Model(&Poi{}).Delete(&Poi{})

	fmt.Println("Add mock data")
	var pois = []Poi{
        Poi{Name: "Lanchonete", X: 27, Y: 12},
        Poi{Name: "Posto", X: 31, Y: 18},
        Poi{Name: "Joalheria", X: 15, Y: 12},
        Poi{Name: "Floricultura", X: 19, Y: 21},
        Poi{Name: "Pub", X: 12, Y: 8},
        Poi{Name: "Supermercado", X: 23, Y: 6},
        Poi{Name: "Churrascaria", X: 28, Y: 2},
    }

    for _, poi := range pois{
    	db.Create(&poi)
    }
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	xy_mock()
	handleRequests()
}