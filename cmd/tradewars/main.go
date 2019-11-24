package main

import (
    "log"
    "net/http"
    "os"
    "github.com/joho/godotenv"
)

func main() {
  fs := http.FileServer(http.Dir("./internal/ui/static/"))
  mux := http.NewServeMux()
  mux.HandleFunc("/players", playersHandler)
  mux.HandleFunc("/", homeHandler)
  mux.HandleFunc("/index.html", homeHandler)
  mux.HandleFunc("/map.html", mapHandler)
  mux.HandleFunc("/map", mapHandler)
  mux.HandleFunc("/trade.html", tradeHandler)
  mux.HandleFunc("/chat.html", chatHandler)
  mux.Handle("/static/", http.StripPrefix("/static", fs))

  godotenv.Load()

  log.Println("Starting server on :" + os.Getenv("PORT"))
  err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)
  log.Fatal(err)
}
