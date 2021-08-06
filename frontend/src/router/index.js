import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/profile',
    name: 'About',
    component: About
  },
  {
    path: '/workouts',
    name: 'Workouts',
    component: About
  },
  {
    path: '/exercises',
    name: 'Exercises',
    component: About
  },
  {
    path: '/results',
    name: 'Results',
    component: About
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
