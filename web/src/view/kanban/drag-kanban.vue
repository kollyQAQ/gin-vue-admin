<template>
  <div class="components-container board">
    <Kanban :key="1" :list="inbox" :group="group" class="kanban inbox" header-text="Inbox" />
    <Kanban :key="2" :list="todo" :group="group" class="kanban todo" header-text="Todo" />
    <Kanban :key="3" :list="doing" :group="group" class="kanban working" header-text="Working" />
    <Kanban :key="4" :list="done" :group="group" class="kanban done" header-text="Done" />
  </div>
</template>
<script>
  import Kanban from '@/components/Kanban'
  import {
    queryPlanList,
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
        done: []
      }
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

