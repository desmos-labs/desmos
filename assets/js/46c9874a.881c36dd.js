"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[84966],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>d});var r=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var a=r.createContext({}),c=function(e){var t=r.useContext(a),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(a.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},f=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,a=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),f=c(n),d=o,m=f["".concat(a,".").concat(d)]||f[d]||u[d]||i;return n?r.createElement(m,s(s({ref:t},p),{},{components:n})):r.createElement(m,s({ref:t},p))}));function d(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,s=new Array(i);s[0]=f;var l={};for(var a in t)hasOwnProperty.call(t,a)&&(l[a]=t[a]);l.originalType=e,l.mdxType="string"==typeof e?e:o,s[1]=l;for(var c=2;c<i;c++)s[c]=n[c];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}f.displayName="MDXCreateElement"},91821:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>f,frontMatter:()=>s,metadata:()=>a,toc:()=>p});n(67294);var r=n(3905);function o(){return o=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},o.apply(this,arguments)}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}const s={id:"genesis-file",title:"Genesis File",sidebar_label:"Genesis File",slug:"genesis-file"},l="Genesis file",a={unversionedId:"testnet/join-public/genesis-file",id:"version-4.2.0/testnet/join-public/genesis-file",title:"Genesis File",description:"To configure a full node for the testnet you need to use the following seed nodes. If you are looking for mainnet seed nodes, please refer to this instead.",source:"@site/versioned_docs/version-4.2.0/05-testnet/03-join-public/02-genesis-file.md",sourceDirName:"05-testnet/03-join-public",slug:"/testnet/join-public/genesis-file",permalink:"/4.2.0/testnet/join-public/genesis-file",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/05-testnet/03-join-public/02-genesis-file.md",tags:[],version:"4.2.0",sidebarPosition:2,frontMatter:{id:"genesis-file",title:"Genesis File",sidebar_label:"Genesis File",slug:"genesis-file"},sidebar:"docs",previous:{title:"Setup",permalink:"/4.2.0/testnet/join-public/setup"},next:{title:"Seeds",permalink:"/4.2.0/testnet/join-public/seeds"}},c={},p=[],u={toc:p};function f(e){var{components:t}=e,n=i(e,["components"]);return(0,r.kt)("wrapper",o({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",o({},{id:"genesis-file"}),"Genesis file"),(0,r.kt)("admonition",o({},{title:"Testnet only   ",type:"caution"}),(0,r.kt)("p",{parentName:"admonition"},"To configure a full node for the ",(0,r.kt)("strong",{parentName:"p"},"testnet")," you need to use the following seed nodes. If you are looking for mainnet seed nodes, please refer to ",(0,r.kt)("a",o({parentName:"p"},{href:"/4.2.0/mainnet/genesis-file"}),"this")," instead.")),(0,r.kt)("p",null,"To connect to the ",(0,r.kt)("inlineCode",{parentName:"p"},"morpheus")," testnet, you will need the corresponding genesis file of each testnet. Visit the ",(0,r.kt)("a",o({parentName:"p"},{href:"https://github.com/desmos-labs/morpheus"}),"testnet repo")," and download the correct genesis file by running the following command."),(0,r.kt)("pre",null,(0,r.kt)("code",o({parentName:"pre"},{className:"language-bash"}),"# Download the existing genesis file for the testnet\n# Replace <chain-id> with the id of the testnet you would like to join\ncurl https://raw.githubusercontent.com/desmos-labs/morpheus/master/<chain-id>/genesis.json > $HOME/.desmos/config/genesis.json\n")))}f.isMDXComponent=!0}}]);