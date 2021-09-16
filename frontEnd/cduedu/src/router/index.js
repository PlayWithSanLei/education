import { createRouter, createWebHistory } from 'vue-router'
const Signup = () => import('../components/Signup')
const Login = () => import('../components/Login')
const Home = () => import('../components/Home')

const routes = [
  {
    path: '',
    redirect: '/login'
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/home',
    name: 'Home',
    component: Home
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

/*
* 添加路由守卫,对token判断
* */
router.beforeEach((to, from, next) => {
  if (to.path === '/login' || to.path === '/signup') {
    next();
  } else {
    let token  = localStorage.getItem('token');
    if (token === null || token === '' || token === undefined) {
      next('/login')
    } else {
      next()
    }
  }
})

export default router
