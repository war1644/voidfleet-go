window.onload = ()=>{
    Lib.eventDelegate();
    Lib.htmlLoad('html/start.html')
    Game.bindEvent()
};


let Game = {
    start:()=>{
        Lib.htmlLoad('html/home.html')
        let imgDom = document.querySelector("#image");
        let loop = ()=>{
            Lib.curl('/frame', (data)=>{
                imgDom.setAttribute('src', data);
            });
        };

        document.body.onkeydown = ()=>{
            if ( event.keyCode === 13 ) {
                event.preventDefault();
            }
            Lib.curl('/key?event='+event.which);
        };
        requestAnimationFrame(loop);
    },
    bindEvent:()=>{
        Event.add("#save1",Game.start);
        Event.add("#save2",Game.start);
    },
};