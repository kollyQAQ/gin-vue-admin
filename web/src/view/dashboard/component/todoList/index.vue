<template>
    <section class="todoapp">
        <!-- header -->
        <header class="header">
            <svg-icon icon-class="edit" class-name="card-panel-icon" style="padding-left: 14px;"/>
            <input class="new-todo" autocomplete="off" placeholder="请输入待办事项" @keyup.enter="addTodo">
        </header>
        <!-- main section -->
        <section v-show="todos.length" class="main">
            <ul class="todo-list">
                <todo
                        v-for="(todo, index) in filteredTodos"
                        :key="index"
                        :todo="todo"
                        @toggleTodo="toggleTodo"
                        @editTodo="editTodo"
                        @deleteTodo="deleteTodo"
                />
            </ul>
        </section>
        <!-- footer -->
        <footer v-show="todos.length" class="footer">
      <span class="todo-count">
        <strong>{{ remaining }}</strong>
        {{ remaining | pluralize('item') }} left
      </span>
            <ul class="filters">
                <li v-for="(val, key) in filters" :key="key">
                    <a :class="{ selected: visibility === key }" @click.prevent="visibility = key">{{ key | capitalize }}</a>
                </li>
            </ul>
        </footer>
    </section>
</template>

<script>
    import Todo from './Todo.vue'
    import {
        queryTodoList,
        addTodo,
        updateTodo,
        deleteTodo,
    } from '@/api/stat'

    const STORAGE_KEY = 'todos'
    const filters = {
        all: todos => todos,
        active: todos => todos.filter(todo => !todo.done),
        completed: todos => todos.filter(todo => todo.done)
    }
    export default {
        components: { Todo },
        filters: {
            pluralize: (n, w) => n === 1 ? w : w + 's',
            capitalize: s => s.charAt(0).toUpperCase() + s.slice(1)
        },
        data() {
            return {
                visibility: 'active',
                filters,
                // todos: JSON.parse(window.localStorage.getItem(STORAGE_KEY)) || defalutList
                todos: []
            }
        },
        computed: {
            allChecked() {
                return this.todos.every(todo => todo.done)
            },
            filteredTodos() {
                return filters[this.visibility](this.todos)
            },
            remaining() {
                return this.todos.filter(todo => !todo.done).length
            }
        },
        methods: {
            setLocalStorage() {
                window.localStorage.setItem(STORAGE_KEY, JSON.stringify(this.todos))
            },
            async addTodo(e) {
                const text = e.target.value
                if (text.trim()) {
                    const res = await addTodo({
                        text: text,
                        status: 0,
                    })
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '添加成功',
                            showClose: true
                        })
                        const res = await queryTodoList();
                        this.todos = res.data;
                    }
                }
                e.target.value = ''
            },
            async toggleTodo(val) {
                val.done = !val.done
                val.status = (val.status + 1) % 2
                const res = await updateTodo(val)
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '修改成功',
                        showClose: true
                    })
                }
            },
            async deleteTodo(todo) {
                const res = await deleteTodo(todo)
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '删除成功',
                        showClose: true
                    })
                    const res = await queryTodoList();
                    this.todos = res.data;
                }
            },
            async editTodo({ todo, value }) {
                todo.text = value
                const res = await updateTodo(todo)
                if (res.code === 0) {
                    this.$message({
                        type: 'success',
                        message: '修改成功',
                        showClose: true
                    })
                }
            },
            clearCompleted() {
                this.todos = this.todos.filter(todo => !todo.done)
                this.setLocalStorage()
            },
            toggleAll({ done }) {
                this.todos.forEach(todo => {
                    todo.done = done
                    this.setLocalStorage()
                })
            }
        },
        async created(){
            const res = await queryTodoList();
            this.todos = res.data;
        }
    }
</script>

<style lang="scss">
    @import './index.scss';
</style>
