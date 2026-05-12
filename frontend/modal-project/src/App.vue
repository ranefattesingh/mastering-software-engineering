<template>
  <h1>{{ message }}</h1>
  <br>
  <input type="text" ref="name">
  <button @click="handleClick">Click Me</button>
  <br>
  <Modal />
</template>

<script>
import Modal from './components/Modal.vue'

export default {
  name: 'App',
  components: { Modal },
  data() {
    return {
      message: 'My First Vue App'
    }
  },
  methods: {
    handleClick() {
      this.message = `Hello, ${this.$refs.name.value}!`
      this.$refs.name.classList.add('active')
      this.$refs.name.focus()
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1 {
  border-bottom: 1px solid #ddd;
  display: inline-block;
  padding-bottom: 10px;
}

</style>
