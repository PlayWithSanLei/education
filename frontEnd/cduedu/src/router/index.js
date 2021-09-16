import { createRouter, createWebHistory } from 'vue-router'
const Signup = () => import('../components/Signup')
const Login = () => import('../components/Login')
const Home = () => import('../components/Home')
const User = () => import('../components/User')
const Org = () => import('../components/Org')
const Acc = () => import('../components/Acc')
const Map = () => import('../components/Map')

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
    redirect: '/user',
    component: Home,
    children: [
      {
        path: '/user',
        name: 'User',
        component: User
      },
      {
        path: '/org',
        name: 'Org',
        component: Org
      },
      {
        path: '/acc',
        name: 'Acc',
        component: Acc
      },
      {
        path: '/map',
        name: 'Map',
        component: Map
      }
    ]

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
