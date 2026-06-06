<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod';
import { z } from 'zod'
import { useForm, Field as VeeField } from 'vee-validate'
import Card from '@/components/ui/card/Card.vue';
import CardHeader from '@/components/ui/card/CardHeader.vue';
import CardTitle from '@/components/ui/card/CardTitle.vue';
import CardDescription from '@/components/ui/card/CardDescription.vue';
import CardContent from '@/components/ui/card/CardContent.vue';
import Field from '@/components/ui/field/Field.vue';
import FieldLabel from '@/components/ui/field/FieldLabel.vue';

const registrationFormSchema = toTypedSchema(
    z.object({
        fullName: z
            .string()
            .min(3, 'Full name must be atleast 3 characters long'),
        email: z.email('Enter a valid email address'),
        password: z
            .string()
            .min(8, 'Password must be at least 8 characters long'),
        confirmPassword: z.string()
    }).refine(d => d.password === d.confirmPassword, {
        message: 'Passwords do not match',
        path: ['ConfirmPassword']
    })
)

const { handleSubmit, resetForm } = useForm({
    validationSchema: registrationFormSchema,
    initialValues: {
        fullName: '',
        email: '',
        password: '',
        confirmPassword: ''
    }
})

const onSubmit = handleSubmit(values => {
    console.log('Form Submitted', values)
    resetForm()
})

</script>

<template>
    <div class="flex justify-center items-center min-h-screen p-4">
        <Card class="w-full max-w-md">
            <CardHeader class="text-center">
                <CardTitle class="text-xl">Registration Page</CardTitle>
                <CardDescription class="">Enter your details to register your account</CardDescription>
            </CardHeader>

            <CardContent>
                <form @submit="onSubmit" id="registrationForm">
                    <FieldGroup>
                        <VeeField v-slot="{ field, errors }" name="fullName">
                            <Field :data-invalid="!!errors.length">
                                <FieldLabel for="registrationFormFullName">
                                    Full Name
                                </FieldLabel>
                                <Input  id="registrationFormFullName"
                                        v-bind="field"
                                        placeholder="Your Name" 
                                        autocomplete="off"
                                        :aria-invalid="!!errors.length"
                                />
                                <FieldError v-if="errors.length" :errors="errors" />
                            </Field>
                        </VeeField>

                        <VeeField v-slot="{ field, errors }" name="email">
                            <Field :data-invalid="!!errors.length">
                                <FieldLabel for="registrationFormEmail">
                                    Email
                                </FieldLabel>
                                <Input  id="registrationFormEmail"
                                        v-bind="field"
                                        placeholder="your.name@example.com" 
                                        autocomplete="off"
                                        :aria-invalid="!!errors.length"
                                />
                                <FieldError v-if="errors.length" :errors="errors" />
                            </Field>
                        </VeeField>
                    </FieldGroup>
                </form>
            </CardContent>
        </Card>
    </div>
</template>