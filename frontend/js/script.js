var x = 0;
var text = [
    {bus_id: 1, position: 3, congestion: 4, direction: 0},
    {bus_id: 2, position: 4, congestion: 1, direction: 1}
];
var bus_stop_name = ["正門停留所", "体育館停留所", "エッグドーム停留所", "中央停留所", "経済学部棟停留所", "スポーツ健康学部棟"]
var cvs = document.getElementById("cv");
var ctx = cvs.getContext("2d");

function loadContents(){
    var p = document.getElementById('tmp');
    //p.innerHTML = x++;
    $.ajax({
        type: "get",
        url: "/api/bus",
        data: {data: 'data'},
        dataType: "json",
        success: function(data) {
            //p.innerHTML = x + ": succeeded.";
            text = JSON.stringify(data);
        },
        error: function(){
            //p.innerHTML = x + ": failed.";
        }
    })
};

function setCongetionColor(congestion){
    //alert(text[id].congestion == 5);
    switch(congestion) {
        case 1:
            ctx.fillStyle = '#00ff00';
            break;
        case 2:
            ctx.fillStyle = '#80ff00';
            break;
        case 3:
            ctx.fillStyle = '#ffff00';
            break;
        case 4:
            ctx.fillStyle = '#ff8000';
            break;
        case 5:
            ctx.fillStyle = '#ff0000';
            break;
        default:
            ctx.fillStyle = '#ffffff';
    }
}

function drawBackGround(){
    ctx.lineWidth = 5;
    ctx.beginPath();
    ctx.moveTo(199,1);
    ctx.lineTo(199,401);
    ctx.stroke();
    ctx.beginPath();
    ctx.moveTo(250,1);
    ctx.lineTo(250,401);
    ctx.stroke();

    ctx.strokeStyle = 'black';
    ctx.fillStyle = 'white';
    ctx.lineWidth = 2;
    for (var i = 0; i < 6; i++) {
        ctx.moveTo(199,25 + 70*i + 1);
        ctx.arc(199, 25 + 70*i + 1, 25, 0, Math.PI*2, true);
        ctx.stroke();
        ctx.fill();
        ctx.moveTo(250,25 + 70*i + 1);
        ctx.arc(250, 25 + 70*i + 1, 25, 0, Math.PI*2, true);
        ctx.stroke();
        ctx.fill();
    }
    ctx.lineWidth = 1;
    ctx.strokeRect(1,1,275,401);

    ctx.fillStyle = "black";
    ctx.font = "bold 15px 'ＭＳ Ｐゴシック'";
    ctx.textAlign = "center";
    for(var i =0; i < 6;i++){
        ctx.fillText(bus_stop_name[i], 100, 30 + 70 * i);
    }
}

function drawBus(waku, id){
    var p = document.getElementById('tmp');
    var matchData = text.filter(function(item, index){
        if (item.bus_id == id) return true;});
    setCongetionColor(matchData[0].congestion);
    p.innerHTML = matchData[0].bus_id;
    if(matchData[0].direction == 0){
        ctx.beginPath();
        ctx.moveTo(200 + 50 * waku, (matchData[0].position-1)*70 + 1);
        ctx.lineTo(180 + 50 * waku, (matchData[0].position-1)*70+40 + 1);
        ctx.lineTo(220 + 50 * waku, (matchData[0].position-1)*70+40 + 1);
        ctx.closePath();
    }else if(matchData[0].direction != 0){
        ctx.beginPath();
        ctx.moveTo(200 + 50 * waku, (matchData[0].position-1)*70+50,50 + 1);
        ctx.lineTo(220 + 50 * waku, (matchData[0].position-1)*70+10,70 + 1);
        ctx.lineTo(180 + 50 * waku, (matchData[0].position-1)*70+10,30 + 1);
        ctx.closePath();
    }
    ctx.fill();
    ctx.stroke();
}

function drawImage(){
    ctx.clearRect(0, 0, 400, 400);
    drawBackGround();
    drawBus(0, 1);
    drawBus(1, 2);
    //ctx.arc(25 + (text[0].position-1)*100, 50, 25, 0, Math.PI*2, true);
    //ctx.fill();
}

function action(){
    loadContents();
    drawImage();
}

$(function() {
    action();    // 初期表示時の初回呼び出し
    setInterval("action()", 10000); // 10 秒間隔で実行
});
