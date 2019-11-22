package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    log.Println("This is the home function")

    if r.URL.Path != "/" && r.URL.Path != "/index.html"{
        http.NotFound(w, r)
        return
    }

    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles("internal/ui/html/index.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // We then use the Execute() method on the template set to write the template
    // content as the response body. The last parameter to Execute() represents any
    // dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func grid(w http.ResponseWriter, r *http.Request) {
    log.Println("This is the grid function")

    if r.URL.Path != "/grid.html" {
        http.NotFound(w, r)
        return
    }

    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles("internal/ui/html/grid.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // We then use the Execute() method on the template set to write the template
    // content as the response body. The last parameter to Execute() represents any
    // dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func trade(w http.ResponseWriter, r *http.Request) {
    log.Println("This is the trade function")

    if r.URL.Path != "/trade.html" {
        http.NotFound(w, r)
        return
    }

    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles("internal/ui/html/trade.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // We then use the Execute() method on the template set to write the template
    // content as the response body. The last parameter to Execute() represents any
    // dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func chat(w http.ResponseWriter, r *http.Request) {
    log.Println("This is the chat function")

    if r.URL.Path != "/chat.html" {
        http.NotFound(w, r)
        return
    }

    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles("internal/ui/html/chat.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // We then use the Execute() method on the template set to write the template
    // content as the response body. The last parameter to Execute() represents any
    // dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    w.Write([]byte("Create a new snippet..."))
}