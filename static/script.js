async function sendMessage() {
    var inputID = document.getElementById("inputID").value;
    var inputa = document.getElementById("inputName").value;
    var inputb = document.getElementById("inputSurname").value;
    var inputc = document.getElementById("inputMark").value;
    console.log(inputID + inputa + inputb + inputc)
    let response = await fetch('/student/create',{
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ id: parseInt(inputID), name: inputa, surname: inputb, mark: parseInt(inputc) })
      });
}

async function getMessage() {
    let response = await fetch('/student');  
    if (response.ok){
        let json = await response.json()
        let element = document.getElementById("table");
        element.innerText = ""
        console.log(json.length)
        for(var i=0;i<json.length; i++) {
            element.innerText = element.innerText + "\n" + json[i].id + "  " 
                                                         + json[i].surname + "  " 
                                                         + json[i].name + "  "  
                                                         + json[i].mark
            console.log(json[i].id + "  " +  json[i].name + "  " + json[i].mark)
        };

    } else{
        alert("Error: " + response.json())
    }  
}

async function delMessage() {
    var inputID = document.getElementById("inputID").value;
    let response = await fetch('/student/delete',{
        method: 'DELETE',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ id: parseInt(inputID)})
      });
}