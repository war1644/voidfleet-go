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
        console.log(JSON.stringify(json));
        Game.playerStatusDomPrecess(json.Player);
        Game.gameMessageDomProcess();
        Game.galaxyMapDomProcess(json.Galaxy);
        Game.starMapDomProcess(json.Galaxy.Current);
    },
    keyListen:()=>{
        document.body.onkeydown = (event)=>{
            // if ( event.key === "enter" ) {
            //     event.preventDefault();
            // }
            Lib.get('/event?event='+event.key);
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
        gameItemListDom = document.querySelector(".game_item_list");

        //DOM载入结束后，获取数据
        Lib.getJson('/frame',Game.update);
    },
    playerStatusDomPrecess:(data)=>{
        // for (let k in data) {
        //     data[k]
        // }
        // <p class="small">旗舰：<span>帝国战机</span></p>
        // <p class="small">装甲：<span>100</span></p>
        // <p class="small">护盾：<span>1000</span></p>
        // <p class="small">声望：<span>0</span></p>
        // <p class="small">僚机：<span>0</span></p>
        // <p class="small">金钱：<span>10000</span></p>
        // <p class="small">能量电池：<span>0</span></p>
    },
    gameMessageDomProcess:(data)=>{

    // <p class="x-small"><span class="badge badge-info badge-pill">信息</span>正在停靠空间站</p>
    //     <p class="x-small"><span class="badge badge-info badge-pill">信息</span>到达天狼星区殖民星球1</p>
    //     <p class="x-small"><span class="badge badge-primary badge-pill">新闻</span>海盗正在攻击天狼星区巡逻队</p>
    //     <p class="x-small"><span class="badge badge-success badge-pill">任务</span>运送能量电池完成</p>
    },
    starMapDomProcess:(data)=>{

        /*
        * <span class="xx-small badge badge-info move-dom id_gate" id="">跳跃门</span>
                        <span class="xx-small badge badge-info move-dom id_star1" id="">殖民星球1</span>
                        <span class="xx-small badge badge-info move-dom id_star2" id="">殖民星球2</span>
                        <span class="xx-small badge badge-info move-dom id_star3" id="">殖民星球3</span>
                        <span class="xx-small badge badge-info move-dom id_station" id="">空间站A</span>
                        <span class="xx-small badge badge-info move-dom id_pirate_base" id="">海盗基地</span>
                        <span class="xx-small badge badge-info move-dom id_outpost" id="">军事前哨</span>
                        <span class="xx-small badge badge-warning move-dom id_main_fleet" id="">主力舰队</span>
                        <span class="xx-small badge badge-warning move-dom id_outpost_fleet" id="">巡逻舰队</span>
        * */
    },
    galaxyMapDomProcess:(data)=>{

        /*<span class="x-small badge badge-info move-dom id_tiannang" id="">
                            <b>天狼星区</b>
                        </span>
                            <span class="x-small badge badge-info move-dom id_renma" id="">
                            <b>人马座</b>
                        </span>
                            <span class="x-small badge badge-info move-dom id_taiyang" id="">
                            <b>太阳系</b>
                        </span>
                            <span class="x-small badge badge-info move-dom id_beiluo" id="">
                            <b>北落师门</b>
                        </span>
                            <span class="x-small badge badge-info move-dom id_beiji" id="">
                            <b>北极星区</b>
                        </span>
                            <span class="x-small badge badge-info move-dom id_x" id="">
                            <b>X星区</b>
                        </span>*/
    }
};

let Player = {

};

let Galaxy = {

};

let Planet = {

};