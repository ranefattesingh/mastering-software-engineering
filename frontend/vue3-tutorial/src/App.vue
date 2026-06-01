<template>
  <h1 :class="titleClass">{{ message }}</h1>
  <p>Count is: {{ counter.count }}</p>

  <div :id="dynamicId"></div>
  <button @click="increment">{{ count }}</button>
  <input type="text" :value="text" @input="onInput">
  <p>{{ text }}</p>

  <button @click="toggle">Toggle</button>
  <h1 v-if="awesome">Vue is awesome!</h1>
  <h1 v-else>Oh no 😢</h1>

  <form @submit.prevent="addTodo">
    <input v-model="newTodo" required placeholder="new todo">
    <button>Add Todo</button>
  </form>
  <ul>
    <li v-for="todo in filteredTodos" :key="todo.id">
      <input type="checkbox" v-model="todo.done">
      <span :class="{ done: todo.done }">{{ todo.text }}</span>
      <button @click="removeTodo(todo)">X</button>
    </li>
  </ul>

    <button @click="hideCompleted = !hideCompleted">
    {{ hideCompleted ? 'Show all' : 'Hide completed' }}
  </button>

  <p ref="pElementRef">hello</p>
  <ChildComponent :msg="greeting" @response="(msg) => childMsg = msg"/>
  <ChildComponent>
     This is some slot content!
  </ChildComponent>
  <p>{{ childMsg }}</p>
</template>

<script setup>
  import { ref, reactive, computed, onMounted, watch } from 'vue'
import ChildComponent from './components/ChildComponent.vue'

  const counter = reactive({
    count: 0
  })

  const dynamicId = ref('')

  const titleClass = ref('title')
  const count = ref(0)
  const text = ref('')

  console.log(counter.count)
  counter.count++

  const message = ref('Hello World!')

  console.log(message.value)

  message.value = 'My New Title'

  function increment() {
    count.value++
  }

  function onInput(e) {
    text.value = e.target.value
  }

  const awesome = ref(true)

  function toggle() {
    awesome.value = !awesome.value
  }


  let id = 0

  const newTodo = ref('')
  const todos = ref([
    { id: id++, text: 'Learn HTML' },
    { id: id++, text: 'Learn JavaScript' },
    { id: id++, text: 'Learn Vue' }
  ])
  const hideCompleted = ref(false)
  const filteredTodos = computed(() => {
    return hideCompleted.value
      ? todos.value.filter((t) => !t.done)
      : todos.value
  })

  function addTodo() {
    todos.value.push({ id: id++, text: newTodo.value })
    newTodo.value = ''
  }

  function removeTodo(todo) {
    todos.value = todos.value.filter((t) => t !== todo)
  }

  const pElementRef = ref(null)
  onMounted(() => {
    pElementRef.value.textContent = "newly updated text"
  })

  watch(count, (newCount) => {
    // yes, console.log() is a side effect
    console.log(`new count is: ${newCount}`)
  })

  const greeting = ref('Hello from parent')
  const childMsg = ref('No child msg yet')

</script>

<style>
.title {
  color: red;
}

.done {
  text-decoration: line-through;
}
</style>