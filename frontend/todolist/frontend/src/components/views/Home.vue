<script setup lang="ts">
import { ref } from 'vue';
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import type { Todo } from '@/types/todo';

    const todos  = ref<Todo[]>([
        {
            id: 1,
            title:  'clean bathroom',
            createDate: '9/6/2026',
            dueDate: '09-06-2026',
            completed:  false
        },
        {
            id: 2,
            title:  'change  bedsheets',
            createDate: '9/6/2026',
            dueDate: '09-06-2026',
            completed:  true
        },
        {
            id: 3,
            title:  'wash bedsheets',
            createDate: '9/6/2026',
            dueDate: '09-06-2026',
            completed:  true
        }
    ])

    const isDueToday = (todo: Todo): boolean => {
        const today = new Date()
        const mm: string = String(today.getMonth() + 1).padStart(2, '0'); // Months are 0-indexed
        const dd: string = String(today.getDate()).padStart(2, '0');
        const yyyy: number = today.getFullYear();
        const formattedDate: string = `${dd}-${mm}-${yyyy}`;
        console.log(formattedDate)
        return formattedDate == todo.dueDate && !todo.completed
    }

    const isCompleted = (todo: Todo): boolean => todo.completed
    
    
</script>

<template> 
    <div class="container mx-auto">
        <Table class="table-auto md:table-fixed">
            <TableCaption>Your Todolist</TableCaption>
            <TableHeader class="text-cente">
                <TableRow >
                    <TableHead class="text-center">ID</TableHead>
                    <TableHead class="text-center">Title</TableHead>
                    <TableHead class="text-center">Created On</TableHead>
                    <TableHead class="text-center">Due On</TableHead>
                    <TableHead class="text-center">Completed</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody class="text-center">
                <TableRow 
                    v-for="todo in todos" 
                    :key="todo.id"
                        :class="isCompleted(todo) 
    ? 'bg-green-200 hover:bg-green-300' 
    : (isDueToday(todo) ? 'bg-red-200 hover:bg-red-300' : '')"
                    >
                    <TableCell>{{ todo.id }}</TableCell>
                    <TableCell>{{ todo.title }}</TableCell>
                    <TableCell>{{ todo.createDate }}</TableCell>
                    <TableCell>{{ todo.dueDate }}</TableCell>
                    <TableCell>{{ todo.completed }}</TableCell>
                </TableRow>
            </TableBody>
            <TableFooter>
                <TableRow>
                    <TableCell>
                        Total Completed
                    </TableCell>
                    <TableCell></TableCell>
                    <TableCell></TableCell>
                    <TableCell></TableCell>
                    <TableCell>{{ todos.filter(todo => todo.completed).length }}</TableCell>
                </TableRow>
            </TableFooter>
        </Table>
    </div>
</template>