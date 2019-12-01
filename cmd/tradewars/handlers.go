package main

import (
    s "github.com/clay-town/trade-wars/internal/tradewars"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"
    "encoding/json"
    "strings"
  //  "github.com/gorilla/mux"
)

func updatePlayerLocation(w http.ResponseWriter, r *http.Request) {
  //Update location of players ship!
  callsign := r.URL.Query().Get("callsign")
  direction := r.URL.Query().Get("dir")
  for i := 0; i < len(jsonShips.Ships); i++ {
      if callsign == jsonShips.Ships[i].Callsign {
        origin := jsonShips.Ships[i].Location
        oldLoc := strings.Split(jsonShips.Ships[i].Location, ":")
        newLoc := spliceAndAdjustLocation(oldLoc, direction)
        jsonShips.Ships[i].Location = newLoc// set new location here
        //jsonStations.Stations[0]
        //stationArray := []s.Station{}
        stationArray := []string{}
        shipArray := []string{}
        dataArray := [][]string{}

        shipArray = append(shipArray, origin)
        shipArray = append(shipArray, newLoc)
        for i := 0; i < len(jsonStations.Stations); i++ {
          stationArray = append(stationArray, jsonStations.Stations[i].Designation)
          stationArray = append(stationArray, jsonStations.Stations[i].Location)
        }
        dataArray = append(dataArray, shipArray)
        dataArray = append(dataArray, stationArray)

        //dataArray = append(dataArray, station)
        json.NewEncoder(w).Encode(dataArray) // return
      }
  }
}

func createNewUser(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
    }
    if r.Method == "POST" {
        // need to fix the starting equipment
        var newShip s.Ship
        var cargo s.Cargo
        var cargos []s.Cargo

        cargos = append(cargos, cargo)
        newShip.Callsign = r.FormValue("callsign")
        newShip.Location = "4:0"
        newShip.Cubits = 400
        newShip.Cargos = cargos
        log.Println(newShip)
        jsonShips.Ships = append(jsonShips.Ships, newShip)
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
    callsign := r.URL.Query().Get("callsign")
    log.Println(callsign)
    for i := 0; i < len(jsonShips.Ships); i++ {
        if callsign == jsonShips.Ships[i].Callsign {
          // match found : set cookie
          cookie := http.Cookie {
                  Name: "callsign",
                  Value: callsign,
                  Expires: time.Now().AddDate(0, 0, 1),
              Path: "/",
            }
          http.SetCookie(w, &cookie)
          w.WriteHeader(http.StatusOK)
          return
        }
    }
    // No match found
    w.WriteHeader(http.StatusInternalServerError)
}

func spliceAndAdjustLocation(oldLoc []string, direction string) string{
  intArr := []int{}
  for _, x := range oldLoc {  // convert string array into int array
    y, err := strconv.Atoi(x)
    if err != nil{
      panic(err)
    }
    intArr = append(intArr, y)
  }
  switch direction {
  case "up":
    intArr[1]--
    if intArr[1] == -1 {
      intArr[1] = 9
    }
  case "down":
    intArr[1]++
    if intArr[1] == 10 {
      intArr[1] = 0
    }
  case "left":
    intArr[0]--
    if intArr[0] == -1 {
      intArr[0] = 9
    }
  case "right":
    intArr[0]++
    if intArr[0] == 10 {
      intArr[0] = 0
    }
  }
  newLoc := strconv.Itoa(intArr[0]) + ":" + strconv.Itoa(intArr[1])
  return newLoc
}

func returnPlayerInformation(w http.ResponseWriter, r *http.Request) {
  // returns information for ship matching the users callsign
  callsign := r.URL.Query().Get("callsign")
  for i := 0; i < len(jsonShips.Ships); i++ {
      if callsign == jsonShips.Ships[i].Callsign {
        json.NewEncoder(w).Encode(jsonShips.Ships[i])
      }
  }
}

func returnStationInformation(w http.ResponseWriter, r *http.Request) {
    // returns information for all of the stations
    json.NewEncoder(w).Encode(jsonStations.Stations)
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
    ts, err := template.ParseFiles("internal/ui/html/map.tmpl")
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
    log.Println(jsonShips)
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
