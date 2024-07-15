"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[14681],{3905:(e,r,n)=>{n.d(r,{Zo:()=>p,kt:()=>f});var t=n(67294);function i(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function o(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);r&&(t=t.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,t)}return n}function s(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?o(Object(n),!0).forEach((function(r){i(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function l(e,r){if(null==e)return{};var n,t,i=function(e,r){if(null==e)return{};var n,t,i={},o=Object.keys(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||(i[n]=e[n]);return i}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var a=t.createContext({}),c=function(e){var r=t.useContext(a),n=r;return e&&(n="function"==typeof e?e(r):s(s({},r),e)),n},p=function(e){var r=c(e.components);return t.createElement(a.Provider,{value:r},e.children)},u={inlineCode:"code",wrapper:function(e){var r=e.children;return t.createElement(t.Fragment,{},r)}},d=t.forwardRef((function(e,r){var n=e.components,i=e.mdxType,o=e.originalType,a=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),d=c(n),f=i,m=d["".concat(a,".").concat(f)]||d[f]||u[f]||o;return n?t.createElement(m,s(s({ref:r},p),{},{components:n})):t.createElement(m,s({ref:r},p))}));function f(e,r){var n=arguments,i=r&&r.mdxType;if("string"==typeof e||i){var o=n.length,s=new Array(o);s[0]=d;var l={};for(var a in r)hasOwnProperty.call(r,a)&&(l[a]=r[a]);l.originalType=e,l.mdxType="string"==typeof e?e:i,s[1]=l;for(var c=2;c<o;c++)s[c]=n[c];return t.createElement.apply(null,s)}return t.createElement.apply(null,n)}d.displayName="MDXCreateElement"},35696:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>c,contentTitle:()=>l,default:()=>d,frontMatter:()=>s,metadata:()=>a,toc:()=>p});n(67294);var t=n(3905);function i(){return i=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var n=arguments[r];for(var t in n)Object.prototype.hasOwnProperty.call(n,t)&&(e[t]=n[t])}return e},i.apply(this,arguments)}function o(e,r){if(null==e)return{};var n,t,i=function(e,r){if(null==e)return{};var n,t,i={},o=Object.keys(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||(i[n]=e[n]);return i}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}const s={id:"chain-links",title:"Query chain links",sidebar_label:"Chain links",slug:"chain-links"},l=void 0,a={unversionedId:"developers/queries/profiles/chain-links",id:"version-2.3/developers/queries/profiles/chain-links",title:"Query chain links",description:"Query chain links",source:"@site/versioned_docs/version-2.3/02-developers/04-queries/profiles/chain-links.md",sourceDirName:"02-developers/04-queries/profiles",slug:"/developers/queries/profiles/chain-links",permalink:"/2.3/developers/queries/profiles/chain-links",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-2.3/02-developers/04-queries/profiles/chain-links.md",tags:[],version:"2.3",frontMatter:{id:"chain-links",title:"Query chain links",sidebar_label:"Chain links",slug:"chain-links"},sidebar:"version-2.3/docs",previous:{title:"Blocked users",permalink:"/2.3/developers/queries/profiles/blocked-users"},next:{title:"Application links",permalink:"/2.3/developers/queries/profiles/application-link"}},c={},p=[{value:"Query chain links",id:"query-chain-links",level:2}],u={toc:p};function d(e){var{components:r}=e,n=o(e,["components"]);return(0,t.kt)("wrapper",i({},u,n,{components:r,mdxType:"MDXLayout"}),(0,t.kt)("h2",i({},{id:"query-chain-links"}),"Query chain links"),(0,t.kt)("p",null,"This query allows you to retrieve the chain links with the optional user ",(0,t.kt)("inlineCode",{parentName:"p"},"address"),"."),(0,t.kt)("p",null,(0,t.kt)("strong",{parentName:"p"},"CLI")),(0,t.kt)("pre",null,(0,t.kt)("code",i({parentName:"pre"},{className:"language-bash"}),"desmos query profiles chain-links [[address]]\n\n# Example\n# desmos query chain-links\n# desmos query chain-links desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud\n")))}d.isMDXComponent=!0}}]);