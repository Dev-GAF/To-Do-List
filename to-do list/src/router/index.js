import { createRouter, createWebHistory } from 'vue-router';
import DefaultLayout from '../layouts/DefaultLayout.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {  
          path: '',
          name: 'home',
          component: () => import('../views/Home.vue'),
        },
        {
          path: 'note/:id?',
          name: 'note',
          component: () => import('../views/Note.vue'),
        },
        {
          path: 'checklist',
          name: 'checklist',
          component: () => import('../views/CheckList.vue'),
        }
      ]
    }
  ]
})

export default router
