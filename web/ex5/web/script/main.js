var display = function() {
  var xmlhttp = new XMLHttpRequest();
  var url = "./statdata";

  xmlhttp.onreadystatechange = function() {
      if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
          var data = JSON.parse(xmlhttp.responseText);
          publish(data);
      }
  };
  xmlhttp.open("GET", url, true);
  xmlhttp.send();
}

var publish = function(data) {
  var out = "<h4>Selected: " + data.selected + "</h4><ul><li>";
  data.factors = data.factors || []
  if ( data.factors.length == 0 ) {
    out += "Prime"
  } else {
    out += data.factors.join("</li><li>")
  }
  out += "</li></ul>"
  document.getElementById("demo").innerHTML = out;
}
