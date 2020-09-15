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
    name: 'DragKanbanDemo',
    components: {
      Kanban
    },
    data() {
      return {
        group: 'mission',
        inbox: [
          { name: 'Mission', id: 1 },
          { name: 'Mission', id: 2 },
          { name: 'Mission', id: 3 },
          { name: 'Mission', id: 4 }
        ],
        todo: [
          { name: 'Mission', id: 5 },
          { name: 'Mission', id: 6 },
          { name: 'Mission', id: 7 }
        ],
        doing: [
          { name: 'Mission', id: 8 },
          { name: 'Mission', id: 9 },
          { name: 'Mission', id: 10 }
        ],
        done: [
          { name: 'Mission', id: 11 },
          { name: 'Mission', id: 12 },
          { name: 'Mission', id: 13 },
          { name: 'Mission', id: 14 }
        ]
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
  width: 1200px;
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

