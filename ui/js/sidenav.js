function HideElements(parentname){
    var nodes = document.getElementById(parentname).getElementsByTagName("div");
    for (var i=0; i<nodes.length; i++){
        child = nodes[i];
        if (child.id != "Console"){
            document.getElementById(child.id+"Container").style.visibility = "hidden";
        }
    }
    GetServerSettings()
}

function SideNavSelect(item){
    var nodes = document.getElementById("SideNav").getElementsByTagName("div");
    for (var i=0; i<nodes.length; i++){
        var child = nodes[i]
        if (child.id == item){
            document.getElementById(item+"Container").style.visibility = "visible";
            document.getElementById(child.id).style.backgroundColor = "grey";
        }else{
            document.getElementById(child.id+"Container").style.visibility = "hidden";
            document.getElementById(child.id).style.backgroundColor = "rgba(0,0,10, 0.85)";
        }
    }

    
}