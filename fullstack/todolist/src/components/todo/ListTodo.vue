<script setup lang="ts">

import Todo from '@/components/todo/Todo.vue';
import type { TodoItem } from '@/types/todo-item';
import {onMounted } from 'vue';

interface Props {
    todos: TodoItem[]
}

const props = withDefaults(defineProps<Props>(), {
  todos: () => [] // Array defaults must be returned from a factory function
})

const emit = defineEmits<{
    (checkboxClick: 'toggle-status', id: number): void
    (deleteTodo: 'delete-todo', id: number): void
    (deleteTodo: 'edit-save-todo', id: number, newTitle: string): void
    (moveTodoUp: 'move-todo-up', id: number): void
    (moveTodoDown: 'move-todo-down', id: number): void
}>()

onMounted(() => {
    // Make backend call to webserver to fetch all todos
    console.log(props.todos, props.todos.length)
})

</script>
<template>
    <Todo 
        v-for="todo in todos" 
        :key="todo.Id" 
        :todo="todo"
        :total="todos.length" 
        @toggle-status="(id) => emit('toggle-status', id)"
        @delete-todo="(id) => emit('delete-todo', id)"
        @edit-save-todo="(id, newTitle) => emit('edit-save-todo', id, newTitle)"
        @move-todo-up="(id) => emit('move-todo-up', id)"
        @move-todo-down="(id) => emit('move-todo-down', id)"
    />
</template>