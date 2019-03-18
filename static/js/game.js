window.onload = ()=>{
    Game.init();
    let params = Lib.getUrlParams("t");
    if (!params) params = 'start';
    Lib.loadHtml(`html/${params}.html`);
};


let Game = {
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
        Lib.loadHtml('html/home.html');
        // Lib.loadHtml('html/fight.html',document.body,()=>{
        //     imgDom = document.querySelector(".fight-screen");
        //     alert(JSON.stringify(imgDom));
        // });
        // let loop = ()=>{
            Lib.getJson('/frame',Game.update);
        // };
        // Lib.loop(loop);
    },
    update:(json)=>{
        console.log(JSON.stringify(json))
    },
    keyListen:()=>{
        document.body.onkeydown = (event)=>{
            // if ( event.key === "enter" ) {
            //     event.preventDefault();
            // }
            Lib.get('/key?event='+event.key);
        };
    },
    bindEvent:()=>{
        Event.add("#save1",Game.start);
        Event.add("#save2",Game.start);

    },
    menuEvent:()=>{
        //
        Event.add("#save2",Game.start);
    },
    domList:()=>{
        playerStatusDom = document.querySelector(".player_status");
        gameMessageDom = document.querySelector(".game_message");
        starMapDom = document.querySelector(".star_map");
        galaxyMapDom = document.querySelector(".galaxy_map");
        factionMapDom = document.querySelector(".faction_map");
        menuListDom = document.querySelector(".menu_list");
        gameItemListDom = document.querySelector(".game_item_list");
        // <p class="small">旗舰：<span>帝国战机</span></p>
        // <p class="small">装甲：<span>100</span></p>
        // <p class="small">护盾：<span>1000</span></p>
        // <p class="small">声望：<span>0</span></p>
        // <p class="small">僚机：<span>0</span></p>
        // <p class="small">金钱：<span>10000</span></p>
        // <p class="small">能量电池：<span>0</span></p>
    }
};