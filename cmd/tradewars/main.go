package main

import (
    s "github.com/clay-town/trade-wars/internal/tradewars"
    "log"
	"net/http"
	"os"
    "github.com/joho/godotenv"
    "io/ioutil"
    "encoding/json"
)
var jsonShips s.Ships

func main() {
    unmarshalJSONFile()

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
    //mux.HandleFunc("/returnUserInfo", returnUserInfo)
    mux.Handle("/static/", http.StripPrefix("/static", fs))

	godotenv.Load()

	log.Println("Starting server on " + os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)
	log.Fatal(err)
}

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
