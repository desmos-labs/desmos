"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[61519],{3905:(e,t,r)=>{r.d(t,{Zo:()=>c,kt:()=>u});var n=r(67294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var i=n.createContext({}),p=function(e){var t=n.useContext(i),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},c=function(e){var t=p(e.components);return n.createElement(i.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,o=e.originalType,i=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),f=p(r),u=a,m=f["".concat(i,".").concat(u)]||f[u]||d[u]||o;return r?n.createElement(m,s(s({ref:t},c),{},{components:r})):n.createElement(m,s({ref:t},c))}));function u(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=r.length,s=new Array(o);s[0]=f;var l={};for(var i in t)hasOwnProperty.call(t,i)&&(l[i]=t[i]);l.originalType=e,l.mdxType="string"==typeof e?e:a,s[1]=l;for(var p=2;p<o;p++)s[p]=r[p];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},26194:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>p,contentTitle:()=>l,default:()=>f,frontMatter:()=>s,metadata:()=>i,toc:()=>c});r(67294);var n=r(3905);function a(){return a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},a.apply(this,arguments)}function o(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}const s={id:"post-var-samp-fields",title:"post_var_samp_fields",hide_table_of_contents:!1},l=void 0,i={unversionedId:"graphql/objects/post-var-samp-fields",id:"version-4.2.0/graphql/objects/post-var-samp-fields",title:"post_var_samp_fields",description:"aggregate var_samp on columns",source:"@site/versioned_docs/version-4.2.0/07-graphql/objects/post-var-samp-fields.mdx",sourceDirName:"07-graphql/objects",slug:"/graphql/objects/post-var-samp-fields",permalink:"/4.2.0/graphql/objects/post-var-samp-fields",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/objects/post-var-samp-fields.mdx",tags:[],version:"4.2.0",frontMatter:{id:"post-var-samp-fields",title:"post_var_samp_fields",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"post_var_pop_fields",permalink:"/4.2.0/graphql/objects/post-var-pop-fields"},next:{title:"post_variance_fields",permalink:"/4.2.0/graphql/objects/post-variance-fields"}},p={},c=[{value:"Fields",id:"fields",level:3},{value:"<code>id</code> (<code>Float</code>)",id:"id-float",level:4},{value:"<code>subspace_id</code> (<code>Float</code>)",id:"subspace_id-float",level:4}],d={toc:c};function f(e){var{components:t}=e,r=o(e,["components"]);return(0,n.kt)("wrapper",a({},d,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("p",null,"aggregate var_samp on columns"),(0,n.kt)("pre",null,(0,n.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"type post_var_samp_fields {\n  id: Float\n  subspace_id: Float\n}\n")),(0,n.kt)("h3",a({},{id:"fields"}),"Fields"),(0,n.kt)("h4",a({},{id:"id-float"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"id"))," (",(0,n.kt)("a",a({parentName:"h4"},{href:"../scalars/float"}),(0,n.kt)("inlineCode",{parentName:"a"},"Float")),")"),(0,n.kt)("h4",a({},{id:"subspace_id-float"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace_id"))," (",(0,n.kt)("a",a({parentName:"h4"},{href:"../scalars/float"}),(0,n.kt)("inlineCode",{parentName:"a"},"Float")),")"))}f.isMDXComponent=!0}}]);