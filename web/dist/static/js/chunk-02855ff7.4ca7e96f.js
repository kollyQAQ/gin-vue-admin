(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-02855ff7","chunk-2d0c0dd4"],{4431:function(e,t,r){"use strict";r.r(t),r.d(t,"getCollectionList",function(){return o}),r.d(t,"getCollectionListV2",function(){return n}),r.d(t,"getCollectionById",function(){return l}),r.d(t,"updateCollection",function(){return a}),r.d(t,"createCollection",function(){return s}),r.d(t,"deleteCollection",function(){return u});var i=r("b775"),o=function(e){return Object(i["default"])({url:"/zhihu/getCollectionList",method:"post",data:e})},n=function(e){return Object(i["default"])({url:"/zhihu/getCollectionListV2",method:"post",data:e})},l=function(e){return Object(i["default"])({url:"/zhihu/getCollectionById",method:"post",data:e})},a=function(e){return Object(i["default"])({url:"/zhihu/updateCollection",method:"post",data:e})},s=function(e){return Object(i["default"])({url:"/zhihu/createCollection",method:"post",data:e})},u=function(e){return Object(i["default"])({url:"/zhihu/deleteCollection",method:"post",data:e})}},"4f31":function(e,t,r){"use strict";r.r(t);var i=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",[r("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0}},[r("el-form-item",[r("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.openDialog("add")}}},[e._v("新增收藏")])],1)],1),e._v(" "),r("el-table",{staticStyle:{width:"100%"},attrs:{data:e.list,border:"",fit:"","highlight-current-row":""}},[r("el-table-column",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{align:"center",label:"ID",width:"65","element-loading-text":"请给我点时间！"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.ID))])]}}])}),e._v(" "),r("el-table-column",{attrs:{"min-width":"250px",label:"Title"},scopedSlots:e._u([{key:"default",fn:function(t){var i=t.row;return[r("a",{staticStyle:{color:"#606266"},attrs:{href:i.url,target:"_blank"}},[e._v("\n          "+e._s(i.title)+"\n        ")])]}}])}),e._v(" "),r("el-table-column",{attrs:{width:"150px",align:"center",label:"Author"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.author))])]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"优先级","min-width":"80",prop:"priority"}}),e._v(" "),r("el-table-column",{attrs:{fixed:"right",label:"操作",width:"200"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{size:"small",type:"primary",icon:"el-icon-edit"},on:{click:function(r){return e.editApi(t.row)}}},[e._v("编辑")]),e._v(" "),r("el-button",{attrs:{size:"small",type:"danger",icon:"el-icon-delete"},on:{click:function(r){return e.deleteApi(t.row)}}},[e._v("删除")])]}}])})],1),e._v(" "),r("el-dialog",{attrs:{"before-close":e.closeDialog,title:e.dialogTitle,visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[r("el-form",{ref:"apiForm",attrs:{inline:!0,model:e.form,rules:e.rules,"label-width":"80px"}},[r("el-form-item",{attrs:{label:"类型",prop:"type"}},[r("el-select",{attrs:{placeholder:"请选择"},model:{value:e.form.type,callback:function(t){e.$set(e.form,"type",t)},expression:"form.type"}},e._l(e.typeOptions,function(e){return r("el-option",{key:e.value,attrs:{label:""+e.label,value:e.value}})}),1)],1),e._v(" "),r("el-form-item",{attrs:{label:"内容"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.title,callback:function(t){e.$set(e.form,"title",t)},expression:"form.title"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"链接"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.url,callback:function(t){e.$set(e.form,"url",t)},expression:"form.url"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"作者"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.author,callback:function(t){e.$set(e.form,"author",t)},expression:"form.author"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"优先级"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.priority,callback:function(t){e.$set(e.form,"priority",t)},expression:"form.priority"}})],1)],1),e._v(" "),r("div",{staticClass:"edit_container"},[r("quill-editor",{ref:"myQuillEditor",attrs:{options:{}},model:{value:e.form.remark,callback:function(t){e.$set(e.form,"remark",t)},expression:"form.remark"}})],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:e.closeDialog}},[e._v("取 消")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:e.enterDialog}},[e._v("确 定")])],1)],1)],1)},o=[],n=(r("96cf"),r("3b8d")),l=r("4431"),a={props:{type:{type:String,default:"文章"}},data:function(){return{list:null,listQuery:{type:this.type,sort:"+id"},loading:!1,dialogFormVisible:!1,dialogTitle:"新增收藏",form:{type:this.type,title:"",url:"",author:"",priority:"0",remark:""},typeOptions:[{value:"用户",label:"用户"},{value:"专栏",label:"专栏"},{value:"文章",label:"文章"},{value:"回答",label:"回答"},{value:"话术",label:"话术"}],rules:{type:[{required:!0,message:"请输入type",trigger:"blur"}],title:[{required:!0,message:"请输入title",trigger:"blur"}],url:[{required:!1,message:"请选择url",trigger:"blur"}],author:[{required:!1,message:"请输入author",trigger:"blur"}]}}},created:function(){this.getList()},methods:{getList:function(){var e=this;this.loading=!0,this.$emit("create"),Object(l["getCollectionListV2"])(this.listQuery).then(function(t){e.list=t.data.list,e.loading=!1})},initForm:function(){this.$refs.apiForm.resetFields(),this.form={type:this.type,title:"",url:"",author:"",priority:"0",remark:""}},closeDialog:function(){this.initForm(),this.dialogFormVisible=!1},openDialog:function(e){switch(e){case"add":this.dialogTitle="新增收藏";break;case"edit":this.dialogTitle="编辑收藏";break;default:break}this.dialogType=e,this.dialogFormVisible=!0},editApi:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(t){var r;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,Object(l["getCollectionById"])({id:t.ID});case 2:r=e.sent,this.form=r.data.collection,this.openDialog("edit");case 5:case"end":return e.stop()}},e,this)}));function t(t){return e.apply(this,arguments)}return t}(),deleteApi:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(t){var r=this;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:this.$confirm("此操作将删除此收藏, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(Object(n["a"])(regeneratorRuntime.mark(function e(){var i;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,Object(l["deleteCollection"])(t);case 2:i=e.sent,0==i.code&&(r.$message({type:"success",message:"删除成功!"}),r.getList());case 4:case"end":return e.stop()}},e)}))).catch(function(){r.$message({type:"info",message:"已取消删除"})});case 1:case"end":return e.stop()}},e,this)}));function t(t){return e.apply(this,arguments)}return t}(),enterDialog:function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:this.$refs.apiForm.validate(function(){var e=Object(n["a"])(regeneratorRuntime.mark(function e(r){var i,o;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:if(!r){e.next=20;break}e.t0=t.dialogType,e.next="add"===e.t0?4:"edit"===e.t0?11:18;break;case 4:return e.next=6,Object(l["createCollection"])(t.form);case 6:return i=e.sent,0==i.code&&t.$message({type:"success",message:"添加成功",showClose:!0}),t.getList(),t.closeDialog(),e.abrupt("break",20);case 11:return e.next=13,Object(l["updateCollection"])(t.form);case 13:return o=e.sent,0==o.code&&t.$message({type:"success",message:"编辑成功",showClose:!0}),t.getList(),t.closeDialog(),e.abrupt("break",20);case 18:return t.$message({type:"error",message:"未知操作",showClose:!0}),e.abrupt("break",20);case 20:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}());case 1:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()}},s=a,u=r("2877"),c=Object(u["a"])(s,i,o,!1,null,null,null);t["default"]=c.exports}}]);