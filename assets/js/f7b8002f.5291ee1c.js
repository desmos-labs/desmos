"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[11834],{3905:(e,r,t)=>{t.d(r,{Zo:()=>g,kt:()=>f});var n=t(67294);function a(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function o(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function l(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?o(Object(t),!0).forEach((function(r){a(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function s(e,r){if(null==e)return{};var t,n,a=function(e,r){if(null==e)return{};var t,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(a[t]=e[t]);return a}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var p=n.createContext({}),i=function(e){var r=n.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):l(l({},r),e)),t},g=function(e){var r=i(e.components);return n.createElement(p.Provider,{value:r},e.children)},c={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},d=n.forwardRef((function(e,r){var t=e.components,a=e.mdxType,o=e.originalType,p=e.parentName,g=s(e,["components","mdxType","originalType","parentName"]),d=i(t),f=a,u=d["".concat(p,".").concat(f)]||d[f]||c[f]||o;return t?n.createElement(u,l(l({ref:r},g),{},{components:t})):n.createElement(u,l({ref:r},g))}));function f(e,r){var t=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var o=t.length,l=new Array(o);l[0]=d;var s={};for(var p in r)hasOwnProperty.call(r,p)&&(s[p]=r[p]);s.originalType=e,s.mdxType="string"==typeof e?e:a,l[1]=s;for(var i=2;i<o;i++)l[i]=t[i];return n.createElement.apply(null,l)}return n.createElement.apply(null,t)}d.displayName="MDXCreateElement"},8184:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>i,contentTitle:()=>s,default:()=>d,frontMatter:()=>l,metadata:()=>p,toc:()=>g});t(67294);var n=t(3905);function a(){return a=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},a.apply(this,arguments)}function o(e,r){if(null==e)return{};var t,n,a=function(e,r){if(null==e)return{};var t,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(a[t]=e[t]);return a}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}const l={id:"poll-answer-aggregate",title:"poll_answer_aggregate",hide_table_of_contents:!1},s=void 0,p={unversionedId:"graphql/objects/poll-answer-aggregate",id:"version-4.2.0/graphql/objects/poll-answer-aggregate",title:"poll_answer_aggregate",description:'aggregated selection of "poll_answer"',source:"@site/versioned_docs/version-4.2.0/07-graphql/objects/poll-answer-aggregate.mdx",sourceDirName:"07-graphql/objects",slug:"/graphql/objects/poll-answer-aggregate",permalink:"/4.2.0/graphql/objects/poll-answer-aggregate",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/objects/poll-answer-aggregate.mdx",tags:[],version:"4.2.0",frontMatter:{id:"poll-answer-aggregate",title:"poll_answer_aggregate",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"poll_answer_aggregate_fields",permalink:"/4.2.0/graphql/objects/poll-answer-aggregate-fields"},next:{title:"poll_answer_max_fields",permalink:"/4.2.0/graphql/objects/poll-answer-max-fields"}},i={},g=[{value:"Fields",id:"fields",level:3},{value:"<code>aggregate</code> (<code>poll_answer_aggregate_fields</code>)",id:"aggregate-poll_answer_aggregate_fields",level:4},{value:"<code>nodes</code> (<code>[poll_answer!]!</code>)",id:"nodes-poll_answer",level:4}],c={toc:g};function d(e){var{components:r}=e,t=o(e,["components"]);return(0,n.kt)("wrapper",a({},c,t,{components:r,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'aggregated selection of "poll_answer"'),(0,n.kt)("pre",null,(0,n.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"type poll_answer_aggregate {\n  aggregate: poll_answer_aggregate_fields\n  nodes: [poll_answer!]!\n}\n")),(0,n.kt)("h3",a({},{id:"fields"}),"Fields"),(0,n.kt)("h4",a({},{id:"aggregate-poll_answer_aggregate_fields"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"aggregate"))," (",(0,n.kt)("a",a({parentName:"h4"},{href:"../objects/poll-answer-aggregate-fields"}),(0,n.kt)("inlineCode",{parentName:"a"},"poll_answer_aggregate_fields")),")"),(0,n.kt)("h4",a({},{id:"nodes-poll_answer"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"nodes"))," (",(0,n.kt)("a",a({parentName:"h4"},{href:"../objects/poll-answer"}),(0,n.kt)("inlineCode",{parentName:"a"},"[poll_answer!]!")),")"))}d.isMDXComponent=!0}}]);