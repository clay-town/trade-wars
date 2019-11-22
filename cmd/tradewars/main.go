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
  mux.HandleFunc("/", home)
  mux.HandleFunc("/grid.html", grid)
  mux.HandleFunc("/trade.html", trade)
  mux.HandleFunc("/chat.html", chat)
  mux.Handle("/static/", http.StripPrefix("/static", fs))

  godotenv.Load()

  log.Println("Starting server on :" + os.Getenv("PORT"))
  err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)
  log.Fatal(err)
}
