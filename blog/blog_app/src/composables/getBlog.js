import { ref } from "vue"

const getBlog = (id) => {
    const blog = ref(null)
    const error = ref(null)

    const load = async () => {
        try {
            let data = await fetch('http://localhost:8000/blogs/' + id)
            if (!data.ok) {
                throw Error('that post does not exists')
            }
            blog.value = await data.json()
        } catch (err) {
            error.value = err.message
            console.log(error.value)
        }
    }

    return { blog, error, load }
}

export default getBlog