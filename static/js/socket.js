const Socket = {
    WS: null,
    init :()=>{
        //请求服务器握手
        Socket.WS = new WebSocket("ws://127.0.0.1:8416");
        //握手成功触发
        Socket.WS.onopen = function(){
            console.log("握手成功");
            if(Socket.WS.readyState===1){
                Socket.WS.send(name+"加入房间");
            }
        };
        //服务器发送消息，触发
        Socket.WS.onmessage = function(e){
            //e就是服务器发送的信息
            console.log( JSON.parse(e.data) );
            //接下来爱干嘛干嘛去、
            //...code...
        };
        //出错时触发方法
        Socket.WS.onerror = function(e){
            console.log("error:"+e);
        };
        Socket.WS.onclose = function (e) {
            console.log("onclose:",e);
        }
    }

};