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
      display('Passwords do not match');
      return;
    } else if (username.length < 1) {
      display('Username not filled in');
      return;
    } else if (password.length < 1) {
      display('Password not filled in');
      return;
    } else if (cpassword.length < 1) {
      display('Check password not filled in');
      return;
    } else if (recaptcha.length < 1) {
      recaptcha = 'none';
    }

    var url = 'api/adduser/' + username + '/' + password + '/' + recaptcha;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (xhttp.readyState == 4) {
        try {
          var data = JSON.parse(xhttp.responseText);
          if (typeof data['Error'] === 'string') {
            display('Error: ' + data.Error);
            console.log(data);
          } else {
            display('Success');
          }
        } catch (e) {
          console.log(xhttp.responseText);
          console.log(e);
          display('Unknown Error');
        }
      }
    };
    xhttp.open('GET', url, true);
    xhttp.send();
  };
};
