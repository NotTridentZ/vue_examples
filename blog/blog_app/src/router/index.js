import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CreateView from '../views/CreateView.vue'
import BlogDetails from '../blogs/BlogDetails.vue'
import NotFoundView from '../views/NotFoundView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/create',
    name: 'Create',
    component: CreateView
  },
  {
    path: '/blogs/:id',
    name: 'BlogDetails',
    component: BlogDetails,
    props: true
  },
  // redirect
  {
    path: '/all-blogs',
    redirect: '/'
  },
  // 404 catchall
  {
    path: '/:catchAll(.*)',
    name: 'NotFound',
    component: NotFoundView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
