package main

import (
  s "github.com/clay-town/trade-wars/internal/tradewars"
  "log"
	"net/http"
	"os"
  "github.com/joho/godotenv"
  "io/ioutil"
  "encoding/json"
  //"github.com/gorilla/mux"
)

// this would be easier if we first outline reusable and returnable data structures
var jsonShips s.Ships
var jsonStations s.Stations

func main() {
  unmarshalJSONFile()
  unmarshalStations()
  fs := http.FileServer(http.Dir("./internal/ui/static/"))
  mux := http.NewServeMux()
	mux.HandleFunc("/players", playersHandler)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/index.html", homeHandler)
	mux.HandleFunc("/map.html", mapHandler)
	mux.HandleFunc("/map", mapHandler)
	mux.HandleFunc("/trade.html", tradeHandler)
	mux.HandleFunc("/chat.html", chatHandler)
  mux.HandleFunc("/createNewUser", createNewUser)
  mux.HandleFunc("/stationInformation", returnStationInformation)
  mux.HandleFunc("/playerInformation", returnPlayerInformation)
  mux.HandleFunc("/updatePlayerLocation", updatePlayerLocation)
  mux.HandleFunc("/nearby", nearbyHandler)
  mux.HandleFunc("/updateonline", updateOnlineHandler)



  mux.Handle("/static/", http.StripPrefix("/static", fs))
	godotenv.Load()
	log.Println("Starting server on " + os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)
	log.Fatal(err)
}
 // refactor these functions into one another
func unmarshalJSONFile() {
    jsonFile, err := os.Open("internal/tradewars/data.json")
    if err!= nil{
        log.Println(err)
    }
    byteValue, _ := ioutil.ReadAll(jsonFile)
    defer jsonFile.Close()
    json.Unmarshal(byteValue, &jsonShips)
    for i := 0; i < len(jsonShips.Ships); i++ {
        log.Println("Ship Location: " + jsonShips.Ships[i].Location)
    }
}

func unmarshalStations() {
    jsonFile, err := os.Open("internal/tradewars/data.json")
    if err!= nil{
        log.Println(err)
    }
    byteValue, _ := ioutil.ReadAll(jsonFile)
    defer jsonFile.Close()
    json.Unmarshal(byteValue, &jsonStations)
    for i := 0; i < len(jsonStations.Stations); i++ {
        log.Println("Station Name: " + jsonStations.Stations[i].Designation)
    }
}
