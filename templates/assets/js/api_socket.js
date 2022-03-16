    let socket;
    let connection;
    let message;
    let flagIndex;
    let flagAdvance;
    let table;
    let tableDataInit;

    function sendMessage(){
        let msg = {
            "clientId": "",
            "message": message
        }
        message = ""
        socket.send(JSON.stringify(msg));
    }

    function changeFindex(){
        flagIndex = true;
        flagAdvance = false;
    }

    function changeFadvance(){
        flagIndex = false;
        flagAdvance = true;
    }

    function showMessages(data){
        console.log(data)
        if(flagIndex && data){
            setDataNumber("dcNumber", data.dcNumber);
            setDataNumber("loadingTime", data.loadingTime);
            setDataNumber("ramConsumption", data.ramConsumption);
            setDataNumber("requestTime", data.requestTime);
            setDataNumber("mostUsedDc", data.mostUsedDc);
            setDataNumber("averageDcProperties", data.averageDcProperties);
        }
        if(flagAdvance){
            if (table){
                updateTable(data)
            }else {
                createTable(data);
            }
        }
    }

    function setDataNumber(id, number){
        let numberElement = $("#"+id);
        if(numberElement.text() != number){
            numberElement.text(number);
            numberElement.css('color', getRandomColor());
            numberElement.toggleClass('tracking-in-expand');
        }
    }

    function getRandomColor(){
        let colors = ['#dce2e7', 'rgba(229,108,188,0.84)', 'rgba(131,213,198,0.75)', 'rgba(144,239,136,0.75)'];
        return  colors[Math.floor(Math.random() * colors.length)];
    }

    function connect(){
        let hostname = location.hostname;
        if(!socket){
            console.log("Opening connection...");
            socket = new WebSocket("ws://"+hostname+":9090/api/v1/socket?clientId=random");
            socket.onopen = () => {
                console.log("Connection established!");
                connection = true;
            }
        }else {
            console.log("connection already made");
        }
        socket.onmessage = (msg)=>{
            console.log("Receiving data!");
            showMessages(JSON.parse(msg.data));
        }
        return socket;
    }

    function createTable(data){
        console.log('createTable')
        table = $('#tableData').DataTable({
            "data": data.dcList,
            "columns": [
                {"data": "dcNumber"},
                {"data": "features"},
                {"data": "averageRequestTime"},
                {"data": "lastRequest"},
                {"data": "lastUpdated"},
            ]
        });
    }

    function updateTable(data) {
        console.log('updateTable');
        data.dcList.forEach(dc => {
            let index = table.column(0).data().indexOf(dc.dcNumber);
            if (index >= 0) {
                table.row(index).data(dc).draw();
                tuti(index);
            } else {
                table.row.add(dc).draw();
            }
        })
    }

    function tuti(rowIndex){
        for (let i = 1; i < 5; i++) {
            let celEl = table.cell(rowIndex, i).node();
            celEl.style.color = getRandomColor();
            console.log(celEl);
        }
    }
