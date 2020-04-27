const Lib = {
    host:"http://127.0.0.1:1212",
    isChrome:false,
    OSX:false,
    frames:0,
    loop:(callback,fps=30)=>{
        let gameLoop = ()=>{
            if(Lib.frames % fps === 0){
                callback();
            }
            Lib.frames ++;
            window.requestAnimationFrame(gameLoop);
        };
        gameLoop();
    },

    getUrlParams:(index,search=location.href)=>{
        let hashes = search.slice(search.indexOf('?') + 1).split('&');
        let params = {};
        hashes.map(hash => {
            let [key, val] = hash.split('=');
            params[key] = decodeURIComponent(val)
        });
        return params[index]
    },
    loadHtml:(url,dom=gameWindowDom,callback=()=>{})=>{
        fetch(url).then(data => data.text()).then(data =>{
            dom.innerHTML = data;
            callback()
        })
    },

    load:(url,callback=(data)=>{})=>{
        fetch(url).then(data => data.text()).then(data => callback(data))
    },
    getJson:(url,callback)=>{
        fetch(url).then(data => data.json()).then(data => callback(data));
    },
    get:(url,callback=(data)=>{})=>{
        fetch(url).then(data => callback(data));
    },

    curl(url, data = {},type='GET') {
        return fetch(url, {
            method: type, // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
            // credentials: "same-origin", // include, *same-origin, omit
            headers: {
                "Content-Type": "application/json",
                // "Content-Type": "application/x-www-form-urlencoded",
            },
            redirect: "follow", // manual, *follow, error
            referrer: "no-referrer", // no-referrer, *client
            body: JSON.stringify(data), // body data type must match "Content-Type" header
        }).then(response => response.json()); // parses response to JSON
    },

    getJsonLength: (json)=>{
        const keys = Object.keys(json);
        return keys.length;
    },

    eventDelegate: ()=>{
        document.body.onclick = (e)=>{
            e = e || window.event;
            let target = e.target || e.srcElement;
            // console.log("target: ",target);
            for (let k of Event.keys) {
                if (target.matches(k)){
                    Event.on(k,e)
                }
            }
        }
    },
    env:()=>{
        let u = window.navigator.userAgent; // 通过navigator.userAgent获取当前浏览器的信息
        Lib.isChrome = u.indexOf('Chrome') > -1; //Chrome
        Lib.OSX = !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/);//OSX
    },
};

const Event = {
    keys:[],
    events: {},
    add:(key,callback)=>{
        Event.events[key] = callback;
        Event.keys.push(key);
    },
    on:(key,e)=>{
        return Event.events[key](e)
    },
};
