//refactor this into Go
var html = "<table id='map'>";
var shipName1 = "<img src='../static/img/spaceShip.jpg'>"
var stationName1 = "<img src='../static/img/spaceStation.jpg'>"
//var stationLocation = "x"+Math.floor(Math.random() * 10)+"y"+Math.floor(Math.random() * 10);
//var shipLocation = "x"+Math.floor(Math.random() * 10)+"y"+Math.floor(Math.random() * 10);
var td = "<img src='../static/img/space.jpg'>";

for (var y = 0; y < 10; y++) {
    html+="<tr>";
    for (var x = 0; x < 10; x++) {
        html+="<td"+' id=x'+x+'y'+y+">"+td+"</td>";
    }
}
updateLocalPlayerInformation('clay-town');
document.getElementById("tableLocation").innerHTML = html

function updateLocalPlayerInformation(callsign){
  var request = new XMLHttpRequest()
  request.open('POST', '/playerInformation?callsign='+callsign)
  request.onload = function() {
    var data = JSON.parse(this.response)
    console.log(data.location)
    document.getElementById(data.location).innerHTML = shipName1
    //populate cargo and bucks divs here 
  }
  request.send()
}

//make call to db to retrieve map information from server
function retrieveStationLocations(){
  var request = new XMLHttpRequest()
  request.open('POST', '/stationInformation', true)
  request.onload = function() {

  }
  request.send()
}
