(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-b5be7064","chunk-2d0aa631"],{1194:function(e,r,t){e.exports=t.p+"static/img/logo_login.7918108b.png"},"322f":function(e,r,t){},"3d01":function(e,r,t){"use strict";t.r(r);var s=function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("div",{staticClass:"user-layout-wrapper",attrs:{id:"userLayout"}},[t("div",{staticClass:"container"},[e._m(0),e._v(" "),t("div",{staticClass:"main"},[t("el-form",{ref:"registerForm",attrs:{model:e.registerForm,rules:e.rules},nativeOn:{keyup:function(r){return!r.type.indexOf("key")&&e._k(r.keyCode,"enter",13,r.key,"Enter")?null:e.submitForm(r)}}},[t("el-form-item",{attrs:{prop:"username"}},[t("el-input",{attrs:{placeholder:"请输入用户名"},model:{value:e.registerForm.username,callback:function(r){e.$set(e.registerForm,"username",r)},expression:"registerForm.username"}},[t("i",{staticClass:"el-input__icon el-icon-user",attrs:{slot:"suffix"},slot:"suffix"})])],1),e._v(" "),t("el-form-item",{attrs:{prop:"password"}},[t("el-input",{attrs:{type:"lock"===e.lock?"password":"text",placeholder:"请输入密码"},model:{value:e.registerForm.password,callback:function(r){e.$set(e.registerForm,"password",r)},expression:"registerForm.password"}},[t("i",{class:"el-input__icon el-icon-"+e.lock,attrs:{slot:"suffix"},on:{click:e.changeLock},slot:"suffix"})])],1),e._v(" "),t("el-form-item",{staticStyle:{position:"relative"},attrs:{prop:"rePassword"}},[t("el-input",{attrs:{type:"lock"===e.lock?"password":"text",placeholder:"请再次输入密码"},model:{value:e.registerForm.rePassword,callback:function(r){e.$set(e.registerForm,"rePassword",r)},expression:"registerForm.rePassword"}},[t("i",{class:"el-input__icon el-icon-"+e.lock,attrs:{slot:"suffix"},on:{click:e.changeLock},slot:"suffix"})])],1),e._v(" "),t("el-form-item",[t("el-button",{staticStyle:{width:"100%"},attrs:{type:"primary"},on:{click:e.submitForm}},[e._v("注 册")])],1)],1)],1)])])},n=[function(){var e=this,r=e.$createElement,s=e._self._c||r;return s("div",{staticClass:"top"},[s("div",{staticClass:"desc"},[s("img",{staticClass:"logo_login",attrs:{src:t("1194"),alt:""}})]),e._v(" "),s("div",{staticClass:"header"},[s("a",{attrs:{href:"/"}},[s("span",{staticClass:"title"},[e._v("Zhihu-Data")])])])])}],o=(t("8e6e"),t("ac6a"),t("456d"),t("96cf"),t("3b8d")),i=t("bd86"),a=t("2f62"),c=t("c24f");function l(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);r&&(s=s.filter(function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable})),t.push.apply(t,s)}return t}function u(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?l(t,!0).forEach(function(r){Object(i["a"])(e,r,t[r])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):l(t).forEach(function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))})}return e}var p={name:"Register",data:function(){var e=this,r=function(r,t,s){if(t!=e.registerForm.password)return s(new Error("两次密码不同"));s()},t=function(e,r,t){if(r.length<5||r.length>12)return t(new Error("请输入正确的用户名"));t()},s=function(e,r,t){if(r.length<6||r.length>12)return t(new Error("请输入正确的密码"));t()};return{lock:"lock",registerForm:{username:"",password:"",rePassword:""},rules:{username:[{validator:t,trigger:"blur"}],password:[{validator:s,trigger:"blur"}],rePassword:[{validator:r,trigger:"blur"}]}}},methods:u({},Object(a["b"])("user",["LoginIn"]),{submitForm:function(){var e=Object(o["a"])(regeneratorRuntime.mark(function e(){var r=this;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:this.$refs.registerForm.validate(function(){var e=Object(o["a"])(regeneratorRuntime.mark(function e(t){var s;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:if(!t){e.next=7;break}return e.next=3,Object(c["register"])(r.registerForm);case 3:s=e.sent,0==s.code&&(r.$message({type:"success",message:"注册成功",showClose:!0}),r.$router.push({name:"login"})),e.next=9;break;case 7:return r.$message({type:"error",message:"请正确填写注册信息",showClose:!0}),e.abrupt("return",!1);case 9:case"end":return e.stop()}},e)}));return function(r){return e.apply(this,arguments)}}());case 1:case"end":return e.stop()}},e,this)}));function r(){return e.apply(this,arguments)}return r}(),changeLock:function(){"lock"===this.lock?this.lock="unlock":this.lock="lock"}})},f=p,m=(t("45bc"),t("2877")),g=Object(m["a"])(f,s,n,!1,null,"0c82f0ce",null);r["default"]=g.exports},"45bc":function(e,r,t){"use strict";var s=t("322f"),n=t.n(s);n.a}}]);