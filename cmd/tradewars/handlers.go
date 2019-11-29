package main

import (
    s "github.com/clay-town/trade-wars/internal/tradewars"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"
    "io/ioutil"
    "encoding/json"
)

var jsonShips = s.Ships{
    {
        Callsign:   "aluminum-raccoon",
        Location:   "x4y8",
    },
}

func createNewUser(w http.ResponseWriter, r *http.Request){
    var newShip s.Ship
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Error in user creation")
    }
    json.Unmarshal(reqBody, &newShip)
    jsonShips = append(jsonShips, newShip)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newShip)
}

// func returnUserInfo(w http.ResponseWriter, r *http.Request) {
//     userID := mux.Vars(r)["id"]
//
//     for _, Ship := range jsonShips {
//         if Ship.Callsign == userID {
//             json.NewEncoder(w).Encode(Ship)
//         }
//     }
// }

//Add function to update ships information

func playersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        ts, err := template.ParseFiles("internal/ui/html/index.html")
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }

        err = ts.Execute(w, nil)
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }

    } else if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
        callsign := r.Form.Get("callsign")
        cookie := http.Cookie {
            Name: "callsign",
            Value: callsign,
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",
        }
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/map.html", http.StatusSeeOther)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
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

func mapHandler(w http.ResponseWriter, r *http.Request) {
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error: Could not obtain callsign from cookie", 500)
        return
    }
    callsign := cookie.Value
    if r.URL.Path != "/map.html" && r.URL.Path != "/map"{
        http.NotFound(w, r)
        return
    }
    htmlCallsign := map[string]interface{}{"callsign": callsign}
    // Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message and use
    // the http.Error() function to send a generic 500 Internal Server Error
    // response to the user.
    ts, err := template.ParseFiles("internal/ui/html/map.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }
    // We then use the Execute() method on the template set to write the template
    // content as the response body. The last parameter to Execute() represents any
    // dynamic data that we want to pass in, which for now we'll leave as nil.
    err = ts.Execute(w, htmlCallsign)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func tradeHandler(w http.ResponseWriter, r *http.Request) {
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error: Could not obtain callsign from cookie", 500)
        return
    }
    callsign := cookie.Value
    log.Println("from trade handler - Callsign: " + callsign)
    if r.URL.Path != "/trade.html" {
        http.NotFound(w, r)
        return
    }
    htmlCallsign := map[string]interface{}{"callsign": callsign}
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
    err = ts.Execute(w, htmlCallsign)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error: Could not obtain callsign from cookie", 500)
        return
    }
    callsign := cookie.Value
    log.Println("from chat handler - Callsign: " + callsign)

    if r.URL.Path != "/chat.html" {
        http.NotFound(w, r)
        return
    }
    htmlCallsign := map[string]interface{}{"callsign": callsign}
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
    err = ts.Execute(w, htmlCallsign)
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
