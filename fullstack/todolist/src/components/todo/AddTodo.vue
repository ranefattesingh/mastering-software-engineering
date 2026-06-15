<script setup lang="ts">
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { ref } from 'vue';

const title =  ref<string>('')

const emit = defineEmits<{
    (addTodo: 'add-todo', item: string) : void
}>()

const handleSubmitClick =() => {
    if (!title.value.trim()) return

    emit('add-todo', title.value)

    title.value = ''
}

</script>
<template>
    <form @submit.prevent="handleSubmitClick">
        <Input class="my-4" placeholder="What's up in your mind?" v-model="title"/>
        <Button class="w-full bg-blue-600 hover:bg-blue-700">Add</Button>
    </form>
</template>