<template>
    <h1>I am a child component</h1>
    <p>{{ props.msg }}</p>
     <slot>Fallback content</slot>
</template>

<script setup>
  const props = defineProps({
    msg: String
  })

  const emit = defineEmits(['response'])
  emit('response', 'hello from child')
</script>