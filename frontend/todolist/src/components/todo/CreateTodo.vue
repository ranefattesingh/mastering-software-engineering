<script setup lang="ts">
import type { DateValue } from '@internationalized/date'
import { getLocalTimeZone, today } from '@internationalized/date'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Field, FieldGroup, FieldLabel, FieldError } from '@/components/ui/field';
import { Input } from '@/components/ui/input';
import { InputGroup } from '@/components/ui/input-group';
import { Textarea } from '@/components/ui/textarea';
import { Button } from '@/components/ui/button';
import { useForm, Field as VeeField } from 'vee-validate'
import { Calendar } from '@/components/ui/calendar';
import Popover from '@/components/ui/popover/Popover.vue';
import { PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { ChevronDownIcon } from '@lucide/vue';

import { ref, type Ref } from 'vue'
import { z } from 'zod'
import type { TodoItem } from '@/types/todo-item';
import { toTypedSchema } from '@vee-validate/zod';


const date = ref(today(getLocalTimeZone())) as Ref<DateValue>
const open = ref(false)

const formSchema = z.object({
    id: z.number(),
    title: z.string(),
    description: z.string(),
    dueDate: z.string(),
    dueTime: z.string(),
    isCompleted:  z.boolean(),
    createDate: z.string()
})

const { handleSubmit, resetForm } = useForm({
    validationSchema: toTypedSchema(formSchema),
    initialValues: {
        id: -1,
        title: '',
        description: '',
        dueDate: date.value.toString(),
        dueTime: '',
        isCompleted: false,
        createDate: today(getLocalTimeZone()).toString()
    }
})

const handleSubmitClick = handleSubmit((data) => {
    console.log(data)

    resetForm()
})

</script>

<template>
        <Card class="w-full sm:max-w-md text-zinc-800">
            <CardHeader>
                <CardTitle class="text-blue-500">Create New Todo</CardTitle>
                <CardDescription class="text-indigo-400">
                    Create your next todo item.
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form id="form-create-todo" @submit.prevent="handleSubmitClick">
                    <FieldGroup>
                        <!-- Title of the todo -->
                         <VeeField name="title" v-slot="{ field, errors }">
                            <Field :data-invalid="!!errors.length">
                                <FieldLabel>
                                    Title
                                </FieldLabel>
                                <Input v-bind="field" :aria-invalid="!!errors.length" />
                                <FieldError v-if="errors.length" :errors="errors" />
                            </Field>
                        </VeeField>

                        <!-- Description of the todo -->
                         <VeeField name="description" v-slot="{ field, errors }">
                            <Field :data-invalid="!!errors.length">
                                <FieldLabel>
                                    Description
                                </FieldLabel>
                                <InputGroup>
                                    <Textarea v-bind="field" :aria-invalid="!!errors.length" />
                                     <FieldError v-if="errors.length" :errors="errors" />
                                </InputGroup>
                            </Field>
                        </VeeField>

                        <div class="flex gap-4">
                            <div class="flex-1">
                                <Field>
                                    <FieldLabel>Due Date</FieldLabel>
                                    <Popover v-model:open="open">
                                        <PopoverTrigger as-child>
                                            <Button variant="outline">
                                                {{ date ? date.toDate(getLocalTimeZone()).toLocaleDateString() : "Select date" }}
                                                <ChevronDownIcon />
                                            </Button>
                                        </PopoverTrigger>
                                        <PopoverContent>
                                            <Calendar 
                                                v-model="date"
                                                @update:model-value="(v) => {
                                                    if (v) {
                                                        date = v
                                                        open = false
                                                    }
                                                }"
                                            />
                                        </PopoverContent>
                                    </Popover>
                                </Field>
                            </div>
                             
                            <VeeField name="dueDate" v-slot="{ field, errors }">
                                <div class="flex-1">
                                    <Field>
                                        <FieldLabel>Time</FieldLabel>
                                        <Input v-bind="field" :aria-invalid="!!errors.length" />
                                        <FieldError v-if="errors.length" :errors="errors" />
                                    </Field>
                                </div>
                            </VeeField> 

                        </div>
                    </FieldGroup>
                </form>
            </CardContent>
            <CardFooter>
                <Field orientation="horizontal" class="flex justify-end">
                    <Button form="form-create-todo" type="button" variant="outline" @click="resetForm">Reset</Button>
                    <Button form="form-create-todo" type="submit" class="bg-blue-500 hover:bg-blue-600">Create</Button>
                </Field>
            </CardFooter>
        </Card>
</template>
