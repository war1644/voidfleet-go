let Lib = {
    events:{},
    htmlLoad:(url,dom=document.body,callback=undefined)=>{
        fetch(url).then(data => data.text()).then(data => {
            dom.innerHTML = data
        }).then(callback)
    },

    // load:(url, callback=undefined)=>{
    //     fetch(url).then(response => response.json()).then(callback)
    // },
    //
    get:(url)=>{
        return fetch(url).then(response => response.json())
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
        })
            .then(response => response.json()); // parses response to JSON
    },

    getJsonLength: (json)=>{
        const keys = Object.keys(json);
        return keys.length;
    },

    eventDelegate: ()=>{
        document.body.onclick = (ev)=>{
            ev = ev || window.event;
            let target = ev.target || ev.srcElement;
            for (let k of Event.keys) {
                if (target.matches(k)){
                    Event.on(k)
                }
            }
        }
    },
};

let Event = {
    keys:[],
    events: {},
    add:(key,callback)=>{
        Event.events[key] = callback;
        Event.keys.push(key);
    },
    on:(key)=>{
        return Event.events[key]()
    },
};
