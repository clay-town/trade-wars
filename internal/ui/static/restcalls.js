
function testRestCall() {
    var request = new XMLHttpRequest()
    request.open('GET', 'https://ghibliapi.herokuapp.com/films', true)
    request.onload = function() {
      var data = JSON.parse(this.response)
      if (request.status >= 200 && request.status < 400){
        data.forEach(movie => {
          console.log(movie.title)
        })
      } else {
        console.log('error')
      }
    }
    request.send()
}

function createNewUser(){
  alert("create")
  var request = new XMLHttpRequest()
  request.open('POST', '/createNewUser', true)
  request.onload = function() {
    alert("request.onload")
  }
  request.send()
}

function userLogin() {
  var request = new XMLHttpRequest()
  request.open('POST', '/players', true)
  request.onload = function(){
  //  var data = JSON.parse(this.response)

  }
  request.send()
}

// This function listens for the user to click the login button 
$(document).ready(function(){
  $("#one").click(function(){
    var request = new XMLHttpRequest()
    var callsign = $('#callsigninput').val();
    request.open('POST', '/players?callsign='+callsign, true)
    request.onload = function(response){
      console.log(response.currentTarget.status)
      var status = response.currentTarget.status;
      if status == 200 { // success
          //redirect page to map
          window.location = "/map.html"
          //initilize map population functions
      } else if (status == 500 ) { //failure
          //insert prompt for user to create a new handle
      }
    }
    request.send()
  });
});
