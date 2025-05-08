import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/note',
      name: 'note',
      component: () => import('../views/Note.vue'),
    },
    {
      path: '/checklist',
      name: 'checklist',
      component: () => import('../views/CheckList.vue'),
    }
  ],
})

export default router
