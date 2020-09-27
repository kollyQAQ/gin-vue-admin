<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input placeholder="输入任务" v-model="newPlan" style="width: 400px;"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="addPlan()" type="primary">新增任务</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="components-container board">
      <Kanban :key="1" :list="inbox" :group="group" class="kanban inbox" header-text="Inbox" status="0"/>
      <Kanban :key="2" :list="todo" :group="group" class="kanban todo" header-text="Todo" status="1"/>
      <Kanban :key="3" :list="doing" :group="group" class="kanban working" header-text="Working" status="2"/>
      <Kanban :key="4" :list="done" :group="group" class="kanban done" header-text="Done" status="3"/>
    </div>
  </div>
</template>
<script>
  import Kanban from '@/components/Kanban'
  import {
    queryPlanList,
    addPlan,
  } from '@/api/stat'

  export default {
    name: 'DragKanban',
    components: {
      Kanban
    },
    data() {
      return {
        group: 'mission',
        inbox: [],
        todo: [],
        doing: [],
        done: [],
        newPlan: ''
      }
    },
    methods: {
      async addPlan() {
        if (this.newPlan === "") {
          this.$message({
            type: 'danget',
            message: '请输入内容',
            showClose: true
          })
          return
        }
        const res = await addPlan({"name":this.newPlan})
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '添加成功',
            showClose: true
          })
        }
        const res1 = await queryPlanList();
        this.inbox = res1.data.inbox;
      },
    },
    async created(){
      const res = await queryPlanList();
      this.inbox = res.data.inbox;
      this.todo = res.data.todo;
      this.doing = res.data.doing;
      this.done = res.data.done;
    }
  }
</script>
<style lang="scss">
.board {
  width: 1100px;
  margin-left: 20px;
  display: flex;
  justify-content: space-around;
  flex-direction: row;
  align-items: flex-start;
}
.kanban {
  &.inbox {
  .board-column-header {
    background: grey;
    }
  }
  &.todo {
    .board-column-header {
      background: #4A9FF9;
    }
  }
  &.working {
    .board-column-header {
      background: #f9944a;
    }
  }
  &.done {
    .board-column-header {
      background: #2ac06d;
    }
  }
}
</style>

