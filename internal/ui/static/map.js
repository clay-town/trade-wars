// refactor structs so that they carry their own image
var html = newTable()
var shipName1 = "<img src='../static/img/spaceShip.jpg'>"
var stationName1 = "<img src='../static/img/spaceStation.jpg'>"

function moveShip(callsign, direction) {
  var request = new XMLHttpRequest()
  request.open('POST', '/updatePlayerLocation?callsign='+callsign+"&dir="+direction, true)
  request.onload = function(){
    var data = JSON.parse(this.response)
    // data is a 2d array. the 0th Element :[previousLocation, newlocation]
    //                     the 1st Element :[stationName, stationLocation,stationName, stationLocation]
    updateMap(data)
  }
  request.send()
}

function updateMap(data) {
  newLoc = data[0][1]
  oldLoc = data[0][0]
  stationLoc1 = data[1][1]
  stationLoc2 = data[1][3]
  stationName1 = data[1][0]
  stationName2 = data[1][2]
  //check to see if station exists first, don't paint over station
  if (newLoc == stationLoc1) {
      document.getElementById(oldLoc).innerHTML = "<img src='../static/img/space.jpg'>"
      //insert station name & trade button into Dom
      document.getElementById("nearby").innerHTML = stationName1
  } else if (newLoc == stationLoc2) {
    document.getElementById(oldLoc).innerHTML = "<img src='../static/img/space.jpg'>"
    //insert station name & trade button into Dom
    document.getElementById("nearby").innerHTML = stationName2
  } else if (oldLoc == stationLoc1 || oldLoc == stationLoc2) {
      // leaving a station
      //paint ship in new location, don't paint over old location
      document.getElementById(newLoc).innerHTML = shipName1  //move ship to new location on grid
      document.getElementById("nearby").innerHTML = "The Great Expanse"
  } else {
    document.getElementById(newLoc).innerHTML = shipName1  //move ship to new location on grid
    document.getElementById(oldLoc).innerHTML = "<img src='../static/img/space.jpg'>" //replace old location with empty space
  }
}

// This function listens for the user to click the login button
$(document).ready(function(){
  $("#one").click(function(){
    var request = new XMLHttpRequest()
    var callsign = $('#callsigninput').val();
    request.open('POST', '/players?callsign='+callsign, true)
    request.onload = function(response){
      var status = response.currentTarget.status;
      if (status == 200) { // success
          window.location = "/map.html";
      } else if (status == 500 ) { //failure
          //insert prompt for user to create a new handle
      }
    }
    request.send()
  });
});

function insertMap(callsign){
  document.getElementById("tableLocation").innerHTML = html // place grid in DOM
  updateLocalSpaceStationInformation();                     // place stations on map
  updateLocalPlayerInformation(callsign);
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
    cargo = ""
    //testCargo = data["cargo"]
    for (i = 0; i<data["cargos"].length; i++){
      cargo += data.cargos[i].item + ": " + data.cargos[i].quantitiy + "<br>"
    }
    document.getElementById("cargo").innerHTML = cargo
    document.getElementById("currentLocation").innerHTML = data.location+"]"
  }
  request.send()
}
