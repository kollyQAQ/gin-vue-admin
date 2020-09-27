<template>
  <div class="board-column">
    <div class="board-column-header">
      {{ headerText }}
    </div>
    <draggable
      :list="list"
      v-bind="$attrs"
      class="board-column-content"
      :set-data="setData"
    >
      <div v-for="element in list" :key="element.id" class="board-item">
        <p v-if="element.level===2" style="color:blue;">{{ element.name }}</p>
        <p v-else-if="element.level===3" style="color:red;">{{ element.name }}</p>
        <p v-else>{{ element.name }}</p>
      </div>
    </draggable>
  </div>
</template>

<script>
  import draggable from 'vuedraggable'
  import {
    updatePlan,
  } from '@/api/stat'

  export default {
    name: 'DragKanbanDemo',
    components: {
      draggable
    },
    props: {
      headerText: {
        type: String,
        default: 'Header'
      },
      options: {
        type: Object,
        default() {
          return {}
        }
      },
      status: {
        type: String,
        default: '0'
      },
      list: {
        type: Array,
        default() {
          return []
        }
      }
    },
    async updated() {
      console.log(this.headerText, this.status)
      console.log(this.list)
      const res = await updatePlan({"status":this.status, "data":this.list})
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '更新成功',
          showClose: true
        })
      }
    },
    methods: {
      setData(dataTransfer) {
        // to avoid Firefox bug
        // Detail see : https://github.com/RubaXa/Sortable/issues/1012
        dataTransfer.setData('Text', '')
      }
    }
  }
</script>
<style lang="scss" scoped>
.board-column {
  min-width: 280px;
  min-height: 100px;
  height: auto;
  overflow: hidden;
  background: #f0f0f0;
  border-radius: 3px;

  .board-column-header {
    height: 50px;
    line-height: 50px;
    overflow: hidden;
    padding: 0 20px;
    text-align: center;
    background: #333;
    color: #fff;
    border-radius: 3px 3px 0 0;
  }

  .board-column-content {
    height: auto;
    overflow: hidden;
    border: 10px solid transparent;
    min-height: 60px;
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
    align-items: center;

    .board-item {
      cursor: pointer;
      width: 100%;
      height: 64px;
      margin: 5px 0;
      background-color: #fff;
      text-align: left;
      line-height: 54px;
      padding: 5px 10px;
      box-sizing: border-box;
      box-shadow: 0px 1px 3px 0 rgba(0, 0, 0, 0.2);
    }
  }
}
</style>

