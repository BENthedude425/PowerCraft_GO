async function ServerAPI(url, console, callback){
    var request = new XMLHttpRequest();
    if (console == true){
        request.open("GET", "MCcommand="+url, true);
    }else{
        request.open("GET", url, true);
    }

    request.onreadystatechange = function (){
        if (request.readyState == 4){
            callback(request.responseText);
        } 
    }

    request.send();
    
}

function hash(string) {
    const utf8 = new TextEncoder().encode(string);
    return crypto.subtle.digest('SHA-256', utf8).then((hashBuffer) => {
      const hashArray = Array.from(new Uint8Array(hashBuffer));
      const hashHex = hashArray
        .map((bytes) => bytes.toString(16).padStart(2, '0'))
        .join('');
      return hashHex;
    });
  }

async function GetLiveServerFeed(feed){
    var Console = document.getElementById("ConsoleText");

    if (feed != "timeout"){ 
        Console.innerText += feed;
        Console.scrollIntoView(false);
    }
    
    texthash = await hash(Console.innerText);
    console.log(Console.innerText);
    ServerAPI('ConsoleFeed-hash' + texthash + "-ln" + Console.innerText.length, false, GetLiveServerFeed);
}

async function FirstRequest(){
    ServerAPI('ConsoleFeed-first', false, GetLiveServerFeed);
}