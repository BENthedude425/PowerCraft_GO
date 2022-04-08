function SetCookies(username, password){
    var request = new XMLHttpRequest();
    var username = document.getElementById("usernamefield").value;
    var password = document.getElementById("passwordfield").value;
    var data = "username="+username+"&password="+password;
    request.open("POST",data,true);
    request.send("");

    request.onreadystatechange = function(){
        if(request.readyState == 4){
            console.log(request.responseText);
            var text = request.responseText;
            var obj = JSON.parse(text);
            
            token = obj["Token"]
            const d = new Date();
            d.setTime(d.getTime() + (7*24*60*60*1000));
            let expires = "expires=" + d.toUTCString();
            document.cookie = "Token=" + token + ";" + expires 

            document.location = obj["Redirect"]
        }
    }

}