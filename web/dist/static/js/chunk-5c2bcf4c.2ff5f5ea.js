(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5c2bcf4c","chunk-5c0bd654","chunk-74d08cbf","chunk-2d216b1c","chunk-2d0a4474"],{"0680":function(e,t,r){"use strict";r.r(t),r.d(t,"getDict",function(){return o});r("96cf");var a=r("3b8d"),n=r("4360"),o=function(){var e=Object(a["a"])(regeneratorRuntime.mark(function e(t){return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,n["store"].dispatch("dictionary/getDictionary",t);case 2:return e.abrupt("return",n["store"].getters["dictionary/getDictionary"][t]);case 3:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}()},"10df":function(e,t,r){"use strict";var a=r("1eab"),n=r.n(a);n.a},"1eab":function(e,t,r){},"2a18":function(e,t,r){"use strict";r.r(t),r.d(t,"toUpperCase",function(){return a}),r.d(t,"toSQLLine",function(){return n}),r.d(t,"toHump",function(){return o});r("a481");var a=function(e){return e[0]?e.replace(e[0],e[0].toUpperCase()):""},n=function(e){return"ID"==e?"ID":e.replace(/([A-Z])/g,"_$1").toLowerCase()},o=function(e){return e.replace(/\_(\w)/g,function(e,t){return t.toUpperCase()})}},a481:function(e,t,r){"use strict";var a=r("cb7c"),n=r("4bf8"),o=r("9def"),i=r("4588"),s=r("0390"),c=r("5f1b"),u=Math.max,l=Math.min,d=Math.floor,f=/\$([$&`']|\d\d?|<[^>]*>)/g,h=/\$([$&`']|\d\d?)/g,p=function(e){return void 0===e?e:String(e)};r("214f")("replace",2,function(e,t,r,g){return[function(a,n){var o=e(this),i=void 0==a?void 0:a[t];return void 0!==i?i.call(a,o,n):r.call(String(o),a,n)},function(e,t){var n=g(r,e,this,t);if(n.done)return n.value;var d=a(e),f=String(this),h="function"===typeof t;h||(t=String(t));var m=d.global;if(m){var v=d.unicode;d.lastIndex=0}var _=[];while(1){var y=c(d,f);if(null===y)break;if(_.push(y),!m)break;var w=String(y[0]);""===w&&(d.lastIndex=s(f,o(d.lastIndex),v))}for(var k="",O=0,j=0;j<_.length;j++){y=_[j];for(var S=String(y[0]),I=u(l(i(y.index),f.length),0),z=[],x=1;x<y.length;x++)z.push(p(y[x]));var C=y.groups;if(h){var D=[S].concat(z,I,f);void 0!==C&&D.push(C);var G=String(t.apply(void 0,D))}else G=b(S,f,I,z,C,t);I>=O&&(k+=f.slice(O,I)+G,O=I+S.length)}return k+f.slice(O)}];function b(e,t,a,o,i,s){var c=a+e.length,u=o.length,l=h;return void 0!==i&&(i=n(i),l=f),r.call(s,l,function(r,n){var s;switch(n.charAt(0)){case"$":return"$";case"&":return e;case"`":return t.slice(0,a);case"'":return t.slice(c);case"<":s=i[n.slice(1,-1)];break;default:var l=+n;if(0===l)return r;if(l>u){var f=d(l/10);return 0===f?r:f<=u?void 0===o[f-1]?n.charAt(1):o[f-1]+n.charAt(1):r}s=o[l-1]}return void 0===s?"":s})}})},b016:function(e,t,r){"use strict";r.r(t);r("8e6e"),r("ac6a"),r("456d");var a=r("bd86"),n=(r("96cf"),r("3b8d")),o=r("0680");function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter(function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable})),r.push.apply(r,a)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(r,!0).forEach(function(t){Object(a["a"])(e,t,r[t])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(r).forEach(function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))})}return e}t["default"]={data:function(){return{page:1,total:20,pageSize:20,tableData:[],searchInfo:{}}},methods:{filterDict:function(e,t){var r=this[t+"Options"]&&this[t+"Options"].filter(function(t){return t.value==e});return r&&r[0]&&r[0].label},getDict:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(t){var r;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,Object(o["getDict"])(t);case 2:r=e.sent,this[t+"Options"]=r;case 4:case"end":return e.stop()}},e,this)}));function t(t){return e.apply(this,arguments)}return t}(),handleSizeChange:function(e){this.pageSize=e,this.getTableData()},handleCurrentChange:function(e){this.page=e,this.getTableData()},getTableData:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(){var t,r,a,n=arguments;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return t=n.length>0&&void 0!==n[0]?n[0]:this.page,r=n.length>1&&void 0!==n[1]?n[1]:this.pageSize,e.next=4,this.listApi(s({page:t,pageSize:r},this.searchInfo));case 4:a=e.sent,this.tableData=a.data.list,this.total=a.data.total,this.page=a.data.page,this.pageSize=a.data.pageSize;case 9:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()}}},c40e:function(e,t,r){"use strict";r.r(t),r.d(t,"getHighRateGoodsList",function(){return n}),r.d(t,"getFirstCategory",function(){return o}),r.d(t,"getSecondCategory",function(){return i}),r.d(t,"getThirdCategory",function(){return s}),r.d(t,"getGoodsList",function(){return c}),r.d(t,"getGoodsCategory",function(){return u}),r.d(t,"getGoodsById",function(){return l}),r.d(t,"createGoods",function(){return d}),r.d(t,"updateGoods",function(){return f}),r.d(t,"deleteGoods",function(){return h});var a=r("b775"),n=function(e){return Object(a["default"])({url:"/zhihu/getHighRateGoodsList",method:"post",data:e})},o=function(e){return Object(a["default"])({url:"/zhihu/getGoodsCategory",method:"post",data:e})},i=function(e){return Object(a["default"])({url:"/zhihu/getGoodsCategory",method:"post",data:e})},s=function(e){return Object(a["default"])({url:"/zhihu/getGoodsCategory",method:"post",data:e})},c=function(e){return Object(a["default"])({url:"/zhihu/getGoodsList",method:"post",data:e})},u=function(e){return Object(a["default"])({url:"/zhihu/getGoodsCategory",method:"post",data:e})},l=function(e){return Object(a["default"])({url:"/zhihu/getGoodsById",method:"post",data:e})},d=function(e){return Object(a["default"])({url:"/zhihu/createGoods",method:"post",data:e})},f=function(e){return Object(a["default"])({url:"/zhihu/updateGoods",method:"post",data:e})},h=function(e){return Object(a["default"])({url:"/zhihu/deleteGoods",method:"post",data:e})}},fea9:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",[r("div",{staticClass:"search-term"},[r("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:e.searchInfo}},[r("el-form-item",{attrs:{label:"sku"}},[r("el-input",{attrs:{placeholder:"sku"},model:{value:e.searchInfo.sku_id,callback:function(t){e.$set(e.searchInfo,"sku_id",t)},expression:"searchInfo.sku_id"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"商品名"}},[r("el-input",{attrs:{placeholder:"商品名"},model:{value:e.searchInfo.sku_name,callback:function(t){e.$set(e.searchInfo,"sku_name",t)},expression:"searchInfo.sku_name"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"一级类目"}},[r("el-select",{attrs:{clearable:"",placeholder:"请选择"},model:{value:e.searchInfo.category_one,callback:function(t){e.$set(e.searchInfo,"category_one",t)},expression:"searchInfo.category_one"}},e._l(e.methodOptions,function(e){return r("el-option",{key:e.value,attrs:{label:""+e.label,value:e.value}})}),1)],1),e._v(" "),r("el-form-item",{attrs:{label:"二级类目"}},[r("el-select",{attrs:{clearable:"",placeholder:"请选择"},model:{value:e.searchInfo.category_two,callback:function(t){e.$set(e.searchInfo,"category_two",t)},expression:"searchInfo.category_two"}},e._l(e.methodOptions,function(e){return r("el-option",{key:e.value,attrs:{label:""+e.label,value:e.value}})}),1)],1),e._v(" "),r("el-form-item",{attrs:{label:"三级类目"}},[r("el-select",{attrs:{clearable:"",placeholder:"请选择"},model:{value:e.searchInfo.category_three,callback:function(t){e.$set(e.searchInfo,"category_three",t)},expression:"searchInfo.category_three"}},e._l(e.methodOptions,function(e){return r("el-option",{key:e.value,attrs:{label:""+e.label,value:e.value}})}),1)],1),e._v(" "),r("el-form-item",[r("el-button",{attrs:{type:"primary"},on:{click:e.onSubmit}},[e._v("查询")])],1)],1)],1),e._v(" "),r("el-table",{attrs:{data:e.tableData,border:"",stripe:""},on:{"sort-change":e.sortChange}},[r("el-table-column",{attrs:{label:"id","min-width":"50",prop:"ID",sortable:"custom"}}),e._v(" "),r("el-table-column",{attrs:{label:"sku","min-width":"80",prop:"sku_id",sortable:"custom"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("a",{staticStyle:{color:"#606266"},attrs:{href:"http://item.jd.com/"+t.row.sku_id+".html",target:"_blank"}},[e._v("\n            "+e._s(t.row.sku_id)+"\n          ")])]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"商品名","min-width":"200",prop:"goods_name",sortable:"custom"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("a",{staticStyle:{color:"#606266"},attrs:{href:t.row.url,target:"_blank"}},[e._v("\n            "+e._s(t.row.goods_name)+"\n          ")])]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"价格","min-width":"50",prop:"price",sortable:"custom"}}),e._v(" "),r("el-table-column",{attrs:{label:"佣金比","min-width":"50",prop:"fee_rate",sortable:"custom"},scopedSlots:e._u([{key:"default",fn:function(t){return[t.row.fee_rate>=20?r("p",{staticStyle:{"font-size":"16px",color:"red"}},[e._v(e._s(t.row.fee_rate))]):t.row.fee_rate>=10?r("p",{staticStyle:{"font-size":"16px",color:"orange"}},[e._v(e._s(t.row.fee_rate))]):r("p",[e._v(e._s(t.row.fee_rate))])]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"佣金","min-width":"50",prop:"fee",sortable:"custom"}}),e._v(" "),r("el-table-column",{attrs:{label:"一级类目","min-width":"50",prop:"category_one",sortable:"custom"}}),e._v(" "),r("el-table-column",{attrs:{label:"二级类目","min-width":"50",prop:"category_two",sortable:"custom"}}),e._v(" "),r("el-table-column",{attrs:{label:"三级类目","min-width":"50",prop:"category_three",sortable:"custom"}}),e._v(" "),r("el-table-column",{attrs:{label:"是否自营","min-width":"50",prop:"jd_sale",sortable:"custom"},scopedSlots:e._u([{key:"default",fn:function(t){return[1==t.row.jd_sale?r("p",[r("i",{staticClass:"el-icon-success",staticStyle:{"font-size":"20px",color:"green"}})]):e._e(),e._v(" "),0==t.row.jd_sale?r("p",[r("i",{staticClass:"el-icon-error"})]):e._e()]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"最后更新时间","min-width":"100",prop:"update_time"}})],1),e._v(" "),r("el-pagination",{style:{float:"right",padding:"20px"},attrs:{"current-page":e.page,"page-size":e.pageSize,"page-sizes":[20,50,100],total:e.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":e.handleCurrentChange,"size-change":e.handleSizeChange}})],1)},n=[],o=(r("96cf"),r("3b8d")),i=r("c40e"),s=r("b016"),c=r("2a18"),u={name:"Goods",mixins:[s["default"]],data:function(){return{listApi:i["getHighRateGoodsList"],methodOptions:[],type:"",rules:{sku_id:[{required:!0,message:"请输入sku_id",trigger:"blur"}],sku_name:[{required:!1,message:"请输入商品名称",trigger:"blur"}]}}},methods:{sortChange:function(e){var t=e.prop,r=e.order;t&&(this.searchInfo.orderKey=Object(c["toSQLLine"])(t),this.searchInfo.desc="descending"==r),this.getTableData()},onSubmit:function(){this.page=1,this.pageSize=20,this.getTableData()}},created:function(){var e=Object(o["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return this.getTableData(),e.next=3,Object(i["getThirdCategory"])();case 3:t=e.sent,this.methodOptions=t.data;case 5:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()},l=u,d=(r("10df"),r("2877")),f=Object(d["a"])(l,a,n,!1,null,"155cdec8",null);t["default"]=f.exports}}]);