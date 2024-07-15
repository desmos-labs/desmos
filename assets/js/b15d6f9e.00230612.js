"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[89063],{3905:(e,r,t)=>{t.d(r,{Zo:()=>u,kt:()=>b});var n=t(67294);function s(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function o(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function i(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?o(Object(t),!0).forEach((function(r){s(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function a(e,r){if(null==e)return{};var t,n,s=function(e,r){if(null==e)return{};var t,n,s={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(s[t]=e[t]);return s}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(s[t]=e[t])}return s}var p=n.createContext({}),c=function(e){var r=n.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):i(i({},r),e)),t},u=function(e){var r=c(e.components);return n.createElement(p.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},l=n.forwardRef((function(e,r){var t=e.components,s=e.mdxType,o=e.originalType,p=e.parentName,u=a(e,["components","mdxType","originalType","parentName"]),l=c(t),b=s,m=l["".concat(p,".").concat(b)]||l[b]||d[b]||o;return t?n.createElement(m,i(i({ref:r},u),{},{components:t})):n.createElement(m,i({ref:r},u))}));function b(e,r){var t=arguments,s=r&&r.mdxType;if("string"==typeof e||s){var o=t.length,i=new Array(o);i[0]=l;var a={};for(var p in r)hasOwnProperty.call(r,p)&&(a[p]=r[p]);a.originalType=e,a.mdxType="string"==typeof e?e:s,i[1]=a;for(var c=2;c<o;c++)i[c]=t[c];return n.createElement.apply(null,i)}return n.createElement.apply(null,t)}l.displayName="MDXCreateElement"},51738:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>c,contentTitle:()=>a,default:()=>l,frontMatter:()=>i,metadata:()=>p,toc:()=>u});t(67294);var n=t(3905);function s(){return s=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},s.apply(this,arguments)}function o(e,r){if(null==e)return{};var t,n,s=function(e,r){if(null==e)return{};var t,n,s={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(s[t]=e[t]);return s}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(s[t]=e[t])}return s}const i={id:"subspace-user-permission-order-by",title:"subspace_user_permission_order_by",hide_table_of_contents:!1},a=void 0,p={unversionedId:"graphql/inputs/subspace-user-permission-order-by",id:"version-4.2.0/graphql/inputs/subspace-user-permission-order-by",title:"subspace_user_permission_order_by",description:'Ordering options when selecting data from "subspaceuserpermission".',source:"@site/versioned_docs/version-4.2.0/07-graphql/inputs/subspace-user-permission-order-by.mdx",sourceDirName:"07-graphql/inputs",slug:"/graphql/inputs/subspace-user-permission-order-by",permalink:"/4.2.0/graphql/inputs/subspace-user-permission-order-by",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/inputs/subspace-user-permission-order-by.mdx",tags:[],version:"4.2.0",frontMatter:{id:"subspace-user-permission-order-by",title:"subspace_user_permission_order_by",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"subspace_user_permission_min_order_by",permalink:"/4.2.0/graphql/inputs/subspace-user-permission-min-order-by"},next:{title:"_text_comparison_exp",permalink:"/4.2.0/graphql/inputs/text-comparison-exp"}},c={},u=[{value:"Fields",id:"fields",level:3},{value:"<code>permissions</code> (<code>order_by</code>)",id:"permissions-order_by",level:4},{value:"<code>section</code> (<code>subspace_section_order_by</code>)",id:"section-subspace_section_order_by",level:4},{value:"<code>user_address</code> (<code>order_by</code>)",id:"user_address-order_by",level:4}],d={toc:u};function l(e){var{components:r}=e,t=o(e,["components"]);return(0,n.kt)("wrapper",s({},d,t,{components:r,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'Ordering options when selecting data from "subspace_user_permission".'),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-graphql"}),"input subspace_user_permission_order_by {\n  permissions: order_by\n  section: subspace_section_order_by\n  user_address: order_by\n}\n")),(0,n.kt)("h3",s({},{id:"fields"}),"Fields"),(0,n.kt)("h4",s({},{id:"permissions-order_by"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"permissions"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../enums/order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by")),")"),(0,n.kt)("h4",s({},{id:"section-subspace_section_order_by"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"section"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../inputs/subspace-section-order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace_section_order_by")),")"),(0,n.kt)("h4",s({},{id:"user_address-order_by"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"user_address"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../enums/order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by")),")"))}l.isMDXComponent=!0}}]);