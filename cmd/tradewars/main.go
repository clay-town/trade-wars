package main

import (
    "log"
    "net/http"
    "os"
    //"github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {


  // // The following chunk of code is for 'mux'
  // mux := http.NewServeMux()
  // mux.HandleFunc("/", home)
  // mux.HandleFunc("/grid.html", grid)
  // mux.HandleFunc("/trade.html", trade)
  // mux.HandleFunc("/chat.html", chat)
  // err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), mux)


  // // This chunk of code is for 'myMux'
  //myMux := mux.NewRouter()
  //myMux.HandleFunc("/", home)
  //myMux.HandleFunc("/index.html", home)
  //myMux.HandleFunc("/grid.html", grid)
  //myMux.HandleFunc("/trade.html", trade)
  //myMux.HandleFunc("/chat.html", chat)
  //http.Handle("/", myMux)
  //err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), myMux)


  // // This chunk of code uses HandleFunc
  //http.HandleFunc("/", home)
  //http.HandleFunc("/index/", home)
  //http.HandleFunc("/grid/", grid)
  //http.HandleFunc("/trade/", trade)
  //http.HandleFunc("/chat/", chat)


  // This chunk of code uses handle
  fs := http.FileServer(http.Dir("internal/ui"))
  http.Handle("/", fs)
  http.Handle("/index.html", fs)
  http.Handle("/grid.html", fs)
  http.Handle("/trade.html", fs)
  http.Handle("/chat.html", fs)

  godotenv.Load()

  log.Println("Starting server on :" + os.Getenv("PORT"))
  err := http.ListenAndServe(os.Getenv("CHROMEHOST") + ":" + os.Getenv("PORT"), nil)



  //http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


  log.Fatal(err)
}
