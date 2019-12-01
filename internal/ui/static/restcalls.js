
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

$(document).ready(function(){
  $("#one").click(function(){
    var request = new XMLHttpRequest()
    var callsign = $('#callsigninput').val();
    request.open('POST', '/players?callsign='+callsign, true)
    request.onload = function(response){
      console.log(response)
      console.log(response.currentTarget.status)
      //var data = JSON.parse(this.response)
      // check headers:
      // success: redirect
    //  window.location = "/map.html"
    }

    request.send()
  });
});
