
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
