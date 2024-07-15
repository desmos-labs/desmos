"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[72045],{3905:(t,e,a)=>{a.d(e,{Zo:()=>s,kt:()=>o});var r=a(67294);function n(t,e,a){return e in t?Object.defineProperty(t,e,{value:a,enumerable:!0,configurable:!0,writable:!0}):t[e]=a,t}function l(t,e){var a=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),a.push.apply(a,r)}return a}function i(t){for(var e=1;e<arguments.length;e++){var a=null!=arguments[e]?arguments[e]:{};e%2?l(Object(a),!0).forEach((function(e){n(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):l(Object(a)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}function p(t,e){if(null==t)return{};var a,r,n=function(t,e){if(null==t)return{};var a,r,n={},l=Object.keys(t);for(r=0;r<l.length;r++)a=l[r],e.indexOf(a)>=0||(n[a]=t[a]);return n}(t,e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(t);for(r=0;r<l.length;r++)a=l[r],e.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(t,a)&&(n[a]=t[a])}return n}var d=r.createContext({}),m=function(t){var e=r.useContext(d),a=e;return t&&(a="function"==typeof t?t(e):i(i({},e),t)),a},s=function(t){var e=m(t.components);return r.createElement(d.Provider,{value:e},t.children)},k={inlineCode:"code",wrapper:function(t){var e=t.children;return r.createElement(r.Fragment,{},e)}},g=r.forwardRef((function(t,e){var a=t.components,n=t.mdxType,l=t.originalType,d=t.parentName,s=p(t,["components","mdxType","originalType","parentName"]),g=m(a),o=n,N=g["".concat(d,".").concat(o)]||g[o]||k[o]||l;return a?r.createElement(N,i(i({ref:e},s),{},{components:a})):r.createElement(N,i({ref:e},s))}));function o(t,e){var a=arguments,n=e&&e.mdxType;if("string"==typeof t||n){var l=a.length,i=new Array(l);i[0]=g;var p={};for(var d in e)hasOwnProperty.call(e,d)&&(p[d]=e[d]);p.originalType=t,p.mdxType="string"==typeof t?t:n,i[1]=p;for(var m=2;m<l;m++)i[m]=a[m];return r.createElement.apply(null,i)}return r.createElement.apply(null,a)}g.displayName="MDXCreateElement"},32471:(t,e,a)=>{a.r(e),a.d(e,{assets:()=>m,contentTitle:()=>p,default:()=>g,frontMatter:()=>i,metadata:()=>d,toc:()=>s});a(67294);var r=a(3905);function n(){return n=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var a=arguments[e];for(var r in a)Object.prototype.hasOwnProperty.call(a,r)&&(t[r]=a[r])}return t},n.apply(this,arguments)}function l(t,e){if(null==t)return{};var a,r,n=function(t,e){if(null==t)return{};var a,r,n={},l=Object.keys(t);for(r=0;r<l.length;r++)a=l[r],e.indexOf(a)>=0||(n[a]=t[a]);return n}(t,e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(t);for(r=0;r<l.length;r++)a=l[r],e.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(t,a)&&(n[a]=t[a])}return n}const i={id:"events",title:"Events",sidebar_label:"Events",slug:"events"},p="Events",d={unversionedId:"developers/modules/profiles/events",id:"developers/modules/profiles/events",title:"Events",description:"The profiles module emits the following events:",source:"@site/docs/02-developers/02-modules/profiles/05-events.md",sourceDirName:"02-developers/02-modules/profiles",slug:"/developers/modules/profiles/events",permalink:"/developers/modules/profiles/events",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/02-developers/02-modules/profiles/05-events.md",tags:[],version:"current",sidebarPosition:5,frontMatter:{id:"events",title:"Events",sidebar_label:"Events",slug:"events"},sidebar:"docs",previous:{title:"Messages",permalink:"/developers/modules/profiles/messages"},next:{title:"Parameters",permalink:"/developers/modules/profiles/params"}},m={},s=[{value:"Handlers",id:"handlers",level:2},{value:"MsgSaveProfile",id:"msgsaveprofile",level:3},{value:"MsgDeleteProfile",id:"msgdeleteprofile",level:2},{value:"MsgRequestDTagTransfer",id:"msgrequestdtagtransfer",level:2},{value:"MsgCancelDTagTransferRequest",id:"msgcanceldtagtransferrequest",level:2},{value:"MsgAcceptDTagTransferRequest",id:"msgacceptdtagtransferrequest",level:2},{value:"MsgRefuseDTagTransferRequest",id:"msgrefusedtagtransferrequest",level:2},{value:"MsgLinkChainAccount",id:"msglinkchainaccount",level:2},{value:"MsgUnlinkChainAccount",id:"msgunlinkchainaccount",level:2},{value:"MsgSetDefaultExternalAddress",id:"msgsetdefaultexternaladdress",level:2},{value:"MsgLinkApplication",id:"msglinkapplication",level:2},{value:"MsgUnlinkApplication",id:"msgunlinkapplication",level:2},{value:"Keeper",id:"keeper",level:2},{value:"Chain Link Saved",id:"chain-link-saved",level:3},{value:"Application Link Saved",id:"application-link-saved",level:3},{value:"IBC",id:"ibc",level:2},{value:"Received link chain account IBC packet",id:"received-link-chain-account-ibc-packet",level:3},{value:"Received oracle response IBC packet",id:"received-oracle-response-ibc-packet",level:3}],k={toc:s};function g(t){var{components:e}=t,a=l(t,["components"]);return(0,r.kt)("wrapper",n({},k,a,{components:e,mdxType:"MDXLayout"}),(0,r.kt)("h1",n({},{id:"events"}),"Events"),(0,r.kt)("p",null,"The profiles module emits the following events:"),(0,r.kt)("h2",n({},{id:"handlers"}),"Handlers"),(0,r.kt)("h3",n({},{id:"msgsaveprofile"}),"MsgSaveProfile"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_profile"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profile_dtag"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{profileDTag}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_profile"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profile_creator"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_profile"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profile_creation_time"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{profileCreationTime}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgSaveProfile")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgdeleteprofile"}),"MsgDeleteProfile"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_profile"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profile_creator"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgDeleteProfile")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgrequestdtagtransfer"}),"MsgRequestDTagTransfer"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"requested_dtag_transfer"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"dtag_to_trade"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{dTagToTrade}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"requested_dtag_transfer"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestSenderAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"requested_dtag_transfer"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_receiver"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestReceiverAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgRequestDTagTransfer")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestSenderAddress}")))),(0,r.kt)("h2",n({},{id:"msgcanceldtagtransferrequest"}),"MsgCancelDTagTransferRequest"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"canceled_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestSenderAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"canceled_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_receiver"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestReceiverAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgCancelDTagTransferRequest")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgacceptdtagtransferrequest"}),"MsgAcceptDTagTransferRequest"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"accepted_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"dtag_to_trade"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{dTagToTrade}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"accepted_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"new_dtag"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{newDTag}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"accepted_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestSenderAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"accepted_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_receiver"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestReceiverAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgAcceptDTagTransferRequest")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgrefusedtagtransferrequest"}),"MsgRefuseDTagTransferRequest"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"refused_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestSenderAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"refused_dtag_transfer_request"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_receiver"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestReceiverAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgRefuseDTagTransferRequest")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msglinkchainaccount"}),"MsgLinkChainAccount"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_account_target"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{targetAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_source_chain_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{chainName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_account_owner"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_creation_time"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{creationTime}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgLinkChainAccount")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgunlinkchainaccount"}),"MsgUnlinkChainAccount"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_account_target"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{targetAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_source_chain_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{chainName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_account_owner"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgUnlinkChainAccount")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgsetdefaultexternaladdress"}),"MsgSetDefaultExternalAddress"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"set_default_external_address"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_chain_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{chainName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"set_default_external_address"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_external_address"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{externalAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"set_default_external_address"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_owner"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{chainLinkOwner}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgSetDefaultAddress")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msglinkapplication"}),"MsgLinkApplication"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"user"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{applicationName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_username"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{applicationUsername}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"created_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_link_creation_time"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{creationTime}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgLinkApplication")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"msgunlinkapplication"}),"MsgUnlinkApplication"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"user"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{applicationName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"deleted_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_username"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{applicationUsername}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"action"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"desmos.profiles.v3.MsgUnlinkApplication")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"message"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"sender"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")))),(0,r.kt)("h2",n({},{id:"keeper"}),"Keeper"),(0,r.kt)("h3",n({},{id:"chain-link-saved"}),"Chain Link Saved"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_owner"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_chain_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{chainName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_chain_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_external_address"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{externalAddress}")))),(0,r.kt)("h3",n({},{id:"application-link-saved"}),"Application Link Saved"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"user"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{applicationName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"saved_application_link"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"application_username"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{applicationUsername}")))),(0,r.kt)("h2",n({},{id:"ibc"}),"IBC"),(0,r.kt)("h3",n({},{id:"received-link-chain-account-ibc-packet"}),"Received link chain account IBC packet"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_link_chain_account_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_link_chain_account_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_owner"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{userAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_link_chain_account_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_chain_name"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{chainName}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_link_chain_account_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"chain_link_external_address"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{externalAddress}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_link_chain_account_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"success"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"true")))),(0,r.kt)("h3",n({},{id:"received-oracle-response-ibc-packet"}),"Received oracle response IBC packet"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Type")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Key")),(0,r.kt)("th",n({parentName:"tr"},{align:"left"}),(0,r.kt)("strong",{parentName:"th"},"Attribute Value")))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_oracle_response_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"module"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"profiles")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_oracle_response_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"client_id"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{clientID}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_oracle_response_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"request_id"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{requestID}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_oracle_response_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"resolve_status"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"{resolveStatus}")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"received_oracle_response_packet"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"success"),(0,r.kt)("td",n({parentName:"tr"},{align:"left"}),"true")))))}g.isMDXComponent=!0}}]);