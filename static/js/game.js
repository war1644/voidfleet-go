window.onload = ()=>{
    Lib.eventDelegate();
    let params = Lib.getUrlParams("t");
    if (!params) params = 'start';
    Lib.loadHtml(`html/${params}.html`);
    Game.bindEvent();
    Game.keyListen();
    // str2.replace(/\[.*?\]/g, function (value) {
    //     let key = value.substring(1, value.length - 1);
    //     return emojiMap[key];
    // });

};


let Game = {
    start:()=>{
        let imgDom;
        // Lib.loadHtml('html/fight.html',document.body,()=>{
        //     imgDom = document.querySelector(".fight-screen");
        //     alert(JSON.stringify(imgDom));
        // });
        let loop = ()=>{
            // Lib.get('/frame',(data)=>{
            //     imgDom.setAttribute('src', data);
            // });
            Lib.loadHtml('/frame');
        };
        Lib.loopExec(loop);
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
        Event.add("#save2",Game.start);
        Event.add("#save2",Game.start);
        Event.add("#save2",Game.start);
    }
};