import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/Base.vue'
import Login from '../views/Login.vue'
import SignUp from '../views/SignUp.vue'


const routes = [
  { path: '/', component: HomeView },
  {path: '/signin', component: Login},
  {path: '/signup', component: SignUp},

]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router