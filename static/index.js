var messageField;
var usernameField;
var passwordField;
var cpasswordField;
var addButton;

var display = function (message) {
  messageField.innerText = message;
}


window.onload = function () {
  messageField = document.getElementById('message');
  usernameField = document.getElementById('username');
  passwordField = document.getElementById('password');
  cpasswordField = document.getElementById('cpassword');
  addButton = document.getElementById('add');
  
  addButton.onclick = function () {
    var username = usernameField.value;
    var password = passwordField.value;
    var cpassword = cpasswordField.value;
    var recaptcha = grecaptcha.getResponse();
    
    if (password !== cpassword) {
      display("Passwords do not match");
      return;
    } else if (username.length < 1) {
      display("Username not filled in");
      return;
    } else if (password.length < 1) {
      display("Password not filled in");
      return;
    } else if (cpassword.length < 1) {
      display("Check password not filled in");
      return;
    }

    var url = "api/adduser/" + username + "/" + password + "/" + recaptcha;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (xhttp.readyState == 4) {
        display(xhttp.responseText);
      }
    };
    xhttp.open("GET", url, true);
    xhttp.send();
  };
};
