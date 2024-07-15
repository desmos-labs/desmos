"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[7386],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>b});var r=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,l=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),u=c(n),b=o,m=u["".concat(l,".").concat(b)]||u[b]||d[b]||a;return n?r.createElement(m,s(s({ref:t},p),{},{components:n})):r.createElement(m,s({ref:t},p))}));function b(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,s=new Array(a);s[0]=u;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:o,s[1]=i;for(var c=2;c<a;c++)s[c]=n[c];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},13958:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>i,default:()=>u,frontMatter:()=>s,metadata:()=>l,toc:()=>p});n(67294);var r=n(3905);function o(){return o=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},o.apply(this,arguments)}function a(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}const s={id:"observe-data",title:"Observing data",sidebar_label:"Observing data",slug:"observe-data"},i="Observing new data",l={unversionedId:"developers/observe-data",id:"developers/observe-data",title:"Observing data",description:"Introduction",source:"@site/docs/02-developers/05-observe-data.md",sourceDirName:"02-developers",slug:"/developers/observe-data",permalink:"/developers/observe-data",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/02-developers/05-observe-data.md",tags:[],version:"current",lastUpdatedAt:1669690424,formattedLastUpdatedAt:"Nov 29, 2022",sidebarPosition:5,frontMatter:{id:"observe-data",title:"Observing data",sidebar_label:"Observing data",slug:"observe-data"},sidebar:"docs",previous:{title:"Query data",permalink:"/developers/query-data"},next:{title:"F.A.Q",permalink:"/developers/faq"}},c={},p=[{value:"Introduction",id:"introduction",level:2},{value:"Websocket",id:"websocket",level:2},{value:"Events",id:"events",level:3},{value:"Example",id:"example",level:4}],d={toc:p};function u(e){var{components:t}=e,n=a(e,["components"]);return(0,r.kt)("wrapper",o({},d,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",o({},{id:"observing-new-data"}),"Observing new data"),(0,r.kt)("h2",o({},{id:"introduction"}),"Introduction"),(0,r.kt)("p",null,"Aside from querying data, you can also observe new data as its inserted inside the chain itself. In this way, you will be notified as soon as a transaction is properly executed without having to constantly polling the chain state by yourself. "),(0,r.kt)("h2",o({},{id:"websocket"}),"Websocket"),(0,r.kt)("p",null,"All the live data observation is done though the usage of a ",(0,r.kt)("a",o({parentName:"p"},{href:"https://en.wikipedia.org/wiki/WebSocket"}),"websocket"),". The endpoint of such websocket is the following: "),(0,r.kt)("pre",null,(0,r.kt)("code",o({parentName:"pre"},{}),"ws://lcd-endpoint/websocket\n\n# Example\n# ws://morpheus.desmos.network/websocket\n")),(0,r.kt)("h3",o({},{id:"events"}),"Events"),(0,r.kt)("p",null,"In order to subscribe to specific events, you will need to send one or more messages to the websocket once you opened a connection to it. Such messages need to contain the following JSON object and must be string encoded: "),(0,r.kt)("pre",null,(0,r.kt)("code",o({parentName:"pre"},{className:"language-json"}),'{\n  "jsonrpc": "2.0",\n  "method": "subscribe",\n  "id": "0",\n  "params": {\n    "query": "tm.event=\'eventCategory\' AND eventType.eventAttribute=\'attributeValue\'"\n  }\n}\n')),(0,r.kt)("p",null,"The ",(0,r.kt)("inlineCode",{parentName:"p"},"query")," field can have the following values: "),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"tm.event='NewBlock'")," if you want to observe each new block that is created (even empty ones);"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"tm.event='Tx'")," if you want to subscribe to all new transactions;"),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("inlineCode",{parentName:"li"},"message.action='<action>'")," if you want to subscribe to events emitted when a specific message is sent to the chain.\nIn this case, please refer to the ",(0,r.kt)("inlineCode",{parentName:"li"},"Message action")," section on each transaction message\nspecification page to know what is the type associated to each message.")),(0,r.kt)("p",null,"Please note that if you want to subscribe to multiple events you will need to send multiple query messages upon connecting to the websocket. "),(0,r.kt)("h4",o({},{id:"example"}),"Example"),(0,r.kt)("pre",null,(0,r.kt)("code",o({parentName:"pre"},{className:"language-json"}),'{\n  "jsonrpc": "2.0",\n  "method": "subscribe",\n  "id": "0",\n  "params": {\n    "query": "message.action=\'save_profile\'"\n  }\n}\n')))}u.isMDXComponent=!0}}]);