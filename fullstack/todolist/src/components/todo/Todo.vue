<script setup lang="ts">
import {
  Item,
  ItemActions,
  ItemContent,
  ItemTitle,
} from '@/components/ui/item'
import { Checkbox } from '@/components/ui/checkbox'
import { Button } from '@/components/ui/button';
import { ChevronUp, ChevronDown, SquarePen, Trash2 } from '@lucide/vue'
import type { TodoItem } from '@/types/todo-item';

interface Props {
    todo?: TodoItem | null
}

const props = withDefaults(defineProps<Props>(), {
    todo: null
})

const handleStatusChange = () => {
    const id = props.todo?.Id ?? 0
    emit('toggle-status', id)
}

const emit = defineEmits<{
    (checkboxClick: 'toggle-status', id: number): void
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
                <ItemTitle :class="[isCompleted() ? 'line-through': '', 'px-2']">{{ todo?.Title }}</ItemTitle>
            </div>
        </ItemContent>
        <ItemActions>
            <Button 
                variant="outline" size="sm" 
                class="text-blue-600 hover:text-blue-700 hover:bg-blue-50 border-blue-600 hover:border-blue-700"
            >
                <SquarePen />
            </Button>
            <Button 
                variant="outline" size="sm"
                class="text-red-600 hover:text-red-700 hover:bg-red-50 border-red-600 hover:border-red-700"
            >
                <Trash2 />
            </Button>
            <Button 
                variant="outline" size="sm"
                  class="text-indigo-600 hover:text-indigo-700 hover:bg-indigo-50 border-indigo-600 hover:border-indigo-700"
            >
                <ChevronUp />
            </Button>
            <Button 
                variant="outline" size="sm"
                  class="text-indigo-600 hover:text-indigo-700 hover:bg-indigo-50 border-indigo-600 hover:border-indigo-700"

            >
                <ChevronDown />
            </Button>
        </ItemActions>
    </Item>
</template>