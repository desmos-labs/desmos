"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[41862],{3905:(e,r,t)=>{t.d(r,{Zo:()=>i,kt:()=>b});var n=t(67294);function s(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function a(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function o(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?a(Object(t),!0).forEach((function(r){s(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function u(e,r){if(null==e)return{};var t,n,s=function(e,r){if(null==e)return{};var t,n,s={},a=Object.keys(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||(s[t]=e[t]);return s}(e,r);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(s[t]=e[t])}return s}var p=n.createContext({}),c=function(e){var r=n.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):o(o({},r),e)),t},i=function(e){var r=c(e.components);return n.createElement(p.Provider,{value:r},e.children)},l={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},m=n.forwardRef((function(e,r){var t=e.components,s=e.mdxType,a=e.originalType,p=e.parentName,i=u(e,["components","mdxType","originalType","parentName"]),m=c(t),b=s,g=m["".concat(p,".").concat(b)]||m[b]||l[b]||a;return t?n.createElement(g,o(o({ref:r},i),{},{components:t})):n.createElement(g,o({ref:r},i))}));function b(e,r){var t=arguments,s=r&&r.mdxType;if("string"==typeof e||s){var a=t.length,o=new Array(a);o[0]=m;var u={};for(var p in r)hasOwnProperty.call(r,p)&&(u[p]=r[p]);u.originalType=e,u.mdxType="string"==typeof e?e:s,o[1]=u;for(var c=2;c<a;c++)o[c]=t[c];return n.createElement.apply(null,o)}return n.createElement.apply(null,t)}m.displayName="MDXCreateElement"},79442:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>c,contentTitle:()=>u,default:()=>m,frontMatter:()=>o,metadata:()=>p,toc:()=>i});t(67294);var n=t(3905);function s(){return s=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},s.apply(this,arguments)}function a(e,r){if(null==e)return{};var t,n,s=function(e,r){if(null==e)return{};var t,n,s={},a=Object.keys(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||(s[t]=e[t]);return s}(e,r);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(s[t]=e[t])}return s}const o={id:"subspace-user-group-member-aggregate",title:"subspace_user_group_member_aggregate",hide_table_of_contents:!1},u=void 0,p={unversionedId:"graphql/subscriptions/subspace-user-group-member-aggregate",id:"version-4.2.0/graphql/subscriptions/subspace-user-group-member-aggregate",title:"subspace_user_group_member_aggregate",description:'fetch aggregated fields from the table: "subspaceusergroup_member"',source:"@site/versioned_docs/version-4.2.0/07-graphql/subscriptions/subspace-user-group-member-aggregate.mdx",sourceDirName:"07-graphql/subscriptions",slug:"/graphql/subscriptions/subspace-user-group-member-aggregate",permalink:"/4.2.0/graphql/subscriptions/subspace-user-group-member-aggregate",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/subscriptions/subspace-user-group-member-aggregate.mdx",tags:[],version:"4.2.0",frontMatter:{id:"subspace-user-group-member-aggregate",title:"subspace_user_group_member_aggregate",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"subspace_user_group_aggregate",permalink:"/4.2.0/graphql/subscriptions/subspace-user-group-aggregate"},next:{title:"subspace_user_group_member",permalink:"/4.2.0/graphql/subscriptions/subspace-user-group-member"}},c={},i=[{value:"Arguments",id:"arguments",level:3},{value:"<code>distinct_on</code> (<code>[subspace_user_group_member_select_column!]</code>)",id:"distinct_on-subspace_user_group_member_select_column",level:4},{value:"<code>limit</code> (<code>Int</code>)",id:"limit-int",level:4},{value:"<code>offset</code> (<code>Int</code>)",id:"offset-int",level:4},{value:"<code>order_by</code> (<code>[subspace_user_group_member_order_by!]</code>)",id:"order_by-subspace_user_group_member_order_by",level:4},{value:"<code>where</code> (<code>subspace_user_group_member_bool_exp</code>)",id:"where-subspace_user_group_member_bool_exp",level:4},{value:"Type",id:"type",level:3},{value:"<code>subspace_user_group_member_aggregate</code>",id:"subspace_user_group_member_aggregate",level:4}],l={toc:i};function m(e){var{components:r}=e,t=a(e,["components"]);return(0,n.kt)("wrapper",s({},l,t,{components:r,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'fetch aggregated fields from the table: "subspace_user_group_member"'),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-graphql"}),"subspace_user_group_member_aggregate(\n  distinct_on: [subspace_user_group_member_select_column!]\n  limit: Int\n  offset: Int\n  order_by: [subspace_user_group_member_order_by!]\n  where: subspace_user_group_member_bool_exp\n): subspace_user_group_member_aggregate!\n")),(0,n.kt)("h3",s({},{id:"arguments"}),"Arguments"),(0,n.kt)("h4",s({},{id:"distinct_on-subspace_user_group_member_select_column"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"distinct_on"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../enums/subspace-user-group-member-select-column"}),(0,n.kt)("inlineCode",{parentName:"a"},"[subspace_user_group_member_select_column!]")),")"),(0,n.kt)("p",null,"distinct select on columns"),(0,n.kt)("h4",s({},{id:"limit-int"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"limit"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../scalars/int"}),(0,n.kt)("inlineCode",{parentName:"a"},"Int")),")"),(0,n.kt)("p",null,"limit the number of rows returned"),(0,n.kt)("h4",s({},{id:"offset-int"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"offset"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../scalars/int"}),(0,n.kt)("inlineCode",{parentName:"a"},"Int")),")"),(0,n.kt)("p",null,"skip the first n rows. Use only with order_by"),(0,n.kt)("h4",s({},{id:"order_by-subspace_user_group_member_order_by"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../inputs/subspace-user-group-member-order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"[subspace_user_group_member_order_by!]")),")"),(0,n.kt)("p",null,"sort the rows by one or more columns"),(0,n.kt)("h4",s({},{id:"where-subspace_user_group_member_bool_exp"}),(0,n.kt)("a",s({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"where"))," (",(0,n.kt)("a",s({parentName:"h4"},{href:"../inputs/subspace-user-group-member-bool-exp"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace_user_group_member_bool_exp")),")"),(0,n.kt)("p",null,"filter the rows returned"),(0,n.kt)("h3",s({},{id:"type"}),"Type"),(0,n.kt)("h4",s({},{id:"subspace_user_group_member_aggregate"}),(0,n.kt)("a",s({parentName:"h4"},{href:"../objects/subspace-user-group-member-aggregate"}),(0,n.kt)("inlineCode",{parentName:"a"},"subspace_user_group_member_aggregate"))),(0,n.kt)("p",null,'aggregated selection of "subspace_user_group_member"'))}m.isMDXComponent=!0}}]);