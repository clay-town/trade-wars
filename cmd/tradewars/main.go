package main

import (
    "log"
    "net/http"
    "os"
    //"github.com/gorilla/mux"
)

func main() {

  //mux := http.NewServeMux()
  //mux.HandleFunc("/", home)
  //mux.HandleFunc("/grid.html", grid)
  //mux.HandleFunc("/trade.html", trade)
  //mux.HandleFunc("/chat.html", chat)
  //mux.HandleFunc("/click.go", click)
  //mux.HandleFunc("/snippet", showSnippet)
  //mux.HandleFunc("/snippet/create", createSnippet)

  //http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  //myMux := mux.NewRouter()
  //myMux.HandleFunc("/", home)
  //myMux.HandleFunc("/index.html", home)
  //myMux.HandleFunc("/grid.html", grid)
  //myMux.HandleFunc("/trade.html", trade)
  //myMux.HandleFunc("/chat.html", chat)
  //myMux.HandleFunc("/click.go", click)
  //myMux.HandleFunc("/snippet", showSnippet)
  //myMux.HandleFunc("/snippet/create", createSnippet)
  //http.Handle("/", myMux)

  //http.HandleFunc("/", home)
  //http.HandleFunc("/index/", home)
  //http.HandleFunc("/grid/", grid)
  //http.HandleFunc("/trade/", trade)
  //http.HandleFunc("/chat/", chat)

  fs := http.FileServer(http.Dir("internal/ui"))
  http.Handle("/", fs)
  http.Handle("/index.html", fs)
  http.Handle("/grid.html", fs)
  http.Handle("/trade.html", fs)
  http.Handle("/chat.html", fs)

  log.Println("Starting server on :" + os.Getenv("PORT"))
  err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), nil)
  //err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)
  //err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), myMux)
  log.Fatal(err)
}