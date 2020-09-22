<template>
  <div>
    <div class="tab-container">
      <el-tag>mounted times ：{{ createdTimes }}</el-tag>
      <el-alert :closable="false" style="width:200px;display:inline-block;vertical-align: middle;margin-left:30px;" title="Tab with keep-alive" type="success" />
      <el-tabs v-model="activeName" style="margin-top:15px;" type="border-card">
          <el-tab-pane v-for="item in tabMapOptions" :key="item.key" :label="item.label" :name="item.key">
              <keep-alive>
                  <tab-pane v-if="activeName==item.key" :type="item.key" @create="showCreatedTimes" />
              </keep-alive>
          </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
import TabPane from './components/TabPane'

export default {
  name: 'Collection',
    components: { TabPane },
    data() {
    return {
        tabMapOptions: [
            { label: '文章', key: '文章' },
            { label: '话术', key: '话术' },
            { label: '用户', key: '用户' },
            { label: '专栏', key: '专栏' },
            { label: '回答', key: '回答' },
            { label: '工具', key: '工具' }
        ],
        activeName: '文章',
        createdTimes: 0,
    }
  },
    watch: {
        activeName(val) {
            this.$router.push(`${this.$route.path}?tab=${val}`)
        }
    },
  methods: {
      showCreatedTimes() {
          this.createdTimes = this.createdTimes + 1
      }
  },
  created(){
    this.getTableData()
      // init the default selected tab
      const tab = this.$route.query.tab
      if (tab) {
          this.activeName = tab
      }
  }
}
</script>
<style scoped lang="scss">
    .button-box {
      padding: 10px 20px;
      .el-button {
        float: right;
      }
    }
    .el-tag--mini {
      margin-left: 5px;
    }
    .warning {
      color: #dc143c;
    }
    /*.tab-container {*/
    /*    margin: 30px;*/
    /*}*/
</style>
