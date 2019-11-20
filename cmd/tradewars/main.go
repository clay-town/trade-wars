package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/grid.html", grid)
    mux.HandleFunc("/trade.html", trade)
    mux.HandleFunc("/chat.html", chat)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("Starting server on :" + os.Getenv("PORT"))
    err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)
    log.Fatal(err)
}

/*func main() {
  fs := http.FileServer(http.Dir("internal/ui"))
  http.Handle("/", fs)
  http.Handle("/grid.html", fs)
  http.Handle("/trade.html", fs)
  http.Handle("/chat.html", fs)

  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/grid.html", grid)
  mux.HandleFunc("/trade.html", trade)
  mux.HandleFunc("/chat.html", chat)

  log.Println("Starting server on :4000")
  err := http.ListenAndServe("0.0.0.0:4000", nil)
  log.Fatal(err)
}*/