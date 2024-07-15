"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[49376],{3905:(e,t,r)=>{r.d(t,{Zo:()=>u,kt:()=>d});var n=r(67294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function l(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?l(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):l(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function o(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},l=Object.keys(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var c=n.createContext({}),p=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},u=function(e){var t=p(e.components);return n.createElement(c.Provider,{value:t},e.children)},i={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,l=e.originalType,c=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),m=p(r),d=a,f=m["".concat(c,".").concat(d)]||m[d]||i[d]||l;return r?n.createElement(f,s(s({ref:t},u),{},{components:r})):n.createElement(f,s({ref:t},u))}));function d(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=r.length,s=new Array(l);s[0]=m;var o={};for(var c in t)hasOwnProperty.call(t,c)&&(o[c]=t[c]);o.originalType=e,o.mdxType="string"==typeof e?e:a,s[1]=o;for(var p=2;p<l;p++)s[p]=r[p];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},97369:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>p,contentTitle:()=>o,default:()=>m,frontMatter:()=>s,metadata:()=>c,toc:()=>u});r(67294);var n=r(3905);function a(){return a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},a.apply(this,arguments)}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},l=Object.keys(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)r=l[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}const s={id:"subspace-free-text-params-select-column",title:"subspace_free_text_params_select_column",hide_table_of_contents:!1},o=void 0,c={unversionedId:"graphql/enums/subspace-free-text-params-select-column",id:"version-4.2.0/graphql/enums/subspace-free-text-params-select-column",title:"subspace_free_text_params_select_column",description:'select columns of table "subspacefreetext_params"',source:"@site/versioned_docs/version-4.2.0/07-graphql/enums/subspace-free-text-params-select-column.mdx",sourceDirName:"07-graphql/enums",slug:"/graphql/enums/subspace-free-text-params-select-column",permalink:"/4.2.0/graphql/enums/subspace-free-text-params-select-column",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/enums/subspace-free-text-params-select-column.mdx",tags:[],version:"4.2.0",frontMatter:{id:"subspace-free-text-params-select-column",title:"subspace_free_text_params_select_column",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"reports_params_select_column",permalink:"/4.2.0/graphql/enums/reports-params-select-column"},next:{title:"subspace_registered_reaction_params_select_column",permalink:"/4.2.0/graphql/enums/subspace-registered-reaction-params-select-column"}},p={},u=[{value:"Values",id:"values",level:3},{value:"<code>enabled</code>",id:"enabled",level:4},{value:"<code>max_length</code>",id:"max_length",level:4},{value:"<code>reg_ex</code>",id:"reg_ex",level:4},{value:"<code>subspace_id</code>",id:"subspace_id",level:4}],i={toc:u};function m(e){var{components:t}=e,r=l(e,["components"]);return(0,n.kt)("wrapper",a({},i,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'select columns of table "subspace_free_text_params"'),(0,n.kt)("pre",null,(0,n.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"enum subspace_free_text_params_select_column {\n  enabled\n  max_length\n  reg_ex\n  subspace_id\n}\n")),(0,n.kt)("h3",a({},{id:"values"}),"Values"),(0,n.kt)("h4",a({},{id:"enabled"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"enabled"))),(0,n.kt)("p",null,"column name"),(0,n.kt)("h4",a({},{id:"max_length"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"max_length"))),(0,n.kt)("p",null,"column name"),(0,n.kt)("h4",a({},{id:"reg_ex"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"reg_ex"))),(0,n.kt)("p",null,"column name"),(0,n.kt)("h4",a({},{id:"subspace_id"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace_id"))),(0,n.kt)("p",null,"column name"))}m.isMDXComponent=!0}}]);