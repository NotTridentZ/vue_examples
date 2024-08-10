import { ref } from "vue"

const getBlogs = () => {
    const blogs = ref([])
    const error = ref(null)

    const load = async () => {
        try {
            let data = await fetch('http://localhost:8000/blogs')
            if (!data.ok) {
                throw Error('could not fetch data')
            }
            blogs.value = await data.json()
        } catch (err) {
            error.value = err.message
            console.log(error.value)
        }
    }

    return { blogs, error, load }
}

export default getBlogs