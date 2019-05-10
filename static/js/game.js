window.onload = ()=>{
    Game.init();
    let params = Lib.getUrlParams("t");
    if (!params) params = 'start';
    Lib.loadHtml(`html/${params}.html`);
};


let Game = {
    mutationObserver:{},
    init:()=>{
        Lib.env();
        if (!Lib.isChrome){
            //console output
            window.console.log = (s)=>{
                external.invoke('{"type":"log","data":"'+JSON.stringify(s)+'"}')
            };
            window.console.debug=(s)=>{
                external.invoke('{"type":"debug","data":"'+JSON.stringify(s)+'"}')
            };
            window.console.warn=(s)=>{
                external.invoke('{"type":"warn","data":"'+JSON.stringify(s)+'"}')
            };
            window.console.error=(s)=>{
                external.invoke('{"type":"error","data":"'+JSON.stringify(s)+'"}')
            };
        }
        console.log("test log");
        console.log(window.navigator.userAgent);
        Lib.eventDelegate();
        Game.bindEvent();
        Game.keyListen();
    },
    start:()=>{
        Lib.loadHtml('html/home.html',document.body,Game.actionDomInit);
        // Lib.loadHtml('html/fight.html',document.body,()=>{
        //     imgDom = document.querySelector(".fight-screen");
        //     alert(JSON.stringify(imgDom));
        // });
        // let loop = ()=>{
        //     Lib.getJson('/frame',Game.update);
        // };
        // Lib.loop(loop);
    },
    update:(json)=>{
        console.log(json);
        Game.playerStatusDomPrecess(json.Player);
        Game.gameMessageDomProcess(json);
        Game.galaxyMapDomProcess(json.Galaxy);
        Game.starMapDomProcess(json.Galaxy.Current);
        Game.cargoListDomProcess(json.Player.Cargo);
        Game.fleetListDomProcess(json.Player.Fleet);
        // Lib.getJson(Lib.host+'/event?event=msg',Game.gameMessageDomProcess);
    },
    keyListen:()=>{
        document.body.onkeydown = (event)=>{
            // if ( event.key === "enter" ) {
            //     event.preventDefault();
            // }
            // Lib.get('/event?event='+event.key);
        };
    },
    bindEvent:()=>{
        // Game.mutationObserver = new MutationObserver((mutations)=>{
        //     mutations.forEach((mutation)=>{
        //         console.log(mutation);
        //     });
        // });
        //
        // // 开始监听页面根元素 HTML 变化。
        // Game.mutationObserver.observe(document.body, {
        //     attributes: true,
        //     characterData: true,
        //     childList: true,
        //     subtree: true,
        //     attributeOldValue: true,
        //     characterDataOldValue: true
        // });


        Event.add("#save1",Game.start);
        Event.add("#save2",Game.start);

    },
    menuEvent:()=>{
        //
        // Event.add("#save2",Game.start);
    },
    actionDomInit:()=>{
        playerStatusDom = document.querySelector(".player_status");
        gameMessageDom = document.querySelector(".game_message");
        starMapDom = document.querySelector(".star_map");
        galaxyMapDom = document.querySelector(".galaxy_map");
        factionMapDom = document.querySelector(".faction_map");
        menuListDom = document.querySelector(".menu_list");
        cargoListDom = document.querySelector(".cargo_list");
        fleetListDom = document.querySelector(".fleet_list");

        //DOM载入结束后，获取数据
        Lib.getJson(Lib.host+'/event?event=start',Game.update);
    },
    playerStatusDomPrecess:(data)=>{
        let dom = '';
        // for (let k in data) {
        dom += `
        <p class="small">旗舰：<span>${data.Ship.Name}</span></p>
        <p class="small">装甲：<span>${data.Ship.HP}</span></p>
        <p class="small">护盾：<span>${data.Ship.EP}</span></p>
        <p class="small">声望：<span>${data.Reputation}</span></p>
        <p class="small">僚机：<span>${data.ShipsCount}</span></p>
        <p class="small">金钱：<span>${data.Money}</span></p>`;
        playerStatusDom.innerHTML = dom;

    },
    gameMessageDomProcess:(data)=>{
        let dom = "",events = data.Msg;
        for (let k in events) {
            let msg = events[k];
            dom += `
        <p class="small"><span class="badge badge-${msg.MsgType} badge-pill">${msg.MsgTypePill}</span>${msg.MsgText}</p>
        `;
        }
        // let domNode = document.createTextNode(dom);
        // gameMessageDom.appendChild(domNode)
        gameMessageDom.innerHTML = dom;
    // <p class="x-small"><span class="badge badge-info badge-pill">信息</span>正在停靠空间站</p>
    //     <p class="x-small"><span class="badge badge-info badge-pill">信息</span>到达天狼星区殖民星球1</p>
    //     <p class="x-small"><span class="badge badge-primary badge-pill">新闻</span>海盗正在攻击天狼星区巡逻队</p>
    //     <p class="x-small"><span class="badge badge-success badge-pill">任务</span>运送能量电池完成</p>
    },
    starMapDomProcess:(data)=>{
        let dom = "",star = data;
        for (let k in data) {
            let info = data[k];
            dom += `
        <span class="small badge badge-warning move-dom" style="left: ${info.X}vw;top: ${info.Y}vh;">${info.Name}</span>
        `;
        }
        starMapDom.innerHTML = dom;
    },
    galaxyMapDomProcess:(data)=>{
        let dom = "",galaxy = data.NameList;
        for (let k in galaxy) {
            let info = galaxy[k];
            dom += `
        <span class="x-small badge badge-warning move-dom" style="left: ${info[1]}vw;top: ${info[2]}vh;">
             <b>${info[0]}</b>
        </span>
        `;
        }
        galaxyMapDom.innerHTML = dom;

    },

    cargoListDomProcess:(data)=>{
        let dom = "",items = data;
        for (let k in items) {
            let info = items[k];
            dom += `
        <li class="list-group-item list-group-item-action">
                        <div class="row">
                            <div class="col-10 border">
                                <p>${info.Name}</p>
                                <span>x ${info.Quantity}</span>
                            </div>
                            <div class="col-2 border align-content-center"><span class="badge badge-pill badge-warning">卖</span><span class="badge badge-pill badge-warning">扔</span></div>
                        </div>
                    </li>
        `;
        }
        cargoListDom.innerHTML = dom;
    },
    fleetListDomProcess:(data)=>{
        let dom = "",items = data;
        for (let k in items) {
            let info = items[k];
            dom += `
        <li class="list-group-item list-group-item-action">
                        <div class="row">
                            <div class="col-10 border">
                                <p>${info.Name}</p>
<!--                                <span>x ${info.Price}</span>-->
                            </div>
                        </div>
                    </li>
        `;
        }
        fleetListDom.innerHTML = dom;
    }
};

let Player = {

};

let Galaxy = {

};

let Planet = {

};