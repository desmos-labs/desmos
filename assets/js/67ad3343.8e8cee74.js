"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[22462],{3905:(e,r,n)=>{n.d(r,{Zo:()=>s,kt:()=>f});var t=n(67294);function o(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function a(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);r&&(t=t.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,t)}return n}function i(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?a(Object(n),!0).forEach((function(r){o(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function l(e,r){if(null==e)return{};var n,t,o=function(e,r){if(null==e)return{};var n,t,o={},a=Object.keys(e);for(t=0;t<a.length;t++)n=a[t],r.indexOf(n)>=0||(o[n]=e[n]);return o}(e,r);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(t=0;t<a.length;t++)n=a[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var p=t.createContext({}),c=function(e){var r=t.useContext(p),n=r;return e&&(n="function"==typeof e?e(r):i(i({},r),e)),n},s=function(e){var r=c(e.components);return t.createElement(p.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return t.createElement(t.Fragment,{},r)}},u=t.forwardRef((function(e,r){var n=e.components,o=e.mdxType,a=e.originalType,p=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=c(n),f=o,m=u["".concat(p,".").concat(f)]||u[f]||d[f]||a;return n?t.createElement(m,i(i({ref:r},s),{},{components:n})):t.createElement(m,i({ref:r},s))}));function f(e,r){var n=arguments,o=r&&r.mdxType;if("string"==typeof e||o){var a=n.length,i=new Array(a);i[0]=u;var l={};for(var p in r)hasOwnProperty.call(r,p)&&(l[p]=r[p]);l.originalType=e,l.mdxType="string"==typeof e?e:o,i[1]=l;for(var c=2;c<a;c++)i[c]=n[c];return t.createElement.apply(null,i)}return t.createElement.apply(null,n)}u.displayName="MDXCreateElement"},69576:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>c,contentTitle:()=>l,default:()=>u,frontMatter:()=>i,metadata:()=>p,toc:()=>s});n(67294);var t=n(3905);function o(){return o=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var n=arguments[r];for(var t in n)Object.prototype.hasOwnProperty.call(n,t)&&(e[t]=n[t])}return e},o.apply(this,arguments)}function a(e,r){if(null==e)return{};var n,t,o=function(e,r){if(null==e)return{};var n,t,o={},a=Object.keys(e);for(t=0;t<a.length;t++)n=a[t],r.indexOf(n)>=0||(o[n]=e[n]);return o}(e,r);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(t=0;t<a.length;t++)n=a[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}const i={id:"chain-link-proof-max-order-by",title:"chain_link_proof_max_order_by",hide_table_of_contents:!1},l=void 0,p={unversionedId:"graphql/inputs/chain-link-proof-max-order-by",id:"version-4.2.0/graphql/inputs/chain-link-proof-max-order-by",title:"chain_link_proof_max_order_by",description:'order by max() on columns of table "chainlinkproof"',source:"@site/versioned_docs/version-4.2.0/07-graphql/inputs/chain-link-proof-max-order-by.mdx",sourceDirName:"07-graphql/inputs",slug:"/graphql/inputs/chain-link-proof-max-order-by",permalink:"/4.2.0/graphql/inputs/chain-link-proof-max-order-by",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/inputs/chain-link-proof-max-order-by.mdx",tags:[],version:"4.2.0",frontMatter:{id:"chain-link-proof-max-order-by",title:"chain_link_proof_max_order_by",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"chain_link_proof_bool_exp",permalink:"/4.2.0/graphql/inputs/chain-link-proof-bool-exp"},next:{title:"chain_link_proof_min_order_by",permalink:"/4.2.0/graphql/inputs/chain-link-proof-min-order-by"}},c={},s=[{value:"Fields",id:"fields",level:3},{value:"<code>plain_text</code> (<code>order_by</code>)",id:"plain_text-order_by",level:4},{value:"<code>signature</code> (<code>order_by</code>)",id:"signature-order_by",level:4}],d={toc:s};function u(e){var{components:r}=e,n=a(e,["components"]);return(0,t.kt)("wrapper",o({},d,n,{components:r,mdxType:"MDXLayout"}),(0,t.kt)("p",null,'order by max() on columns of table "chain_link_proof"'),(0,t.kt)("pre",null,(0,t.kt)("code",o({parentName:"pre"},{className:"language-graphql"}),"input chain_link_proof_max_order_by {\n  plain_text: order_by\n  signature: order_by\n}\n")),(0,t.kt)("h3",o({},{id:"fields"}),"Fields"),(0,t.kt)("h4",o({},{id:"plain_text-order_by"}),(0,t.kt)("a",o({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"plain_text"))," (",(0,t.kt)("a",o({parentName:"h4"},{href:"../enums/order-by"}),(0,t.kt)("inlineCode",{parentName:"a"},"order_by")),")"),(0,t.kt)("h4",o({},{id:"signature-order_by"}),(0,t.kt)("a",o({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"signature"))," (",(0,t.kt)("a",o({parentName:"h4"},{href:"../enums/order-by"}),(0,t.kt)("inlineCode",{parentName:"a"},"order_by")),")"))}u.isMDXComponent=!0}}]);