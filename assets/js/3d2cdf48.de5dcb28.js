"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[90617],{3905:(e,t,i)=>{i.d(t,{Zo:()=>c,kt:()=>m});var n=i(67294);function a(e,t,i){return t in e?Object.defineProperty(e,t,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[t]=i,e}function r(e,t){var i=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),i.push.apply(i,n)}return i}function o(e){for(var t=1;t<arguments.length;t++){var i=null!=arguments[t]?arguments[t]:{};t%2?r(Object(i),!0).forEach((function(t){a(e,t,i[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(i)):r(Object(i)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(i,t))}))}return e}function l(e,t){if(null==e)return{};var i,n,a=function(e,t){if(null==e)return{};var i,n,a={},r=Object.keys(e);for(n=0;n<r.length;n++)i=r[n],t.indexOf(i)>=0||(a[i]=e[i]);return a}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(n=0;n<r.length;n++)i=r[n],t.indexOf(i)>=0||Object.prototype.propertyIsEnumerable.call(e,i)&&(a[i]=e[i])}return a}var s=n.createContext({}),p=function(e){var t=n.useContext(s),i=t;return e&&(i="function"==typeof e?e(t):o(o({},t),e)),i},c=function(e){var t=p(e.components);return n.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var i=e.components,a=e.mdxType,r=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),u=p(i),m=a,h=u["".concat(s,".").concat(m)]||u[m]||d[m]||r;return i?n.createElement(h,o(o({ref:t},c),{},{components:i})):n.createElement(h,o({ref:t},c))}));function m(e,t){var i=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var r=i.length,o=new Array(r);o[0]=u;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var p=2;p<r;p++)o[p]=i[p];return n.createElement.apply(null,o)}return n.createElement.apply(null,i)}u.displayName="MDXCreateElement"},72266:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>p,contentTitle:()=>l,default:()=>u,frontMatter:()=>o,metadata:()=>s,toc:()=>c});i(67294);var n=i(3905);function a(){return a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var i=arguments[t];for(var n in i)Object.prototype.hasOwnProperty.call(i,n)&&(e[n]=i[n])}return e},a.apply(this,arguments)}function r(e,t){if(null==e)return{};var i,n,a=function(e,t){if(null==e)return{};var i,n,a={},r=Object.keys(e);for(n=0;n<r.length;n++)i=r[n],t.indexOf(i)>=0||(a[i]=e[i]);return a}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(n=0;n<r.length;n++)i=r[n],t.indexOf(i)>=0||Object.prototype.propertyIsEnumerable.call(e,i)&&(a[i]=e[i])}return a}const o={},l="ADR 004: Expiration of application links",s={unversionedId:"architecture/adr-004-application-links-expiration",id:"architecture/adr-004-application-links-expiration",title:"ADR 004: Expiration of application links",description:"Changelog",source:"@site/docs/architecture/adr-004-application-links-expiration.md",sourceDirName:"architecture",slug:"/architecture/adr-004-application-links-expiration",permalink:"/architecture/adr-004-application-links-expiration",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/architecture/adr-004-application-links-expiration.md",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"ADR 003: Remove custom JSON tags from Proto files",permalink:"/architecture/adr-003-remove-custom-json-tags-from-profo-files"},next:{title:"ADR 006: Support multisig chain link",permalink:"/architecture/adr-005-support-multisig-chain-link"}},p={},c=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Abstract",id:"abstract",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Backwards Compatibility",id:"backwards-compatibility",level:3},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"Test Cases optional",id:"test-cases-optional",level:2},{value:"References",id:"references",level:2}],d={toc:c};function u(e){var{components:t}=e,i=r(e,["components"]);return(0,n.kt)("wrapper",a({},d,i,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",a({},{id:"adr-004-expiration-of-application-links"}),"ADR 004: Expiration of application links"),(0,n.kt)("h2",a({},{id:"changelog"}),"Changelog"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"September 20th, 2021: Initial draft;"),(0,n.kt)("li",{parentName:"ul"},"September 21th, 2021: Moved from DRAFT to PROPOSED;"),(0,n.kt)("li",{parentName:"ul"},"December  22th, 2021: First review;"),(0,n.kt)("li",{parentName:"ul"},"January   04th, 2022: Second review;"),(0,n.kt)("li",{parentName:"ul"},"January   10th, 2022: Third review;"),(0,n.kt)("li",{parentName:"ul"},"January   12th, 2022: Fourth review.")),(0,n.kt)("h2",a({},{id:"status"}),"Status"),(0,n.kt)("p",null,"ACCEPTED Implemented"),(0,n.kt)("h2",a({},{id:"abstract"}),"Abstract"),(0,n.kt)("p",null,"Currently, when a user links a centralized application with their Desmos profile, the created link contains a timestamp of when it has been created.",(0,n.kt)("br",{parentName:"p"}),"\n",'Since centralized applications\' username can be switched and sold, we SHOULD implement an "expiration date" system on links.\nThis means that after a certain amount of time passes, the link will be automatically marked deleted and the user has to connect it again in order to keep it valid.'),(0,n.kt)("h2",a({},{id:"context"}),"Context"),(0,n.kt)("p",null,"Desmos ",(0,n.kt)("inlineCode",{parentName:"p"},"x/profiles")," module gives users the possibility to link their Desmos profile with both centralized application and\nother blockchains accounts. By doing this, a user can be verified as the owner of those accounts and prove to the system\nthat they're not impersonating anyone else. This verification however remains valid only if the user\nnever trades or sells their centralized-app username to someone else. If they do, the link to such username MUST be invalidated.\nUnfortunately for us, it's not possible to understand when this happens since it's off-chain.\nTo prevent this situation, an \"expiration time\" SHOULD be added to the ",(0,n.kt)("inlineCode",{parentName:"p"},"ApplicationLink")," object."),(0,n.kt)("h2",a({},{id:"decision"}),"Decision"),(0,n.kt)("p",null,"To implement the link expiration we will act as follow:"),(0,n.kt)("ol",null,(0,n.kt)("li",{parentName:"ol"},"extend the ",(0,n.kt)("inlineCode",{parentName:"li"},"ApplicationLink")," structure by adding an ",(0,n.kt)("inlineCode",{parentName:"li"},"ExpiratonTime")," field that represents the time when the link will expire and will be deleted from the store;"),(0,n.kt)("li",{parentName:"ol"},"save a reference of the expiring link inside the store using a prefix and a ",(0,n.kt)("inlineCode",{parentName:"li"},"time.Time")," value which will make it easy to iterate over the expired links;"),(0,n.kt)("li",{parentName:"ol"},"add a new keeper method that allows to iterate over the expired links;"),(0,n.kt)("li",{parentName:"ol"},"inside the ",(0,n.kt)("inlineCode",{parentName:"li"},"BeginBlock")," method, iterate over all the expired links and delete them from the store.")),(0,n.kt)("p",null,"We will also need to introduce a new ",(0,n.kt)("inlineCode",{parentName:"p"},"x/profiles")," on chain parameter named ",(0,n.kt)("inlineCode",{parentName:"p"},"AppLinkParams")," which contains\nthe default validity duration of all the app links. The parameter will be later used inside the ",(0,n.kt)("inlineCode",{parentName:"p"},"StartProfileConnection"),"\nto calculate the estimated expiration time of each ",(0,n.kt)("inlineCode",{parentName:"p"},"AppLink"),"."),(0,n.kt)("h2",a({},{id:"consequences"}),"Consequences"),(0,n.kt)("h3",a({},{id:"backwards-compatibility"}),"Backwards Compatibility"),(0,n.kt)("p",null,"This update will affect the ",(0,n.kt)("inlineCode",{parentName:"p"},"ApplicationLink")," object by adding the new ",(0,n.kt)("inlineCode",{parentName:"p"},"ExpirationTime"),"\nfield and breaking the compatibility with the previous versions of the software. To allow\na smooth update and overcome these compatibility issues, we need to set up a proper migration\nfrom the previous versions to the one that will include the additions contained in this ADR."),(0,n.kt)("h3",a({},{id:"positive"}),"Positive"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Considerably reduce the possibility of impersonation of entities and users of centralized apps;")),(0,n.kt)("h3",a({},{id:"negative"}),"Negative"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"By adding the extra ",(0,n.kt)("inlineCode",{parentName:"li"},"ExpirationTime")," field we are going to raise the overall ",(0,n.kt)("inlineCode",{parentName:"li"},"AppLinks")," handling complexity. Since we're performing an iteration over all the expired link references inside at the start of each block this can require an amount of time that SHOULD be studied with benchmark tests during the implementation.")),(0,n.kt)("h3",a({},{id:"neutral"}),"Neutral"),(0,n.kt)("p",null,"(none known)"),(0,n.kt)("h2",a({},{id:"test-cases-optional"}),"Test Cases ","[optional]"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"The expired ",(0,n.kt)("inlineCode",{parentName:"li"},"AppLinks")," are deleted correctly when they expire.")),(0,n.kt)("h2",a({},{id:"references"}),"References"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Issue ",(0,n.kt)("a",a({parentName:"li"},{href:"https://github.com/desmos-labs/desmos/issues/516"}),"#516")),(0,n.kt)("li",{parentName:"ul"},"PR ",(0,n.kt)("a",a({parentName:"li"},{href:"https://github.com/desmos-labs/desmos/pull/562"}),"#562"))))}u.isMDXComponent=!0}}]);