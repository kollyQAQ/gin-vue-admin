<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  components: {}
}
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
}

/*滚动条整体样式*/
::-webkit-scrollbar{
  width: 8px;/*竖向滚动条的宽度*/
  height: 8px;/*横向滚动条的高度*/
}
::-webkit-scrollbar-thumb{/*滚动条里面的小方块*/
  background: #666666;
  border-radius: 5px;
}
::-webkit-scrollbar-track{/*滚动条轨道的样式*/
  background: #ccc;
  border-radius: 5px;
}

.el-dialog {
  max-height: calc(100vh - 20vh);
  overflow: hidden;
}

.el-dialog__body {
  height: calc(100vh - 20vh - 200px);
  overflow-y: scroll;
}
</style>
