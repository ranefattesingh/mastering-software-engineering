const app = Vue.createApp({
    // Setup any function or data
    // define template in string => html
    // template: '<h2> This is a template </h2>'
    data() {
        return {
            title: 'This is a book title',
            author: 'John Doe',
            age: 30,
            showBooks: true,
        }
    },
    methods: {
        changeTitle() {
            console.log("you clicked the button ")
            this.title = 'This is a new book title'
        },
        toggleShowBooks() {
            this.showBooks = !this.showBooks
        }
    }
})

app.mount('#app')