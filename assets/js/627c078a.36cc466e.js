"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[34476],{3905:(e,t,r)=>{r.d(t,{Zo:()=>c,kt:()=>u});var n=r(67294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var i=n.createContext({}),d=function(e){var t=n.useContext(i),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},c=function(e){var t=d(e.components);return n.createElement(i.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,i=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),f=d(r),u=o,v=f["".concat(i,".").concat(u)]||f[u]||p[u]||a;return r?n.createElement(v,s(s({ref:t},c),{},{components:r})):n.createElement(v,s({ref:t},c))}));function u(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,s=new Array(a);s[0]=f;var l={};for(var i in t)hasOwnProperty.call(t,i)&&(l[i]=t[i]);l.originalType=e,l.mdxType="string"==typeof e?e:o,s[1]=l;for(var d=2;d<a;d++)s[d]=r[d];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},74593:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>d,contentTitle:()=>l,default:()=>f,frontMatter:()=>s,metadata:()=>i,toc:()=>c});r(67294);var n=r(3905);function o(){return o=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},o.apply(this,arguments)}function a(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}const s={id:"post-stddev-fields",title:"post_stddev_fields",hide_table_of_contents:!1},l=void 0,i={unversionedId:"graphql/objects/post-stddev-fields",id:"version-4.2.0/graphql/objects/post-stddev-fields",title:"post_stddev_fields",description:"aggregate stddev on columns",source:"@site/versioned_docs/version-4.2.0/07-graphql/objects/post-stddev-fields.mdx",sourceDirName:"07-graphql/objects",slug:"/graphql/objects/post-stddev-fields",permalink:"/4.2.0/graphql/objects/post-stddev-fields",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/objects/post-stddev-fields.mdx",tags:[],version:"4.2.0",frontMatter:{id:"post-stddev-fields",title:"post_stddev_fields",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"post_reference",permalink:"/4.2.0/graphql/objects/post-reference"},next:{title:"post_stddev_pop_fields",permalink:"/4.2.0/graphql/objects/post-stddev-pop-fields"}},d={},c=[{value:"Fields",id:"fields",level:3},{value:"<code>id</code> (<code>Float</code>)",id:"id-float",level:4},{value:"<code>subspace_id</code> (<code>Float</code>)",id:"subspace_id-float",level:4}],p={toc:c};function f(e){var{components:t}=e,r=a(e,["components"]);return(0,n.kt)("wrapper",o({},p,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("p",null,"aggregate stddev on columns"),(0,n.kt)("pre",null,(0,n.kt)("code",o({parentName:"pre"},{className:"language-graphql"}),"type post_stddev_fields {\n  id: Float\n  subspace_id: Float\n}\n")),(0,n.kt)("h3",o({},{id:"fields"}),"Fields"),(0,n.kt)("h4",o({},{id:"id-float"}),(0,n.kt)("a",o({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"id"))," (",(0,n.kt)("a",o({parentName:"h4"},{href:"../scalars/float"}),(0,n.kt)("inlineCode",{parentName:"a"},"Float")),")"),(0,n.kt)("h4",o({},{id:"subspace_id-float"}),(0,n.kt)("a",o({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace_id"))," (",(0,n.kt)("a",o({parentName:"h4"},{href:"../scalars/float"}),(0,n.kt)("inlineCode",{parentName:"a"},"Float")),")"))}f.isMDXComponent=!0}}]);