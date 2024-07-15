"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[80062],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>u});var a=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var i=a.createContext({}),d=function(e){var t=a.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=d(e.components);return a.createElement(i.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,i=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),m=d(n),u=r,k=m["".concat(i,".").concat(u)]||m[u]||c[u]||o;return n?a.createElement(k,l(l({ref:t},p),{},{components:n})):a.createElement(k,l({ref:t},p))}));function u(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,l=new Array(o);l[0]=m;var s={};for(var i in t)hasOwnProperty.call(t,i)&&(s[i]=t[i]);s.originalType=e,s.mdxType="string"==typeof e?e:r,l[1]=s;for(var d=2;d<o;d++)l[d]=n[d];return a.createElement.apply(null,l)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},95990:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>d,contentTitle:()=>s,default:()=>m,frontMatter:()=>l,metadata:()=>i,toc:()=>p});n(67294);var a=n(3905);function r(){return r=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var a in n)Object.prototype.hasOwnProperty.call(n,a)&&(e[a]=n[a])}return e},r.apply(this,arguments)}function o(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}const l={id:"create-local",title:"Create a local testnet",sidebar_label:"create-local",slug:"create-local"},s="Create a local testnet",i={unversionedId:"testnet/create-local",id:"version-4.2.0/testnet/create-local",title:"Create a local testnet",description:"You need to install desmos before going further.",source:"@site/versioned_docs/version-4.2.0/05-testnet/02-create-local.md",sourceDirName:"05-testnet",slug:"/testnet/create-local",permalink:"/4.2.0/testnet/create-local",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/05-testnet/02-create-local.md",tags:[],version:"4.2.0",sidebarPosition:2,frontMatter:{id:"create-local",title:"Create a local testnet",sidebar_label:"create-local",slug:"create-local"},sidebar:"docs",previous:{title:"Overview",permalink:"/4.2.0/testnet/overview"},next:{title:"Setup",permalink:"/4.2.0/testnet/join-public/setup"}},d={},p=[{value:"Creating a single node testnet",id:"creating-a-single-node-testnet",level:2},{value:"Creating a multi node testnet",id:"creating-a-multi-node-testnet",level:2},{value:"Configuration",id:"configuration",level:4},{value:"Logging",id:"logging",level:4},{value:"Keys &amp; Accounts",id:"keys--accounts",level:4}],c={toc:p};function m(e){var{components:t}=e,n=o(e,["components"]);return(0,a.kt)("wrapper",r({},c,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",r({},{id:"create-a-local-testnet"}),"Create a local testnet"),(0,a.kt)("admonition",r({},{title:"Required Desmos executables  ",type:"warning"}),(0,a.kt)("p",{parentName:"admonition"},"You need to ",(0,a.kt)("a",r({parentName:"p"},{href:"/4.2.0/install"}),"install ",(0,a.kt)("inlineCode",{parentName:"a"},"desmos"))," before going further.  ")),(0,a.kt)("p",null,"There are two types of local testnets:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("a",r({parentName:"p"},{href:"#creating-a-single-node-testnet"}),"Single node testnet"),", which allows you to have a faster testnet with only one validator running on your machine. ")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("p",{parentName:"li"},(0,a.kt)("a",r({parentName:"p"},{href:"#creating-a-multi-node-testnet"}),"Multi-node testnet"),", which requires you to have ",(0,a.kt)("a",r({parentName:"p"},{href:"https://docker.io"}),"Docker")," installed to run 4 validator nodes locally on your machine. "))),(0,a.kt)("h2",r({},{id:"creating-a-single-node-testnet"}),"Creating a single node testnet"),(0,a.kt)("p",null,"To create a single node local testnet, run the following commands:"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("p",{parentName:"li"},"Create a local key. Replace ",(0,a.kt)("inlineCode",{parentName:"p"},"<your-key-name>")," with whatever name you prefer."),(0,a.kt)("pre",{parentName:"li"},(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),"desmos keys add <your-key-name>\n")),(0,a.kt)("p",{parentName:"li"},"You will be required to input a password. Please make sure you use one that you will remember later. You should now see an output like"),(0,a.kt)("pre",{parentName:"li"},(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),'$ desmos keys add jack --dry-run\n\n- name: jack\n  type: local\n  address: desmos1qdv08q76fmfwwzrxcqs78z6pzfxe88cgc5a3tk\n  pubkey: desmospub1addwnpepq2j9a35spphh6q529y2thg8tjw9l2c32hck98fnmu99sxpw9a9aegugm6xs\n  mnemonic: ""\n  threshold: 0\n  pubkeys: []\n\n\n**Important** write this mnemonic phrase in a safe place.\nIt is the only way to recover your account if you ever forget your password.\n\nglory discover erosion mention grow prosper supreme term nephew venue pear eternal budget rely outdoor lobster strong sign space make soccer medal tuition patrol\n')),(0,a.kt)("p",{parentName:"li"},"Make sure you save the shown mnemonic phrase in some safe place as it might return useful in the future. ")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("p",{parentName:"li"},"Initialize the testnet"),(0,a.kt)("pre",{parentName:"li"},(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),"desmos init testnet --chain-id testnet\ndesmos add-genesis-account <your-key-name> 100000000000stake\ndesmos gentx <your-key-name> 1000000000stake --chain-id testnet\ndesmos collect-gentxs\n")),(0,a.kt)("p",{parentName:"li"},"During the procedure you will be asked to input the same key password you have set inside point 1. ")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("p",{parentName:"li"},"Start the testnet.",(0,a.kt)("br",{parentName:"p"}),"\n","Once you have completed all the steps, you are ready to start your local testnet by running: "),(0,a.kt)("pre",{parentName:"li"},(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),"desmos start\n")))),(0,a.kt)("h2",r({},{id:"creating-a-multi-node-testnet"}),"Creating a multi node testnet"),(0,a.kt)("p",null,"To create a local multi node testnet, you can simply run the following command: "),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),"make localnet-start\n")),(0,a.kt)("p",null,"This command creates a 4-node network using the ",(0,a.kt)("inlineCode",{parentName:"p"},"desmoslabs/desmosnode")," image. The ports for each node are found in this\ntable:"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",r({parentName:"tr"},{align:null}),"Node ID"),(0,a.kt)("th",r({parentName:"tr"},{align:null}),"P2P Port"),(0,a.kt)("th",r({parentName:"tr"},{align:null}),"RPC Port"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"desmosnode0")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26656")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26657"))),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"desmosnode1")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26659")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26660"))),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"desmosnode2")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26661")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26662"))),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"desmosnode3")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26663")),(0,a.kt)("td",r({parentName:"tr"},{align:null}),(0,a.kt)("inlineCode",{parentName:"td"},"26664"))))),(0,a.kt)("p",null,"To update the binary, just rebuild it and restart the nodes:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{}),"make build-linux localnet-start\n")),(0,a.kt)("h4",r({},{id:"configuration"}),"Configuration"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"make localnet-start")," creates files for a 4-node testnet in ",(0,a.kt)("inlineCode",{parentName:"p"},"./build")," by calling the ",(0,a.kt)("inlineCode",{parentName:"p"},"desmos testnet")," command. This outputs a handful of files in the ",(0,a.kt)("inlineCode",{parentName:"p"},"./build")," directory:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),"$ tree -L 2 build/\nbuild/\n\u251c\u2500\u2500 desmos\n\u251c\u2500\u2500 gentxs\n\u2502   \u251c\u2500\u2500 node0.json\n\u2502   \u251c\u2500\u2500 node1.json\n\u2502   \u251c\u2500\u2500 node2.json\n\u2502   \u2514\u2500\u2500 node3.json\n\u251c\u2500\u2500 node0\n\u2502   \u251c\u2500\u2500 desmos\n\u2502   \u2502   \u251c\u2500\u2500 key_seed.json\n\u2502   \u2502   \u2514\u2500\u2500 keys\n\u2502   \u2514\u2500\u2500 desmos\n\u2502       \u251c\u2500\u2500 ${LOG:-desmos.log}\n\u2502       \u251c\u2500\u2500 config\n\u2502       \u2514\u2500\u2500 data\n\u251c\u2500\u2500 node1\n\u2502   \u251c\u2500\u2500 desmos\n\u2502   \u2502   \u2514\u2500\u2500 key_seed.json\n\u2502   \u2514\u2500\u2500 desmos\n\u2502       \u251c\u2500\u2500 ${LOG:-desmos.log}\n\u2502       \u251c\u2500\u2500 config\n\u2502       \u2514\u2500\u2500 data\n\u251c\u2500\u2500 node2\n\u2502   \u251c\u2500\u2500 desmos\n\u2502   \u2502   \u2514\u2500\u2500 key_seed.json\n\u2502   \u2514\u2500\u2500 desmos\n\u2502       \u251c\u2500\u2500 ${LOG:-desmos.log}\n\u2502       \u251c\u2500\u2500 config\n\u2502       \u2514\u2500\u2500 data\n\u2514\u2500\u2500 node3\n    \u251c\u2500\u2500 desmos\n    \u2502   \u2514\u2500\u2500 key_seed.json\n    \u2514\u2500\u2500 desmos\n        \u251c\u2500\u2500 ${LOG:-desmos.log}\n        \u251c\u2500\u2500 config\n        \u2514\u2500\u2500 data\n")),(0,a.kt)("p",null,"Each ",(0,a.kt)("inlineCode",{parentName:"p"},"./build/nodeN")," directory is mounted to the ",(0,a.kt)("inlineCode",{parentName:"p"},"/desmos")," directory in each container."),(0,a.kt)("h4",r({},{id:"logging"}),"Logging"),(0,a.kt)("p",null,"Logs are saved under each ",(0,a.kt)("inlineCode",{parentName:"p"},"./build/nodeN/desmos/desmos.log"),". You can also watch logs directly via Docker, for example:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{}),"docker logs -f desmosnode0\n")),(0,a.kt)("h4",r({},{id:"keys--accounts"}),"Keys & Accounts"),(0,a.kt)("p",null,"To interact with ",(0,a.kt)("inlineCode",{parentName:"p"},"desmos")," and start querying state or creating txs, you use the\n",(0,a.kt)("inlineCode",{parentName:"p"},"desmos")," directory of any given node as your ",(0,a.kt)("inlineCode",{parentName:"p"},"home"),", for example:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{className:"language-bash"}),"desmos keys list --home ./build/node0/desmos\n")),(0,a.kt)("p",null,"Now that accounts exists, you may create new accounts and send those accounts funds!"),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Note"),": Each node's seed is located at ",(0,a.kt)("inlineCode",{parentName:"p"},"./build/nodeN/desmos/key_seed.json")," and can be restored to the CLI using the ",(0,a.kt)("inlineCode",{parentName:"p"},"desmos keys add --restore")," command"))}m.isMDXComponent=!0}}]);