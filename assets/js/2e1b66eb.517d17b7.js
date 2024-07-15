"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[78812],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>m});var r=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},v=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,s=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),v=c(n),m=o,f=v["".concat(s,".").concat(m)]||v[m]||p[m]||a;return n?r.createElement(f,i(i({ref:t},u),{},{components:n})):r.createElement(f,i({ref:t},u))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,i=new Array(a);i[0]=v;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:o,i[1]=l;for(var c=2;c<a;c++)i[c]=n[c];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}v.displayName="MDXCreateElement"},48394:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>v,frontMatter:()=>i,metadata:()=>s,toc:()=>u});n(67294);var r=n(3905);function o(){return o=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},o.apply(this,arguments)}function a(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}const i={id:"overview",title:"Overview",sidebar_label:"Overview",slug:"overview"},l="Testnets Overview",s={unversionedId:"testnet/overview",id:"version-4.2.0/testnet/overview",title:"Overview",description:"Testnets (from the words test- and nets-, networks) are the way we at Desmos use to test all the features our blockchain before launching them publicly.",source:"@site/versioned_docs/version-4.2.0/05-testnet/01-overview.md",sourceDirName:"05-testnet",slug:"/testnet/overview",permalink:"/4.2.0/testnet/overview",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/05-testnet/01-overview.md",tags:[],version:"4.2.0",sidebarPosition:1,frontMatter:{id:"overview",title:"Overview",sidebar_label:"Overview",slug:"overview"},sidebar:"docs",previous:{title:"Validator FAQ",permalink:"/4.2.0/validators/validator-faq"},next:{title:"create-local",permalink:"/4.2.0/testnet/create-local"}},c={},u=[{value:"Public testnet",id:"public-testnet",level:2},{value:"Local testnet",id:"local-testnet",level:2}],p={toc:u};function v(e){var{components:t}=e,n=a(e,["components"]);return(0,r.kt)("wrapper",o({},p,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",o({},{id:"testnets-overview"}),"Testnets Overview"),(0,r.kt)("p",null,"Testnets (from the words ",(0,r.kt)("em",{parentName:"p"},"test-")," and ",(0,r.kt)("em",{parentName:"p"},"nets-"),", networks) are the way we at Desmos use to test all the features our blockchain before launching them publicly. "),(0,r.kt)("p",null,"In other words, a testnet is the playground that you can use to start learning about Desmos, its features and how you can use them to create your own decentralized social networks or social enabled app. "),(0,r.kt)("p",null,"There are two different types of testnets: "),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Public testnets"),(0,r.kt)("li",{parentName:"ul"},"Local testnets")),(0,r.kt)("h2",o({},{id:"public-testnet"}),"Public testnet"),(0,r.kt)("p",null,"A Public testnet is a preview of what the Desmos mainnet will actually be. "),(0,r.kt)("p",null,"Testnet's ",(0,r.kt)("a",o({parentName:"p"},{href:"/4.2.0/validators/overview"}),"validators")," are publicly known and every developer can write and read transactions from them. "),(0,r.kt)("p",null,"Public testnets are the battlefields on which you can test the integration of your app (or your validator's setup) without worrying too much about security, but being sure to always be up-to-date with the latest stable changes."),(0,r.kt)("admonition",o({},{title:"Joining the public testnet  ",type:"tip"}),(0,r.kt)("p",{parentName:"admonition"},"If you want to know more about how to join the currently running public testnet, please refer to the ",(0,r.kt)("a",o({parentName:"p"},{href:"/4.2.0/testnet/join-public/setup"}),(0,r.kt)("em",{parentName:"a"},"Join the public testnet"))," page.  ")),(0,r.kt)("h2",o({},{id:"local-testnet"}),"Local testnet"),(0,r.kt)("p",null,"A Local testnet exist only on the machine that is used to running it. This means that none, except you, can actually access the data you store. "),(0,r.kt)("p",null,"Local testnets are perfect if you want to quickly setup a Desmos blockchain instance without worrying too much about setting up a public full-node machine. "),(0,r.kt)("p",null,"They also might be particularly useful to developers that do not want to write on the public net but want first to try out their app's integration locally to make sure everything works properly. "),(0,r.kt)("admonition",o({},{title:"Creating a local testnet  ",type:"tip"}),(0,r.kt)("p",{parentName:"admonition"},"If you want to know more about how creating a local testnet, please refer to the ",(0,r.kt)("a",o({parentName:"p"},{href:"/4.2.0/testnet/create-local"}),(0,r.kt)("em",{parentName:"a"},"Create a local testnet"))," page.  ")))}v.isMDXComponent=!0}}]);