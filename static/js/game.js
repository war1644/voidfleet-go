window.onload = ()=>{
    Game.init();
    // let params = Lib.getUrlParams("t");
    // if (!params) params = 'home';
    // Lib.loadHtml(`html/${params}.html`);
    Game.start();
};


const Game = {
    mask:false,
    showMask:()=>{
        if(Game.mask)
        {
            maskDom.style.display = 'none';
        }else{
            maskDom.style.display = 'flex';
        }
        Game.mask = !Game.mask;
    },
    mutationObserver:{},
    init:()=>{
        Lib.env();
        // if (!Lib.isChrome){
        //     //console output
        //     window.console.log = (s)=>{
        //         external.invoke('{"type":"log","data":"'+JSON.stringify(s)+'"}')
        //     };
        //     window.console.debug=(s)=>{
        //         external.invoke('{"type":"debug","data":"'+JSON.stringify(s)+'"}')
        //     };
        //     window.console.warn=(s)=>{
        //         external.invoke('{"type":"warn","data":"'+JSON.stringify(s)+'"}')
        //     };
        //     window.console.error=(s)=>{
        //         external.invoke('{"type":"error","data":"'+JSON.stringify(s)+'"}')
        //     };
        // }
        console.log("test log");
        console.log(window.navigator.userAgent);
        Lib.eventDelegate();
        Game.bindEvent();
        Game.keyListen();
        UI.init();
    },
    start:()=>{
        Lib.loadHtml('html/home.html',gameWindowDom,Game.actionDomInit);
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
        maskDom = document.querySelector(".mask");

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
        <p class="small"><span class="badge badge_${msg.MsgType} badge_pill">${msg.MsgTypePill}</span>${msg.MsgText}</p>
        `;
        }
        gameMessageDom.innerHTML += dom;
    //badge_info 信息
    //badge_primary
    //badge_success 任务
    },
    starMapEvent:()=>{
        Event.add(".move-dom",(e)=>{
            let dom = e.target;
            console.log("点击了move-dom",e.target);


        });
    },
    starMapDomProcess:(data)=>{
        let dom = "" ;
        for (let k in data) {
            if (data.hasOwnProperty(k)){
                let info = data[k];
                dom += `
        <span class="badge badge_warning move-dom" data-name="${info.Name}" style="left: ${info.X}vw;top: ${info.Y}vh;">${info.Name}</span>
        `;
            }
        }
        starMapDom.innerHTML = dom;
        Game.starMapEvent();
    },
    galaxyMapDomProcess:(data)=>{
        let dom = "",galaxy = data.NameList;
        for (let k in galaxy) {
            let info = galaxy[k];
            dom += `
        <span class="badge badge_warning move-dom" style="left: ${info[1]}vw;top: ${info[2]}vh;">
             ${info[0]}
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
                <li class="d_flex list-group-item list-group-item-action">
                    <div>
                        ${info.Name}
                        <span class="item_quantity">* ${info.Quantity}</span>
                    </div>
                    <div class="">
                        <span class="badge badge_pill badge_warning">使用</span>
                        <span class="badge badge_pill badge_warning">丢弃</span>
                    </div>
                </li>`;
        }
        cargoListDom.innerHTML = dom;
    },
    fleetListDomProcess:(data)=>{
        let dom = "",items = data;
        for (let k in items) {
            let info = items[k];
            dom += `
<li class="d_flex list-group-item list-group-item-action">
                    <div>
                        ${info.Name}
                        <span class="item_quantity">* ${info.Quantity}</span>
                        <span class="item_price">* ${info.Price}</span>
                    </div>
                    <div class="">
                        <span class="badge badge_pill badge_warning">使用</span>
                        <span class="badge badge_pill badge_warning">丢弃</span>
                    </div>
                </li>`;
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