//refactor this into Go
var html = "<table id='map'>";
var shipName1 = "<img src='../static/img/spaceShip.jpg'>"
var stationName1 = "<img src='../static/img/spaceStation.jpg'>"
var stationLocation = "x"+Math.floor(Math.random() * 10)+"y"+Math.floor(Math.random() * 10);
var shipLocation = "x"+Math.floor(Math.random() * 10)+"y"+Math.floor(Math.random() * 10);
var td = "<img src='../static/img/space.jpg'>";

for (var y = 0; y < 10; y++) {
    html+="<tr>";
    for (var x = 0; x < 10; x++) {
        html+="<td"+' id=x'+x+'y'+y+">"+td+"</td>";
    }
}
document.getElementById("tableLocation").innerHTML = html
document.getElementById(stationLocation).innerHTML = stationName1
document.getElementById(shipLocation).innerHTML = shipName1
