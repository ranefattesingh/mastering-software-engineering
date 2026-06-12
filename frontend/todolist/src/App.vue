<script setup lang="ts">
import type { DateValue } from '@internationalized/date'
import { getLocalTimeZone, today } from '@internationalized/date'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Field, FieldGroup, FieldLabel } from '@/components/ui/field';
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


const date = ref(today(getLocalTimeZone())) as Ref<DateValue>
const open = ref(false)

</script>

<template>
    <div class="flex h-screen items-center justify-center pb-24">
        <Card class="w-full sm:max-w-md text-zinc-800">
            <CardHeader>
                <CardTitle class="text-blue-500">Create New Todo</CardTitle>
                <CardDescription class="text-indigo-400">
                    Create your next todo item.
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form>
                    <FieldGroup>
                        <!-- Title of the todo -->
                        <Field>
                            <FieldLabel>
                                Title
                            </FieldLabel>
                            <Input />
                        </Field>

                        <!-- Description of the todo -->
                        <Field>
                            <FieldLabel>
                                Description
                            </FieldLabel>
                            <InputGroup>
                                <Textarea />
                            </InputGroup>
                        </Field>

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
                            <div class="flex-1">
                                <Field>
                                    <FieldLabel>Time</FieldLabel>
                                    <Input />
                                </Field>
                            </div>
                        </div>
                    </FieldGroup>
                </form>
            </CardContent>
            <CardFooter>
                <Field orientation="horizontal" class="flex justify-end">
                    <Button type="button" variant="outline">Reset</Button>
                    <Button type="button" class="bg-blue-500 hover:bg-blue-600">Create</Button>
                </Field>
            </CardFooter>
        </Card>
    </div>
</template>
