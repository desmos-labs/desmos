"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[93355],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),p=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=p(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},f=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),f=p(n),m=a,u=f["".concat(s,".").concat(m)]||f[m]||d[m]||l;return n?r.createElement(u,o(o({ref:t},c),{},{components:n})):r.createElement(u,o({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,o=new Array(l);o[0]=f;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:a,o[1]=i;for(var p=2;p<l;p++)o[p]=n[p];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}f.displayName="MDXCreateElement"},61969:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>i,default:()=>f,frontMatter:()=>o,metadata:()=>s,toc:()=>c});n(67294);var r=n(3905);function a(){return a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},a.apply(this,arguments)}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}const o={id:"poll-answer-aggregate-fields",title:"poll_answer_aggregate_fields",hide_table_of_contents:!1},i=void 0,s={unversionedId:"graphql/objects/poll-answer-aggregate-fields",id:"version-4.2.0/graphql/objects/poll-answer-aggregate-fields",title:"poll_answer_aggregate_fields",description:'aggregate fields of "poll_answer"',source:"@site/versioned_docs/version-4.2.0/07-graphql/objects/poll-answer-aggregate-fields.mdx",sourceDirName:"07-graphql/objects",slug:"/graphql/objects/poll-answer-aggregate-fields",permalink:"/4.2.0/graphql/objects/poll-answer-aggregate-fields",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/objects/poll-answer-aggregate-fields.mdx",tags:[],version:"4.2.0",frontMatter:{id:"poll-answer-aggregate-fields",title:"poll_answer_aggregate_fields",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"dtag_transfer_requests",permalink:"/4.2.0/graphql/objects/dtag-transfer-requests"},next:{title:"poll_answer_aggregate",permalink:"/4.2.0/graphql/objects/poll-answer-aggregate"}},p={},c=[{value:"Fields",id:"fields",level:3},{value:"<code>count</code> (<code>Int!</code>)",id:"count-int",level:4},{value:"<code>max</code> (<code>poll_answer_max_fields</code>)",id:"max-poll_answer_max_fields",level:4},{value:"<code>min</code> (<code>poll_answer_min_fields</code>)",id:"min-poll_answer_min_fields",level:4}],d={toc:c};function f(e){var{components:t}=e,n=l(e,["components"]);return(0,r.kt)("wrapper",a({},d,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("p",null,'aggregate fields of "poll_answer"'),(0,r.kt)("pre",null,(0,r.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"type poll_answer_aggregate_fields {\n  count(\n  columns: [poll_answer_select_column!]\n  distinct: Boolean\n): Int!\n  max: poll_answer_max_fields\n  min: poll_answer_min_fields\n}\n")),(0,r.kt)("h3",a({},{id:"fields"}),"Fields"),(0,r.kt)("h4",a({},{id:"count-int"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"count"))," (",(0,r.kt)("a",a({parentName:"h4"},{href:"../scalars/int"}),(0,r.kt)("inlineCode",{parentName:"a"},"Int!")),")"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("h5",a({parentName:"li"},{id:"columns-poll_answer_select_column"}),(0,r.kt)("a",a({parentName:"h5"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"columns"))," (",(0,r.kt)("a",a({parentName:"h5"},{href:"../enums/poll-answer-select-column"}),(0,r.kt)("inlineCode",{parentName:"a"},"[poll_answer_select_column!]")),")"))),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("h5",a({parentName:"li"},{id:"distinct-boolean"}),(0,r.kt)("a",a({parentName:"h5"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"distinct"))," (",(0,r.kt)("a",a({parentName:"h5"},{href:"../scalars/boolean"}),(0,r.kt)("inlineCode",{parentName:"a"},"Boolean")),")"))),(0,r.kt)("h4",a({},{id:"max-poll_answer_max_fields"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"max"))," (",(0,r.kt)("a",a({parentName:"h4"},{href:"../objects/poll-answer-max-fields"}),(0,r.kt)("inlineCode",{parentName:"a"},"poll_answer_max_fields")),")"),(0,r.kt)("h4",a({},{id:"min-poll_answer_min_fields"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"min"))," (",(0,r.kt)("a",a({parentName:"h4"},{href:"../objects/poll-answer-min-fields"}),(0,r.kt)("inlineCode",{parentName:"a"},"poll_answer_min_fields")),")"))}f.isMDXComponent=!0}}]);