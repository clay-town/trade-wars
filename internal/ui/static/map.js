// refactor structs so that they carry their own image
var html = newTable()
var shipName1 = "<img src='../static/img/spaceShip.jpg'>"
var stationName1 = "<img src='../static/img/spaceStation.jpg'>"

document.getElementById("tableLocation").innerHTML = html // place grid in DOM
updateLocalPlayerInformation('clay-town');                // place player information in DOM
updateLocalSpaceStationInformation();                     // place stations on map

function moveShip(callsign, direction) {
  var request = new XMLHttpRequest()
  request.open('POST', '/updatePlayerLocation?callsign='+callsign+"&dir="+direction, true)
  request.onload = function(){
    var data = JSON.parse(this.response)
    //check to see if station exists first, don't paint over station
    document.getElementById(data[1]).innerHTML = shipName1  //move ship to new location on grid
    document.getElementById(data[0]).innerHTML = "<img src='../static/img/space.jpg'>" //replace old location with empty space
    console.log(data)
  }
  request.send()
}

function newTable() {
  var td = "<img src='../static/img/space.jpg'>";
  html = "<table id='map'>";
  for (var y = 0; y < 10; y++) { // build HTML table for grid map
      html+="<tr>";
      for (var x = 0; x < 10; x++) {
          html+="<td"+' id='+x+':'+y+">"+td+"</td>";
      }
  }
  return html
}

function updateLocalSpaceStationInformation(){
  var request = new XMLHttpRequest()
  request.open('POST', '/stationInformation', true)
  request.onload = function() {
    var data = JSON.parse(this.response)
    for (i = 0; i<data.length; i++) {
      document.getElementById(data[i].location).innerHTML = stationName1
    }
  }
  request.send()
}

function updateLocalPlayerInformation(callsign){
  var request = new XMLHttpRequest()
  request.open('POST', '/playerInformation?callsign='+callsign)
  request.onload = function() {
    var data = JSON.parse(this.response)
    document.getElementById(data.location).innerHTML = shipName1
    document.getElementById("cubits").innerHTML = data["cubits"]
    console.log(data["cargos"].length)
    cargo = ""
    for (i = 0; i<data["cargos"].length; i++){
      cargo += data.cargos[i].item + ": " + data.cargos[i].quantitiy + "<br>"
    }
    document.getElementById("cargo").innerHTML = cargo
    //"Dog treats: 12 <br> Concert Tickets: 2 <br> Fuel: 3"
    //populate cargo and bucks divs here
  }
  request.send()
}
