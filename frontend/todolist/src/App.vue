<script setup lang="ts">
import CreateTodo from '@/components/todo/CreateTodo.vue';
import TodolistTable from '@/components/todo/TodolistTable.vue';

</script>

<template>
    <div class="md:flex m-5 gap-5">
        <div class="flex-1">
            <CreateTodo />
        </div>
        <div class="flex-2">
            <TodolistTable />
        </div>
    </div>
</template>