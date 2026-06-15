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

function handleUpdateTodo(id: number, newTitle: string) {
    const index = todos.value.findIndex(todo => todo.Id === id)
    if (isNaN(index) && index > 0) return


    todos.value[index].Title = newTitle
}

function handleDeleteTodo(id: number) {
    const index = todos.value.findIndex(todo => todo.Id === id)
    if (isNaN(index) && index > 0) return


    todos.value.splice(index, 1)
}

function handleMoveTodoUp(id: number) {
    const index = todos.value.findIndex(todo => todo.Id === id)
    if (index <= 0) return
    handleMoveTodo(index, index-1)
}

function handleMoveTodoDown(id: number) {
    const index = todos.value.findIndex(todo => todo.Id === id)
    if (index >= todos.value.length - 1) return
    handleMoveTodo(index, index+1)
}

function handleMoveTodo(src: number, dest: number) {
    const [movedItem] = todos.value.splice(src, 1)
    todos.value.splice(dest, 0, movedItem)
    const srcOrder = todos.value[src].Order
    todos.value[src].Order = todos.value[dest].Order 
    todos.value[dest].Order = srcOrder

    console.log(todos.value)

    
}



</script>

<template>
    <div class="sm:grid sm:grid-cols-9">
        <div class="sm:col-span-3"></div>
        <div  class="sm:col-span-4 sm:grid sm:grid-cols-3">
            <div class="sm:col-span-2">
                <AddTodo @add-todo="handleAddTodo"/>
            </div>
            <div class="col-span-1"></div>
        </div>

        <div class="col-span-3"></div>

        <div class="col-span-4">
            <ListTodo 
                :todos="todos" 
                @toggle-status="handleStatusChange" 
                @delete-todo="handleDeleteTodo"
                @edit-save-todo="handleUpdateTodo"
                @move-todo-up="handleMoveTodoUp"
                @move-todo-down="handleMoveTodoDown"
            />
        </div>
    </div>
</template>