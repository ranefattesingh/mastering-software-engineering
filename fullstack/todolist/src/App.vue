<script setup lang="ts">
import AddTodo  from '@/components/todo/AddTodo.vue';
import ListTodo from '@/components/todo/ListTodo.vue';
import type { TodoItem } from '@/types/todo-item';
import { ref } from 'vue';

const todos = ref<TodoItem[]>([])

function handleAddTodo(newTodo: string) {
    const todo: TodoItem = {
        Id: todos.value.length + 1,
        Title: newTodo,
        Order: todos.value.length,
        Status: false
    }

    todos.value.push(todo)
}

function handleStatusChange(id: number) {
    for (let i = 0; i < todos.value.length; i++) {
        if (todos.value[i].Id === id) {
            todos.value[i].Status = !todos.value[i].Status
        }
    }
}

function deleteTodo(id: number) {
    const index = todos.value.findIndex(todo => todo.Id === id)
    if (isNaN(index) && index > 0) return


    todos.value.splice(index, 1)
}

</script>

<template>
    <div class="grid grid-cols-9">
        <div class="col-span-3"></div>
        <div  class="col-span-4 grid grid-cols-3">
            <div class="col-span-2">
                <AddTodo @add-todo="handleAddTodo"/>
            </div>
            <div class="col-span-1"></div>
        </div>

        <div class="col-span-3"></div>

        <div class="col-span-4">
            <ListTodo 
                :todos="todos" 
                @toggle-status="handleStatusChange" 
                @delete-todo="deleteTodo"
            />
        </div>
    </div>
</template>