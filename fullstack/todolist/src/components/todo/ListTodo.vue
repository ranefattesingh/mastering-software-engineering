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
        @toggle-status="(id) => emit('toggle-status', id)"
    />
</template>