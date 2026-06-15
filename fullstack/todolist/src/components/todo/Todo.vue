<script setup lang="ts">
import {
  Item,
  ItemActions,
  ItemContent,
  ItemTitle,
} from '@/components/ui/item'
import { Checkbox } from '@/components/ui/checkbox'
import { Button } from '@/components/ui/button';
import { ChevronUp, ChevronDown, SquarePen, Trash2, SaveIcon } from '@lucide/vue'
import type { TodoItem } from '@/types/todo-item';
import { Input } from '@/components/ui/input';
import { ref } from 'vue';

interface Props {
    todo?: TodoItem | null
    total: number
}

const props = withDefaults(defineProps<Props>(), {
    todo: null
})

const isEditing = ref(false)
const newTitle = ref(props.todo?.Title ?? '')


const handleStatusChange = () => {
    // Call HTTP API PATCH endpoint
    const id = props.todo?.Id ?? 0
    emit('toggle-status', id)
}

const handleDeleteTodo = () => {
    // Call HTTP API DELETE endpoint
    const id = props.todo?.Id ?? 0
    emit('delete-todo', id)
}

const toggleEdit = () => isEditing.value = !isEditing.value

const  handleSaveClick = () => {
    // Call HTTP API PUT endpoint
    const id = props.todo?.Id ?? 0
    console.log(newTitle)
    emit('edit-save-todo', id, newTitle.value)
    toggleEdit()
}

const handleMoveUp = () => {
    // Call HTTP API PUT endpoint
    const id = props.todo?.Id ?? 0
    emit('move-todo-up', id)
}

const handleMoveDown = () => {
    // Call HTTP API PUT endpoint
    const id = props.todo?.Id ?? 0
    emit('move-todo-down', id)
}


const emit = defineEmits<{
    (checkboxClick: 'toggle-status', id: number): void
    (deleteClick: 'delete-todo', id: number): void
    (editSaveClick: 'edit-save-todo', id: number, newTitle: string): void
    (moveTodoUp: 'move-todo-up', id: number): void
    (moveTodoUp: 'move-todo-down', id: number): void
}>()

function isCompleted(): boolean {
    const status = props.todo?.Status ?? false

    return status
}

</script>

<template>
    <Item variant="outline" class="my-4">
        <ItemContent>
            <div class="flex items-center">
                <Checkbox id="terms"
                    :model-value="props.todo?.Status ?? false"
                    @update:model-value="handleStatusChange"
                    class="-4 w-4 appearance-none rounded-full border border-gray-400 checked:bg-blue-600"
                />
                <ItemTitle :class="[isCompleted() ? 'line-through': '', 'px-2']">
                    <span v-if="!isEditing">{{ todo?.Title }}</span>
                    <Input v-else 
                        v-model="newTitle"
                    />
                </ItemTitle>
            </div>
        </ItemContent>
        <ItemActions>
            <Button 
                @click="toggleEdit"
                v-if="!isEditing"
                variant="outline" size="sm" 
                class="text-blue-600 hover:text-blue-700 hover:bg-blue-50 border-blue-600 hover:border-blue-700"
            >
                <SquarePen/>
            </Button>
            <Button 
                @click="handleSaveClick"
                v-else 
                variant="outline" size="sm" 
                class="text-green-600 hover:text-green-700 hover:bg-green-50 border-green-600 hover:border-green-700"
            >
                <SaveIcon />
            </Button>
            <Button 
                @click="handleDeleteTodo"
                variant="outline" size="sm"
                class="text-red-600 hover:text-red-700 hover:bg-red-50 border-red-600 hover:border-red-700"
            >
                <Trash2 />
            </Button>
            <Button 
                :disabled="(total === 1) || (todo?.Order == 0)"
                @click="handleMoveUp"
                variant="outline" size="sm"
                  class="text-indigo-600 hover:text-indigo-700 hover:bg-indigo-50 border-indigo-600 hover:border-indigo-700"
            >
                <ChevronUp />
            </Button>
            <Button 
                :disabled="(total == 1) || (todo?.Order == total-1)"
                @click="handleMoveDown"
                variant="outline" size="sm"
                  class="text-indigo-600 hover:text-indigo-700 hover:bg-indigo-50 border-indigo-600 hover:border-indigo-700"

            >
                <ChevronDown />
            </Button>
        </ItemActions>
    </Item>
</template>