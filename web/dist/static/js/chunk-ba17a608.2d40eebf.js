(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ba17a608"],{"0a49":function(t,e,s){var i=s("9b43"),n=s("626a"),r=s("4bf8"),o=s("9def"),a=s("cd1c");t.exports=function(t,e){var s=1==t,h=2==t,u=3==t,c=4==t,l=6==t,f=5==t||l,m=e||a;return function(e,a,y){for(var d,v,p=r(e),b=n(p),g=i(a,y,3),x=o(b.length),$=0,A=s?m(e,x):h?m(e,0):void 0;x>$;$++)if((f||$ in b)&&(d=b[$],v=g(d,$,p),t))if(s)A[$]=v;else if(v)switch(t){case 3:return!0;case 5:return d;case 6:return $;case 2:A.push(d)}else if(c)return!1;return l?-1:u||c?c:A}}},1169:function(t,e,s){var i=s("2d95");t.exports=Array.isArray||function(t){return"Array"==i(t)}},"20d6":function(t,e,s){"use strict";var i=s("5ca1"),n=s("0a49")(6),r="findIndex",o=!0;r in[]&&Array(1)[r](function(){o=!1}),i(i.P+i.F*o,"Array",{findIndex:function(t){return n(this,t,arguments.length>1?arguments[1]:void 0)}}),s("9c6c")(r)},"4bde":function(t,e,s){"use strict";var i=s("f198"),n=s.n(i);n.a},"586c":function(t,e,s){"use strict";s.r(e);var i=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"router-history"},[s("el-tabs",{attrs:{closable:!(1==t.historys.length&&"dashboard"==this.$route.name),type:"card"},on:{"tab-click":t.changeTab,"tab-remove":t.removeTab},nativeOn:{contextmenu:function(e){return e.preventDefault(),t.openContextMenu(e)}},model:{value:t.activeValue,callback:function(e){t.activeValue=e},expression:"activeValue"}},t._l(t.historys,function(t){return s("el-tab-pane",{key:t.name,attrs:{label:t.meta.title,name:t.name,tab:t}})}),1),t._v(" "),s("ul",{directives:[{name:"show",rawName:"v-show",value:t.contextMenuVisible,expression:"contextMenuVisible"}],staticClass:"contextmenu",style:{left:t.left+"px",top:t.top+"px"}},[s("li",{on:{click:t.closeAll}},[t._v("关闭所有")]),t._v(" "),s("li",{on:{click:t.closeLeft}},[t._v("关闭左侧")]),t._v(" "),s("li",{on:{click:t.closeRight}},[t._v("关闭右侧")]),t._v(" "),s("li",{on:{click:t.closeOther}},[t._v("关闭其他")])])],1)},n=[],r=(s("20d6"),s("28a5"),s("7f7f"),{name:"HistoryComponent",data:function(){return{historys:[],activeValue:"dashboard",contextMenuVisible:!1,left:0,top:0,isCollapse:!1,isMobile:!1,rightActive:""}},created:function(){var t=this;this.$bus.on("mobile",function(e){t.isMobile=e}),this.$bus.on("collapse",function(e){t.isCollapse=e});var e=[{name:"dashboard",meta:{title:"仪表盘"}}];this.historys=JSON.parse(sessionStorage.getItem("historys"))||e,this.setTab(this.$route)},beforeDestroy:function(){this.$bus.off("collapse"),this.$bus.off("mobile")},methods:{openContextMenu:function(t){if(1==this.historys.length&&"dashboard"==this.$route.name)return!1;var e;t.srcElement.id&&(this.contextMenuVisible=!0,e=this.isCollapse?54:220,this.isMobile&&(e=0),this.left=t.clientX-e,this.top=t.clientY+10,this.rightActive=t.srcElement.id.split("-")[1])},closeAll:function(){this.historys=[{name:"dashboard",meta:{title:"仪表盘"}}],this.$router.push({name:"dashboard"}),this.contextMenuVisible=!1,sessionStorage.setItem("historys",JSON.stringify(this.historys))},closeLeft:function(){var t,e=this,s=this.historys.findIndex(function(s){return s.name==e.rightActive&&(t=s),s.name==e.rightActive}),i=this.historys.findIndex(function(t){return t.name==e.activeValue});this.historys.splice(0,s),s>i&&this.$router.push(t),sessionStorage.setItem("historys",JSON.stringify(this.historys))},closeRight:function(){var t,e=this,s=this.historys.findIndex(function(s){return s.name==e.rightActive&&(t=s),s.name==e.rightActive}),i=this.historys.findIndex(function(t){return t.name==e.activeValue});this.historys.splice(s+1,this.historys.length),s<i&&this.$router.push(t),sessionStorage.setItem("historys",JSON.stringify(this.historys))},closeOther:function(){var t,e=this;this.historys=this.historys.filter(function(s){return s.name==e.rightActive&&(t=s),s.name==e.rightActive}),this.$router.push(t),sessionStorage.setItem("historys",JSON.stringify(this.historys))},setTab:function(t){if(!this.historys.some(function(e){return e.name==t.name})){var e={};e.name=t.name,e.meta=t.meta,e.query=t.query,e.params=t.params,this.historys.push(e)}this.activeValue=this.$route.name},changeTab:function(t){var e=t.$attrs.tab;this.$router.push({name:e.name,query:e.query,params:e.params})},removeTab:function(t){var e=this.historys.findIndex(function(e){return e.name==t});this.$route.name==t&&(1==this.historys.length?this.$router.push({name:"dashboard"}):e<this.historys.length-1?this.$router.push({name:this.historys[e+1].name,query:this.historys[e+1].query,params:this.historys[e+1].params}):this.$router.push({name:this.historys[e-1].name,query:this.historys[e-1].query,params:this.historys[e-1].params})),this.historys.splice(e,1)}},watch:{contextMenuVisible:function(){var t=this;this.contextMenuVisible?document.body.addEventListener("click",function(){t.contextMenuVisible=!1}):document.body.removeEventListener("click",function(){t.contextMenuVisible=!1})},$route:function(t){this.historys=this.historys.filter(function(t){return!t.meta.hidden}),this.setTab(t),sessionStorage.setItem("historys",JSON.stringify(this.historys))}}}),o=r,a=(s("4bde"),s("2877")),h=Object(a["a"])(o,i,n,!1,null,null,null);e["default"]=h.exports},cd1c:function(t,e,s){var i=s("e853");t.exports=function(t,e){return new(i(t))(e)}},e853:function(t,e,s){var i=s("d3f4"),n=s("1169"),r=s("2b4c")("species");t.exports=function(t){var e;return n(t)&&(e=t.constructor,"function"!=typeof e||e!==Array&&!n(e.prototype)||(e=void 0),i(e)&&(e=e[r],null===e&&(e=void 0))),void 0===e?Array:e}},f198:function(t,e,s){}}]);