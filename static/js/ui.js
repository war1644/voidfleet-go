const UI = {
    div: null,
    isShow: false,
    mapChange:(e)=>{
        console.log(e.target.dataset.menu);

    },
    menuChange:(e)=>{

    },





    msgBoxInit: () => {

    },

    showMsg: (text) => {

    },
    init: () => {
        Event.add('.map_change',UI.mapChange);
        Event.add('.menu_change',UI.menuChange);

    },
    show: () => {
        UI.div.style.display = "block";
        UI.isShow = true;
    },
    hidde: () => {
        UI.div.style.display = "none";
        UI.isShow = false;
    },
    btnBox: () => {
        let btns = [
            `<button style="font-size:40px'" onclick="UI.event({type:'退出'})">退出</button>`,
            `<button style="font-size:40px'" onclick="UI.event({type:'舰队'})">舰队</button>`,
            `<button style="font-size:40px'" onclick="UI.event({type:'货舱'})">货舱</button>`,
            `<button style="font-size:40px'" onclick="UI.event({type:'系统'})">系统</button>`,
        ];
        UI.footer.innerHTML = btns.join('');
    },
    event: ({ type }) => {
        console.log(type);
        switch (type) {
            case "菜单":
                UI.menusBox();
                UI.show();
                break;
            case "退出":
                UI.hidde();
                break;
            case "舰队":
                UI.fleetBox();
                break;
            case "货舱":
                UI.cargoBox();
                break;
            case "系统":
                UI.systemBox();
                break;
            case "商店":
                UI.goodsBox();
                break;
            case "造船厂":
                UI.shipsBox();
                break;
            case "港口":
                UI.dockBox();
                break;
            case "防卫基地":
                UI.defendCenterBox();
                break;
            case "酒吧":
                UI.barBox();
                break;
            default:
                break;
        }
    },
    menusBox: () => {
        UI.header.innerHTML = `${Game.menuTitle}`;
        const list = Game.planetMenuList;
        const cardList = [];
        list.forEach((v) => {
            const card = `<div class="card" onclick="UI.event({type:${v.name}})">
                <h2>${v.name}</h2>
                <div class="banner"></div>
                <p>${v.descript}</p>
            </div>`;
            cardList.push(card);
        });
        UI.content.innerHTML = cardList.join('');
        UI.btnBox();
    },
    listBox: ({title,list}) => {
        UI.header.innerHTML = title;
        const cardList = [];
        list.forEach((v,i) => {
            const card = `<div class="card" onclick="UI.goodsEvent({type:'buy',name:${v.name},index:${i}})">
                <h2>${v.name} $${v.price} *${v.count}</h2>
                <div class="banner"></div>
                <p>${v.descript}</p>
            </div>`;
            cardList.push(card);
        });
        UI.content.innerHTML = cardList.join('');
        UI.btnBox();
    },
    goodsBox:()=>{
        listBox({title:'商店',list:Game.shopList});
    },
    goodsEvent:({type,index})=>{
        console.log(`购买了${index},`)

    },


}