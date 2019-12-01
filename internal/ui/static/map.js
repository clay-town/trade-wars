// refactor structs so that they carry their own image
var html = newTable()
var shipName1 = "<img src='../static/img/spaceShip.jpg'>"
var stationName1 = "<img src='../static/img/spaceStation.jpg'>"

document.getElementById("tableLocation").innerHTML = html // place grid in DOM
updateLocalPlayerInformation('clay-town');                // place player information in DOM
updateLocalSpaceStationInformation();                     // place stations on map

moveShip('clay-town','left')

function moveShip(callsign, direction) {
  // make call to server to update ship location
  // replace current location with empty image
  // repaint ship location on map
  var request = new XMLHttpRequest()
  request.open('POST', '/updatePlayerLocation?callsign='+callsign+"&dir="+direction, true)
  request.onload = function(){
    // check return status
  }
  request.send()
}

function newTable() {
  var td = "<img src='../static/img/space.jpg'>";
  html = "<table id='map'>";
  for (var y = 0; y < 10; y++) { // build HTML table for grid map
      html+="<tr>";
      for (var x = 0; x < 10; x++) {
          html+="<td"+' id=x'+x+'y'+y+">"+td+"</td>";
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
    //populate cargo and bucks divs here
  }
  request.send()
}
