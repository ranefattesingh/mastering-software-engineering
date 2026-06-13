<script setup lang="ts">
import { Table, TableHead, TableHeader, TableRow } from '@/components/ui/table';

</script>

<template>
    <Table class="border rounded-md">
        <TableHeader>
            <TableRow>
                <TableHead>ID</TableHead>
                <TableHead>Title</TableHead>
                <TableHead>Created On</TableHead>
                <TableHead>Due On</TableHead>
                <TableHead>Status</TableHead>
            </TableRow>
        </TableHeader>
    </Table>
</template>