function sendGetRequest() {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/register", true);
    xhr.onreadystatechange = function() {
      if (xhr.readyState === 4 && xhr.status === 200) {
        console.log(xhr.responseText);
        document.open();
        document.write(xhr.responseText);
        document.close();
      }
    };
    xhr.send();
    history.pushState({}, null, '/register');
  }

function sendGetLoginStudent() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/Student/Login", true);
  xhr.onreadystatechange = function() {
      console.log(xhr.responseText);
        document.open();
        document.write(xhr.responseText);
        document.close();
  };
    xhr.send();
    history.pushState({}, null, '/Student/Login');
}

function sendGetRegisterStudent() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/Student/Register", true);
  xhr.onreadystatechange = function() {
      console.log(xhr.responseText);
        document.open();
        document.write(xhr.responseText);
        document.close();
  };
    xhr.send();
    history.pushState({}, null, '/Student/Register');
}

