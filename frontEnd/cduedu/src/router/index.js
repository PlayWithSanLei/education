import { createRouter, createWebHistory } from 'vue-router'
const Signup = () => import('../components/Signup')

const routes = [
  {
    path: '',
    redirect: '/signup'
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
