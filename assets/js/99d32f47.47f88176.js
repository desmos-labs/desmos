"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[29348],{3905:(e,r,t)=>{t.d(r,{Zo:()=>c,kt:()=>u});var n=t(67294);function i(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function o(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function a(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?o(Object(t),!0).forEach((function(r){i(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function l(e,r){if(null==e)return{};var t,n,i=function(e,r){if(null==e)return{};var t,n,i={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(i[t]=e[t]);return i}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var s=n.createContext({}),p=function(e){var r=n.useContext(s),t=r;return e&&(t="function"==typeof e?e(r):a(a({},r),e)),t},c=function(e){var r=p(e.components);return n.createElement(s.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},f=n.forwardRef((function(e,r){var t=e.components,i=e.mdxType,o=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),f=p(t),u=i,h=f["".concat(s,".").concat(u)]||f[u]||d[u]||o;return t?n.createElement(h,a(a({ref:r},c),{},{components:t})):n.createElement(h,a({ref:r},c))}));function u(e,r){var t=arguments,i=r&&r.mdxType;if("string"==typeof e||i){var o=t.length,a=new Array(o);a[0]=f;var l={};for(var s in r)hasOwnProperty.call(r,s)&&(l[s]=r[s]);l.originalType=e,l.mdxType="string"==typeof e?e:i,a[1]=l;for(var p=2;p<o;p++)a[p]=t[p];return n.createElement.apply(null,a)}return n.createElement.apply(null,t)}f.displayName="MDXCreateElement"},77906:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>p,contentTitle:()=>l,default:()=>f,frontMatter:()=>a,metadata:()=>s,toc:()=>c});t(67294);var n=t(3905);function i(){return i=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},i.apply(this,arguments)}function o(e,r){if(null==e)return{};var t,n,i=function(e,r){if(null==e)return{};var t,n,i={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(i[t]=e[t]);return i}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}const a={id:"profile-relationship",title:"profile_relationship",hide_table_of_contents:!1},l=void 0,s={unversionedId:"graphql/objects/profile-relationship",id:"version-4.2.0/graphql/objects/profile-relationship",title:"profile_relationship",description:'columns and relationships of "profile_relationship"',source:"@site/versioned_docs/version-4.2.0/07-graphql/objects/profile-relationship.mdx",sourceDirName:"07-graphql/objects",slug:"/graphql/objects/profile-relationship",permalink:"/4.2.0/graphql/objects/profile-relationship",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/objects/profile-relationship.mdx",tags:[],version:"4.2.0",frontMatter:{id:"profile-relationship",title:"profile_relationship",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"profile_relationship_min_fields",permalink:"/4.2.0/graphql/objects/profile-relationship-min-fields"},next:{title:"profile",permalink:"/4.2.0/graphql/objects/profile"}},p={},c=[{value:"Fields",id:"fields",level:3},{value:"<code>receiver</code> (<code>profile!</code>)",id:"receiver-profile",level:4},{value:"<code>receiver_address</code> (<code>String!</code>)",id:"receiver_address-string",level:4},{value:"<code>sender</code> (<code>profile!</code>)",id:"sender-profile",level:4},{value:"<code>sender_address</code> (<code>String!</code>)",id:"sender_address-string",level:4},{value:"<code>subspace</code> (<code>String!</code>)",id:"subspace-string",level:4}],d={toc:c};function f(e){var{components:r}=e,t=o(e,["components"]);return(0,n.kt)("wrapper",i({},d,t,{components:r,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'columns and relationships of "profile_relationship"'),(0,n.kt)("pre",null,(0,n.kt)("code",i({parentName:"pre"},{className:"language-graphql"}),"type profile_relationship {\n  receiver: profile!\n  receiver_address: String!\n  sender: profile!\n  sender_address: String!\n  subspace: String!\n}\n")),(0,n.kt)("h3",i({},{id:"fields"}),"Fields"),(0,n.kt)("h4",i({},{id:"receiver-profile"}),(0,n.kt)("a",i({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"receiver"))," (",(0,n.kt)("a",i({parentName:"h4"},{href:"../objects/profile"}),(0,n.kt)("inlineCode",{parentName:"a"},"profile!")),")"),(0,n.kt)("p",null,"An object relationship"),(0,n.kt)("h4",i({},{id:"receiver_address-string"}),(0,n.kt)("a",i({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"receiver_address"))," (",(0,n.kt)("a",i({parentName:"h4"},{href:"../scalars/string"}),(0,n.kt)("inlineCode",{parentName:"a"},"String!")),")"),(0,n.kt)("h4",i({},{id:"sender-profile"}),(0,n.kt)("a",i({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"sender"))," (",(0,n.kt)("a",i({parentName:"h4"},{href:"../objects/profile"}),(0,n.kt)("inlineCode",{parentName:"a"},"profile!")),")"),(0,n.kt)("p",null,"An object relationship"),(0,n.kt)("h4",i({},{id:"sender_address-string"}),(0,n.kt)("a",i({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"sender_address"))," (",(0,n.kt)("a",i({parentName:"h4"},{href:"../scalars/string"}),(0,n.kt)("inlineCode",{parentName:"a"},"String!")),")"),(0,n.kt)("h4",i({},{id:"subspace-string"}),(0,n.kt)("a",i({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace"))," (",(0,n.kt)("a",i({parentName:"h4"},{href:"../scalars/string"}),(0,n.kt)("inlineCode",{parentName:"a"},"String!")),")"))}f.isMDXComponent=!0}}]);