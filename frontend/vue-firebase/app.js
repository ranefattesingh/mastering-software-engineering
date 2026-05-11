const app = Vue.createApp({
    // Setup any function or data
    // define template in string => html
    // template: '<h2> This is a template </h2>'
    data() {
        return {
            // Tutorial 2
            title: 'This is a book title',
            author: 'John Doe',
            age: 30,
            showBooks: true,

            // Tutorial 3
            x: 0,
            y: 0,
            books: [
                { title: 'Book One', author: 'John Doe', img: 'assets/1.jpg', isFav: true },
                { title: 'Book Two', author: 'Jane Smith', img: 'assets/2.jpg', isFav: false },
                { title: 'Book Three', author: 'John Doe', img: 'assets/3.jpg', isFav: true }
            ],
            url: 'https://github.com/ranefattesingh'
        }
    },
    methods: {
        // Tutorial 2
        changeTitle() {
            console.log("you clicked the button ")
            this.title = 'This is a new book title'
        },
        toggleShowBooks() {
            this.showBooks = !this.showBooks
        },

        // Tutorial 3
        handleEvent(e, data) {
            console.log(e, e.type)
            if (data) {
                console.log(data)
            }
        },
        handleMousemove(e) {
            this.x = e.offsetX
            this.y = e.offsetY 
        },
        toggleFav(book) {
            book.isFav = !book.isFav
        }
    },
    computed: {
        filteredBooks() {
            return this.books.filter(book => book.isFav)
        }
    }
})

app.mount('#app')