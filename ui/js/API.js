function API(method, action, data){
    request = new XMLHttpRequest();
    request.open(method, action);
    request.setRequestHeader("Content-type" , "application/x-www-form-urlencoded")
    if (method == "GET"){
        request.send(data);
    }else if (method == "POST"){
        var formContainer = document.getElementById(data);
        var formValues = []
        for(i=0; i < formContainer.childElementCount; i++){
            if (formContainer.children[i].tagName == "INPUT"){
                value = formContainer.children[i].value;
                name = formContainer.children[i].name;
                formValues.push(name + "=" + value);
            }
        }
        
        request.send(formValues.join("&"));
    }
}