"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[63245],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>d});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),p=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},s=function(e){var t=p(e.components);return r.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,c=e.parentName,s=i(e,["components","mdxType","originalType","parentName"]),m=p(n),d=a,f=m["".concat(c,".").concat(d)]||m[d]||u[d]||l;return n?r.createElement(f,o(o({ref:t},s),{},{components:n})):r.createElement(f,o({ref:t},s))}));function d(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,o=new Array(l);o[0]=m;var i={};for(var c in t)hasOwnProperty.call(t,c)&&(i[c]=t[c]);i.originalType=e,i.mdxType="string"==typeof e?e:a,o[1]=i;for(var p=2;p<l;p++)o[p]=n[p];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},10022:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>i,default:()=>m,frontMatter:()=>o,metadata:()=>c,toc:()=>s});n(67294);var r=n(3905);function a(){return a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},a.apply(this,arguments)}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}const o={id:"application-link-select-column",title:"application_link_select_column",hide_table_of_contents:!1},i=void 0,c={unversionedId:"graphql/enums/application-link-select-column",id:"version-4.2.0/graphql/enums/application-link-select-column",title:"application_link_select_column",description:'select columns of table "application_link"',source:"@site/versioned_docs/version-4.2.0/07-graphql/enums/application-link-select-column.mdx",sourceDirName:"07-graphql/enums",slug:"/graphql/enums/application-link-select-column",permalink:"/4.2.0/graphql/enums/application-link-select-column",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/enums/application-link-select-column.mdx",tags:[],version:"4.2.0",frontMatter:{id:"application-link-select-column",title:"application_link_select_column",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"application_link_oracle_request_select_column",permalink:"/4.2.0/graphql/enums/application-link-oracle-request-select-column"},next:{title:"chain_link_chain_config_select_column",permalink:"/4.2.0/graphql/enums/chain-link-chain-config-select-column"}},p={},s=[{value:"Values",id:"values",level:3},{value:"<code>application</code>",id:"application",level:4},{value:"<code>creation_time</code>",id:"creation_time",level:4},{value:"<code>result</code>",id:"result",level:4},{value:"<code>state</code>",id:"state",level:4},{value:"<code>user_address</code>",id:"user_address",level:4},{value:"<code>username</code>",id:"username",level:4}],u={toc:s};function m(e){var{components:t}=e,n=l(e,["components"]);return(0,r.kt)("wrapper",a({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("p",null,'select columns of table "application_link"'),(0,r.kt)("pre",null,(0,r.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"enum application_link_select_column {\n  application\n  creation_time\n  result\n  state\n  user_address\n  username\n}\n")),(0,r.kt)("h3",a({},{id:"values"}),"Values"),(0,r.kt)("h4",a({},{id:"application"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"application"))),(0,r.kt)("p",null,"column name"),(0,r.kt)("h4",a({},{id:"creation_time"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"creation_time"))),(0,r.kt)("p",null,"column name"),(0,r.kt)("h4",a({},{id:"result"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"result"))),(0,r.kt)("p",null,"column name"),(0,r.kt)("h4",a({},{id:"state"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"state"))),(0,r.kt)("p",null,"column name"),(0,r.kt)("h4",a({},{id:"user_address"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"user_address"))),(0,r.kt)("p",null,"column name"),(0,r.kt)("h4",a({},{id:"username"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"username"))),(0,r.kt)("p",null,"column name"))}m.isMDXComponent=!0}}]);