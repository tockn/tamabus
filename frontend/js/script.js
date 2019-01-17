var x = 0;
var text = [
  {"bus_id": 1, "position": 3, "congestion": 5, "direction": 0},
  {"bus_id": 2, "position": 5, "congestion": 1, "direction": 1}
];
var cvs = document.getElementById("cv");
var ctx = cvs.getContext("2d");

var t = document.getElementById('tmp');
//t.innerHTML = text[0].bus_id;

function loadContents(){
  var p = document.getElementById('tmp');
  p.innerHTML = x++;
  $.ajax({
    type: "get",
    url: "/api/bus",
    data: {data: 'data'},
    dataType: "json",
    success: function(data) {
      p.innerHTML = x + ": succeeded.";
      text = JSON.stringify(data);
  },
    error: function(){
      p.innerHTML = x + ": failed.";
    }
  })
};

function setCongetionColor(id){
  //alert(text[id].congestion == 5);
  switch(text[id].congestion) {
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
  ctx.moveTo(50,0);
  ctx.lineTo(50,400);
  ctx.stroke();
  ctx.strokeRect(1,1,99,399);
  ctx.strokeStyle = 'black';
  ctx.fillStyle = 'white';
  ctx.lineWidth = 2;
  for (var i = 0; i < 6; i++) {
    ctx.moveTo(50,25 + 70*i);
    ctx.arc(50, 25 + 70*i, 25, 0, Math.PI*2, true);
    ctx.stroke();
    ctx.fill();
  }
}

function drawBus(direction, id){
  setCongetionColor(id);
  if(text[id].direction == 0){
    ctx.beginPath();
    ctx.moveTo(50, (text[id].position-1)*70);
    ctx.lineTo(30, (text[id].position-1)*70+40);
    ctx.lineTo(70, (text[id].position-1)*70+40);
    ctx.closePath();
  }else if(text[id].direction == 1){
    ctx.beginPath();
    ctx.moveTo(50, (text[id].position-1)*70+50,50);
    ctx.lineTo(30, (text[id].position-1)*70+10,30);
    ctx.lineTo(70, (text[id].position-1)*70+10,70);
    ctx.closePath();
  }
    ctx.fill();
    ctx.stroke();
}

function drawImage(){
  ctx.clearRect(0, 0, 400, 400);
  drawBackGround();
  drawBus(0, 0);
  drawBus(0, 1);
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
